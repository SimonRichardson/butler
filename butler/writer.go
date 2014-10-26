package butler

import "github.com/SimonRichardson/butler/generic"

type Writer struct {
	Run func() (generic.Any, []generic.Any)
}

func (w Writer) Of(x generic.Any) Writer {
	return Writer{
		Run: func() (generic.Any, []generic.Any) {
			return x, []generic.Any{}
		},
	}
}

func (w Writer) Chain(f func(generic.Any) Writer) Writer {
	return Writer{
		Run: func() (generic.Any, []generic.Any) {
			a, b := w.Run()
			x, _ := f(a).Run()
			return x, append(b, a)
		},
	}
}

func (w Writer) Map(f func(generic.Any) generic.Any) Writer {
	return w.Chain(func(x generic.Any) Writer {
		return Writer{
			Run: func() (generic.Any, []generic.Any) {
				return f(x), []generic.Any{}
			},
		}
	})
}
