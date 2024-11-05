package hagl_test

import (
	"fmt"
	assert "github.com/stretchr/testify/require"
	"strings"
	"testing"

	. "github.com/gschier/hagl"
)

func TestElement_HTML(t *testing.T) {
	t.Run("generates simple element", func(t *testing.T) {
		root := Div().Text("Hello World!")
		assert.Equal(t, "<div>Hello World!</div>", root.HTML())
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
		}, "\n"), root.HTMLPretty())
	})
}

func TestElement_Attr(t *testing.T) {
	t.Run("adds attr", func(t *testing.T) {
		root := Div().Attr("style", "display: block")
		assert.Equal(t, `<div style="display: block"></div>`, root.HTML())
	})

	t.Run("overwrites attribute", func(t *testing.T) {
		root := Div().Attr("id", "1").Attr("id", "2")
		assert.Equal(t, `<div id="2"></div>`, root.HTML())
	})
}

func TestElement_Class(t *testing.T) {
	t.Run("adds classes", func(t *testing.T) {
		root := Button().Class("btn", "btn--primary")
		assert.Equal(t, `<button class="btn btn--primary"></button>`, root.HTML())
	})

	t.Run("adds duplicate classes", func(t *testing.T) {
		root := Div().Class("btn", "btn--primary", "btn")
		assert.Equal(t, `<div class="btn btn--primary"></div>`, root.HTML())
	})

	t.Run("appends to manually set attr", func(t *testing.T) {
		root := Button().Attr("class", "btn").Class("btn--primary")
		assert.Equal(t, `<button class="btn btn--primary"></button>`, root.HTML())
	})
}

func TestElement_Style(t *testing.T) {
	t.Run("adds style", func(t *testing.T) {
		root := Button().Style("background", "red")
		assert.Equal(t, `<button style="background:red"></button>`, root.HTML())
	})

	t.Run("adds multiple styles", func(t *testing.T) {
		root := Button().Style("background", "red").Style("color", "white")
		assert.Equal(t, `<button style="background:red;color:white"></button>`, root.HTML())
	})

	t.Run("doesn't overwrite style", func(t *testing.T) {
		root := Button().Style("background", "red").Style("background", "blue")
		assert.Equal(t, `<button style="background:red;background:blue"></button>`, root.HTML())
	})
}

func TestElement_HTMLPretty(t *testing.T) {
	t.Run("pre", func(t *testing.T) {
		root := Pre().Text("function foo() {\n  return 'Hello World!';\n}")
		assert.Equal(t, strings.Join([]string{
			"<pre>function foo() {",
			"  return 'Hello World!';",
			"}</pre>",
		}, "\n"), root.HTMLPretty())
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
			"    <li>This is a really long string that will get wrapped because it's too long.</li>",
			"  </ul>",
			"  <pre>function foo() {",
			"  return 'Hello World!';",
			"}</pre>",
			"  <pre><div>foo</div>Bar<h2>woo!</h2><!-- That was cool --></pre>",
			"</div>",
		}, "\n"), root.HTMLPretty())
	})
}

func TestEl(t *testing.T) {
	t.Run("test component", func(t *testing.T) {
		Btn := func() Node {
			return Button().Class("btn btn--primary").Attr("type", "button")
		}

		root := Div().Children(Btn().Text("Click Me!"))
		assert.Equal(t, `<div><button class="btn btn--primary" type="button">Click Me!</button></div>`, root.HTML())
	})
}

func TestFragmentEl(t *testing.T) {
	t.Run("works with no children", func(t *testing.T) {
		root := Hr().Class("red").Attr("type", "foo")
		assert.Equal(t, `<hr class="red" type="foo"/>`, root.HTML())
	})

	t.Run("works with children", func(t *testing.T) {
		root := Hr().Class("red").Attr("type", "foo").Text("foo")
		assert.Equal(t, `<hr class="red" type="foo">foo</hr>`, root.HTML())
	})
}

func TestSelfClosingEl(t *testing.T) {
	t.Run("works with no children", func(t *testing.T) {
		root := Hr().Class("red").Attr("type", "foo")
		assert.Equal(t, `<hr class="red" type="foo"/>`, root.HTML())
	})

	t.Run("works with children", func(t *testing.T) {
		root := Hr().Class("red").Attr("type", "foo").Text("foo")
		assert.Equal(t, `<hr class="red" type="foo">foo</hr>`, root.HTML())
	})
}

func TestComment(t *testing.T) {
	t.Run("basic example", func(t *testing.T) {
		root := Comment("This is a comment")
		assert.Equal(t, "<!-- This is a comment -->", root.HTML())
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
		}, "\n"), root.HTMLPretty())
	})
}

func TestFragment(t *testing.T) {
	t.Run("basic example", func(t *testing.T) {
		root := Fragment().Children(
			Div().Children(Span().Text("foo")),
			Div(),
			Div(),
		)
		assert.Equal(t, "<div><span>foo</span></div><div></div><div></div>", root.HTML())
	})

	t.Run("basic example wrapped", func(t *testing.T) {
		root := Div().Children(
			Fragment().Children(
				Div().Children(Span().Text("foo")),
				Div(),
				Div(),
			),
		)
		assert.Equal(t, "<div><div><span>foo</span></div><div></div><div></div></div>", root.HTML())
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
		}, "\n"), root.HTMLPretty())
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
		}, "\n"), root.HTMLPretty())
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

		assert.Equal(t, `<ul><li>foo</li><li>baz</li></ul>`, root.HTML())
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
		r = root.HTMLPretty()
	}

	result = r
	fmt.Printf("Result %v\n", result)
}
