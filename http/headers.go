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
	)

	return h.name.Build().
		Chain(g.Get()).
		Chain(g.Merge(h.value.Build())).
		Chain(constant(g.StateT_.Of(h))).
		Chain(modify(api(h.Api)))
}

func (h Header) String() string {
	return fmt.Sprintf("%s: %s", h.name, h.value)
}
