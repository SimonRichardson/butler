package butler

import (
	g "github.com/SimonRichardson/butler/generic"
)

type service struct {
	request  request
	response response
}

func Service(request, response builder) service {
	return service{
		request:  Request(request.list),
		response: Response(response.list),
	}
}

func (s service) Build() g.StateT {
	return s.request.Build().
		Chain(g.Get()).
		Chain(g.Merge(s.response.Build()))
}
