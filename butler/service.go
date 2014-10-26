package butler

import "github.com/SimonRichardson/butler/generic"

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

func (s service) Then(f func(...generic.Any) Result) service {
	return s
}
