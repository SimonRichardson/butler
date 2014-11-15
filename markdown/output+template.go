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

func templateRoute(list g.List) []mark {
	var (
		method  = getMethod(list).GetOrElse(g.Constant(DefaultMethod))
		path    = getRoute(list).GetOrElse(g.Constant(DefaultPath))
		headers = getHeaders(list).Map(func(x g.Any) g.Any {
			return ul(inline(str(x.(http.Header).String())))
		})
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
				append([]mark{str("Headers")}, toMarks(headers)...)...,
			),
			ul(
				str("Body"),
			),
		),
		ul(
			str("Response"),
			ul(
				str("Headers"),
			),
			ul(
				str("Body"),
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
