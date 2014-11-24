package butler

import g "github.com/SimonRichardson/butler/generic"

type Server struct {
	requests  g.List
	responses g.List
}

func (s Server) Requests() g.List {
	return s.requests
}

func (s Server) Responses() g.List {
	return s.responses
}

func (s Server) concat(a Server) Server {
	return Server{
		requests:  s.requests.Concat(a.requests),
		responses: s.responses.Concat(a.responses),
	}
}

type server struct {
	x g.Either
}

func (s server) AndThen(x service) g.Either {
	return s.x.Chain(func(y g.Any) g.Either {
		a := y.(Server)
		return Compile(x).Bimap(
			g.Constant(a),
			func(y g.Any) g.Any {
				b := y.(Server)
				return a.concat(b)
			},
		)
	})
}

func Compile(x service) g.Either {
	// TODO Make this take multiple services.
	return server{
		x: x.Compile().Map(func(x g.Any) g.Any {
			tup := g.AsTuple2(x)
			return Server{
				requests:  g.List_.Of(tup.Fst()),
				responses: g.List_.Of(tup.Snd()),
			}
		}),
	}
}
