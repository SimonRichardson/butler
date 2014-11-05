package generic

type Writer struct {
	Run func() (Any, []Any)
}

func (w Writer) Chain(f func(Any) Writer) Writer {
	return Writer{
		Run: func() (Any, []Any) {
			a, b := w.Run()
			x, _ := f(a).Run()
			return x, append(b, a)
		},
	}
}

func (w Writer) Map(f func(Any) Any) Writer {
	return w.Chain(func(x Any) Writer {
		return Writer{
			Run: func() (Any, []Any) {
				return f(x), []Any{}
			},
		}
	})
}

func (w Writer) Tell(x Any) Writer {
	return Writer{
		Run: func() (Any, []Any) {
			_, y := w.Run()
			return Empty{}, append(y, x)
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
		Run: func() (Any, []Any) {
			return x, []Any{}
		},
	}
}
