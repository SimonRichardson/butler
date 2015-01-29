package butler

import (
	g "github.com/SimonRichardson/butler/generic"
)

type response struct {
	list g.List
}

func ServiceResponse(list g.List) response {
	return response{
		list: list,
	}
}

func (r response) Build() g.StateT {
	return g.StateT_.Of(g.Empty{})
}
