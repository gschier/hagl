package go_temper

// Elements copied from Elm's HTML package:
//   https://package.elm-lang.org/packages/elm/html/latest/Html

// Headers

func H1() *Element {
	return newTagEl("h1")
}

func H2() *Element {
	return newTagEl("h2")
}

func H3() *Element {
	return newTagEl("h3")
}

func H4() *Element {
	return newTagEl("h4")
}

func H5() *Element {
	return newTagEl("h5")
}

func H6() *Element {
	return newTagEl("h6")
}

// Grouping Content

func Div() *Element {
	return newTagEl("div")
}

func P() *Element {
	return newTagEl("p")
}

func Hr() *Element {
	return newSelfClosingEl("hr")
}

func Pre() *Element {
	return newPreserveWhitespaceEl("pre")
}

func Blockquote() *Element {
	return newTagEl("blockquote")
}

// Text

func Span() *Element {
	return newTagEl("span")
}

func A() *Element {
	return newTagEl("a")
}

func Code() *Element {
	return newTagEl("code")
}

func Em() *Element {
	return newTagEl("em")
}

func Strong() *Element {
	return newTagEl("em")
}

func I() *Element {
	return newTagEl("i")
}

func B() *Element {
	return newTagEl("b")
}

func U() *Element {
	return newTagEl("u")
}

func Sub() *Element {
	return newTagEl("sub")
}

func Sup() *Element {
	return newTagEl("sup")
}

func Br() *Element {
	return newSelfClosingEl("br")
}

// Lists

func Ol() *Element {
	return newTagEl("ol")
}

func Ul() *Element {
	return newTagEl("ul")
}

func Li() *Element {
	return newTagEl("li")
}

func Dl() *Element {
	return newTagEl("dl")
}

func Dt() *Element {
	return newTagEl("dt")
}

func Dd() *Element {
	return newTagEl("dd")
}

// Embedded Content

func Img() *Element {
	return newSelfClosingEl("img")
}

func IFrame() *Element {
	return newTagEl("iframe")
}

func Canvas() *Element {
	return newTagEl("canvas")
}

func Math() *Element {
	return newTagEl("math")
}

// Inputs

func Form() *Element {
	return newTagEl("form")
}

func Input() *Element {
	return newSelfClosingEl("input")
}

func Textarea() *Element {
	return newTagEl("textarea")
}

func Button() *Element {
	return newTagEl("button")
}

func Select() *Element {
	return newTagEl("select")
}

func Option() *Element {
	return newTagEl("option")
}

func Fieldset() *Element {
	return newTagEl("fieldset")
}

func Legend() *Element {
	return newTagEl("legend")
}

func Label() *Element {
	return newTagEl("label")
}

func Datalist() *Element {
	return newTagEl("datalist")
}

func Optgroup() *Element {
	return newTagEl("optgroup")
}

func Output() *Element {
	return newTagEl("output")
}

func Progress() *Element {
	return newTagEl("progress")
}

func Meter() *Element {
	return newTagEl("meter")
}

// Sections

func Section() *Element {
	return newTagEl("section")
}

func Nav() *Element {
	return newTagEl("nav")
}

func Article() *Element {
	return newTagEl("article")
}

func Aside() *Element {
	return newTagEl("aside")
}

func Header() *Element {
	return newTagEl("header")
}

func Footer() *Element {
	return newTagEl("footer")
}

func Address() *Element {
	return newTagEl("address")
}

func Main() *Element {
	return newTagEl("main")
}

// Figures

func Figure() *Element {
	return newTagEl("figure")
}

func Figcaption() *Element {
	return newTagEl("figcaption")
}

// Tables

func Table() *Element {
	return newTagEl("table")
}

func Caption() *Element {
	return newTagEl("caption")
}

func Colgroup() *Element {
	return newTagEl("colgroup")
}

func Col() *Element {
	return newTagEl("col")
}

func Tbody() *Element {
	return newTagEl("tbody")
}

func Thead() *Element {
	return newTagEl("thead")
}

func Tfoot() *Element {
	return newTagEl("tfoot")
}

func Tr() *Element {
	return newTagEl("tr")
}

func Td() *Element {
	return newTagEl("td")
}

func Th() *Element {
	return newTagEl("th")
}

// Audio/Video

func Audio() *Element {
	return newSelfClosingEl("audio")
}

func Video() *Element {
	return newSelfClosingEl("video")
}

func Source() *Element {
	return newSelfClosingEl("source")
}

func Track() *Element {
	return newSelfClosingEl("track")
}

// Embedded Objects

func Embed() *Element {
	return newSelfClosingEl("embed")
}

func Object() *Element {
	return newSelfClosingEl("object")
}

func Param() *Element {
	return newSelfClosingEl("param")
}

// Text Edits

func Ins() *Element {
	return newTagEl("ins")
}

func Del() *Element {
	return newTagEl("del")
}

// Semantic Text

func Small() *Element {
	return newTagEl("small")
}

func Cite() *Element {
	return newTagEl("cite")
}

func Dfn() *Element {
	return newTagEl("dfn")
}

func Abbr() *Element {
	return newTagEl("abbr")
}

func Time() *Element {
	return newTagEl("time")
}

func Var() *Element {
	return newTagEl("var")
}

func Samp() *Element {
	return newTagEl("samp")
}

func Kbd() *Element {
	return newTagEl("kbd")
}

func S() *Element {
	return newTagEl("s")
}

func Q() *Element {
	return newTagEl("q")
}

// Less-common Text

func Mark() *Element {
	return newTagEl("mark")
}

func Ruby() *Element {
	return newTagEl("ruby")
}

func Rt() *Element {
	return newTagEl("rt")
}

func Rp() *Element {
	return newTagEl("rp")
}

func Bdi() *Element {
	return newTagEl("bdi")
}

func Bdo() *Element {
	return newTagEl("bdo")
}

func Wbr() *Element {
	return newTagEl("wbr")
}

// Interactive Elements

func Details() *Element {
	return newTagEl("details")
}

func Summary() *Element {
	return newTagEl("summary")
}

func Menuitem() *Element {
	return newTagEl("menuitem")
}

func Menu() *Element {
	return newTagEl("menu")
}

// Special

// Text is special element that renders text
func Text(text string) *Element {
	el := newEl()
	el.text = text
	return el
}

// Node allows defining a custom element
func Node(tag string) *Element {
	return newTagEl(tag)
}

// Fragment is a special element that renders children without needing
// a wrapper element.
func Fragment() *Element {
	el := newEl()
	return el
}
