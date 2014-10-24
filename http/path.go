package http

import "github.com/SimonRichardson/butler/doc"

type Route struct {
	doc.Api
	path String
}

func NewRoute(path string) Route {
	return Route{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected route %s"),
			doc.NewInlineText("Unexpected route %s"),
		)),
		path: NewString(path),
	}
}

func Path(path string) Route {
	return NewRoute(path)
}
