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
	states := r.list.Map(func(x generic.Any) generic.Any {
		return x.(Build).Build()
	})
	// Go from List<WriterT<Either<Http, []Doc>>> -> WriterT<Either<Http, []Doc>>
	// Do this by chaining all the items into one state
	var rec func(List, State) State
	rec = func(states List, state State) State {
        return state.
	}
	return rec(states, State{}.Of(Empty{}))
}
