package butler

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
	return NewMethod(MDelete)
}

func Get() Api {
	return NewMethod(MGet)
}

func Head() Api {
	return NewMethod(MHead)
}

func Options() Api {
	return NewMethod(MOptions)
}

func Patch() Api {
	return NewMethod(MPatch)
}

func Post() Api {
	return NewMethod(MPost)
}

func Put() Api {
	return NewMethod(MPut)
}

func Trace() Api {
	return NewMethod(MTrace)
}
