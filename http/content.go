package http

import (
	"reflect"

	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/output"
)

type ContentEncoder struct {
	doc.Api
	encoder output.Encoder
}

func Content(encoder output.Encoder) ContentEncoder {
	return ContentEncoder{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected content encoder `%s`"),
			doc.NewInlineText("Unexpected content encoder `%s`"),
		)),
		encoder: encoder,
	}
}

func (c ContentEncoder) Build() g.StateT {
	var (
		always = func(x g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				var (
					encoder = asContentEncoder(b).encoder
					name    = reflect.TypeOf(encoder).String()
				)
				return g.Either_.Of(name)
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
		Chain(modify(api(c.Api))).
		Chain(finalise(c))
}
