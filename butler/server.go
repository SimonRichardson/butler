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
			isNonEmpty = func(a g.List) g.Option {
				return a.Head().Chain(
					func(x g.Any) g.Option {
						return g.AsOption(x)
					},
				)
			}
			overflow = func(matched g.List, parts []string) func(g.Any) g.Option {
				return func(x g.Any) g.Option {
					return g.Option_.FromBool(len(parts) <= matched.Size(), g.Option_.Of(x))
				}
			}
			compute = func(matched g.List) func(g.Any) g.Option {
				return func(x g.Any) g.Option {
					var (
						preRemove = func(x g.Any) bool {
							return g.Option_.ToBool(g.AsOption(x))
						}
						extract = func(x g.Any) g.Any {
							return g.AsOption(x).Map(
								func(y g.Any) g.Any {
									return g.AsTuple2(y).Snd()
								},
							)
						}
						postRemove = func(x g.Any) bool {
							return g.AsOption(x).Fold(
								func(x g.Any) g.Any {
									return g.Option_.ToBool(g.AsList(x).Head())
								},
								g.Constant(false),
							).(bool)
						}
						match = func(x g.Any) g.Any {
							return g.AsOption(x).Chain(
								func(a g.Any) g.Option {
									return g.AsList(a).Find(func(b g.Any) bool {
										// TODO: We should use the compiled service!
										return true
									})
								},
							)
						}

						result = matched.
							Filter(preRemove).
							Map(extract).
							Filter(postRemove).
							Map(match).
							Head()
					)

					return g.AsOption(result)
				}
			}
		)

		return func(w http.ResponseWriter, r *http.Request) {
			var (
				path     = r.URL.Path
				last     = path[len(path)-1]
				redirect = g.Option_.FromBool(string(last) == "/", last)
				parts    = strings.Split(path, "/")
			)

			x := redirect.Fold(
				redirectService(r),
				func() g.Any {
					matched := g.Walker_.Match(s.routes, condition(parts))

					return isNonEmpty(matched).
						Chain(overflow(matched, parts)).
						Chain(compute(matched)).
						GetOrElse(notFoundService(r))
				},
			)

			fmt.Println("Fin", x)
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
					return getRoute(requests).Fold(
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
