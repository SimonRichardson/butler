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

func (h headerType) IsInline() bool {
	return true
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

func (h header) IsInline() bool {
	return false
}

func (h header) Children() g.Option {
	return g.Option_.Of(h.nodes)
}

func (h header) String() string {
	return ""
}

func h1(val marks) header {
	return header{
		nodes: fromMarks(append([]marks{H1}, val)),
	}
}

func h2(val marks) header {
	return header{
		nodes: fromMarks(append([]marks{H2}, val)),
	}
}

func h3(val marks) header {
	return header{
		nodes: fromMarks(append([]marks{H3}, val)),
	}
}

func h4(val marks) header {
	return header{
		nodes: fromMarks(append([]marks{H4}, val)),
	}
}

func h5(val marks) header {
	return header{
		nodes: fromMarks(append([]marks{H5}, val)),
	}
}

func h6(val marks) header {
	return header{
		nodes: fromMarks(append([]marks{H6}, val)),
	}
}
