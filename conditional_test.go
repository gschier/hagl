package hagl_test

import (
	assert "github.com/stretchr/testify/require"
	"testing"

	. "github.com/gschier/hagl"
)

func TestSwitch(t *testing.T) {
	t.Run("matches default case", func(t *testing.T) {
		r := Div().Children(
			Switch("foo").
				Case("bar", func() Node { return Text("bar") }).
				Case("baz", func() Node { return Text("baz") }).
				Default(func() Node { return Text("default") }).
				GetNode(),
		)
		assert.Equal(t, "<div>default</div>", r.ToHTML())
	})

	t.Run("matches case", func(t *testing.T) {
		r := Div().Children(
			Switch("baz").
				Case("bar", func() Node { return Text("bar") }).
				Case("baz", func() Node { return Text("baz") }).
				Default(func() Node { return Text("default") }).
				GetNode(),
		)
		assert.Equal(t, "<div>baz</div>", r.ToHTML())
	})

	t.Run("matches default but no default set", func(t *testing.T) {
		r := Div().Children(
			Switch("foo").
				Case("bar", func() Node { return Text("bar") }).
				Case("baz", func() Node { return Text("baz") }).
				GetNode(),
		)
		assert.Equal(t, "<div></div>", r.ToHTML())
	})
}
