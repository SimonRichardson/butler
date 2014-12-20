package http

import (
	"reflect"

	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/io"
)

type ContentDecoder struct {
	doc.Api
	decoder io.Decoder
}

func Body(decoder io.Decoder) ContentDecoder {
	return ContentDecoder{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected content decoder `%s` with keys `%s`"),
			doc.NewInlineText("Unexpected content decoder `%s` with keys `%s`"),
		)),
		decoder: decoder,
	}
}

func (c ContentDecoder) Build() g.StateT {
	var (
		always = func(x g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				var (
					decoder = AsContentDecoder(b)
					name    = reflect.TypeOf(decoder.decoder).String()
				)
				return g.NewTuple2(decoder, name)
			}
		}
		values = func(x g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				var (
					tup = g.AsTuple2(b)
					fst = AsContentDecoder(tup.Fst())
				)
				return fst.Keys().Bimap(
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
						tup := g.AsTuple2(a)
						return []g.Any{
							tup.Fst(),
							g.List_.ToSlice(g.AsList(tup.Snd())),
						}
					}
					return api.Run(g.AsEither(a).Bimap(sum, sum))
				}
			}
		}
		finalise = func(c ContentDecoder) func(g.Any) g.StateT {
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

func (c ContentDecoder) Keys() g.Either {
	return c.decoder.Keys()
}
