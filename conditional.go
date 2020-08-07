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

func (lf *ifStatement) Case(v interface{}, n func() Node) *ifStatement {
	lf.cases[v] = n
	return lf
}

func (lf *ifStatement) Default(n func() Node) *ifStatement {
	lf.defaultCase = n
	return lf
}

func (lf *ifStatement) ID(id string) Node {
	panic("ID not supported for Switch")
}

func (lf *ifStatement) Children(child ...Node) Node {
	panic("Children not supported for Switch")
}

func (lf *ifStatement) Range(n int, child func(i int) Node) Node {
	panic("Range not supported for Switch")
}

func (lf *ifStatement) Text(text string) Node {
	panic("Text not supported for Switch")
}

func (lf *ifStatement) Textf(format string, a ...interface{}) Node {
	panic("Textf not supported for Switch")
}

func (lf *ifStatement) Attr(name, value string) Node {
	panic("Attr not supported for Switch")
}

func (lf *ifStatement) Class(cls ...string) Node {
	panic("Class not supported for Switch")
}

func (lf *ifStatement) Style(name, value string) Node {
	panic("Style not supported for Switch")
}

func (lf *ifStatement) HTML() string {
	return lf.base.HTML()
}

func (lf *ifStatement) HTMLPretty() string {
	return lf.base.HTMLPretty()
}

func (lf *ifStatement) Write(w io.Writer) (int, error) {
	return lf.base.Write(w)
}

func (lf *ifStatement) GetNode() *RawNode {
	var node Node
	if n, ok := lf.cases[lf.v]; ok {
		node = n()
	}

	if node == nil && lf.defaultCase != nil {
		node = lf.defaultCase()
	}

	if node == nil {
		return Fragment().GetNode()
	}

	return node.GetNode()
}
