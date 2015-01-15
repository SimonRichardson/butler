package http

import g "github.com/SimonRichardson/butler/generic"

type Result interface {
	Builder() g.Any
	Api() string
	Matcher() g.StateT
}

type result struct {
	build g.Any
	api   string
	match g.StateT
}

func NewResult(build g.Any, api string, match g.StateT) Result {
	return result{
		build: build,
		api:   api,
		match: match,
	}
}

func (r result) Builder() g.Any {
	return r.build
}

func (r result) Api() string {
	return r.api
}

func (r result) Matcher() g.StateT {
	return r.match
}

var (
	Result_ = result_{}
)

type result_ struct{}

func (r result_) FromTuple3(x g.Tuple3) Result {
	var (
		build = x.Fst()
		api   = x.Snd().(string)
		match = g.AsStateT(x.Trd())
	)
	return NewResult(build, api, match)
}
