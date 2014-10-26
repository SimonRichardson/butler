package butler

import "github.com/SimonRichardson/butler/generic"

type response struct {
	list generic.List
}

func Response(list generic.List) response {
	return response{
		list: list,
	}
}

func (r response) Build() generic.Any {
	return r.list.Map(func(x generic.Any) generic.Any {
		return x.(Build).Build()
	})
}
