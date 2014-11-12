package markdown

import (
	"fmt"

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

	tree := g.NewTreeNode(
		document(),
		g.List_.To(
			g.Tree_.Of(h1("Butler")),
			g.Tree_.Of(hr1()),
			g.Tree_.Of(hr2()),
		),
	)

	/*
		request := server.Describe()

		getMethod(request).Map(func(x g.Any) g.Any {
			list = g.NewCons(h2(x.(http.Method).String()), list)
			return x
		})

		getRoute(request).Map(func(x g.Any) g.Any {
			list = g.NewCons(h2(x.(http.Route).String()), list)
			return x
		})
	*/

	result := tree.FoldLeft("", func(a, b g.Any) g.Any {
		return fmt.Sprintf("%s%s", a.(string), b.(marks).String(DefaultIndent))
	})
	return []byte(result.(string)), nil
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
