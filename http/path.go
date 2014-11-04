package http

import (
	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
)

type Route struct {
	doc.Api
	path String
}

func NewRoute(path string) Route {
	return Route{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected route `%s`"),
			doc.NewInlineText("Unexpected route `%s`"),
		)),
		path: NewString(path, pathChar()),
	}
}

func (r Route) Build() g.StateT {
	var ()

	return g.StateT_.Of(r).
		Chain(compose(g.StateT_.Modify)(g.Constant1)).
		Chain(func(a g.Any) g.StateT {
		return g.StateT_.Modify(func(b g.Any) g.Any {
			return a.(g.StateT).ExecState("")
		})
	})

}

func Path(path string) Route {
	return NewRoute(path)
}
