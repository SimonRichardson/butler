package butler

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type service struct {
	request  request
	response response
	callable func() g.Any
}

func Service(request, response Builder) service {
	return service{
		request:  ServiceRequest(request.List()),
		response: ServiceResponse(response.List()),
	}
}

func (s service) Then(f func() g.Any) service {
	return service{
		request:  s.request,
		response: s.response,
		callable: f,
	}
}

func (s service) Compile() g.Either {
	return g.Either_.Of(g.Empty{})
}

func (s service) String() string {
	return getRoute(s.request.list).Fold(
		func(x g.Any) g.Any {
			return fmt.Sprintf("Service(`%s`)", x)
		},
		g.Constant("Service()"),
	).(string)
}
