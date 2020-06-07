package go_temper

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestElement_HTML(t *testing.T) {
	t.Run("generates simple element", func(t *testing.T) {
		root := El("div").Text("Hello World!")
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
		assert.Equal(t, "<div><ul><li>Item 1</li><li>Item 2</li></ul></div>", root.HTMLPretty())
	})
}

func TestElement_Attr(t *testing.T) {
	t.Run("adds attr", func(t *testing.T) {
		root := El("div").Attr("style", "display: block")
		assert.Equal(t, `<div style="display: block"></div>`, root.HTML())
	})

	t.Run("overwrites attribute", func(t *testing.T) {
		root := El("div").Attr("id", "1").Attr("id", "2")
		assert.Equal(t, `<div id="2"></div>`, root.HTML())
	})
}

func TestElement_Class(t *testing.T) {
	t.Run("adds classes", func(t *testing.T) {
		root := El("button").Class("btn", "btn--primary")
		assert.Equal(t, `<button class="btn btn--primary"></button>`, root.HTML())
	})

	t.Run("adds duplicate classes", func(t *testing.T) {
		root := El("div").Class("btn", "btn--primary", "btn")
		assert.Equal(t, `<div class="btn btn--primary"></div>`, root.HTML())
	})

	t.Run("appends to manually set attr", func(t *testing.T) {
		root := El("button").Attr("class", "btn").Class("btn--primary")
		assert.Equal(t, `<button class="btn btn--primary"></button>`, root.HTML())
	})
}

func TestElement_HTMLPretty(t *testing.T) {
	t.Run("pretty HTML", func(t *testing.T) {
		root := Div().Children(
			Ul().Children(
				Li().Text("1"),
				Li().Text("Hello\n\n\n\nWorld!"),
				Li().Text("This is a really long string that will get wrapped because it's too long."),
			),
		)
		assert.Equal(t, strings.Join([]string{
			"<div>",
			"  <ul>",
			"    <li>1</li>",
			"    <li>Hello World!</li>",
			"    <li>",
			"      This is a really long string that will get wrapped because it's too long.",
			"    </li>",
			"  </ul>",
			"</div>",
		}, "\n"), root.HTMLPretty())
	})
}
