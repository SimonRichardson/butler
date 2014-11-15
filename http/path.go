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
		path: NewString(path, PathChar()),
	}
}

func (r Route) Build() g.StateT {
	var (
		api = func(api doc.Api) func(g.Any) func(g.Any) g.Any {
			return func(a g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					return g.AsWriter(b).Chain(func(a g.Any) g.Writer {
						str := g.Either_.Of(singleton(a.(String).value))
						return g.NewWriter(r, singleton(api.Run(str)))
					})
				}
			}
		}
	)

	return r.path.Build().
		Chain(g.Get()).
		Chain(constant(g.StateT_.Of(r))).
		Chain(modify(api(r.Api)))
}

func (r Route) String() string {
	return r.path.String()
}
