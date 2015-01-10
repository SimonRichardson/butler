package http

import (
	"fmt"

	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
)

type MethodType string

const (
	DELETE  MethodType = "delete"
	GET     MethodType = "get"
	HEAD    MethodType = "head"
	OPTIONS MethodType = "options"
	PATCH   MethodType = "patch"
	POST    MethodType = "post"
	PUT     MethodType = "put"
	TRACE   MethodType = "trace"
)

type Method struct {
	doc.Api
	method String
}

func NewMethod(method MethodType) Method {
	return Method{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected method `%s`"),
			doc.NewInlineText("Unexpected method `%s`"),
		)),
		method: NewString(string(method), MethodChar()),
	}
}

func (m Method) Build() g.WriterT {
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
		finalize = func(a Method) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.NewTuple2(a, b)
			}
		}
		matcher = func(method g.WriterT) func(g.Any) g.Any {
			return func(a g.Any) g.Any {
				var (
					match = func(a g.Any) func(g.Any) g.Any {
						return func(b g.Any) g.Any {
							var (
								x    = method.Run().Fst()
								y    = g.AsTuple2(b).Snd()
								part = y.(string)
							)
							return x.Chain(func(x g.Any) g.Either {
								var (
									c = g.AsTuple3(x).Trd()
									d = g.Either_.Of(c)
								)
								return d.Chain(func(a g.Any) g.Either {
									s := g.AsStateT(a)
									return s.EvalState(part)
								})
							}).Bimap(matchPut(a), matchPut(a))
						}
					}
					program = g.StateT_.Of(a).
						Chain(modify(matchGet)).
						Chain(g.Get()).
						Chain(modify(match)).
						Chain(g.Get()).
						Chain(matchFlatten)
				)
				return g.AsTuple2(a).Append(program)
			}
		}

		method = m.method.Build()

		program = method.
			Chain(extract)
	)
	return join(program, api(m.Api), func(x g.Any) []g.Any {
		return singleton(x)
	}).Bimap(
		finalize(m),
		finalize(m),
	).Bimap(
		matcher(method),
		matcher(method),
	)
}

func (m Method) String() string {
	return m.method.String()
}
