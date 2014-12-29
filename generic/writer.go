package generic

type Writer struct {
	Run func() WriterTuple
}

func NewWriter(x Any, y []Any) Writer {
	return Writer{
		Run: func() WriterTuple {
			return NewWriterTuple(x, y)
		},
	}
}

func (w Writer) Chain(f func(Any) Writer) Writer {
	return Writer{
		Run: func() WriterTuple {
			var (
				exe0 = w.Run()
				a    = exe0.Fst()
				b    = exe0.Snd()

				exe1 = f(a).Run()
				x    = exe1.Fst()
				y    = exe1.Snd()
			)
			return NewWriterTuple(x, append(b, y...))
		},
	}
}

func (w Writer) Map(f func(Any) Any) Writer {
	return w.Chain(func(x Any) Writer {
		return Writer{
			Run: func() WriterTuple {
				return NewWriterTuple(f(x), []Any{x})
			},
		}
	})
}

func (w Writer) Tell(x Any) Writer {
	return Writer{
		Run: func() WriterTuple {
			b := w.Run().Snd()
			return NewWriterTuple(Empty{}, append(b, x))
		},
	}
}

// Static methods

var (
	Writer_ = writer{}
)

type writer struct{}

func (w writer) Of(x Any) Writer {
	return Writer{
		Run: func() WriterTuple {
			return NewWriterTuple(x, []Any{})
		},
	}
}
