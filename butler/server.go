package butler

import (
	g "github.com/SimonRichardson/butler/generic"
)

type Server struct {
	requests g.List
}

func (s Server) Run(request RemoteRequest) g.Promise {
	return g.Promise_.Of(request)
}

func (s Server) Requests() g.List {
	return s.requests
}

func Compile(x service) Server {
	var (
		writer = g.Writer_.Of(g.Empty{})
		run    = func(a g.Any) g.Any {
			x, y := g.AsWriter(a).Run()
			return g.NewTuple2(flatten(g.AsTuple2(x)), y)
		}
		request  = g.AsEither(x.Build().ExecState(writer)).Fold(run, run)
		requests = g.AsTuple2(request)
	)
	return Server{
		requests: g.AsList(requests.Fst()),
	}
}

func flatten(a g.Tuple2) g.List {
	var rec func(l g.List, t g.Tuple2) g.List
	rec = func(l g.List, t g.Tuple2) g.List {
		if b, ok := t.Fst().(g.Tuple2); ok {
			return rec(
				g.NewCons(t.Snd(), l),
				b,
			)
		} else {
			return l
		}
	}
	return rec(g.List_.Empty(), a)
}
