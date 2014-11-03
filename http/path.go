package http

import (
	"github.com/SimonRichardson/butler/doc"
	"github.com/SimonRichardson/butler/generic"
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

func (r Route) Build() generic.State {
	var (
		extract = func(x generic.Any) func(func(Route, generic.State) generic.Tuple2) generic.Tuple2 {
			return func(f func(Route, generic.State) generic.Tuple2) generic.Tuple2 {
				tuple := x.(generic.Tuple2)
				route := tuple.Fst().(Route)
				state := tuple.Snd().(generic.State)

				return f(route, state)
			}
		}
		setup = func(x generic.Any) generic.Any {
			return generic.NewTuple2(r, generic.State{})
		}
		use = func(x generic.Any) generic.Any {
			return extract(x)(func(route Route, state generic.State) generic.Tuple2 {
				return generic.NewTuple2(
					route,
					route.path.Build(),
				)
			})
		}
		execute = func(x generic.Any) generic.Any {
			return extract(x)(func(route Route, state generic.State) generic.Tuple2 {
				x := state.EvalState("")
				tuple := x.(generic.Tuple2)

				return generic.NewTuple2(
					route,
					tuple.Snd().(generic.Either),
				)
			})
		}
		api = func(x generic.Any) generic.Any {
			tuple := x.(generic.Tuple2)
			route := tuple.Fst().(Route)

			sum := func(a generic.Any) generic.Any {
				return []generic.Any{a}
			}
			folded := tuple.Snd().(generic.Either).Bimap(sum, sum)

			return generic.NewTuple2(route, route.Api.Run(folded))
		}
	)

	return generic.State_.Of(r).
		Map(setup).
		Map(use).
		Map(execute).
		Map(api)
}

func Path(path string) Route {
	return NewRoute(path)
}
