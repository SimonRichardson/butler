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

// Series of predicates, could give more info via a Option or Either
func AnyChar() func(byte) g.Either {
	return func(r byte) g.Either {
		return g.NewRight(r)
	}
}

func HeaderNameChar() func(byte) g.Either {
	return func(r byte) g.Either {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			fallthrough
		case r >= 32 && r <= 39 || r >= 94 && r <= 96:
			fallthrough
		case r == 42 || r == 43 || r == 45 || r == 46 || r == 124:
			return g.NewRight(r)
		}
		return g.NewLeft(r)
	}
}

func HeaderValueChar() func(byte) g.Either {
	return func(r byte) g.Either {
		switch {
		case r >= 32 && r <= 126:
			return g.NewRight(r)
		}
		return g.NewLeft(r)
	}
}

func PathChar() func(byte) g.Either {
	return func(r byte) g.Either {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			fallthrough
		case r == 47 || r == 58:
			return g.NewRight(r)
		}
		return g.NewLeft(r)
	}
}

func UrlChar() func(byte) g.Either {
	return func(r byte) g.Either {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			return g.NewRight(r)
		}
		return g.NewLeft(r)
	}
}

// Build up the state, so it runs this when required.
// 1) Convert string to list of chars
// 2) Convert character to number
// 3) Check number is in range
// 4) Return either (expected/unexpected)
// StateT<Either<Writer<String, Doc>>>
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
