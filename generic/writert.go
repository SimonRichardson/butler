package generic

type WriterT struct {
	Run func() WriterTTuple
}

func NewWriterT(a Either, b []Any) WriterT {
	return WriterT{
		Run: func() WriterTTuple {
			return NewWriterTTuple(a, b)
		},
	}
}

func (w WriterT) Chain(f func(Any) WriterT) WriterT {
	return WriterT{
		Run: func() WriterTTuple {
			var (
				a = w.Run()
				x = a.Fst().Fold(
					func(a Any) Any {
						return WriterT_.Of(a)
					},
					func(a Any) Any {
						return f(a)
					},
				)
				y = AsWriterT(x).Run()
			)
			return NewWriterTTuple(y.Fst(), append(a.Snd(), y.Snd()...))
		},
	}
}

func (w WriterT) Map(f func(Any) Any) WriterT {
	return w.Chain(func(a Any) WriterT {
		return WriterT_.Of(f(a))
	})
}

func (w WriterT) Tell(x Any) WriterT {
	return WriterT{
		Run: func() WriterTTuple {
			var (
				a = w.Run()
				b = a.Snd()
			)
			return NewWriterTTuple(AsEither(a.Fst()), append(b, x))
		},
	}
}

// Static methods

var (
	WriterT_ = writerT{}
)

type writerT struct{}

func (w writerT) Lift(x Either) WriterT {
	return NewWriterT(x, []Any{})
}

func (w writerT) Of(x Any) WriterT {
	return NewWriterT(Either_.Of(x), []Any{})
}

func (w writerT) Tell(x Any) WriterT {
	return NewWriterT(Either_.Of(Empty{}), []Any{x})
}
