package http

import (
	"reflect"

	"github.com/SimonRichardson/butler/doc"
	"github.com/SimonRichardson/butler/generic"
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

func (m ContentEncoder) Build() generic.State {
	var (
		extract = func(x generic.Any) func(func(ContentEncoder, output.Encoder) generic.Tuple2) generic.Tuple2 {
			return func(f func(ContentEncoder, output.Encoder) generic.Tuple2) generic.Tuple2 {
				tuple := x.(generic.Tuple2)
				encoder := tuple.Fst().(ContentEncoder)
				output := tuple.Snd().(output.Encoder)

				return f(encoder, output)
			}
		}
		setup = func(x generic.Any) generic.Any {
			return generic.NewTuple2(m, m.encoder)
		}
		validate = func(x generic.Any) generic.Any {
			return extract(x)(func(encoder ContentEncoder, output output.Encoder) generic.Tuple2 {
				return generic.NewTuple2(
					encoder,
					generic.Either_.FromBool(output != nil, output),
				)
			})
		}
		api = func(x generic.Any) generic.Any {
			tuple := x.(generic.Tuple2)
			encoder := tuple.Fst().(ContentEncoder)

			sum := func(a generic.Any) generic.Any {
				name := reflect.ValueOf(a).Type().String()
				return []generic.Any{name}
			}
			folded := tuple.Snd().(generic.Either).Bimap(sum, sum)

			return generic.NewTuple2(encoder, encoder.Api.Run(folded))
		}
	)
	return generic.State_.Of(m).
		Map(setup).
		Map(validate).
		Map(api)
}
