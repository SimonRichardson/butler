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
	empty := g.NewTuple2(g.List_.Empty(), g.List_.Empty())
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
				io     = server.IO()

				// get the first one for now!
				tuple     = g.AsTuple2(io.Head().GetOrElse(g.Constant(empty)))
				requests  = g.AsList(tuple.Fst())
				responses = g.AsList(tuple.Snd())
				route     = templateRoute(requests, responses)

				// Concat all the routes together!
				doc = document(append(templateHeader(), append(route, templateFooter()...)...)...)
			)
			return doc.String()
		},
	)
}
