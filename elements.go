package hagl

import (
	"fmt"
	"html"
	"strings"
)

// Elements copied from Elm's HTML package:
//   https://package.elm-lang.org/packages/elm/html/latest/Html

// Document

func Html() Node {
	return newTagNode("html")
}

func Head() Node {
	return newTagNode("head")
}

func Title() Node {
	return newTagNode("title")
}

func Body() Node {
	return newTagNode("body")
}

func Base() Node {
	return newSelfClosingTagNode("base")
}

func Link() Node {
	return newSelfClosingTagNode("link")
}

func Meta() Node {
	return newSelfClosingTagNode("meta")
}

func Script() Node {
	return newPreserveWhitespaceTagNode("script")
}

func Style() Node {
	return newTagNode("style")
}

// Headers

func H1() Node {
	return newTagNode("h1")
}

func H2() Node {
	return newTagNode("h2")
}

func H3() Node {
	return newTagNode("h3")
}

func H4() Node {
	return newTagNode("h4")
}

func H5() Node {
	return newTagNode("h5")
}

func H6() Node {
	return newTagNode("h6")
}

// Grouping HTML

func Div() Node {
	return newTagNode("div")
}

func P() Node {
	return newTagNode("p")
}

func Hr() Node {
	return newSelfClosingTagNode("hr")
}

func Pre() Node {
	return newPreserveWhitespaceTagNode("pre")
}

func Blockquote() Node {
	return newTagNode("blockquote")
}

// Text

func Span() Node {
	return newTagNode("span")
}

func A() Node {
	return newTagNode("a")
}

func Code() Node {
	return newTagNode("code")
}

func Em() Node {
	return newTagNode("em")
}

func Strong() Node {
	return newTagNode("strong")
}

func I() Node {
	return newTagNode("i")
}

func B() Node {
	return newTagNode("b")
}

func U() Node {
	return newTagNode("u")
}

func Sub() Node {
	return newTagNode("sub")
}

func Sup() Node {
	return newTagNode("sup")
}

func Br() Node {
	return newSelfClosingTagNode("br")
}

// Lists

func Ol() Node {
	return newTagNode("ol")
}

func Ul() Node {
	return newTagNode("ul")
}

func Li() Node {
	return newTagNode("li")
}

func Dl() Node {
	return newTagNode("dl")
}

func Dt() Node {
	return newTagNode("dt")
}

func Dd() Node {
	return newTagNode("dd")
}

// Embedded HTML

func Img() Node {
	return newSelfClosingTagNode("img")
}

func Svg() Node {
	return newTagNode("svg")
}

func Path() Node {
	return newTagNode("path")
}

func IFrame() Node {
	return newTagNode("iframe")
}

func Canvas() Node {
	return newTagNode("canvas")
}

func Math() Node {
	return newTagNode("math")
}

// Inputs

func Form() Node {
	return newTagNode("form")
}

func Input() Node {
	return newSelfClosingTagNode("input")
}

func Textarea() Node {
	return newTagNode("textarea")
}

func Button() Node {
	return newTagNode("button")
}

func Select() Node {
	return newTagNode("select")
}

func Option() Node {
	return newTagNode("option")
}

func Fieldset() Node {
	return newTagNode("fieldset")
}

func Legend() Node {
	return newTagNode("legend")
}

func Label() Node {
	return newTagNode("label")
}

func Datalist() Node {
	return newTagNode("datalist")
}

func Optgroup() Node {
	return newTagNode("optgroup")
}

func Output() Node {
	return newTagNode("output")
}

func Progress() Node {
	return newTagNode("progress")
}

func Meter() Node {
	return newTagNode("meter")
}

// Sections

func Section() Node {
	return newTagNode("section")
}

func Nav() Node {
	return newTagNode("nav")
}

func Article() Node {
	return newTagNode("article")
}

func Aside() Node {
	return newTagNode("aside")
}

func Header() Node {
	return newTagNode("header")
}

