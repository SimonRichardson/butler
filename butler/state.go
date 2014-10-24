package butler

import "github.com/SimonRichardson/butler/generic"

type State struct {
	Run func(generic.Any) (generic.Any, generic.Any)
}

func (x State) Of(y generic.Any) State {
	return State{
		Run: func(z generic.Any) (generic.Any, generic.Any) {
			return y, z
		},
	}
}

func (x State) Chain(f func(generic.Any) State) State {
	return State{
		Run: func(s generic.Any) (generic.Any, generic.Any) {
			a, b := x.Run(s)
			return f(a).Run(b)
		},
	}
}

func (x State) Map(f func(generic.Any) generic.Any) State {
	return x.Chain(func(y generic.Any) State {
		return x.Of(f(y))
	})
}

// Derived

func (x State) EvalState(y generic.Any) generic.Any {
	a, _ := x.Run(y)
	return a
}

func (x State) ExecState(y generic.Any) generic.Any {
	_, b := x.Run(y)
	return b
}

func (x State) Get() State {
	return State{
		Run: func(z generic.Any) (generic.Any, generic.Any) {
			return z, z
		},
	}
}

func (x State) Modify(f func(generic.Any) generic.Any) State {
	return State{
		Run: func(z generic.Any) (generic.Any, generic.Any) {
			return nil, f(z)
		},
	}
}

func (x State) Put(a generic.Any, b generic.Any) State {
	return State{
		Run: func(z generic.Any) (generic.Any, generic.Any) {
			return a, b
		},
	}
}
