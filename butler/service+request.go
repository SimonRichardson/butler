package butler

import g "github.com/SimonRichardson/butler/generic"

type request struct {
	list g.List
}

func ServiceRequest(list g.List) request {
	return request{
		list: list,
	}
}

func (r request) Build() g.StateT {
	return g.StateT_.Of(g.Empty{})
}
