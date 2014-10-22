package butler

type MethodType string

const (
	DELETE  MethodType = "DELETE"
	GET     MethodType = "GET"
	HEAD    MethodType = "HEAD"
	OPTIONS MethodType = "OPTIONS"
	PATCH   MethodType = "PATCH"
	POST    MethodType = "POST"
	PUT     MethodType = "PUT"
	TRACE   MethodType = "TRACE"
)

type Method struct {
	method MethodType
}

func NewMethod(method MethodType) Api {
	return NewApi(
		Method{
			method: method,
		},
		NewDocTypes(
			NewInlineText("Expected method %s"),
			NewInlineText("Unexpected method %s"),
		),
	)
}

func Delete() Api {
	return NewMethod(DELETE)
}

func Get() Api {
	return NewMethod(GET)
}

func Head() Api {
	return NewMethod(HEAD)
}

func Options() Api {
	return NewMethod(OPTIONS)
}

func Patch() Api {
	return NewMethod(PATCH)
}

func Post() Api {
	return NewMethod(POST)
}

func Put() Api {
	return NewMethod(PUT)
}

func Trace() Api {
	return NewMethod(TRACE)
}
