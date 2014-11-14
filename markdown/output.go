package markdown

import (
	"github.com/SimonRichardson/butler/butler"
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
)

const (
	DefaultString string = ""
)

type mark interface {
	IsBlock() bool
	Children() g.Option
	String() string
}

type Markdown struct{}

func (m Markdown) Encode(a g.Any) ([]byte, error) {
	return nil, nil
}

func Output(server butler.Server) ([]byte, error) {
	// Build the service and output it as markdown!
	doc := document(append(templateHeader(), templateFooter()...)...)
	return []byte(doc.String()), nil
}

func templateHeader() []mark {
	return []mark{
		h1(link("Butler", "http://github.com/simonrichardson/butler")),
		h4(str("Serving you content in a monadic style.")),
		hr1(),
		ul(
			link("Route definitions", "#routes"),
		),
		h5(str("Routes")),
		p(str("The route definitions for your service are as follows:")),
	}
}

func templateFooter() []mark {
	return []mark{
		hr2(),
		center(link("Served by Butler", "http://github.com/simonrichardson/butler")),
	}
}

func getMethod(x g.List) g.Option {
	return x.Find(func(a g.Any) bool {
		_, ok := a.(http.Method)
		return ok
	})
}

func getRoute(x g.List) g.Option {
	return x.Find(func(a g.Any) bool {
		_, ok := a.(http.Route)
		return ok
	})
}
