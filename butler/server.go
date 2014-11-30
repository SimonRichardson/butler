package butler

import (
	g "github.com/SimonRichardson/butler/generic"
	h "github.com/SimonRichardson/butler/http"
)

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
	//io := g.NewIO(func() g.Any {
	//	return http.NewServeMux()
	//})
	return s.x().Bimap(
		g.Identity(),
		func(a g.Any) g.Any {
			//fmt.Println(a)
			return a
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

func concat(a, b Server) Server {
	return Server{
		list: groupBy(a.list.Concat(b.list)),
	}
}

func groupBy(list g.List) g.List {
	var (
		route = func(a g.List) g.Option {
			return a.Find(func(a g.Any) bool {
				_, ok := a.(h.Route)
				return ok
			})
		}
		group = func(a g.Any) g.Any {
			list := g.AsList(g.AsTuple3(a).Snd())
			return route(list).Map(func(a g.Any) g.Any {
				return a.(h.Route).String()
			}).GetOrElse(g.Constant(""))
		}
	)
	return list.GroupBy(group)
}
