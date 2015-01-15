package http

import (
	"fmt"

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

func (c ContentDecoder) Build() g.WriterT {
	var (
		name = func(a g.Any) g.WriterT {
			var (
				x = AsContentDecoder(a)
				y = x.decoder.String()
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
				y = AsContentDecoder(x.Fst())
				z = y.Keys().Bimap(run(x), run(x))
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
		finalize = func(a ContentDecoder) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.NewTuple2(a, b)
			}
		}
		matcher = func(decoder io.Decoder) func(g.Any) g.Any {
			return func(a g.Any) g.Any {
				var (
					match = func(a g.Any) func(g.Any) g.Any {
						return func(b g.Any) g.Any {
							var (
								get = func(x g.Any) g.Either {
									if a, ok := x.([]byte); ok {
										return g.Either_.Of(a)
									}
									if b, ok := x.(string); ok {
										return g.Either_.Of([]byte(b))
									}
									return g.NewLeft(x)
								}
								x = g.NewTuple2(a, a)
							)
							return get(b).Chain(func(z g.Any) g.Either {
								return decoder.Decode(z.([]byte)).
									Bimap(matchPut(x), matchPut(x))
							})
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
		var (
			serialiseTags = func(a g.List) g.Any {
				var (
					x = a.Map(func(x g.Any) g.Any {
						a := g.AsTuple2(x)
						return fmt.Sprintf("%s[%s]", a.Snd(), a.Fst())
					})
					y = x.ReduceLeft(func(x, y g.Any) g.Any {
						return fmt.Sprintf("%s, %s", x, y)
					})
				)
				return fmt.Sprintf("(%s)", y.GetOrElse(g.Constant("")))
			}
			a = g.AsTuple3(x).Slice()[1:]
			b = serialiseTags(g.AsList(a[1]))
		)
		return []g.Any{a[0], b}
	}).Bimap(
		finalize(c),
		finalize(c),
	).Bimap(
		matcher(c.decoder),
		matcher(c.decoder),
	)
}

func (c ContentDecoder) Keys() g.Either {
	return c.decoder.Keys()
}

func (c ContentDecoder) String() string {
	return c.decoder.String()
}
