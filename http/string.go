package http

import (
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

func (s String) Build() g.StateT {
	var (
		split = func(g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.List_.FromString(b.(String).value)
			}
		}
		first = func(a g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.AsList(b).Map(func(a g.Any) g.Any {
					return []byte(a.(string))[0]
				})
			}
		}
		validate = func(f func(byte) g.Either) func(g.Any) func(g.Any) g.Any {
			return func(x g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					return g.AsList(b).Map(func(a g.Any) g.Any {
						return f(a.(byte))
					})
				}
			}
		}
		fold = func(a g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.AsList(b).FoldLeft(g.Either_.Of(""), func(a, b g.Any) g.Any {
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
				})
			}
		}
		api = func(api doc.Api) func(g.Any) func(g.Any) g.Any {
			return func(g.Any) func(g.Any) g.Any {
				return func(a g.Any) g.Any {
					sum := func(a g.Any) g.Any {
						return singleton(a)
					}
					return api.Run(g.AsEither(a).Bimap(sum, sum))
				}
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
		}
	)

	return g.StateT_.Of(s).
		Chain(modify(g.Constant1)).
		Chain(modify(split)).
		Chain(modify(first)).
		Chain(modify(validate(s.validator))).
		Chain(modify(fold)).
		Chain(modify(api(s.Api))).
		Chain(finalise(s))
}

func (s String) String() string {
	return s.value
}
