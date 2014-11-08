package http

import (
	"strconv"

	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
)

type Query struct {
	doc.Api
	value String
	build func(g.Any) g.Any
}

func NewQuery(value string, build func(g.Any) g.Any) Query {
	return Query{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected query `%s`"),
			doc.NewInlineText("Unexpected query `%s`"),
		)),
		value: NewString(value, UrlChar()),
		build: build,
	}
}

func (q Query) Build() g.StateT {
	var (
		api = func(api doc.Api) func(g.Any) func(g.Any) g.Any {
			return func(a g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					return asWriter(b).Chain(func(a g.Any) g.Writer {
						str := g.Either_.Of(singleton(a.(String).value))
						return g.NewWriter(q, singleton(api.Run(str)))
					})
				}
			}
		}
	)

	return q.value.Build().
		Chain(get()).
		Chain(constant(g.StateT_.Of(q))).
		Chain(modify(api(q.Api)))
}

func QueryInt(name string) Query {
	return NewQuery(name, func(x g.Any) g.Any {
		y, _ := strconv.Atoi(x.(string))
		return y
	})
}

func QueryString(name string) Query {
	return NewQuery(name, g.Identity())
}
