package http

import (
	"fmt"

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
		compile = func(a g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.AsWriter(b).Map(func(a g.Any) g.Any {
					result := compilePath(asString(a))
					fmt.Println(result)
					return g.NewTuple2(a, result)
				})
			}
		}
		api = func(api doc.Api) func(g.Any) func(g.Any) g.Any {
			return func(a g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					return g.AsWriter(b).Chain(func(a g.Any) g.Writer {
						var (
							tuple  = g.AsTuple2(a)
							str    = asString(tuple.Fst())
							single = singleton(str.value)
							either = g.Either_.Of(single)
						)

						return g.NewWriter(r, singleton(api.Run(either)))
					})
				}
			}
		}
	)

	return r.path.Build().
		Chain(g.Get()).
		Chain(modify(compile)).
		Chain(constant(g.StateT_.Of(r))).
		Chain(modify(api(r.Api)))
}

func (r Route) String() string {
	return r.path.String()
}
