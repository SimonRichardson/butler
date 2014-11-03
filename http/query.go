package http

import (
	"strconv"

	"github.com/SimonRichardson/butler/doc"
	"github.com/SimonRichardson/butler/generic"
)

type QueryType string

const (
	QInt    QueryType = "Int"
	QString QueryType = "String"
)

type Query struct {
	doc.Api
	name  QueryType
	value String
	build func(generic.Any) generic.Any
}

func NewQuery(name QueryType, value string, build func(generic.Any) generic.Any) Query {
	return Query{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected query %s"),
			doc.NewInlineText("Unexpected query %s"),
		)),
		name:  name,
		value: NewString(value, urlChar()),
		build: build,
	}
}

func (q Query) Build() generic.State {
	var (
		extract = func(x generic.Any) func(func(Query, generic.State) generic.Tuple2) generic.Tuple2 {
			return func(f func(Query, generic.State) generic.Tuple2) generic.Tuple2 {
				tuple := x.(generic.Tuple2)
				query := tuple.Fst().(Query)
				state := tuple.Snd().(generic.State)

				return f(query, state)
			}
		}
		setup = func(x generic.Any) generic.Any {
			return generic.NewTuple2(q, generic.State{})
		}
		use = func(x generic.Any) generic.Any {
			return extract(x)(func(query Query, state generic.State) generic.Tuple2 {
				return generic.NewTuple2(
					query,
					query.value.Build(),
				)
			})
		}
		execute = func(x generic.Any) generic.Any {
			return extract(x)(func(query Query, state generic.State) generic.Tuple2 {
				x := state.EvalState("")
				tuple := x.(generic.Tuple2)

				return generic.NewTuple2(
					query,
					tuple.Snd().(generic.Either),
				)
			})
		}
		api = func(x generic.Any) generic.Any {
			tuple := x.(generic.Tuple2)
			query := tuple.Fst().(Query)

			sum := func(a generic.Any) generic.Any {
				return []generic.Any{a}
			}
			folded := tuple.Snd().(generic.Either).Bimap(sum, sum)

			return generic.NewTuple2(query, query.Api.Run(folded))
		}
	)

	return generic.State_.Of(q).
		Map(setup).
		Map(use).
		Map(execute).
		Map(api)
}

func QueryInt(name string) Query {
	return NewQuery(QInt, name, func(x generic.Any) generic.Any {
		y, _ := strconv.Atoi(x.(string))
		return y
	})
}

func QueryString(name string) Query {
	return NewQuery(QString, name, generic.Identity())
}
