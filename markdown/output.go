package markdown

import (
	"fmt"

	"github.com/SimonRichardson/butler/butler"
	"github.com/SimonRichardson/butler/generic"
)

type Markdown struct{}

func (m Markdown) Encode(a generic.Any) ([]byte, error) {
	return nil, nil
}

func Output(server butler.Service) ([]byte, error) {
	// Build the service and output it as markdown!
	list := generic.ToList(
		h1("Butler"),
		hr(),
	)
	document := list.FoldLeft("", func(a, b generic.Any) generic.Any {
		return a.(string) + b.(string)
	})
	return []byte(document.(string)), nil
}

type marks interface {
	String() string
}

type headerType string

var (
	H1 headerType = "h1"
	H2 headerType = "h2"
)

func (h headerType) String() string {
	switch h {
	case H1:
		return "#"
	case H2:
		return "##"
	}
	return ""
}

type header struct {
	Type  headerType
	Value string
}

func (h header) String() string {
	return fmt.Sprintf("%s %s\n", h.Type.String(), h.Value)
}

func h1(value string) marks {
	return header{
		Type:  H1,
		Value: value,
	}
}

func h2(value string) marks {
	return header{
		Type:  H1,
		Value: value,
	}
}

type furnitureType string

var (
	HR furnitureType = "hr"
)

func (f furnitureType) String() string {
	switch f {
	case HR:
		return "----\n"
	}
	return ""
}

type furniture struct {
	Type furnitureType
}

func hr() marks {
	return furniture{
		Type: HR,
	}
}
