package hagl_test

import (
	"fmt"
	assert "github.com/stretchr/testify/require"
	"strings"
	"testing"

	. "github.com/gschier/hagl"
)

func TestElement_ToHTML(t *testing.T) {
	t.Run("generates simple element", func(t *testing.T) {
		root := Div().Text("Hello World!")
		assert.Equal(t, "<div>Hello World!</div>", root.ToHTML())
	})

	t.Run("escapes text content", func(t *testing.T) {
		root := Div().Text(`Hello "<World>"`)
		assert.Equal(t, `<div>Hello &#34;&lt;World&gt;&#34;</div>`, root.ToHTML())
	})

	t.Run("leaves raw text content", func(t *testing.T) {
		root := Div().HTMLUnsafe(`Hello <strong>World</strong>`)
		assert.Equal(t, `<div>Hello <strong>World</strong></div>`, root.ToHTML())
	})

	t.Run("generates nested elements", func(t *testing.T) {
		root := Div().Children(
			H1().Text("Hello World!"),
			Ul().Children(
				Li().Text("Item 1"),
				Li().Text("Item 2"),
			),
			Button().Class("btn").Text("Click Me!"),
		)
		assert.Equal(t, strings.Join([]string{
			"<div>",
			"  <h1>Hello World!</h1>",
			"  <ul>",
			"    <li>Item 1</li>",
			"    <li>Item 2</li>",
			"  </ul>",
			"  <button class=\"btn\">Click Me!</button>",
			"</div>",
		}, "\n"), root.ToHTMLPretty())
	})
}

func TestElement_ToText(t *testing.T) {
	t.Run("generates simple element", func(t *testing.T) {
		root := Div().Text("Hello World!")
		assert.Equal(t, "Hello World!", root.ToText())
	})

	t.Run("generates nested elements", func(t *testing.T) {
		root := Div().Children(
			H1().Text("Hello World!"),
			Ul().Children(
				Li().Text("Item 1"),
				Li().Text("Item 2"),
			),
			Ol().Children(
				Li().Text("Item 1"),
				Li().Text("Item 2"),
			),
			P().Text(
				"This is a paragraph.",
				"It is very long so that the text will be wrapped onto a newline.",
				"Did it work?",
			),
			P().Children(
				A().Href("https://yaak.app").Class("btn").Text("Click Me!"),
			),
		)
		assert.Equal(t, strings.Join([]string{
			"Hello World!",
			"",
			" - Item 1",
			" - Item 2",
			"",
			" 1) Item 1",
			" 2) Item 2",
			"",
			"This is a paragraph. It is very long so that the text will be wrapped onto a ",
			"newline. Did it work?",
			"",
			"Click Me! (https://yaak.app)",
		}, "\n"), root.ToText())
	})

	t.Run("doesn't add too much whitespace", func(t *testing.T) {
		root := Div().Children(
			Div().Children(
				Div().Children(
					P().Text("P 1"),
					P().Text("P 2"),
				),
			),
		)
		assert.Equal(t, strings.Join([]string{
			"P 1",
			"",
			"P 2",
		}, "\n"), root.ToText())
	})
}

func TestElement_Attr(t *testing.T) {
	t.Run("adds attr", func(t *testing.T) {
		root := Div().Attr("style", "display: block")
		assert.Equal(t, `<div style="display: block"></div>`, root.ToHTML())
	})

	t.Run("overwrites attr", func(t *testing.T) {
		root := Div().Attr("id", "1").Attr("id", "2")
		assert.Equal(t, `<div id="2"></div>`, root.ToHTML())
	})

	t.Run("escapes attr", func(t *testing.T) {
		root := Div().Attr(`"style <Hello-WORLD>"`, `Hello <string>"World"</strong>"`)
		assert.Equal(t, `<div styleHello-WORLD="Hello &lt;string&gt;&#34;World&#34;&lt;/strong&gt;&#34;"></div>`, root.ToHTML())
	})
}

