package butler

import g "github.com/SimonRichardson/butler/generic"

type Server struct {
	routes    g.Tree
	routeList g.List
}

func (s Server) Routes() g.Tree {
	return s.routes
}

func (s Server) RouteList() g.List {
	return s.routeList
}

func (s Server) Run() g.IO {
	return g.IO_.Of(g.Empty{})
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

func Compile(s service) server {
	return server{
		x: func() g.Either {
			return g.Either_.Of(g.Empty{})
		},
	}
}
