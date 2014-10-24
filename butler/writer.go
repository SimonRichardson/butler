package butler

import "github.com/SimonRichardson/butler/generic"

type Writer struct {
	Run func() Tuple
}

func (w Writer) Of(x generic.Any) Writer {
	return Writer{
		Run: func() Tuple {
			return NewTuple(x, []generic.Any{})
		},
	}
}

func (w Writer) Chain(f func(generic.Any) Writer) Writer {
	return Writer{
		Run: func() Tuple {
			res := w.Run()
			t := f(res._1).Run()
			return NewTuple(t._1, append(res._2, res._1))
		},
	}
}

func (w Writer) Map(f func(generic.Any) generic.Any) Writer {
	return w.Chain(func(x generic.Any) Writer {
		return Writer{
			Run: func() Tuple {
				return NewTuple(f(x), []generic.Any{})
			},
		}
	})
}
