package butler

import g "github.com/SimonRichardson/butler/generic"

type service struct {
	request  request
	response response
	callable func(map[string]g.Any) g.Any
}

func Service(request, response Builder) service {
	return service{
		request:  ServiceRequest(request.List()),
		response: ServiceResponse(response.List()),
	}
}

func (s service) Then(f func(map[string]g.Any) g.Any) service {
	return service{
		request:  s.request,
		response: s.response,
		callable: f,
	}
}

func (s service) Request() g.StateT {
	return s.request.Build()
}

func (s service) Response() g.StateT {
	return s.response.Build()
}
