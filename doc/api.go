package doc

import (
	g "github.com/SimonRichardson/butler/generic"
)

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

func (a Api) Run(e g.Either) g.Either {
	return a.types.Run(e)
}
