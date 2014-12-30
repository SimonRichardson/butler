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
			return g.NewTuple2(a, g.Empty{})
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
		/*
			matcher = func(a g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					return g.AsWriter(b).Map(func(x g.Any) g.Any {

						var (
							match = func(a g.Any) func(g.Any) g.Any {
								return func(b g.Any) g.Any {
									var (
										set    = g.AsSet(b)
										header = AsHeader(a)
										key    = header.Name()
										value  = header.Value()
									)
									return set.Get(key).Chain(
										func(x g.Any) g.Option {
											return g.AsOption(g.AsList(x).Find(func(a g.Any) bool {
												return a.(string) == value
											})).Map(func(a g.Any) g.Any {
												return g.NewTuple2(key, x)
											})
										},
									)
								}
							}
							norm = func(a g.Any) g.StateT {
								return g.StateT{
									Run: func(x g.Any) g.Either {
										o := g.AsOption(x)
										return g.AsEither(o.Fold(
											func(x g.Any) g.Any {
												return g.NewRight(g.NewTuple2(g.Empty{}, o))
											},
											func() g.Any {
												return g.NewLeft(g.NewTuple2(g.Empty{}, o))
											},
										))
									},
								}
							}
							combine = func(a g.Any) func(g.Any) g.Any {
								return func(b g.Any) g.Any {
									return g.AsOption(b).Map(func(c g.Any) g.Any {
										return g.NewTuple2(x, c)
									})
								}
							}
							program = g.StateT_.Of(x).
								Chain(modify(match)).
								Chain(norm).
								Chain(modify(combine))
						)

						return g.NewTuple2(b, program)
					})
				}
			}
		*/

		program = h.name.Build().
			Bimap(wrap, wrap).
			Chain(combine(h.value.Build()))
	)

	return join(program, api(h.Api), func(x g.Any) []g.Any {
		return g.AsTuple2(x).Slice()
	})
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
