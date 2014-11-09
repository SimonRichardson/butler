package generic

type Writer struct {
	Run func() (Any, []Any)
}

func NewWriter(x Any, y []Any) Writer {
	return Writer{
		Run: func() (Any, []Any) {
			return x, y
		},
	}
}

func (w Writer) Chain(f func(Any) Writer) Writer {
	return Writer{
		Run: func() (Any, []Any) {
			a, b := w.Run()
			x, y := f(a).Run()
			return x, append(b, y...)
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
			_, b := w.Run()
			return Empty{}, append(b, x)
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
