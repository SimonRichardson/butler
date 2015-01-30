package generic

type ReaderT struct {
	Run func(Any) StateT
}

func NewReaderT(state StateT) ReaderT {
	return ReaderT{
		Run: func(Any) StateT {
			return state
		},
	}
}

func (r ReaderT) Chain(f func(Any) ReaderT) ReaderT {
	return ReaderT{
		Run: func(x Any) StateT {
			return r.Run(x).Chain(func(y Any) StateT {
				return f(y).Run(x)
			})
		},
	}
}

// Static methods

var (
	ReaderT_ = readerT{}
)

type readerT struct{}

func (r readerT) Lift(s StateT) ReaderT {
	return NewReaderT(s)
}

func (r readerT) Of(a Any) ReaderT {
	return NewReaderT(StateT_.Of(a))
}

func (r readerT) Ask() ReaderT {
	return ReaderT{
		Run: func(a Any) StateT {
			return StateT_.Of(a)
		},
	}
}

func (r readerT) Sequence(ms List) ReaderT {
	return AsReaderT(ms.FoldLeft(r.Of(List_.Empty()), func(x, y Any) Any {
		return AsReaderT(x).Chain(func(xs Any) ReaderT {
			return AsReaderT(y).Chain(func(xt Any) ReaderT {
				return r.Of(NewCons(x, AsList(xs)))
			})
		})
	}))
}
