package hagl

import (
	"fmt"
	"html"
	"io"
	"regexp"
	"strings"
)

var multiWhitespaceRegexp = regexp.MustCompile("\n+")

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
	Text(text string) Node
	Textf(format string, a ...interface{}) Node
	Attr(name, value string) Node
	Class(cls ...string) Node
	Style(name, value string) Node
	HTML() string
	HTMLPretty() string
	Write(w io.Writer) (int, error)

	// GetNode returns the root node. This is for internal use only
	GetNode() *RawNode
}

type RawNode struct {
	// children contains the child elements for the element
	children []Node

	// attrs contains the attributes for the element
	attrs []attr

	// classes hold the class names assigned to the element, which will
	// be converted to attributes when rendered
	classes []string

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
	tab string
}

func (e *RawNode) ID(id string) Node {
	return e.Attr("id", id)
}

func (e *RawNode) GetNode() *RawNode {
	return e
}

func (e *RawNode) Children(child ...Node) Node {
	for _, c := range child {
		// Skip nil Nodes
		if c == nil {
			continue
		}

		// Preformatted nodes have preformatted children
		if e.preformatted {
			c.GetNode().preformatted = true
		}

		e.children = append(e.children, c)
	}
	return e
}

// Range is a convenience used to generate n children based on a factory function.
// the factory will be called n times and will skip any nil children
func (e *RawNode) Range(n int, child func(i int) Node) Node {
	return e.Children(FragmentRange(n, child))
}

// Text is a helper method to add a text node to children
func (e *RawNode) Text(text string) Node {
	return e.Children(Text(text))
}

// Textf is the same as Text, but accepts fmt args
func (e *RawNode) Textf(format string, a ...interface{}) Node {
	return e.Children(Text(fmt.Sprintf(format, a...)))
}

func (e *RawNode) Attr(name, value string) Node {
	for i, a := range e.attrs {
		if a.name == name {
			e.attrs[i] = attr{name: name, value: value}
			return e
		}
	}

	e.attrs = append(e.attrs, attr{name: name, value: value})

	return e
}

func (e *RawNode) Class(cls ...string) Node {
	existingClasses := strings.Fields(e.attr("class"))

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

	e.Attr("class", strings.Join(existingClasses, " "))
	return e
}

// Style is a utility method to append to the style attribute. If a style
// attribute already exists, the new style will be appended.
func (e *RawNode) Style(name, value string) Node {
	str := name + ":" + value
	for i, a := range e.attrs {
		if a.name == "style" {
			e.attrs[i].value += ";" + str
			return e
		}
	}

	// No style attr, so add one
	e.Attr("style", str)
	return e
}

func (e *RawNode) HTML() string {
	return e.html(-1, false)
}

func (e *RawNode) HTMLPretty() string {
	return strings.TrimSpace(e.html(0, true))
}

func (e *RawNode) Write(w io.Writer) (int, error) {
	return w.Write([]byte(e.HTML()))
}

func (e *RawNode) indent(level int, text string) string {
	prefix := strings.Builder{}
	for i := 0; i < level; i++ {
		prefix.WriteString(e.tab)
	}

	prefix.WriteString(text)
	return prefix.String()
}

func (e *RawNode) html(level int, prettify bool) string {
	innerHTML := ""
	onlyTextChildren := true

	// Render children if the element has them
	for _, c := range e.children {
		if c.GetNode().nodeType != textNode {
			onlyTextChildren = false
		}

		if e.preformatted {
			// Indent open tag but nothing else
			// TODO: Figure out what to do with tags inside <pre>
			innerHTML += e.indent(0, c.GetNode().html(0, false))
		} else {
			innerHTML += c.GetNode().html(level+e.indentIncrement, prettify)
		}

		// Add newline after each child if we're prettifying. Note, we don't
		// add one to fragment children because they don't take up space
		if prettify && !e.preformatted && c.GetNode().nodeType != fragmentNode {
			innerHTML += "\n"
		}
	}

	var (
		attrsStr = e.attrsToString()
		prefix   = ""
		suffix   = ""
	)

	if e.nodeType == textNode && e.preformatted {
		// Leave pre-formatted text nodes alone
		innerHTML = e.text
	} else if e.nodeType == textNode {
		// Replace multiple whitespace with single space for text nodes
		innerHTML = multiWhitespaceRegexp.ReplaceAllString(e.text, " ")
	} else if e.nodeType == fragmentNode {
		// No prefix/suffix for fragments
	} else if e.nodeType == commentNode {
		prefix = "<!-- "
		suffix = " -->"
	} else if e.selfClosing && innerHTML == "" {
		prefix = "<" + e.tag + attrsStr
		suffix = "/>"
	} else {
		prefix = "<" + e.tag + attrsStr + ">"
		suffix = "</" + e.tag + ">"
	}

	// Adjust prefix and suffix, depending on what we need

	if !prettify {
		// we're not prettifying, so leave as-is
	} else if prefix == "" && suffix == "" {
		// Not wrapping, so leave as is
	} else if e.preformatted || onlyTextChildren {
		// Put the entire element on one line
		prefix = e.indent(level, prefix)
		innerHTML = strings.TrimSpace(innerHTML)
	} else {
		// Indent, with start, content, end on separate lines
		prefix = e.indent(level, prefix) + "\n"
		suffix = e.indent(level, suffix)
	}

	return prefix + innerHTML + suffix
}

func (e *RawNode) attrsToString() string {
	items := strings.Builder{}
	for _, a := range e.attrs {
		escaped := html.EscapeString(a.value)
		items.WriteString(" ")
		items.WriteString(a.name)
		items.WriteString("=\"")
		items.WriteString(escaped)
		items.WriteString("\"")
	}
	return items.String()
}

func (e *RawNode) attr(name string) string {
	for _, a := range e.attrs {
		if a.name == name {
			return a.value
		}
	}
	return ""
}
