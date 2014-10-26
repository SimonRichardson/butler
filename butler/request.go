package butler

type request struct{}

func Request(list List) request {
	return request{}
}
