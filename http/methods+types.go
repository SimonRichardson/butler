package http

func Delete() Method {
	return NewMethod(DELETE)
}

func Get() Method {
	return NewMethod(GET)
}

func Head() Method {
	return NewMethod(HEAD)
}

func Options() Method {
	return NewMethod(OPTIONS)
}

func Patch() Method {
	return NewMethod(PATCH)
}

func Post() Method {
	return NewMethod(POST)
}

func Put() Method {
	return NewMethod(PUT)
}

func Trace() Method {
	return NewMethod(TRACE)
}
