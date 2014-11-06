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
		api = func(r Route) func(g.Any) func(g.Any) g.Any {
			return func(a g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					fmt.Println(a)
					fmt.Println(b.(g.Writer).Run())
					fmt.Println("-------")
					return b.(g.Writer).Chain(func(a g.Any) g.Writer {

						str := g.NewRight(singleton(a.(String).value))

						writer := g.NewWriter(r, singleton(r.Api.Run(str)))
						fmt.Println(writer.Run())
						return writer
					})
				}
			}
		}
	)

	return r.path.Build().
		Chain(get()).
		Chain(constant(g.StateT_.Of(r))).
		Chain(modify(api(r)))
}

func Path(path string) Route {
	return NewRoute(path)
}
