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
		api = func(x doc.Api) func(g.Either) g.WriterT {
			return func(y g.Either) g.WriterT {
				return g.WriterT_.Lift(x.Run(y)).
					Tell(fmt.Sprintf("Api `%v`", y))
			}
		}
		finalize = func(a String) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.NewTuple2(a, b)
			}
		}
		program = g.WriterT_.Of(s).
			Chain(split).
			Chain(first).
			Chain(validate(s.validator)).
			Chain(flatten)
	)

	return join(program, api(s.Api), func(x g.Any) []g.Any {
		return singleton(x)
	}).Bimap(
		finalize(s),
		finalize(s),
	)
}

func (s String) String() string {
	return s.value
}

// Static methods

var (
	String_ = string_{}
)

type string_ struct{}

func (x string_) Of(v g.Any) String {
	return NewString(v.(string), AnyChar())
}

func (x string_) Empty() String {
	return NewString("", AnyChar())
}
