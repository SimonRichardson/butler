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
				server    = butler.AsServer(x)
				requests  = server.Requests()
				responses = server.Responses()
				route     = templateRoute(requests, responses)
				doc       = document(append(templateHeader(), append(route, templateFooter()...)...)...)
			)
			return doc.String()
		},
	)
}
