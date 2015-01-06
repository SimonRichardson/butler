package http

import (
	"fmt"

	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
)

type Header struct {
	doc.Api
	name  String
	value String
}

func NewHeader(name, value string) Header {
	return Header{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected header `%s` with value `%s`"),
			doc.NewInlineText("Unexpected header with `%s` with value `%s`"),
		)),
		name:  NewString(name, HeaderNameChar()),
		value: NewString(value, HeaderValueChar()),
	}
}

func (h Header) Build() g.WriterT {
	var (
		wrap = func(a g.Any) g.Any {
			b := String_.Empty()
			return g.NewTuple2(a, g.NewTuple2(b, b))
		}
		combine = func(x g.WriterT) func(g.Any) g.WriterT {
			return func(y g.Any) g.WriterT {
				var (
					a   = x.Run()
					b   = a.Fst()
					c   = a.Snd()
					d   = g.AsTuple2(y)
					run = func(f func(g.Any) g.Either) func(g.Any) g.Any {
						return func(a g.Any) g.Any {
							x := d.MapSnd(func(b g.Any) g.Any {
								return a
							})
							return g.NewWriterT(f(x), c).
								Tell(fmt.Sprintf("Combine %v, %v", y, a))
						}
					}
				)
				return g.AsWriterT(b.Fold(
					run(func(x g.Any) g.Either {
						return g.NewLeft(x)
					}),
					run(func(x g.Any) g.Either {
						return g.NewRight(x)
					}),
				))
			}
		}
		api = func(x doc.Api) func(g.Either) g.WriterT {
			return func(y g.Either) g.WriterT {
				return g.WriterT_.Lift(x.Run(y)).
					Tell(fmt.Sprintf("Api `%v`", y))
			}
		}
		finalize = func(a Header) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.NewTuple2(a, b)
			}
		}
		matcher = func(name, value g.WriterT) func(g.Any) g.Any {
			return func(a g.Any) g.Any {
				var (
					match = func(a g.Any) func(g.Any) g.Any {
						return func(b g.Any) g.Any {
							var (
								x     = name.Run().Fst()
								y     = g.AsTuple2(b).Snd()
								parts = y.([]string)
							)
							return x.Chain(func(x g.Any) g.Either {
								var (
									c = g.AsTuple3(x).Trd()
									d = g.Either_.FromBool(len(parts) > 0, c)
								)
								return d.Chain(func(a g.Any) g.Either {
									s := g.AsStateT(a)
									return s.EvalState(parts[0]).Chain(func(a g.Any) g.Either {
										return value.Run().Fst().Chain(func(x g.Any) g.Either {
											var (
												t = g.AsTuple3(x).Trd()
												s = g.AsStateT(t)
											)
											return s.EvalState(parts[1])
										})
									})
								})
							}).Bimap(matchPut(a), matchPut(a))
						}
					}
					program = g.StateT_.Of(a).
						Chain(modify(matchSplit(":"))).
						Chain(g.Get()).
						Chain(modify(match)).
						Chain(g.Get()).
						Chain(matchFlatten)
				)
				return g.AsTuple2(a).Append(program)
			}
		}

		name  = h.name.Build()
		value = h.value.Build()

		program = name.
			Bimap(wrap, wrap).
			Chain(combine(value))
	)

	return join(program, api(h.Api), func(x g.Any) []g.Any {
		var (
			unwrap = func(a g.Any) g.Any {
				var (
					x = g.AsTuple3(a).Fst()
					y = AsString(x)
				)
				return y.String()
			}
		)
		return g.AsTuple2(x).
			Bimap(unwrap, unwrap).
			Slice()
	}).Bimap(
		finalize(h),
		finalize(h),
	).Bimap(
		matcher(name, value),
		matcher(name, value),
	)
}

func (h Header) Name() string {
	return h.name.value
}

func (h Header) Value() string {
	return h.value.value
}

func (h Header) String() string {
	return fmt.Sprintf("%s: %s", h.name, h.value)
}
