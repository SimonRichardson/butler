package butler

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

func (s service) Build() {
	s.request.Build()
}
