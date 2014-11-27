package butler

import (
	"net/http"

	g "github.com/SimonRichardson/butler/generic"
)

type Server struct {
	list g.List
	io   g.IO
}

func (s Server) IO() g.IO {
	return s.io
}

func (s Server) List() g.List {
	return s.list
}

func (s Server) concat(a Server) Server {
	return Server{
		list: s.list.Concat(a.list),
		io:   s.io,
	}
}

type server struct {
	x func(g.IO) g.Either
}

func (s server) AndThen(x service) server {
	return server{
		x: func(io g.IO) g.Either {
			return s.x(io).Chain(func(y g.Any) g.Either {
				a := y.(Server)
				return Compile(x).x(io).Bimap(
					g.Constant1(a),
					func(y g.Any) g.Any {
						b := y.(Server)
						return a.concat(b)
					},
				)
			})
		},
	}
}

func (s server) Run() g.Either {
	io := g.NewIO(func() g.Any {
		return http.NewServeMux()
	})
	return s.x(io)
}

func Compile(x service) server {
	return server{
		x: func(io g.IO) g.Either {
			return x.Compile(io).Map(func(x g.Any) g.Any {
				tup := g.AsTuple2(x)
				return Server{
					list: g.List_.Of(tup),
					io:   io,
				}
			})
		},
	}
}
