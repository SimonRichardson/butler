package http

import (
	"reflect"

	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/io"
)

type ContentEncoder struct {
	doc.Api
	encoder io.Encoder
	hint    func() g.Any
}

func Content(encoder io.Encoder, hint func() g.Any) ContentEncoder {
	return ContentEncoder{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected content encoder `%s` with example output `%s`"),
			doc.NewInlineText("Unexpected content encoder `%s` with example output `%s`"),
		)),
		encoder: encoder,
		hint:    hint,
	}
}

func (c ContentEncoder) Build() g.StateT {
	var (
		always = func(x g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				var (
					encoder = AsContentEncoder(b)
					name    = reflect.TypeOf(encoder.encoder).String()
				)
				return g.NewTuple2(encoder, name)
			}
		}
		values = func(x g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				var (
					tup = g.AsTuple2(b)
					fst = tup.Fst().(ContentEncoder)
				)
				return fst.Generate().Bimap(
					func(x g.Any) g.Any {
						return g.NewTuple2(tup.Snd(), "")
					},
					func(x g.Any) g.Any {
						return g.NewTuple2(tup.Snd(), x)
					},
				)
			}
		}
		api = func(api doc.Api) func(g.Any) func(g.Any) g.Any {
			return func(g.Any) func(g.Any) g.Any {
				return func(a g.Any) g.Any {
					sum := func(a g.Any) g.Any {
						return g.AsTuple2(a).Slice()
					}
					return api.Run(g.AsEither(a).Bimap(sum, sum))
				}
			}
		}
		finalise = func(c ContentEncoder) func(g.Any) g.StateT {
			return func(g.Any) g.StateT {
				return g.StateT{
					Run: func(a g.Any) g.Either {
						cast := func(b g.Any) g.Any {
							x := g.NewWriter(c, singleton(a))
							return g.NewTuple2(g.Empty{}, x)
						}
						return g.AsEither(a).Bimap(cast, cast)
					},
				}
			}
		}
	)
	return g.StateT_.Of(c).
		Chain(modify(g.Constant1)).
		Chain(modify(always)).
		Chain(g.Get()).
		Chain(modify(values)).
		Chain(modify(api(c.Api))).
		Chain(finalise(c))
}

func (c ContentEncoder) Keys() g.Either {
	return c.encoder.Keys(c.hint())
}

func (c ContentEncoder) Generate() g.Either {
	return c.encoder.Generate(c.hint())
}
