package markdown

import (
	"github.com/SimonRichardson/butler/butler"
	g "github.com/SimonRichardson/butler/generic"
)

const (
	DefaultString string = ""

	DefaultMethod string = "GET"
	DefaultPath   string = "/"
)

type mark interface {
	IsBlock() bool
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
				route = templateRoute(butler.AsServer(x).Requests())
				doc   = document(append(templateHeader(), append(route, templateFooter()...)...)...)
			)
			return doc.String()
		},
	)
}

func templateError() []mark {
	return []mark{
		h1(link("Butler", "http://github.com/simonrichardson/butler")),
		h3(str("Serving you content in a monadic style.")),
		hr1(),
		h2(str("ERROR:")),
		p(str("Failed to build the markdown document, errors:")),
	}
}

func templateFailures(x g.List) []mark {
	return asMarks(x.FoldLeft([]mark{}, func(a, b g.Any) g.Any {
		value := g.AsEither(b).Fold(g.Identity(), g.Identity())
		return append(asMarks(a), ul(str(value.(string))))
	}))
}
