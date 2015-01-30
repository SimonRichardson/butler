package http

import (
	"fmt"

	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
)

type QueryType interface {
	Type() string
}

type queryType struct {
	value string
}

func newQueryType(val string) *queryType {
	return &queryType{
		value: val,
	}
}

func (t *queryType) Type() string {
	return t.value
}

var (
	QDate     = newQueryType("date")
	QDateTime = newQueryType("date-time")
	QInt      = newQueryType("int")
	QUint     = newQueryType("uint")
	QString   = newQueryType("string")
)

type RawQuery interface {
	Name() string
	Type() string
}

type Query struct {
	doc.Api
	name      String
	queryType QueryType
	build     func(g.Any) g.Any
}

func NewQuery(name string, queryType QueryType, build func(g.Any) g.Any) Query {
	return Query{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected query `%s` with type `%s`"),
			doc.NewInlineText("Unexpected query `%s` with type `%s`"),
		)),
		name:      NewString(name, QueryChar()),
		queryType: queryType,
		build:     build,
	}
}

func (q Query) Build() g.WriterT {
	var (
		extract = func(a g.Any) g.WriterT {
			var (
				x = AsResult(a)
				y = AsString(x.Builder())
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
		finalize = func(a Query) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.NewTuple2(a, b)
			}
		}
		matcher = func(name g.WriterT) func(g.Any) g.Any {
			return func(a g.Any) g.Any {
				var (
					match = func(a g.Any) func(g.Any) g.Any {
						return func(b g.Any) g.Any {
							var (
								x          = name.Run().Fst()
								y          = g.AsTuple2(b).Snd()
								parts      = y.([]string)
								paramName  = parts[0]
								paramValue = parts[1]
								build      = func(a g.Any) g.Any {
									return q.build(paramValue)
								}
							)
							return x.Chain(func(x g.Any) g.Either {
								var (
									c = AsResult(x).Matcher()
									d = g.Either_.FromBool(len(paramName) > 0, c)
								)
								return d.Chain(func(a g.Any) g.Either {
									s := g.AsStateT(a)
									return s.EvalState(paramName)
								})
							}).
								Map(build).
								Bimap(matchPut(a), matchPut(a))
						}
					}
					program = g.StateT_.Of(a).
						Chain(modify(matchSplit("="))).
						Chain(g.Get()).
						Chain(modify(match)).
						Chain(g.Get()).
						Chain(matchFlatten)
				)
				return Result_.FromTuple3(g.AsTuple2(a).Append(program))
			}
		}

		name = q.name.Build()

		program = name.
			Chain(extract)
	)
	return join(program, api(q.Api), func(x g.Any) []g.Any {
		return append(singleton(x), q.Type())
	}).Bimap(
		finalize(q),
		finalize(q),
	).Bimap(
		matcher(name),
		matcher(name),
	)
}

func (q Query) Name() string {
	return q.name.value
}

func (q Query) Type() string {
	return q.queryType.Type()
}

func (q Query) String() string {
	return fmt.Sprintf("%s[%s]", q.name, q.Type())
}
