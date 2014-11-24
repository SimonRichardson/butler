package butler

import g "github.com/SimonRichardson/butler/generic"

type Server struct {
	io g.List
}

func (s Server) IO() g.List {
	return s.io
}

func (s Server) concat(a Server) Server {
	return Server{
		io: s.io.Concat(a.io),
	}
}

type server struct {
	x func() g.Either
}

func (s server) AndThen(x service) server {
	return server{
		x: func() g.Either {
			return s.x().Chain(func(y g.Any) g.Either {
				a := y.(Server)
				return Compile(x).x().Bimap(
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
	return s.x()
}

func Compile(x service) server {
	return server{
		x: func() g.Either {
			return x.Compile().Map(func(x g.Any) g.Any {
				tup := g.AsTuple2(x)
				return Server{
					io: g.List_.Of(tup),
				}
			})
		},
	}
}
