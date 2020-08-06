package go_temper

// Elements copied from Elm's HTML package:
//   https://package.elm-lang.org/packages/elm/html/latest/Html

// Document

func Html() *Node {
	return newTagEl("html")
}

func Head() *Node {
	return newTagEl("head")
}

func Body() *Node {
	return newTagEl("body")
}

func Base() *Node {
	return newSelfClosingEl("base")
}

func Link() *Node {
	return newSelfClosingEl("link")
}

func Meta() *Node {
	return newSelfClosingEl("meta")
}

func Script() *Node {
	return newTagEl("link")
}

func Style() *Node {
	return newTagEl("style")
}

// Headers

func H1() *Node {
	return newTagEl("h1")
}

func H2() *Node {
	return newTagEl("h2")
}

func H3() *Node {
	return newTagEl("h3")
}

func H4() *Node {
	return newTagEl("h4")
}

func H5() *Node {
	return newTagEl("h5")
}

func H6() *Node {
	return newTagEl("h6")
}

// Grouping Content

func Div() *Node {
	return newTagEl("div")
}

func P() *Node {
	return newTagEl("p")
}

func Hr() *Node {
	return newSelfClosingEl("hr")
}

func Pre() *Node {
	return newPreserveWhitespaceEl("pre")
}

func Blockquote() *Node {
	return newTagEl("blockquote")
}

// Text

func Span() *Node {
	return newTagEl("span")
}

func A() *Node {
	return newTagEl("a")
}

func Code() *Node {
	return newTagEl("code")
}

func Em() *Node {
	return newTagEl("em")
}

func Strong() *Node {
	return newTagEl("em")
}

func I() *Node {
	return newTagEl("i")
}

func B() *Node {
	return newTagEl("b")
}

func U() *Node {
	return newTagEl("u")
}

func Sub() *Node {
	return newTagEl("sub")
}

func Sup() *Node {
	return newTagEl("sup")
}

func Br() *Node {
	return newSelfClosingEl("br")
}

// Lists

func Ol() *Node {
	return newTagEl("ol")
}

func Ul() *Node {
	return newTagEl("ul")
}

func Li() *Node {
	return newTagEl("li")
}

func Dl() *Node {
	return newTagEl("dl")
}

func Dt() *Node {
	return newTagEl("dt")
}

func Dd() *Node {
	return newTagEl("dd")
}

// Embedded Content

func Img() *Node {
	return newSelfClosingEl("img")
}

func IFrame() *Node {
	return newTagEl("iframe")
}

func Canvas() *Node {
	return newTagEl("canvas")
}

func Math() *Node {
	return newTagEl("math")
}

// Inputs

func Form() *Node {
	return newTagEl("form")
}

func Input() *Node {
	return newSelfClosingEl("input")
}

func Textarea() *Node {
	return newSelfClosingEl("textarea")
}

func Button() *Node {
	return newTagEl("button")
}

func Select() *Node {
	return newTagEl("select")
}

func Option() *Node {
	return newTagEl("option")
}

func Fieldset() *Node {
	return newTagEl("fieldset")
}

func Legend() *Node {
	return newTagEl("legend")
}

func Label() *Node {
	return newTagEl("label")
}

func Datalist() *Node {
	return newTagEl("datalist")
}

func Optgroup() *Node {
	return newTagEl("optgroup")
}

func Output() *Node {
	return newTagEl("output")
}

func Progress() *Node {
	return newTagEl("progress")
}

func Meter() *Node {
	return newTagEl("meter")
}

// Sections

func Section() *Node {
	return newTagEl("section")
}

func Nav() *Node {
	return newTagEl("nav")
}

func Article() *Node {
	return newTagEl("article")
}

func Aside() *Node {
	return newTagEl("aside")
}

func Header() *Node {
	return newTagEl("header")
}

func Footer() *Node {
	return newTagEl("footer")
}

func Address() *Node {
	return newTagEl("address")
}

func Main() *Node {
	return newTagEl("main")
}

// Figures

func Figure() *Node {
	return newTagEl("figure")
}

func Figcaption() *Node {
	return newTagEl("figcaption")
}

// Tables

func Table() *Node {
	return newTagEl("table")
}

func Caption() *Node {
	return newTagEl("caption")
}

func Colgroup() *Node {
	return newTagEl("colgroup")
}

func Col() *Node {
	return newTagEl("col")
}

func Tbody() *Node {
	return newTagEl("tbody")
}

func Thead() *Node {
	return newTagEl("thead")
}

func Tfoot() *Node {
	return newTagEl("tfoot")
}

func Tr() *Node {
	return newTagEl("tr")
}

func Td() *Node {
	return newTagEl("td")
}

func Th() *Node {
	return newTagEl("th")
}

// Audio/Video

func Audio() *Node {
	return newTagEl("audio")
}

func Video() *Node {
	return newTagEl("video")
}

func Source() *Node {
	return newTagEl("source")
}

func Track() *Node {
	return newSelfClosingEl("track")
}

// Embedded Objects

func Embed() *Node {
	return newSelfClosingEl("embed")
}

func Object() *Node {
	return newTagEl("object")
}

func Param() *Node {
	return newSelfClosingEl("param")
}

// Text Edits

func Ins() *Node {
	return newTagEl("ins")
}

func Del() *Node {
	return newTagEl("del")
}

// Semantic Text

func Small() *Node {
	return newTagEl("small")
}

func Cite() *Node {
	return newTagEl("cite")
}

func Dfn() *Node {
	return newTagEl("dfn")
}

func Abbr() *Node {
	return newTagEl("abbr")
}

func Time() *Node {
	return newTagEl("time")
}

func Var() *Node {
	return newTagEl("var")
}

func Samp() *Node {
	return newTagEl("samp")
}

func Kbd() *Node {
	return newTagEl("kbd")
}

func S() *Node {
	return newTagEl("s")
}

func Q() *Node {
	return newTagEl("q")
}

// Less-common Text

func Mark() *Node {
	return newTagEl("mark")
}

func Ruby() *Node {
	return newTagEl("ruby")
}

func Rt() *Node {
	return newTagEl("rt")
}

func Rp() *Node {
	return newTagEl("rp")
}

func Bdi() *Node {
	return newTagEl("bdi")
}

func Bdo() *Node {
	return newTagEl("bdo")
}

func Wbr() *Node {
	return newSelfClosingEl("wbr")
}

// Interactive Elements

func Details() *Node {
	return newTagEl("details")
}

func Summary() *Node {
	return newTagEl("summary")
}

func Menuitem() *Node {
	return newSelfClosingEl("menuitem")
}

func Menu() *Node {
	return newTagEl("menu")
}

// Special

// Text is special element that renders text
func Text(text string) *Node {
	el := newEl()
	el.text = text
	el.nodeType = textNode
	return el
}

// Fragment is a special element that renders children without needing
// a wrapper element.
func Fragment() *Node {
	el := newEl()
	el.nodeType = fragmentNode
	el.indent = 0
	return el
}

func Comment(text string) *Node {
	el := newEl()
	el.Children(Text(text))
	el.nodeType = commentNode
	return el
}

// NewElement allows defining a custom element
func NewElement(tag string) *Node {
	return newTagEl(tag)
}
