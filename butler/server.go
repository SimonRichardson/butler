package butler

import (
	"fmt"
	"html"
	"net/http"

	g "github.com/SimonRichardson/butler/generic"
)

type Server struct {
	list g.List
}

func (s Server) List() g.List {
	return s.list
}

type ServerWithIO struct {
	list g.List
	io   g.IO
}

func (s ServerWithIO) IO() g.IO {
	return s.io
}

func (s ServerWithIO) List() g.List {
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
	return s.x().Bimap(
		g.Identity(),
		func(a g.Any) g.Any {
			var (
				io = g.NewIO(func() g.Any {
					return http.NewServeMux()
				})
				list   = AsServer(a).list
				mutate = func(a, b g.Any) g.Any {
					var (
						io    = g.AsIO(a)
						tuple = g.AsTuple2(b)
						route = tuple.Fst().(string)
						// list  = g.AsList(tuple.Snd())
					)
					return io.Map(func(a g.Any) g.Any {
						var (
							mux = a.(*http.ServeMux)
						)
						mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
							fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
						})
						return mux
					})
					/*
						// Create the patterns required here.

							return list.FoldLeft(io, func(x, y g.Any) g.Any {
								var (
									io      = g.AsIO(x)
									tuple   = g.AsTuple3(y)
									service = tuple.Fst().(service)
								)
								return x
							})
					*/
				}
			)
			return ServerWithIO{
				list: list,
				io:   g.AsIO(list.FoldLeft(io, mutate)),
			}
		},
	)
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
