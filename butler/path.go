package butler

type Route struct {
	path Api
}

func NewRoute(path string) Api {
	return NewApi(
		Route{
			path: NewString(path),
		},
		NewDocTypes(
			NewInlineText("Expected route %s"),
			NewInlineText("Unexpected route %s"),
		),
	)
}

func Path(path string) Api {
	return NewRoute(path)
}