func TestElement_Class(t *testing.T) {
	t.Run("adds classes", func(t *testing.T) {
		root := Button().Class("btn", "btn--primary")
		assert.Equal(t, `<button class="btn btn--primary"></button>`, root.ToHTML())
	})

	t.Run("adds duplicate classes", func(t *testing.T) {
		root := Div().Class("btn", "btn--primary", "btn")
		assert.Equal(t, `<div class="btn btn--primary"></div>`, root.ToHTML())
	})

	t.Run("appends to manually set attr", func(t *testing.T) {
		root := Button().Attr("class", "btn").Class("btn--primary")
		assert.Equal(t, `<button class="btn btn--primary"></button>`, root.ToHTML())
	})
}

func TestElement_Style(t *testing.T) {
	t.Run("adds style", func(t *testing.T) {
		root := Button().StyleProperty("background", "red")
		assert.Equal(t, `<button style="background:red"></button>`, root.ToHTML())
	})

	t.Run("adds multiple styles", func(t *testing.T) {
		root := Button().StyleProperty("background", "red").StyleProperty("color", "white")
		assert.Equal(t, `<button style="background:red;color:white"></button>`, root.ToHTML())
	})

	t.Run("doesn't overwrite style", func(t *testing.T) {
		root := Button().StyleProperty("background", "red").StyleProperty("background", "blue")
		assert.Equal(t, `<button style="background:red;background:blue"></button>`, root.ToHTML())
	})
}

func TestElement_HTMLPretty(t *testing.T) {
	t.Run("pre", func(t *testing.T) {
		root := Pre().HTMLUnsafe("function foo() {\n  return 'Hello World!';\n}")
		assert.Equal(t, strings.Join([]string{
			"<pre>function foo() {",
			"  return 'Hello World!';",
			"}</pre>",
		}, "\n"), root.ToHTMLPretty())
	})

	t.Run("pretty HTML", func(t *testing.T) {
		root := Div().Children(
			Ul().Children(
				Li().Text("1"),
				Li().Text("Hello\n\n\n\nWorld!"),
				Li().Text("This is a really long string that will get wrapped because it's too long."),
			),
			Pre().Text("function foo() {\n  return 'Hello World!';\n}"),
			Pre().Children(
				Div().Text("foo"),
				Text("Bar"),
				H2().Text("woo!"),
				Comment("That was cool"),
			),
		)
		assert.Equal(t, strings.Join([]string{
			"<div>",
			"  <ul>",
			"    <li>1</li>",
			"    <li>Hello\n\n\n\nWorld!</li>",
			"    <li>This is a really long string that will get wrapped because it&#39;s too long.</li>",
			"  </ul>",
			"  <pre>function foo() {",
			"  return &#39;Hello World!&#39;;",
			"}</pre>",
			"  <pre><div>foo</div>Bar<h2>woo!</h2><!-- That was cool --></pre>",
			"</div>",
		}, "\n"), root.ToHTMLPretty())
	})
}

func TestEl(t *testing.T) {
	t.Run("test component", func(t *testing.T) {
		Btn := func() Node {
			return Button().Class("btn btn--primary").Attr("type", "button")
		}

		root := Div().Children(Btn().Text("Click Me!"))
		assert.Equal(t, `<div><button class="btn btn--primary" type="button">Click Me!</button></div>`, root.ToHTML())
	})
}

func TestFragmentEl(t *testing.T) {
	t.Run("works with no children", func(t *testing.T) {
		root := Hr().Class("red").Attr("type", "foo")
		assert.Equal(t, `<hr class="red" type="foo"/>`, root.ToHTML())
	})

	t.Run("works with children", func(t *testing.T) {
		root := Hr().Class("red").Attr("type", "foo").Text("foo")
		assert.Equal(t, `<hr class="red" type="foo">foo</hr>`, root.ToHTML())
	})
}

