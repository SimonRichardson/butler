package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type headerType string

var (
	H1 headerType = "#"
	H2 headerType = "##"
	H3 headerType = "###"
	H4 headerType = "####"
	H5 headerType = "#####"
	H6 headerType = "######"
)

func (h headerType) IsBlock() bool {
	return false
}

func (h headerType) Children() g.Option {
	return g.Option_.Empty()
}

func (h headerType) String() string {
	return fmt.Sprintf("%s ", string(h))
}

type header struct {
	nodes g.List
}

func (h header) IsBlock() bool {
	return true
}

func (h header) Children() g.Option {
	return g.Option_.Of(h.nodes)
}

func (h header) String() string {
	return ""
}

func h1(val mark) header {
	return header{
		nodes: g.List_.To(H1, val),
	}
}

func h2(val mark) header {
	return header{
		nodes: g.List_.To(H2, val),
	}
}

func h3(val mark) header {
	return header{
		nodes: g.List_.To(H3, val),
	}
}

func h4(val mark) header {
	return header{
		nodes: g.List_.To(H4, val),
	}
}

func h5(val mark) header {
	return header{
		nodes: g.List_.To(H5, val),
	}
}

func h6(val mark) header {
	return header{
		nodes: g.List_.To(H6, val),
	}
}
