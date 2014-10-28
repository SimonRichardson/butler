package http

import "github.com/SimonRichardson/butler/generic"

func get() func(generic.Any) generic.State {
	return func(x generic.Any) generic.State {
		return generic.State{}.Get()
	}
}

func compose() func(generic.Any) generic.State {
	return func(x generic.Any) generic.State {
		state := generic.State{}
		return state.Modify(func(y generic.Any) generic.Any {
			return x
		})
	}
}
