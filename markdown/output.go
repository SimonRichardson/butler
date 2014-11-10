package markdown

import (
	"fmt"

	"github.com/SimonRichardson/butler/butler"
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
)

type Markdown struct{}

func (m Markdown) Encode(a g.Any) ([]byte, error) {
	return nil, nil
}

func Output(server butler.Server) ([]byte, error) {
	// Build the service and output it as markdown!
	list := g.List_.To(
		h1("Butler"),
		hr(),
	)

	request := server.Describe()

	getRoute(request).Map(func(x g.Any) g.Any {
		list = g.NewCons(h2(x.(http.Route).String()), list)
		return x
	})

	document := list.FoldLeft("", func(a, b g.Any) g.Any {
		return a.(string) + b.(marks).String()
	})
	return []byte(document.(string)), nil
}

func getRoute(x g.List) g.Option {
	return x.Find(func(a g.Any) bool {
		_, ok := a.(http.Route)
		return ok
	})
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
		Type:  H2,
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
		return "----"
	}
	return ""
}

type furniture struct {
	Type furnitureType
}

func (f furniture) String() string {
	return fmt.Sprintf("%s\n\n", f.Type.String())
}

func hr() marks {
	return furniture{
		Type: HR,
	}
}
