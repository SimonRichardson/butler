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
			doc.NewInlineText("Unexpected header with `%s`"),
		)),
		name:  NewString(name, HeaderNameChar()),
		value: NewString(value, HeaderValueChar()),
	}
}

// Build up the state, so it runs this when required.
// 1) Make sure the name is valid
// 2) Make sure the value is valid
func (h Header) Build() g.StateT {
	var (
		api = func(api doc.Api) func(g.Any) func(g.Any) g.Any {
			return func(a g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					return g.AsWriter(b).Chain(func(a g.Any) g.Writer {
						var (
							t = g.AsTuple2(a)
							x = t.Fst().(String).value
							y = t.Snd().(String).value
						)
						str := g.Either_.Of(append(singleton(x), y))
						return g.NewWriter(h, singleton(api.Run(str)))
					})
				}
			}
		}
		matcher = func(a g.Any) func(g.Any) g.Any {
			return func(b g.Any) g.Any {
				return g.AsWriter(b).Map(func(x g.Any) g.Any {

					var (
						match = func(a g.Any) func(g.Any) g.Any {
							return func(b g.Any) g.Any {
								var (
									set    = g.AsSet(b)
									header = AsHeader(a)
									key    = header.name.String()
									value  = header.value.String()
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
						combine = func(a g.Any) func(g.Any) g.Any {
							return func(b g.Any) g.Any {
								return g.NewTuple2(x, b)
							}
						}
					)

					return g.StateT_.Of(x).
						Chain(modify(match)).
						Chain(modify(combine))
				})
			}
		}
	)

	return h.name.Build().
		Chain(g.Get()).
		Chain(g.Merge(h.value.Build())).
		Chain(constant(g.StateT_.Of(h))).
		Chain(modify(api(h.Api))).
		Chain(modify(matcher))
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
