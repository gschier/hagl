package hagl

import (
	"fmt"
	"html"
	"io"
	"strings"
)

func newEl() *RawNode {
	return &RawNode{
		nodeType:        tagNode,
		styles:          make(map[string]string),
		tab:             "  ",
		indentIncrement: 1,
	}
}

func newTagNode(tag string) *RawNode {
	el := newEl()
	el.tag = tag
	return el
}

func newSelfClosingTagNode(tag string) *RawNode {
	el := newTagNode(tag)
	el.selfClosing = true
	return el
}

func newPreserveWhitespaceTagNode(tag string) *RawNode {
	el := newTagNode(tag)
	el.preformatted = true
	return el
}

type attr struct {
	name  string
	value string
}

type nodeType int

const (
	textNode nodeType = iota
	tagNode
	commentNode
	fragmentNode
)

type Node interface {
	ID(id string) Node
	Children(child ...Node) Node
	Range(n int, child func(i int) Node) Node
	Text(text ...string) Node
	Textf(format string, a ...interface{}) Node
	AttrBool(name string) Node
	Attr(name, value string) Node
	AttrIf(cond bool, name, value string) Node
	Class(cls ...string) Node
	ClassIf(condition bool, cls string) Node
	Style(name, value string) Node
	HTML() string
	HTMLPretty() string
	Write(w io.Writer) (int, error)
	WritePretty(w io.Writer) (int, error)
	MustWrite(w io.Writer)
	MustWritePretty(w io.Writer)
	Extend(base Node) Node
	If(c bool) Node

	// Helpers

	Href(value string) Node
	Rel(value string) Node
	Src(value string) Node
	Target(value string) Node
	Name(value string) Node
	Action(value string) Node
	Method(value string) Node
	Alt(value string) Node
	Type(value string) Node
	Title(value string) Node

	// GetNode returns the root node. This is for internal use only
	GetNode() *RawNode
}

type RawNode struct {
	// children contains the child elements for the element
	children []Node

	// attrs contains the attributes for the element
	attrs []attr

	// styles hold the styles assigned to the element, which will
	// be converted to attributes when rendered
	styles map[string]string

	// tag defines the name of the HTML element
	tag string

	// text contains the text contents of the element
	text string

	// nodeType defines the type of node
	nodeType nodeType

	// preformatted is used for tags like <pre> that shouldn't
	// have the whitespace of their children changed in any way
	preformatted bool

	// selfClosing defines whether or not the element can close itself.
	//
	// For example, a horizontal rule or input can <hr/> <input/>
	selfClosing bool

	// indentIncrement specifies the amount of indents the element should add.
	// This is most useful for Fragment, which has no indent.
	indentIncrement int

	// tab specifies the character used to indent
	tab  string
	hide bool
}

func (rn *RawNode) ID(id string) Node {
	return rn.Attr("id", id)
}

func (rn *RawNode) Href(value string) Node {
	return rn.Attr("href", value)
}

func (rn *RawNode) Rel(value string) Node {
	return rn.Attr("rel", value)
}

func (rn *RawNode) Src(value string) Node {
	return rn.Attr("src", value)
}

func (rn *RawNode) Target(value string) Node {
	return rn.Attr("target", value)
}

func (rn *RawNode) Alt(value string) Node {
	return rn.Attr("alt", value)
}

func (rn *RawNode) Type(value string) Node {
	return rn.Attr("type", value)
}

func (rn *RawNode) Title(value string) Node {
	return rn.Attr("title", value)
}

func (rn *RawNode) Name(value string) Node {
	return rn.Attr("name", value)
}

func (rn *RawNode) Action(value string) Node {
	return rn.Attr("action", value)
}

func (rn *RawNode) Method(value string) Node {
	return rn.Attr("method", value)
}

func (rn *RawNode) GetNode() *RawNode {
	return rn
}

func (rn *RawNode) Children(child ...Node) Node {
	for _, c := range child {
		// Skip nil Nodes
		if c == nil {
			continue
		}

		// Preformatted nodes have preformatted children
		if rn.preformatted {
			c.GetNode().preformatted = true
		}

		rn.children = append(rn.children, c)
	}
	return rn
}

func (rn *RawNode) If(c bool) Node {
	if !c {
		rn.hide = true
	}

	return rn
}

func (rn *RawNode) Extend(node Node) Node {
	n := node.GetNode()

	rn.children = append(n.children, rn.children...)
	rn.hide = n.hide

	for _, attr := range n.attrs {
		found := false
		for i, existingAttr := range rn.attrs {
			if existingAttr.name == attr.name && existingAttr.name == "class" {
				rn.attrs[i].value = existingAttr.value + " " + attr.value
				found = true
				break
			}
		}

		if !found {
			rn.attrs = append(rn.attrs, attr)
		}
	}

	rn.text = ""
	rn.Text(n.text)
	rn.Text(rn.text)

	return rn
}

// Range is a convenience used to generate n children based on a factory function.
// the factory will be called n times and will skip any nil children
func (rn *RawNode) Range(n int, child func(i int) Node) Node {
	for i := 0; i < n; i++ {
		rn.Children(child(i))
	}
	return rn
}

// Text is a helper method to add a text node to children
func (rn *RawNode) Text(text ...string) Node {
	return rn.Children(Text(text...))
}

