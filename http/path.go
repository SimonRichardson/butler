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

func (r Route) Build() g.WriterT {
	var (
		extract = func(a g.Any) g.WriterT {
			var (
				x = g.AsTuple3(a)
				y = AsString(x.Fst())
			)
			return g.WriterT_.Of(y.String()).
				Tell(fmt.Sprintf("Extract `%v`", y))
		}
		api = func(x doc.Api) func(g.Either) g.WriterT {
			return func(y g.Either) g.WriterT {
				return g.WriterT_.Lift(x.Run(y)).
					Tell(fmt.Sprintf("Api `%v`", y))
			}
		}
		finalize = func(a Route) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.NewTuple2(a, b)
			}
		}
		program = r.path.Build().
			Chain(extract)
	)

	return join(program, api(r.Api), func(x g.Any) []g.Any {
		return singleton(x)
	}).Bimap(
		finalize(r),
		finalize(r),
	)
}

func (r Route) Route() g.Tree {
	result := compilePath(r.String())
	return g.Tree_.FromList(result)
}

func (r Route) String() string {
	return r.path.String()
}
