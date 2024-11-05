package hagl

import (
	"io"
)

type SwitchStatement struct {
	base        *RawNode
	v           interface{}
	cases       map[interface{}]func() Node
	defaultCase func() Node
}

func Switch(v interface{}) *SwitchStatement {
	return &SwitchStatement{
		base:  Fragment().(*RawNode),
		cases: make(map[interface{}]func() Node),
		v:     v,
	}
}

func (c *SwitchStatement) Case(v interface{}, n func() Node) *SwitchStatement {
	c.cases[v] = n
	return c
}

func (c *SwitchStatement) Default(n func() Node) *SwitchStatement {
	c.defaultCase = n
	return c
}

func (c *SwitchStatement) ID(_ string) Node {
	panic("ID not supported for Switch")
}

func (c *SwitchStatement) Children(_ ...Node) Node {
	panic("Children not supported for Switch")
}

func (c *SwitchStatement) Range(_ int, _ func(i int) Node) Node {
	panic("Range not supported for Switch")
}

func (c *SwitchStatement) Text(_ string) Node {
	panic("Text not supported for Switch")
}

func (c *SwitchStatement) Textf(_ string, _ ...interface{}) Node {
	panic("Textf not supported for Switch")
}

func (c *SwitchStatement) Attr(_, _ string) Node {
	panic("Attr not supported for Switch")
}

func (c *SwitchStatement) Class(_ ...string) Node {
	panic("Class not supported for Switch")
}

func (c *SwitchStatement) Style(_, _ string) Node {
	panic("Style not supported for Switch")
}

func (c *SwitchStatement) HTML() string {
	return c.base.HTML()
}

func (c *SwitchStatement) HTMLPretty() string {
	return c.base.HTMLPretty()
}

func (c *SwitchStatement) Write(w io.Writer) (int, error) {
	return c.base.Write(w)
}

func (c *SwitchStatement) GetNode() *RawNode {
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
