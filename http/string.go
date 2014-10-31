package http

import (
	"github.com/SimonRichardson/butler/doc"
	"github.com/SimonRichardson/butler/generic"
)

type String struct {
	doc.Api
	value     string
	validator func(byte) generic.Either
}

func NewString(value string, validator func(byte) generic.Either) String {
	return String{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("%q"),
			doc.NewInlineText("%q"),
		)),
		value:     value,
		validator: validator,
	}
}

// Series of predicates, could give more info via a Option or Either
func anyChar() func(byte) generic.Either {
	return func(r byte) generic.Either {
		return generic.NewRight(r)
	}
}

func headerNameChar() func(byte) generic.Either {
	return func(r byte) generic.Either {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			fallthrough
		case r >= 32 && r <= 39 || r >= 94 && r <= 96:
			fallthrough
		case r == 42 || r == 43 || r == 45 || r == 46 || r == 124:
			return generic.NewRight(r)
		}
		return generic.NewLeft(r)
	}
}

func headerValueChar() func(byte) generic.Either {
	return func(r byte) generic.Either {
		switch {
		case r >= 32 && r <= 126:
			return generic.NewRight(r)
		}
		return generic.NewLeft(r)
	}
}

func pathChar() func(byte) generic.Either {
	return func(r byte) generic.Either {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			fallthrough
		case r == 47 || r == 58:
			return generic.NewRight(r)
		}
		return generic.NewLeft(r)
	}
}

func urlChar() func(byte) generic.Either {
	return func(r byte) generic.Either {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			return generic.NewRight(r)
		}
		return generic.NewLeft(r)
	}
}

// Build up the state, so it runs this when required.
// 1) Convert string to list of chars
// 2) Convert character to number
// 3) Check number is in range
// 4) Return either (expected/unexpected)
// State<List<Tuple2<String, []Doc>>>
func (s String) Build() generic.State {
	var (
		extract = func(x generic.Any) func(func(String, generic.List) generic.Tuple2) generic.Tuple2 {
			return func(f func(String, generic.List) generic.Tuple2) generic.Tuple2 {
				tuple := x.(generic.Tuple2)
				str := tuple.Fst().(String)
				list := tuple.Snd().(generic.List)

				return f(str, list)
			}
		}
		split = func(x generic.Any) generic.Any {
			s := x.(String)
			return generic.NewTuple2(s, generic.List_.FromString(s.value))
		}
		run = func(x generic.Any) generic.Any {
			return extract(x)(func(str String, list generic.List) generic.Tuple2 {
				return generic.NewTuple2(
					str,
					list.Map(func(a generic.Any) generic.Any {
						return []byte(a.(string))[0]
					}),
				)
			})
		}
		validate = func(x generic.Any) generic.Any {
			return extract(x)(func(str String, list generic.List) generic.Tuple2 {
				f := str.validator

				return generic.NewTuple2(
					str,
					list.Map(func(a generic.Any) generic.Any {
						return f(a.(byte))
					}),
				)
			})
		}
		fold = func(x generic.Any) generic.Any {
			return extract(x)(func(str String, list generic.List) generic.Tuple2 {
				folded := list.FoldLeft(generic.Either_.Of(""), func(a, b generic.Any) generic.Any {
					return a.(generic.Either).Fold(
						func(x generic.Any) generic.Any {
							return generic.NewLeft(x)
						},
						func(x generic.Any) generic.Any {
							sum := func(y generic.Any) generic.Any {
								aa := y.(byte)
								bb := []byte(x.(string))
								return string(append(bb, aa))
							}
							return b.(generic.Either).Bimap(sum, sum)
						},
					)
				})

				return generic.NewTuple2(str, folded)
			})
		}
		api = func(x generic.Any) generic.Any {
			tuple := x.(generic.Tuple2)
			str := tuple.Fst().(String)

			sum := func(a generic.Any) generic.Any {
				return []generic.Any{a}
			}
			folded := tuple.Snd().(generic.Either).Bimap(sum, sum)

			return generic.NewTuple2(str, str.Api.Run(folded))
		}
	)

	return generic.State_.Of(s).
		Map(split).
		Map(run).
		Map(validate).
		Map(fold).
		Map(api)
}
