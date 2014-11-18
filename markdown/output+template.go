package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
)

func templateHeader() []mark {
	return []mark{
		h1(link("Butler", "http://github.com/simonrichardson/butler")),
		h3(str("Serving you content in a monadic style.")),
		hr1(),
		ul(
			link("Route definitions", "#routes"),
		),
		h4(str("Routes")),
		p(str("The route definitions for your service are as follows:")),
	}
}

func templateFooter() []mark {
	return []mark{
		hr2(),
		center(link("Served by Butler", "http://github.com/simonrichardson/butler")),
	}
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

func templateRoute(requests, responses g.List) []mark {
	var (
		method     = getMethod(requests).GetOrElse(g.Constant(DefaultMethod))
		path       = getRoute(requests).GetOrElse(g.Constant(DefaultPath))
		reqHeaders = getHeaders(requests).Map(func(x g.Any) g.Any {
			return ul(inline(str(x.(http.Header).String())))
		})
		resHeaders = getHeaders(responses).Map(func(x g.Any) g.Any {
			return ul(inline(str(x.(http.Header).String())))
		})
		content = getContent(responses).Map(func(x g.Any) g.Any {
			return str("Content")
		}).GetOrElse(g.Constant(nothing())).(mark)
	)
	return []mark{
		h4(
			group(
				str("Route "),
				str(fmt.Sprintf("[%s] ", method)),
				str(fmt.Sprintf("[%s]", path)),
			),
		),
		ul(
			str("Request"),
			ul(
				append([]mark{str("Headers")}, toMarks(reqHeaders)...)...,
			),
			ul(
				str("Body"),
			),
		),
		ul(
			str("Response"),
			ul(
				append([]mark{str("Headers")}, toMarks(resHeaders)...)...,
			),
			ul(
				append([]mark{str("Body")}, content)...,
			),
		),
	}
}

func getMethod(x g.List) g.Option {
	return x.Find(func(a g.Any) bool {
		_, ok := a.(http.Method)
		return ok
	})
}

func getRoute(x g.List) g.Option {
	return x.Find(func(a g.Any) bool {
		_, ok := a.(http.Route)
		return ok
	})
}

func getHeaders(x g.List) g.List {
	return x.Filter(func(a g.Any) bool {
		_, ok := a.(http.Header)
		return ok
	})
}

func getContent(x g.List) g.Option {
	return x.Find(func(a g.Any) bool {
		_, ok := a.(http.ContentEncoder)
		return ok
	})
}
