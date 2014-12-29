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
		name:      NewString(name, UrlChar()),
		queryType: queryType,
		build:     build,
	}
}

func (q Query) Build() g.WriterT {
	return g.WriterT_.Of("Query(???)")
	/*var (
		query = func(t QueryType) func(g.Any) func(g.Any) g.Any {
			return func(g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					return g.NewTuple2(b, t.Type())
				}
			}
		}
		api = func(api doc.Api) func(g.Any) func(g.Any) g.Any {
			return func(a g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					tuple := g.AsTuple2(b)
					return g.AsWriter(tuple.Fst()).Chain(func(a g.Any) g.Writer {
						var (
							name  = singleton(a.(String).value)
							value = tuple.Snd()
							str   = g.Either_.Of(append(name, value))
						)
						return g.NewWriter(q, singleton(api.Run(str)))
					})
				}
			}
		}
	)

	return q.name.Build().
		Chain(g.Get()).
		Chain(modify(query(q.queryType))).
		Chain(constant(g.StateT_.Of(q))).
		Chain(modify(api(q.Api)))
	*/
}

func (q Query) Name() string {
	return q.name.value
}

func (q Query) Type() string {
	return q.queryType.Type()
}

func (q Query) String() string {
	return fmt.Sprintf("%s [%s]", q.name, q.queryType.Type())
}
