package hagl

import (
	"io"
)

var _ Node = new(component)

type component struct {
	base   *RawNode
	render func(children []Node) Node
}

func NewComponent(render func(children []Node) Node) func() Node {
	return func() Node {
		return &component{
			base:   Fragment().GetNode(),
			render: render,
		}
	}
}

func (c *component) ID(id string) Node {
	return c.base.ID(id)
}

func (c *component) Children(child ...Node) Node {
	c.base.Children(child...)
	return c
}

func (c *component) Range(n int, child func(i int) Node) Node {
	c.base.Range(n, child)
	return c
}

func (c *component) Text(text ...string) Node {
	c.base.Text(text...)
	return c
}

func (c *component) Textf(format string, a ...interface{}) Node {
	c.base.Textf(format, a...)
	return c
}

func (c *component) Attr(name, value string) Node {
	c.base.Attr(name, value)
	return c
}

func (c *component) Class(cls ...string) Node {
	c.base.Class(cls...)
	return c
}

func (c *component) ClassIf(condition bool, cls string) Node {
	c.base.ClassIf(condition, cls)
	return c
}

func (c *component) Style(name, value string) Node {
	c.base.Style(name, value)
	return c
}

func (c *component) Href(value string) Node {
	c.base.Href(value)
	return c
}

func (c *component) Rel(value string) Node {
	c.base.Rel(value)
	return c
}

func (c *component) Src(value string) Node {
	c.base.Src(value)
	return c
}

func (c *component) Target(value string) Node {
	c.base.Target(value)
	return c
}

func (c *component) Alt(value string) Node {
	c.base.Alt(value)
	return c
}

func (c *component) Type(value string) Node {
	c.base.Type(value)
	return c
}

func (c *component) Title(value string) Node {
	c.base.Title(value)
	return c
}

func (c *component) If(b bool) Node {
	c.base.If(b)
	return c
}

func (c *component) Extend(base Node) Node {
	c.base.Extend(base)
	return c
}

func (c *component) merge() Node {
	var baseCopy = new(RawNode)
	*baseCopy = *c.base

	// Render the component, passing the children that were added
	n := c.render(baseCopy.children)

	// Update base
	baseCopy.children = []Node{}
	n.Extend(baseCopy)
	return n
}

func (c *component) HTML() string {
	return c.merge().HTML()
}

func (c *component) HTMLPretty() string {
	return c.merge().HTMLPretty()
}

func (c *component) Write(w io.Writer) (int, error) {
	return c.merge().Write(w)
}

func (c *component) WritePretty(w io.Writer) (int, error) {
	return c.merge().WritePretty(w)
}

func (c *component) MustWrite(w io.Writer) {
	c.merge().MustWrite(w)
}

func (c *component) MustWritePretty(w io.Writer) {
	c.merge().MustWritePretty(w)
}

func (c *component) GetNode() *RawNode {
	return c.merge().GetNode()
}
