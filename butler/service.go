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
	)

	return exec(s.request).Fold(
		func(x g.Any) g.Any {
			return g.NewLeft(x)
		},
		func(a g.Any) g.Any {
			b := g.AsTuple2(a)
			return exec(s.response).Map(func(x g.Any) g.Any {
				y := g.AsTuple2(x)
				return g.NewTuple2(
					g.AsList(b.Fst()),
					g.AsList(y.Fst()),
				)
			})
		},
	).(g.Either)
}
