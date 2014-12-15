package butler

import (
	"fmt"
	"net/http"
	"strings"

	g "github.com/SimonRichardson/butler/generic"
	h "github.com/SimonRichardson/butler/http"
)

var (
	notFoundService = func(r *http.Request) func() g.Any {
		return func() g.Any {
			// We build the not found service at run time.
			var (
				request  = Request()
				response = Response().ContentType(r.Header.Get("content-type"))
			)
			return Service(request, response).Then(func() g.Any {
				return error404()
			})
		}
	}
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
		)

		return func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered:", r)
				}
			}()

			var (
				path  = r.URL.Path
				parts = strings.Split(path, "/")

				matched = g.Walker_.Match(s.routes, condition(parts))
			)

			x := isNonEmpty(matched).Map(func(x g.Any) g.Any {
				// TODO (Simon) :
				// 0. Move http.Request over to a better object.
				// 1. Now go through the service, matches all the values
				// 2. Render service if matched
				// 3. No match = 404
				var (
					extract = func(x g.Any) g.Any {
						return g.AsOption(x).Map(
							func(y g.Any) g.Any {
								return g.AsTuple2(y).Snd()
							},
						)
					}
					remove = func(x g.Any) bool {
						return g.Option_.ToBool(g.AsOption(x))
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
						Filter(remove).
						Map(extract).
						Filter(remove).
						Map(match).
						Head()
				)

				return g.AsOption(result)
			}).GetOrElse(notFoundService(r))

			fmt.Println(x)
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
