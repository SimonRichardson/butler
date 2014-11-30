package butler

import (
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
)

func concat(a, b Server) Server {
	return Server{
		list: groupBy(a.list.Concat(b.list)),
	}
}

func groupBy(list g.List) g.List {
	var (
		route = func(a g.List) g.Option {
			return a.Find(func(a g.Any) bool {
				_, ok := a.(http.Route)
				return ok
			})
		}
		group = func(a g.Any) g.Any {
			list := g.AsList(g.AsTuple3(a).Snd())
			return route(list).Map(func(a g.Any) g.Any {
				return a.(http.Route).String()
			}).GetOrElse(g.Constant(""))
		}
	)
	return list.GroupBy(group)
}
