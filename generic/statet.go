package generic

type StateT struct {
	Run func(Any) Either
}

func (s StateT) Chain(f func(Any) StateT) StateT {
	return StateT{
		Run: func(a Any) Either {
			return s.Run(a).Chain(func(b Any) Either {
				tuple := b.(Tuple2)
				return f(tuple.Fst()).Run(tuple.Snd())
			})
		},
	}
}

func (s StateT) Map(f func(Any) Any) StateT {
	return s.Chain(func(a Any) StateT {
		return StateT_.Of(f(a))
	})
}

func (s StateT) EvalState(a Any) Any {
	return s.Run(a).Map(func(b Any) Any {
		return b.(Tuple2).Fst()
	})
}

func (s StateT) ExecState(a Any) Any {
	return s.Run(a).Map(func(b Any) Any {
		return b.(Tuple2).Snd()
	})
}

// Static methods

var (
	StateT_ = stateT{}
)

type stateT struct{}

func (s stateT) Lift(e Either) StateT {
	return StateT{
		Run: func(a Any) Either {
			return e.Map(func(b Any) Any {
				return NewTuple2(b, a)
			})
		},
	}
}

func (s stateT) Of(a Any) StateT {
	return StateT{
		Run: func(b Any) Either {
			return Either_.Of(NewTuple2(a, b))
		},
	}
}

func (s stateT) Get() StateT {
	return StateT{
		Run: func(b Any) Either {
			return Either_.Of(NewTuple2(b, b))
		},
	}
}

func (s stateT) Modify(f func(Any) Any) StateT {
	return StateT{
		Run: func(b Any) Either {
			return Either_.Of(NewTuple2(Empty{}, f(b)))
		},
	}
}

func (s stateT) Put(a Any) StateT {
	return s.Modify(func(Any) Any {
		return a
	})
}
