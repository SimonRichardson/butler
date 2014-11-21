package butler

import g "github.com/SimonRichardson/butler/generic"

type Server struct {
	requests  g.List
	responses g.List
}

func (s Server) Run(request RemoteRequest) g.Promise {
	return g.Promise_.Of(request)
}

func (s Server) Requests() g.List {
	return s.requests
}

func (s Server) Responses() g.List {
	return s.responses
}

func Compile(x service) g.Either {
	// TODO Make this take multiple services.
	return x.Compile().Map(func(x g.Any) g.Any {
		tup := g.AsTuple2(x)
		return Server{
			requests:  g.AsList(tup.Fst()),
			responses: g.AsList(tup.Snd()),
		}
	})
}
