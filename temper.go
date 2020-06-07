package go_temper

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

var multiWhitespaceRegexp = regexp.MustCompile("\n+")

type Element interface {
	HTML() string
	HTMLPretty() string
	Text(text string) Element
	Children(...Element) Element
	Class(...string) Element
	Attr(name, value string) Element
}

func Div() Element {
	return El("div")
}

func Ul() Element {
	return El("ul")
}

func Li() Element {
	return El("li")
}

func El(tag string) Element {
	return &element{
		tag:      tag,
		text:     "",
		children: make([]Element, 0),
		attrs:    make([]*attr, 0),
	}
}

func Text(text string) Element {
	return &element{text: text}
}

type attr struct {
	n string
	v string
}

type element struct {
	children []Element
	attrs    []*attr
	tag      string
	text     string
	classes  []string
}

func (e *element) Children(child ...Element) Element {
	e.children = append(e.children, child...)
	return e
}

func (e *element) Text(text string) Element {
	e.children = []Element{Text(text)}
	return e
}

func (e *element) Attr(name, value string) Element {
	for i, a := range e.attrs {
		if a.n == name {
			e.attrs[i] = &attr{n: name, v: value}
			return e
		}
	}

	e.attrs = append(e.attrs, &attr{n: name, v: value})

	return e
}

func (e *element) Class(cls ...string) Element {
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

func (e *element) HTML() string {
	return e.html(-1, "")
}

func (e *element) HTMLPretty() string {
	return strings.TrimSpace(e.html(0, "  "))
}

func (e *element) html(level int, tab string) string {
	// It's a text tag, so format it and return
	if e.isTextNode() {
		return indent(level, multiWhitespaceRegexp.ReplaceAllString(e.text, " "), tab)
	}

	innerHTML := ""
	for _, c := range e.children {
		innerHTML += c.(*element).html(level+1, tab)
	}
	if strings.HasSuffix(innerHTML, "\n") {
		innerHTML = innerHTML[0 : len(innerHTML)-1]
	}

	attrsStr := e.attrsToString()

	var tagStart string
	var tagEnd = fmt.Sprintf("</%s>", e.tag)
	if attrsStr != "" {
		tagStart = fmt.Sprintf("<%s %s>", e.tag, attrsStr)
	} else {
		tagStart = fmt.Sprintf("<%s>", e.tag)
	}

	// No indent so no pretty
	if tab == "" {
		return tagStart + innerHTML + tagEnd
	}

	// Short text contents, so put on one line
	if e.onlyTextChildren() && len(tagStart)+len(innerHTML) < 60 {
		return indent(level, tagStart+strings.TrimSpace(innerHTML)+tagEnd, tab) + "\n"
	}

	h := indent(level, tagStart, tab) + "\n"
	h += innerHTML + "\n"
	h += indent(level, tagEnd, tab) + "\n"

	return h
}

func (e *element) isTextNode() bool {
	return e.tag == ""
}

func (e *element) onlyTextChildren() bool {
	for _, c := range e.children {
		if !c.(*element).isTextNode() {
			return false
		}
	}
	return true
}

func (e *element) attrsToString() string {
	items := make([]string, 0)
	for _, a := range e.attrs {
		escaped := html.EscapeString(a.v)
		items = append(items, fmt.Sprintf(`%s="%s"`, a.n, escaped))
	}

	return strings.Join(items, " ")
}

func (e *element) attr(name string) string {
	for _, a := range e.attrs {
		if a.n == name {
			return a.v
		}
	}

	return ""
}
