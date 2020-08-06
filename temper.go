package go_temper

import (
	"html"
	"regexp"
	"strings"
)

func newEl() *Element {
	return &Element{
		children: make([]Element, 0),
		attrs:    make([]attr, 0),
	}
}

func newTagEl(tag string) *Element {
	el := newEl()
	el.tag = tag
	return el
}

func newSelfClosingEl(tag string) *Element {
	el := newTagEl(tag)
	el.selfClosing = true
	return el
}

func newPreserveWhitespaceEl(tag string) *Element {
	el := newTagEl(tag)
	el.preserveWhitespace = true
	return el
}

type attr struct {
	name  string
	value string
}

type Element struct {
	// children contains the child elements for the element
	children []Element

	// attrs contains the attributes for the element
	attrs []attr

	// classes hold the class names assigned to the element, which will
	// be converted to attributes when rendered
	classes []string

	// tag defines the name of the HTML element
	tag string

	// text contains the text contents of the element
	text string

	// preserveWhitespace is used for tags like <pre> that shouldn't
	// have the whitespace of their children changed in any way
	preserveWhitespace bool

	// selfClosing defines whether or not the element can close itself.
	//
	// For example, a horizontal rule or input can <hr/> <input/>
	selfClosing bool
}

func (e *Element) ID(id string) *Element {
	return e.Attr("id", id)
}

func (e *Element) Children(child ...*Element) *Element {
	for _, c := range child {
		if e.preserveWhitespace {
			c.preserveWhitespace = true
		}
		e.children = append(e.children, *c)
	}
	return e
}

func (e *Element) Text(text string) *Element {
	e.children = make([]Element, 0) // Clear existing
	return e.Children(Text(text))
}

func (e *Element) Attr(name, value string) *Element {
	for i, a := range e.attrs {
		if a.name == name {
			e.attrs[i] = attr{name: name, value: value}
			return e
		}
	}

	e.attrs = append(e.attrs, attr{name: name, value: value})

	return e
}

func (e *Element) Class(cls ...string) *Element {
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

func (e *Element) HTML() string {
	return e.html(-1, "")
}

func (e *Element) HTMLPretty() string {
	return strings.TrimSpace(e.html(0, "  "))
}

var multiWhitespaceRegexp = regexp.MustCompile("\n+")

func (e *Element) html(level int, tab string) string {
	// It's a text tag, so format it and return
	if e.isTextNode() {
		if e.preserveWhitespace {
			return e.text
		} else {
			normalizedWhitespace := multiWhitespaceRegexp.ReplaceAllString(e.text, " ")
			return indent(level, normalizedWhitespace, tab)
		}
	}

	innerHTML := ""

	// It's a fragment, so just render out the children, indented
	// at a lesser level
	if e.isFragment() {
		for _, c := range e.children {
			innerHTML += indent(level-1, c.html(level, tab), tab)
		}
		return innerHTML
	}

	for _, c := range e.children {
		if c.preserveWhitespace {
			// Indent open tag but nothing else
			innerHTML += indentN(level+1, c.HTML(), tab, 1)
		} else {
			innerHTML += c.html(level+1, tab)
		}
	}

	// If it ends in a newline, remove it
	if strings.HasSuffix(innerHTML, "\n") {
		innerHTML = strings.TrimSuffix(innerHTML, "\n")
	}

	attrsStr := e.attrsToString()
	if e.selfClosing && innerHTML == "" {
		return "<" + e.tag + attrsStr + "/>"
	}

	tagStart := "<" + e.tag + attrsStr + ">"
	tagEnd := "</" + e.tag + ">"

	// No indent so no pretty
	if tab == "" {
		return tagStart + innerHTML + tagEnd
	}

	// Short text contents, so put on one line
	if !e.preserveWhitespace && e.onlyTextChildren() && len(tagStart)+len(innerHTML) < 60 {
		return indent(level, tagStart+strings.TrimSpace(innerHTML)+tagEnd, tab) + "\n"
	}

	return indent(level, tagStart, tab) + "\n" +
		innerHTML + "\n" +
		indent(level, tagEnd, tab) + "\n"
}

func (e *Element) isTextNode() bool {
	return e.tag == "" && len(e.children) == 0
}

func (e *Element) isFragment() bool {
	return e.tag == "" && len(e.children) != 0
}

func (e *Element) onlyTextChildren() bool {
	for _, c := range e.children {
		if !c.isTextNode() {
			return false
		}
	}
	return true
}

func (e *Element) attrsToString() string {
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

func (e *Element) attr(name string) string {
	for _, a := range e.attrs {
		if a.name == name {
			return a.value
		}
	}

	return ""
}
