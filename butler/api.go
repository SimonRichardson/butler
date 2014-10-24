package butler

type Api struct {
	types DocTypes
}

func NewApi(types DocTypes) Api {
	return Api{
		types: types,
	}
}

func (a Api) Doc() DocTypes {
	return a.types
}
