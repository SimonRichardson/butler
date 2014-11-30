package markdown

import (
	"github.com/SimonRichardson/butler/butler"
	g "github.com/SimonRichardson/butler/generic"
)

const (
	DefaultString string = ""

	DefaultMethod      string = "GET"
	DefaultPath        string = "/"
	DefaultContentType string = ""
)

type mark interface {
	IsBlockStart() bool
	IsBlockFinish() bool
	Children() g.Option
	String() string
}

func Output(server g.Either) g.Either {
	// Build the service and output it as markdown!
	return server.Bimap(
		func(x g.Any) g.Any {
			var (
				list     = g.AsTuple2(x).Snd()
				failures = templateFailures(g.List_.To(list.([]g.Any)...))
				doc      = document(append(templateError(), append(failures, templateFooter()...)...)...)
			)
			return doc.String()
		},
		func(x g.Any) g.Any {
			var (
				server = butler.AsServer(x)
				list   = server.List()
				folded = list.FoldLeft([]mark{}, func(a, b g.Any) g.Any {
					var (
						tuple  = g.AsTuple2(b)
						values = g.AsList(tuple.Snd())

						folded = values.FoldLeft([]mark{}, func(a, b g.Any) g.Any {
							var (
								tuple     = g.AsTuple3(b)
								requests  = g.AsList(tuple.Snd())
								responses = g.AsList(tuple.Trd())
							)

							return append(asMarks(a), templateRoute(requests, responses)...)
						})
					)

					return append(asMarks(a), asMarks(folded)...)
				})
				doc = document(append(templateHeader(), append(asMarks(folded), templateFooter()...)...)...)
			)
			return doc.String()
		},
	)
}
