package generic

type WriterT struct {
	Run func() WriterTuple
}

func NewWriterT(a Either, b []Any) WriterT {
	return WriterT{
		Run: func() WriterTuple {
			return NewWriterTuple(a, b)
		},
	}
}

func (w WriterT) Chain(f func(Any) WriterT) WriterT {
	return WriterT{
		Run: func() WriterTuple {
			var (
				x = w.Run()
				y = AsEither(x.Fst()).Chain(func(a Any) Either {
					return AsEither(f(a).Run().Fst())
				})
			)
			return NewWriterTuple(y, x.Snd())
		},
	}
}

func (w WriterT) Map(f func(Any) Any) WriterT {
	return w.Chain(func(a Any) WriterT {
		var (
			x = AsWriterTuple(a)
			y = AsEither(x.Fst()).Map(f)
		)
		return NewWriterT(y, []Any{})
	})
}

func (w WriterT) Tell(x Any) WriterT {
	return WriterT{
		Run: func() WriterTuple {
			b := w.Run().Snd()
			return NewWriterTuple(Either_.Of(Empty{}), append(b, x))
		},
	}
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
