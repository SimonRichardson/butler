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
	return []mark{
		h4(
			renderHeader(requests),
		),
		ul(
			str("Request"),
			ul(
				append([]mark{str("Headers")}, renderRequestHeaders(requests)...)...,
			),
			ul(
				str("Body"),
			),
		),
		ul(
			str("Response"),
			ul(
				append([]mark{str("Headers")}, renderResponseHeaders(responses)...)...,
			),
			ul(
				append([]mark{str("Body")}, renderResponseContent(responses))...,
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

func empty() g.Any {
	return nothing()
}

func renderHeader(requests g.List) mark {
	var (
		method = getMethod(requests).GetOrElse(g.Constant(DefaultMethod))
		path   = getRoute(requests).GetOrElse(g.Constant(DefaultPath))
	)
	return group(
		str("Route "),
		str(fmt.Sprintf("[%s] ", method)),
		str(fmt.Sprintf("[%s]", path)),
	)
}

func renderRequestHeaders(requests g.List) []mark {
	headers := getHeaders(requests).Map(func(x g.Any) g.Any {
		return ul(inline(str(fmt.Sprintf("`%s`", x.(http.Header).String()))))
	})
	return toMarks(headers)
}

func renderRequestContent(requests g.List) mark {
	return getContent(requests).Chain(func(x g.Any) g.Option {
		var (
			encoder  = x.(http.ContentEncoder)
			generate = encoder.Generate()
			toMark   = func(x g.Any) g.Any {
				return g.AsList(x).Map(func(x g.Any) g.Any {
					return ul०p(str(x.(string)))
				})
			}
			toSlice = func(x g.Any) g.Any {
				return g.List_.ToSlice(g.AsList(x))
			}
			toGroup = func(x g.Any) g.Any {
				var (
					val = x.([]g.Any)
					num = len(val)
					res = make([]mark, num, num)
				)
				for k, v := range val {
					res[k] = v.(mark)
				}
				return group(res...)
			}
		)
		return g.Either_.ToOption(generate).
			Map(toMark).
			Map(toSlice).
			Map(toGroup)

	}).GetOrElse(empty).(mark)
}

func renderResponseHeaders(responses g.List) []mark {
	headers := getHeaders(responses).Map(func(x g.Any) g.Any {
		return ul(inline(str(fmt.Sprintf("`%s`", x.(http.Header).String()))))
	})
	return toMarks(headers)
}

func renderResponseContent(responses g.List) mark {
	return getContent(responses).Chain(func(x g.Any) g.Option {
		var (
			encoder  = x.(http.ContentEncoder)
			generate = encoder.Generate()
			toMark   = func(x g.Any) g.Any {
				return multiline(str(x.(string)))
			}
		)
		return g.Either_.ToOption(generate).
			Map(toMark)
	}).GetOrElse(empty).(mark)
}
