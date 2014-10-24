package http

import "github.com/SimonRichardson/butler/doc"

type MethodType string

const (
	MDelete  MethodType = "delete"
	MGet     MethodType = "get"
	MHead    MethodType = "head"
	MOptions MethodType = "options"
	MPatch   MethodType = "patch"
	MPost    MethodType = "post"
	MPut     MethodType = "put"
	MTrace   MethodType = "trace"
)

type Method struct {
	doc.Api
	method MethodType
}

func NewMethod(method MethodType) Method {
	return Method{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected method %s"),
			doc.NewInlineText("Unexpected method %s"),
		)),
		method: method,
	}
}

func Delete() Method {
	return NewMethod(MDelete)
}

func Get() Method {
	return NewMethod(MGet)
}

func Head() Method {
	return NewMethod(MHead)
}

func Options() Method {
	return NewMethod(MOptions)
}

func Patch() Method {
	return NewMethod(MPatch)
}

func Post() Method {
	return NewMethod(MPost)
}

func Put() Method {
	return NewMethod(MPut)
}

func Trace() Method {
	return NewMethod(MTrace)
}
