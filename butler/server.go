package butler

import g "github.com/SimonRichardson/butler/generic"

type Server struct {
	requests g.List
}

func (s Server) Run(request RemoteRequest) g.Promise {
	return g.Promise_.Of(request)
}

func (s Server) Requests() g.List {
	return s.requests
}

func Compile(x service) g.Either {
	var (
		writer = g.Writer_.Of(g.Empty{})
		run    = func(a g.Any) g.Any {
			x, y := g.AsWriter(a).Run()
			return g.NewTuple2(flatten(g.AsTuple2(x)), y)
		}
		request  = g.AsEither(x.Build().ExecState(writer)).Bimap(run, run)
		requests = g.AsEither(request)
	)
	return requests.Map(func(y g.Any) g.Any {
		z := g.AsTuple2(y)
		return Server{
			requests: g.AsList(z.Fst()),
		}
	})
}

func validateRequests(a g.List) g.List {
	return a
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
