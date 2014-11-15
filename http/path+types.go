package http

func Path(path string) Route {
	return NewRoute(path)
}
