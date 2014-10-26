package http

import (
	"github.com/SimonRichardson/butler/doc"
	"github.com/SimonRichardson/butler/generic"
)

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

func (r Route) Build() generic.State {
	return generic.State{}
}

func Path(path string) Route {
	return NewRoute(path)
}
