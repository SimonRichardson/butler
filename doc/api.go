package doc

import "github.com/SimonRichardson/butler/generic"

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

func (a Api) Run(e generic.Either) generic.Either {
	return a.types.Run(e)
}
