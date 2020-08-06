package go_temper

import (
	"html"
	"regexp"
	"strings"
)

var multiWhitespaceRegexp = regexp.MustCompile("\n+")

func newEl() *Node {
	return &Node{
		nodeType:        elementNode,
		tab:             "  ",
		indentIncrement: 1,
	}
}

func newTagEl(tag string) *Node {
	el := newEl()
	el.tag = tag
	return el
}

func newSelfClosingEl(tag string) *Node {
	el := newTagEl(tag)
	el.selfClosing = true
	return el
}

func newPreserveWhitespaceEl(tag string) *Node {
	el := newTagEl(tag)
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
	elementNode
	commentNode
	fragmentNode
)

type Node struct {
	// children contains the child elements for the element
	children []Node

	// attrs contains the attributes for the element
	attrs []attr

	// classes hold the class names assigned to the element, which will
	// be converted to attributes when rendered
	classes []string

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

func (e *Node) ID(id string) *Node {
	return e.Attr("id", id)
}

func (e *Node) Children(child ...*Node) *Node {
	for _, c := range child {
		if e.preformatted {
			c.preformatted = true
		}
		e.children = append(e.children, *c)
	}
	return e
}

// Text is a helper method to add a text node to children
func (e *Node) Text(text string) *Node {
	return e.Children(Text(text))
}

func (e *Node) Attr(name, value string) *Node {
	for i, a := range e.attrs {
		if a.name == name {
			e.attrs[i] = attr{name: name, value: value}
			return e
		}
	}

	e.attrs = append(e.attrs, attr{name: name, value: value})

	return e
}

func (e *Node) Class(cls ...string) *Node {
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

func (e *Node) HTML() string {
	return e.html(-1, false)
}

func (e *Node) HTMLPretty() string {
	return strings.TrimSpace(e.html(0, true))
}

func (e *Node) indent(level int, text string) string {
	prefix := strings.Builder{}
	for i := 0; i < level; i++ {
		prefix.WriteString(e.tab)
	}

	prefix.WriteString(text)
	return prefix.String()
}

func (e *Node) html(level int, prettify bool) string {
	innerHTML := ""
	onlyTextChildren := true

	for i, c := range e.children {
		if c.nodeType != textNode {
			onlyTextChildren = false
		}

		if e.preformatted {
			// Indent open tag but nothing else
			// TODO: Figure out what to do with tags inside <pre>
			innerHTML += e.indent(0, c.html(0, false))
		} else {
			innerHTML += c.html(level+e.indentIncrement, prettify)
		}

		// Remove extra whitespace from last child so we don't get a
		// blank line
		if i == len(e.children)-1 {
			innerHTML = strings.TrimSuffix(innerHTML, "\n")
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
		suffix = suffix + "\n"
		innerHTML = strings.TrimSpace(innerHTML)
	} else {
		// Indent, with start, content, end on separate lines
		prefix = e.indent(level, prefix) + "\n"
		suffix = e.indent(level, suffix) + "\n"
		innerHTML += "\n"
	}

	return prefix + innerHTML + suffix
}

func (e *Node) attrsToString() string {
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

func (e *Node) attr(name string) string {
	for _, a := range e.attrs {
		if a.name == name {
			return a.value
		}
	}

	return ""
}
