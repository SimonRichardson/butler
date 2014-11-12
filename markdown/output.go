package markdown

import (
	"github.com/SimonRichardson/butler/butler"
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
)

const (
	DefaultIndent string = ""
	DefaultString string = ""
)

type marks interface {
	String(indent string) string
}

type Markdown struct{}

func (m Markdown) Encode(a g.Any) ([]byte, error) {
	return nil, nil
}

func Output(server butler.Server) ([]byte, error) {
	// Build the service and output it as markdown!
	list := g.List_.To(
		h1("Butler"),
		hr1(),
	)

	request := server.Describe()

	getMethod(request).Map(func(x g.Any) g.Any {
		list = g.NewCons(h2(x.(http.Method).String()), list)
		return x
	})

	getRoute(request).Map(func(x g.Any) g.Any {
		list = g.NewCons(h2(x.(http.Route).String()), list)
		return x
	})

	document := list.FoldLeft("", func(a, b g.Any) g.Any {
		return a.(string) + b.(marks).String(DefaultIndent)
	})
	return []byte(document.(string)), nil
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
