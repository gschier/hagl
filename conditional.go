package hagl

import (
	"io"
)

type ifStatement struct {
	base        *RawNode
	v           interface{}
	cases       map[interface{}]func() Node
	defaultCase func() Node
}

func Switch(v interface{}) *ifStatement {
	return &ifStatement{
		base:  Fragment().(*RawNode),
		cases: make(map[interface{}]func() Node),
		v:     v,
	}
}

func (c *ifStatement) Case(v interface{}, n func() Node) *ifStatement {
	c.cases[v] = n
	return c
}

func (c *ifStatement) Default(n func() Node) *ifStatement {
	c.defaultCase = n
	return c
}

func (c *ifStatement) ID(_ string) Node {
	panic("ID not supported for Switch")
}

func (c *ifStatement) Children(_ ...Node) Node {
	panic("Children not supported for Switch")
}

func (c *ifStatement) Range(_ int, _ func(i int) Node) Node {
	panic("Range not supported for Switch")
}

func (c *ifStatement) Text(_ string) Node {
	panic("Text not supported for Switch")
}

func (c *ifStatement) Textf(_ string, _ ...interface{}) Node {
	panic("Textf not supported for Switch")
}

func (c *ifStatement) Attr(_, _ string) Node {
	panic("Attr not supported for Switch")
}

func (c *ifStatement) Class(_ ...string) Node {
	panic("Class not supported for Switch")
}

func (c *ifStatement) Style(_, _ string) Node {
	panic("Style not supported for Switch")
}

func (c *ifStatement) HTML() string {
	return c.base.HTML()
}

func (c *ifStatement) HTMLPretty() string {
	return c.base.HTMLPretty()
}

func (c *ifStatement) Write(w io.Writer) (int, error) {
	return c.base.Write(w)
}

func (c *ifStatement) GetNode() *RawNode {
	var node Node
	if n, ok := c.cases[c.v]; ok {
		node = n()
	}

	if node == nil && c.defaultCase != nil {
		node = c.defaultCase()
	}

	if node == nil {
		return Fragment().GetNode()
	}

	return node.GetNode()
}
