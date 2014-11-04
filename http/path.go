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
	var (
		api = func(api doc.Api) func(g.Any) func(g.Any) g.Any {
			return func(a g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					var (
						sum = func(a g.Any) g.Any {
							return []g.Any{a}
						}
						folded = b.(g.Either).Bimap(sum, sum)
					)
					return api.Run(folded)
				}
			}
		}
	)

	return r.path.Build().
		Chain(get()).
		Chain(constant(g.StateT_.Of(r))).
		Chain(modify(api(r.Api)))
}

func Path(path string) Route {
	return NewRoute(path)
}
