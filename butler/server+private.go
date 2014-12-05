package butler

func concat(a, b Server) Server {
	return Server{
		routes:    a.routes.Merge(b.routes),
		routeList: a.routeList.Concat(b.routeList),
	}
}
