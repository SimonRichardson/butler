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
			doc.NewInlineText("Expected route `%s` parts `%s`"),
			doc.NewInlineText("Unexpected route `%s` parts `%s`"),
		)),
		path: NewString(path, PathChar()),
	}
}

func (r Route) Build() g.StateT {
	var (
		compile = func(a g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.AsWriter(b).Map(func(a g.Any) g.Any {
					return g.NewTuple2(a, r.Route())
				})
			}
		}
		api = func(api doc.Api) func(g.Any) func(g.Any) g.Any {
			return func(a g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					return g.AsWriter(b).Chain(func(a g.Any) g.Writer {
						var (
							tuple  = g.AsTuple2(a)
							str    = AsString(tuple.Fst())
							parts  = g.Tree_.ToList(g.AsTree(tuple.Snd()))
							single = append(singleton(str.value), g.List_.ToSlice(parts))
							either = g.Either_.Of(single)
						)

						return g.NewWriter(r, singleton(api.Run(either)))
					})
				}
			}
		}
		matcher = func(a g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.AsWriter(b).Map(func(x g.Any) g.Any {
					var (
						program = g.StateT_.Of(x)
					)
					return g.NewTuple2(b, program)
				})
			}
		}
	)

	return r.path.Build().
		Chain(g.Get()).
		Chain(modify(compile)).
		Chain(constant(g.StateT_.Of(r))).
		Chain(modify(api(r.Api))).
		Chain(modify(matcher))
}

func (r Route) Route() g.Tree {
	result := compilePath(r.String())
	return g.Tree_.FromList(result)
}

func (r Route) String() string {
	return r.path.String()
}
