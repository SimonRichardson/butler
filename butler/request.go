package butler

import "github.com/SimonRichardson/butler/generic"

type request struct {
	list generic.List
}

func Request(list generic.List) request {
	return request{
		list: list,
	}
}

func (r request) Build() generic.Any {
	return r.list.Map(func(x generic.Any) generic.Any {
		return x.(Build).Build()
	})
}
