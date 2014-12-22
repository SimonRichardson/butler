package generic

type WriterT struct {
	Run func() Tuple2
}

func NewWriterT(a Either, b []Any) WriterT {
	return WriterT{
		Run: func() Tuple2 {
			return NewTuple2(a, b)
		},
	}
}

func (w WriterT) Chain(f func(Any) WriterT) WriterT {
	return WriterT{
		Run: func() Tuple2 {
			var (
				x = w.Run()
				y = AsEither(x.Fst()).Chain(func(a Any) Either {
					return AsEither(f(a).Run().Fst())
				})
			)
			return NewTuple2(y, x.Snd())
		},
	}
}

func (w WriterT) Map(f func(Any) Any) WriterT {
	return w.Chain(func(a Any) WriterT {
		var (
			x = AsTuple2(a)
			y = AsEither(x.Fst()).Map(f)
		)
		return NewWriterT(y, []Any{})
	})
}

// Static methods

var (
	WriterT_ = writerT{}
)

type writerT struct{}

func (w writerT) Lift(e Either) WriterT {
	return NewWriterT(e, []Any{})
}

func (w writerT) Of(a Any) WriterT {
	return NewWriterT(Either_.Of(a), []Any{})
}