func Footer() Node {
	return newTagNode("footer")
}

func Address() Node {
	return newTagNode("address")
}

func Main() Node {
	return newTagNode("main")
}

// Figures

func Figure() Node {
	return newTagNode("figure")
}

func Figcaption() Node {
	return newTagNode("figcaption")
}

// Tables

func Table() Node {
	return newTagNode("table")
}

func Caption() Node {
	return newTagNode("caption")
}

func Colgroup() Node {
	return newTagNode("colgroup")
}

func Col() Node {
	return newTagNode("col")
}

func Tbody() Node {
	return newTagNode("tbody")
}

func Thead() Node {
	return newTagNode("thead")
}

func Tfoot() Node {
	return newTagNode("tfoot")
}

func Tr() Node {
	return newTagNode("tr")
}

func Td() Node {
	return newTagNode("td")
}

func Th() Node {
	return newTagNode("th")
}

// Audio/Video

func Audio() Node {
	return newTagNode("audio")
}

func Video() Node {
	return newTagNode("video")
}

func Source() Node {
	return newTagNode("source")
}

func Track() Node {
	return newSelfClosingTagNode("track")
}

// Embedded Objects

func Embed() Node {
	return newSelfClosingTagNode("embed")
}

func Object() Node {
	return newTagNode("object")
}

func Param() Node {
	return newSelfClosingTagNode("param")
}

// Text Edits

func Ins() Node {
	return newTagNode("ins")
}

func Del() Node {
	return newTagNode("del")
}

// Semantic Text

func Small() Node {
	return newTagNode("small")
}

func Cite() Node {
	return newTagNode("cite")
}

func Dfn() Node {
	return newTagNode("dfn")
}

func Abbr() Node {
	return newTagNode("abbr")
}

func Time() Node {
	return newTagNode("time")
}

func Var() Node {
	return newTagNode("var")
}

func Samp() Node {
	return newTagNode("samp")
}

func Kbd() Node {
	return newTagNode("kbd")
}

func S() Node {
	return newTagNode("s")
}

func Q() Node {
	return newTagNode("q")
}

// Less-common Text

func Mark() Node {
	return newTagNode("mark")
}

func Ruby() Node {
	return newTagNode("ruby")
}

func Rt() Node {
	return newTagNode("rt")
}

func Rp() Node {
	return newTagNode("rp")
}

func Bdi() Node {
	return newTagNode("bdi")
}

func Bdo() Node {
	return newTagNode("bdo")
}

func Wbr() Node {
	return newSelfClosingTagNode("wbr")
}

// Interactive Elements

func Details() Node {
	return newTagNode("details")
}

func Summary() Node {
	return newTagNode("summary")
}

func Menuitem() Node {
	return newSelfClosingTagNode("menuitem")
}

func Menu() Node {
	return newTagNode("menu")
}

func Custom(tag string) Node {
	return newTagNode(tag)
}

// Special

// Text is special element that renders text
func Text(text ...string) Node {
	el := newEl()
	escaped := make([]string, len(text))
	for i, t := range text {
		escaped[i] = html.EscapeString(t)
	}
	el.text = strings.Join(escaped, " ")
	el.nodeType = textNode
	return el
}

// Textf is special element that renders formatted text
func Textf(text string, a ...interface{}) Node {
	return Text(fmt.Sprintf(text, a...))
}

// UnsafeText is special element that renders raw text or HTML (unescaped)
func UnsafeText(text ...string) Node {
	el := newEl()
	el.text = strings.Join(text, " ")
	el.nodeType = textNode
	return el
}

// Fragment is a special element that renders children without needing
// a wrapper element.
func Fragment() Node {
	el := newEl()
	el.nodeType = fragmentNode
	el.indentIncrement = 0
	return el
}

func Comment(text string) Node {
	el := newEl()
	el.Children(Text(text))
	el.nodeType = commentNode
	return el
}

// NewElement allows defining a custom element
func NewElement(tag string) Node {
	return newTagNode(tag)
}
