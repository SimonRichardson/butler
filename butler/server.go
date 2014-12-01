package butler

import g "github.com/SimonRichardson/butler/generic"

type Server struct {
	list g.List
}

func (s Server) List() g.List {
	return s.list
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
						return concat(a, b)
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
				return Server{
					list: groupBy(g.List_.Of(x)),
				}
			})
		},
	}
}
