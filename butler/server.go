package butler

import (
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
)

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
	route := func(a g.List) g.Option {
		return a.Find(func(a g.Any) bool {
			_, ok := a.(http.Route)
			return ok
		})
	}
	return server{
		x: func() g.Either {
			return g.AsEither(x.Compile().Fold(
				func(x g.Any) g.Any {
					return g.NewLeft(x)
				},
				func(x g.Any) g.Any {
					// Get the route.
					var (
						tuple     = g.AsTuple3(x)
						requests  = g.AsList(tuple.Snd())
						responses = g.AsList(tuple.Trd())
					)
					return route(requests).Fold(
						func(y g.Any) g.Any {
							route := y.(http.Route)
							return g.NewRight(
								Server{
									routes:    route.Route(),
									routeList: g.List_.Of(g.NewTuple2(requests, responses)),
								},
							)
						},
						g.Constant(g.NewLeft(x)),
					)
				},
			))
		},
	}
}
