package butler

type Api struct {
	input Any
	types DocTypes
}

func NewApi(input Any, types DocTypes) Api {
	return Api{
		input: input,
		types: types,
	}
}
