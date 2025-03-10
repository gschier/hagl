package hagl_test

import (
	. "github.com/gschier/hagl"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestComponent(t *testing.T) {
	t.Run("single component", func(t *testing.T) {
		Layout := NewComponent(func(children []Node) Node {
			return Main().Class("layout").Attr("layout", "xxx").Children(children...)
		})

		El := Layout().Class("added").Attr("added", "yyy").Children(
			H1().Class("h1-class").Text("Hello"),
		)

		assert.Equal(t, `<main class="layout added" layout="xxx" added="yyy"><h1 class="h1-class">Hello</h1></main>`, El.ToHTML())
	})

	t.Run("multiple components", func(t *testing.T) {
		Layout := NewComponent(func(children []Node) Node {
			return Div().Class("layout").Children(children...)
		})

		Heading := NewComponent(func(children []Node) Node {
			return H1().Class("h1").Children(children...)
		})

		assert.Equal(t, `<div class="layout"><h1 class="h1">Home</h1></div>`, Layout().Children(
			Heading().Text("Home"),
		).ToHTML())
	})

	t.Run("correct attrs order", func(t *testing.T) {
		Btn := NewComponent(func(children []Node) Node {
			return Button().Type("button").Children(children...)
		})

		assert.Equal(t,
			`<button type="submit">Submit</button>`,
			Btn().Type("submit").Text("Submit").ToHTML(),
		)
	})
}
