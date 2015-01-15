package http

import (
	"fmt"

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

func (c ContentEncoder) Build() g.WriterT {
	var (
		name = func(a g.Any) g.WriterT {
			var (
				x = AsContentEncoder(a)
				y = x.encoder.String()
			)
			return g.WriterT_.Of(g.NewTuple2(x, y)).
				Tell(fmt.Sprintf("Name `%v`", y))
		}
		keys = func(a g.Any) g.WriterT {
			var (
				run = func(a g.Tuple2) func(g.Any) g.Any {
					return func(b g.Any) g.Any {
						return a.Append(b)
					}
				}
				x = g.AsTuple2(a)
				y = AsContentEncoder(x.Fst())
				z = y.Generate().Bimap(run(x), run(x))
			)
			return g.WriterT_.Lift(z).
				Tell(fmt.Sprintf("Key `%v`", z))
		}
		api = func(x doc.Api) func(g.Either) g.WriterT {
			return func(y g.Either) g.WriterT {
				return g.WriterT_.Lift(x.Run(y)).
					Tell(fmt.Sprintf("Api `%v`", y))
			}
		}
		finalize = func(a ContentEncoder) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.NewTuple2(a, b)
			}
		}
		matcher = func(encoder io.Encoder) func(g.Any) g.Any {
			return func(a g.Any) g.Any {
				var (
					match = func(a g.Any) func(g.Any) g.Any {
						return func(b g.Any) g.Any {
							x := g.NewTuple2(a, a)
							return encoder.Encode(b).
								Bimap(matchPut(x), matchPut(x))
						}
					}
					program = g.StateT_.Of(a).
						Chain(modify(match)).
						Chain(g.Get()).
						Chain(matchFlatten)
				)
				return Result_.FromTuple3(g.AsTuple2(a).Append(program))
			}
		}

		program = g.WriterT_.Of(c).
			Chain(name).
			Chain(keys)
	)
	return join(program, api(c.Api), func(x g.Any) []g.Any {
		return g.AsTuple3(x).Slice()[1:]
	}).Bimap(
		finalize(c),
		finalize(c),
	).Bimap(
		matcher(c.encoder),
		matcher(c.encoder),
	)
}

func (c ContentEncoder) Keys() g.Either {
	return c.encoder.Keys(c.hint())
}

func (c ContentEncoder) Generate() g.Either {
	return c.encoder.Generate(c.hint())
}

func (c ContentEncoder) String() string {
	return c.encoder.String()
}
