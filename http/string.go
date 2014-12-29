package http

import (
	"fmt"

	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
)

type String struct {
	doc.Api
	value     string
	validator func(byte) g.Either
}

func NewString(value string, validator func(byte) g.Either) String {
	return String{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected string `%s`"),
			doc.NewInlineText("Unexpected string `%s`"),
		)),
		value:     value,
		validator: validator,
	}
}

func (s String) Build() g.WriterT {
	var (
		split = func(x g.Any) g.WriterT {
			var (
				a = AsString(x).String()
				b = g.List_.FromString(a)
			)
			return g.WriterT_.Of(b).
				Tell(fmt.Sprintf("Split `%v`", x))
		}
		first = func(x g.Any) g.WriterT {
			var (
				a = g.AsList(x).Map(func(y g.Any) g.Any {
					return []byte(y.(string))[0]
				})
			)
			return g.WriterT_.Of(a).
				Tell(fmt.Sprintf("Select first %v", x))
		}
		validate = func(f func(byte) g.Either) func(g.Any) g.WriterT {
			return func(x g.Any) g.WriterT {
				var (
					a = g.AsList(x).Map(func(a g.Any) g.Any {
						return f(a.(byte))
					})
				)
				return g.WriterT_.Of(a).
					Tell(fmt.Sprintf("Validate each %v", x))
			}
		}
		flatten = func(x g.Any) g.WriterT {
			a := g.AsEither(g.AsList(x).FoldLeft(g.Either_.Of(""), func(a, b g.Any) g.Any {
				return g.AsEither(a).Fold(
					func(x g.Any) g.Any {
						return g.NewLeft(x)
					},
					func(x g.Any) g.Any {
						sum := func(y g.Any) g.Any {
							var (
								aa = y.(byte)
								bb = []byte(x.(string))
							)
							return string(append(bb, aa))
						}
						return g.AsEither(b).Bimap(sum, sum)
					},
				)
			}))
			return g.WriterT_.Lift(a).
				Tell(fmt.Sprintf("Flatten %v", x))
		}
		/*
			api = func(x doc.Api) func(g.Any) g.WriterT {
				return func(y g.Any) g.WriterT {
					var (
						a = x.Run(g.Either_.Of(singleton(y)))
					)
					return g.WriterT_.Lift(a).
						Tell(fmt.Sprintf("Api `%v`", y))
				}
			}

				finalise = func(s String) func(g.Any) g.StateT {
					return func(g.Any) g.StateT {
						return g.StateT{
							Run: func(a g.Any) g.Either {
								cast := func(b g.Any) g.Any {
									x := g.NewWriter(s, singleton(a))
									return g.NewTuple2(g.Empty{}, x)
								}
								return g.AsEither(a).Bimap(cast, cast)
							},
						}
					}
				}*/
		program = g.WriterT_.Of(s).
			Chain(split).
			Chain(first).
			Chain(validate(s.validator)).
			Chain(flatten)
	)

	return program

	//Map(finalise(s))
}

func (s String) String() string {
	return s.value
}
