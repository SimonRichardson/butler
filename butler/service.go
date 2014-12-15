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
	var (
		writer = g.Writer_.Of(g.Empty{})
		run    = func(a g.Any) g.Any {
			x, y := g.AsWriter(a).Run()
			return g.NewTuple2(flatten(g.AsTuple2(x)), y)
		}
		exec = func(b Build) g.Either {
			return g.AsEither(b.Build().ExecState(writer)).Bimap(run, run)
		}
		collate = func(a g.Tuple2) func(g.Any) g.Any {
			return func(x g.Any) g.Any {
				var (
					b         = g.AsTuple2(x)
					requests  = g.AsList(a.Fst())
					responses = g.AsList(b.Fst())
				)
				return g.NewTuple3(
					s,
					requests,
					responses,
				)
			}
		}
	)

	return exec(s.request).Fold(
		func(x g.Any) g.Any {
			return g.NewLeft(x)
		},
		func(a g.Any) g.Any {
			b := g.AsTuple2(a)
			return exec(s.response).
				Map(collate(b))
		},
	).(g.Either)
}

func (s service) String() string {
	return getRoute(s.request.list).Fold(
		func(x g.Any) g.Any {
			return fmt.Sprintf("Service(`%s`)", x)
		},
		g.Constant("Service()"),
	).(string)
}
