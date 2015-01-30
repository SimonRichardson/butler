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
				b = a.Fst().Fold(
					func(a Any) Any {
						// Should always fail!
						return NewWriterT(NewLeft(a), []Any{})
					},
					func(a Any) Any {
						return f(a)
					},
				)
				c = AsWriterT(b).Run()
			)
			return NewWriterTTuple(c.Fst(), append(a.Snd(), c.Snd()...))
		},
	}
}

func (w WriterT) Map(f func(Any) Any) WriterT {
	return w.Chain(func(a Any) WriterT {
		return WriterT_.Of(f(a))
	})
}

func (w WriterT) Bimap(f func(Any) Any, g func(Any) Any) WriterT {
	return WriterT{
		Run: func() WriterTTuple {
			var (
				a = w.Run()
				b = a.Fst().Fold(
					func(a Any) Any {
						return NewWriterT(NewLeft(f(a)), []Any{})
					},
					func(a Any) Any {
						return NewWriterT(NewRight(f(a)), []Any{})
					},
				)
				c = AsWriterT(b).Run()
			)
			return NewWriterTTuple(c.Fst(), append(a.Snd(), c.Snd()...))
		},
	}
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

func (w writerT) Sequence(list List) WriterT {
	var (
		orgin   = WriterT_.Of(Empty{})
		reduced = list.FoldLeft(orgin, func(x, y Any) Any {
			return AsWriterT(x).Chain(func(z Any) WriterT {
				return AsWriterT(y)
			})
		})
	)
	return AsWriterT(reduced)
}