func TestSelfClosingEl(t *testing.T) {
	t.Run("works with no children", func(t *testing.T) {
		root := Hr().Class("red").Attr("type", "foo")
		assert.Equal(t, `<hr class="red" type="foo"/>`, root.ToHTML())
	})

	t.Run("works with children", func(t *testing.T) {
		root := Hr().Class("red").Attr("type", "foo").Text("foo")
		assert.Equal(t, `<hr class="red" type="foo">foo</hr>`, root.ToHTML())
	})
}

func TestComment(t *testing.T) {
	t.Run("basic example", func(t *testing.T) {
		root := Comment("This is a comment")
		assert.Equal(t, "<!-- This is a comment -->", root.ToHTML())
	})

	t.Run("complex example", func(t *testing.T) {
		root := Div().Children(
			Comment("This is in a div"),
			Div().Children(
				Comment("This is an awesome comment"),
			),
		)
		assert.Equal(t, strings.Join([]string{
			"<div>",
			"  <!-- This is in a div -->",
			"  <div>",
			"    <!-- This is an awesome comment -->",
			"  </div>",
			"</div>",
		}, "\n"), root.ToHTMLPretty())
	})
}

func TestFragment(t *testing.T) {
	t.Run("basic example", func(t *testing.T) {
		root := Fragment().Children(
			Div().Children(Span().Text("foo")),
			Div(),
			Div(),
		)
		assert.Equal(t, "<div><span>foo</span></div><div></div><div></div>", root.ToHTML())
	})

	t.Run("basic example wrapped", func(t *testing.T) {
		root := Div().Children(
			Fragment().Children(
				Div().Children(Span().Text("foo")),
				Div(),
				Div(),
			),
		)
		assert.Equal(t, "<div><div><span>foo</span></div><div></div><div></div></div>", root.ToHTML())
	})

	t.Run("pretty basic example", func(t *testing.T) {
		root := Fragment().Children(
			Div().Children(Span().Text("foo")),
			Div(),
			Div(),
		)
		assert.Equal(t, strings.Join([]string{
			"<div>",
			"  <span>foo</span>",
			"</div>",
			"<div></div>",
			"<div></div>",
		}, "\n"), root.ToHTMLPretty())
	})

	t.Run("pretty basic example wrapped", func(t *testing.T) {
		root := Div().Children(
			Fragment().Children(
				Div().Children(Span().Text("foo")),
				Div(),
				Fragment().Children(
					H1().Text("Hi"),
					H2(),
				),
			),
		)
		assert.Equal(t, strings.Join([]string{
			"<div>",
			"  <div>",
			"    <span>foo</span>",
			"  </div>",
			"  <div></div>",
			"  <h1>Hi</h1>",
			"  <h2></h2>",
			"</div>",
		}, "\n"), root.ToHTMLPretty())
	})
}

func TestElement_Map(t *testing.T) {
	t.Run("generates children and skips nil values", func(t *testing.T) {
		items := []string{"foo", "bar", "baz"}
		root := Ul().Range(len(items), func(i int) Node {
			if items[i] == "bar" {
				return nil
			}
			return Li().Text(items[i])
		})

		assert.Equal(t, `<ul><li>foo</li><li>baz</li></ul>`, root.ToHTML())
	})
}

var result string

func BenchmarkHTMLPretty(b *testing.B) {
	root := Div().Children(
		Fragment().Children(
			Div().Class("foo", "bar", "baz").Children(
				Span().Text("THis is a really long string of text and I think"+
					" it makes for a good test. Here are some random characters."+
					" it makes for a good test. Here are some random characters."+
					" it makes for a good test. Here are some random characters."+
					""),
			),
			Div().ID("woo!").Attr("foo", "bar").Attr("baz", "qux").Attr("hi", "there").Children(
				Ul().Children(
					Li().Text("Item 1"),
				),
			),
			Div().Attr("foo", "bar").Children(
				Pre().Text("Hi, there."),
				Div().Text("Hi, there."),
				Span().Text("Hi, there."),
			),
		),
	)

	var r string
	for n := 0; n < b.N; n++ {
		r = root.ToHTMLPretty()
	}

	result = r
	fmt.Printf("Result %v\n", result)
}
