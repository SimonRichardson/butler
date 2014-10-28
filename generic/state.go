package generic

type State struct {
	Run func(Any) (Any, Any)
}

func (x State) Of(y Any) State {
	return State{
		Run: func(z Any) (Any, Any) {
			return y, z
		},
	}
}

func (x State) Chain(f func(Any) State) State {
	return State{
		Run: func(s Any) (Any, Any) {
			a, b := x.Run(s)
			return f(a).Run(b)
		},
	}
}

func (x State) Map(f func(Any) Any) State {
	return x.Chain(func(y Any) State {
		return x.Of(f(y))
	})
}

// Derived

func (x State) EvalState(y Any) Any {
	a, _ := x.Run(y)
	return a
}

func (x State) ExecState(y Any) Any {
	_, b := x.Run(y)
	return b
}

func (x State) Get() State {
	return State{
		Run: func(z Any) (Any, Any) {
			return z, z
		},
	}
}

func (x State) Modify(f func(Any) Any) State {
	return State{
		Run: func(z Any) (Any, Any) {
			return nil, f(z)
		},
	}
}

func (x State) Put(a Any, b Any) State {
	return State{
		Run: func(z Any) (Any, Any) {
			return a, b
		},
	}
}