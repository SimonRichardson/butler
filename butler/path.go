package butler

type Route struct {
	Api
	path String
}

func NewRoute(path string) Route {
	return Route{
		Api: NewApi(NewDocTypes(
			NewInlineText("Expected route %s"),
			NewInlineText("Unexpected route %s"),
		)),
		path: NewString(path),
	}
}

func Path(path string) Route {
	return NewRoute(path)
}
