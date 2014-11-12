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
	return false
}

func (h headerType) Children() g.Option {
	return g.Option_.Empty()
}

func (h headerType) String() string {
	return string(h)
}

type header struct {
	Type  headerType
	Value marks
}

func (h header) IsInline() bool {
	return false
}

func (h header) Children() g.Option {
	return g.Option_.Of(g.List_.Of(h.Value))
}

func (h header) String() string {
	return fmt.Sprintf("%s", h.Type.String())
}

func h1(val marks) header {
	return header{
		Type:  H1,
		Value: val,
	}
}

func h2(val marks) header {
	return header{
		Type:  H2,
		Value: val,
	}
}

func h3(val marks) header {
	return header{
		Type:  H3,
		Value: val,
	}
}

func h4(val marks) header {
	return header{
		Type:  H4,
		Value: val,
	}
}

func h5(val marks) header {
	return header{
		Type:  H5,
		Value: val,
	}
}

func h6(val marks) header {
	return header{
		Type:  H6,
		Value: val,
	}
}
