package butler

import (
	"fmt"
	"net/http"
	"strings"

	g "github.com/SimonRichardson/butler/generic"
	h "github.com/SimonRichardson/butler/http"
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

func (s Server) Run() g.IO {
	return g.NewIO(func() g.Any {
		var (
			condition = func(parts []string) func(g.List, g.Any, int) bool {
				return func(a g.List, b g.Any, c int) bool {
					return g.AsOption(b).Fold(
						func(a g.Any) g.Any {
							if c < len(parts) {
								var (
									tuple = g.AsTuple2(a)
									node  = h.AsPathNode(tuple.Fst())
								)
								return node.Match(parts[c])
							}
							return false
						},
						g.Constant(c == 0),
					).(bool)
				}
			}
			empty = func(a g.List) bool {
				return a.Head().Fold(
					func(x g.Any) g.Any {
						return g.AsOption(x).Fold(
							g.Constant1(false),
							g.Constant(true),
						).(bool)
					},
					g.Constant(true),
				).(bool)
			}
		)

		return func(w http.ResponseWriter, r *http.Request) {
			var (
				path  = r.URL.Path
				parts = strings.Split(path, "/")

				matched = g.Walker_.Match(s.routes, condition(parts))
			)

			if empty(matched) {
				fmt.Fprintln(w, "404")
			} else {
				fmt.Println(matched)
			}
		}
	})
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
	route := func(a g.List) g.Option {
		return a.Find(func(a g.Any) bool {
			_, ok := a.(h.Route)
			return ok
		})
	}
	return server{
		x: func() g.Either {
			return g.AsEither(s.Compile().Fold(
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
							var (
								route  = y.(h.Route)
								mapped = g.Walker_.Map(route.Route(), func(a g.Any, b int, c bool) g.Any {
									return g.NewTuple2(a, g.List_.FromBool(c, s))
								})
							)
							return g.NewRight(
								Server{
									routes:    mapped,
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