// Textf is the same as Text, but accepts fmt args
func (rn *RawNode) Textf(format string, a ...interface{}) Node {
	return rn.Children(Text(fmt.Sprintf(format, a...)))
}

func (rn *RawNode) Attr(name, value string) Node {
	for i, a := range rn.attrs {
		if a.name == name {
			rn.attrs[i] = attr{name: name, value: value}
			return rn
		}
	}

	rn.attrs = append(rn.attrs, attr{name: name, value: value})

	return rn
}

func (rn *RawNode) AttrIf(cond bool, name, value string) Node {
	if cond {
		return rn.Attr(name, value)
	} else {
		return rn
	}
}

func (rn *RawNode) AttrBool(name string) Node {
	return rn.Attr(name, name)
}

func (rn *RawNode) Class(cls ...string) Node {
	existingClasses := strings.Fields(rn.attr("class"))

	for _, newCls := range cls {
		found := false
		for _, c := range existingClasses {
			if c == newCls {
				found = true
				break
			}
		}

		if !found {
			existingClasses = append(existingClasses, newCls)
		}
	}

	rn.Attr("class", strings.Join(existingClasses, " "))
	return rn
}

func (rn *RawNode) ClassIf(condition bool, cls string) Node {
	if !condition {
		return rn
	}

	return rn.Class(cls)
}

// Style is a utility method to append to the style attribute. If a style
// attribute already exists, the new style will be appended.
func (rn *RawNode) Style(name, value string) Node {
	str := name + ":" + value
	for i, a := range rn.attrs {
		if a.name == "style" {
			rn.attrs[i].value += ";" + str
			return rn
		}
	}

	// No style attr, so add one
	rn.Attr("style", str)
	return rn
}

func (rn *RawNode) HTML() string {
	return rn.html(-1, false)
}

func (rn *RawNode) HTMLPretty() string {
	return strings.TrimSpace(rn.html(0, true))
}

func (rn *RawNode) Write(w io.Writer) (int, error) {
	return w.Write([]byte(rn.HTML()))
}

func (rn *RawNode) MustWrite(w io.Writer) {
	_, err := w.Write([]byte(rn.HTML()))
	if err != nil {
		panic(err)
	}
}

func (rn *RawNode) WritePretty(w io.Writer) (int, error) {
	return w.Write([]byte(rn.HTMLPretty()))
}

func (rn *RawNode) MustWritePretty(w io.Writer) {
	_, err := w.Write([]byte(rn.HTMLPretty()))
	if err != nil {
		panic(err)
	}
}

func (rn *RawNode) indent(level int, text string) string {
	prefix := strings.Builder{}
	for i := 0; i < level; i++ {
		prefix.WriteString(rn.tab)
	}

	prefix.WriteString(text)
	return prefix.String()
}

func (rn *RawNode) html(level int, prettify bool) string {
	// Nothing to do for hidden nodes
	if rn.hide {
		return ""
	}

	innerHTML := ""
	onlyTextChildren := true

	// Render children if the element has them
	for _, c := range rn.children {
		if c.GetNode().nodeType != textNode {
			onlyTextChildren = false
		}

		if rn.preformatted {
			// Indent open tag but nothing else
			// TODO: Figure out what to do with tags inside <pre>
			innerHTML += rn.indent(0, c.GetNode().html(0, false))
		} else {
			innerHTML += c.GetNode().html(level+rn.indentIncrement, prettify)
		}

		// Add newline after each child if we're prettifying. Note, we don't
		// add one to fragment children because they don't take up space
		if prettify && !rn.preformatted && c.GetNode().nodeType != fragmentNode {
			innerHTML += "\n"
		}
	}

	var (
		attrsStr = rn.attrsToString()
		prefix   string
		suffix   string
	)

	if rn.nodeType == textNode {
		// Leave pre-formatted text nodes alone
		innerHTML = rn.text
	} else if rn.nodeType == fragmentNode {
		// No prefix/suffix for fragments
	} else if rn.nodeType == commentNode {
		prefix = "<!-- "
		suffix = " -->"
	} else if rn.selfClosing && innerHTML == "" {
		prefix = "<" + rn.tag + attrsStr
		suffix = "/>"
	} else {
		prefix = "<" + rn.tag + attrsStr + ">"
		suffix = "</" + rn.tag + ">"
	}

	// Adjust prefix and suffix, depending on what we need

	if !prettify {
		// we're not prettifying, so leave as-is
	} else if prefix == "" && suffix == "" {
		// Not wrapping, so leave as is
	} else if rn.preformatted || onlyTextChildren {
		// Put the entire element on one line
		prefix = rn.indent(level, prefix)
		innerHTML = strings.TrimSpace(innerHTML)
	} else {
		// Indent, with start, content, end on separate lines
		prefix = rn.indent(level, prefix) + "\n"
		suffix = rn.indent(level, suffix)
	}

	return prefix + innerHTML + suffix
}

func (rn *RawNode) attrsToString() string {
	items := strings.Builder{}
	for _, a := range rn.attrs {
		escaped := html.EscapeString(a.value)
		items.WriteString(" ")
		items.WriteString(a.name)
		items.WriteString("=\"")
		items.WriteString(escaped)
		items.WriteString("\"")
	}
	return items.String()
}

func (rn *RawNode) attr(name string) string {
	for _, a := range rn.attrs {
		if a.name == name {
			return a.value
		}
	}
	return ""
}
