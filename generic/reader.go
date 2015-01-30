package generic

type Reader struct {
	Run func(Any) Any
}

func (r Reader) Chain(f func(Any) Reader) Reader {
	return Reader{
		Run: func(x Any) Any {
			return f(r.Run(x)).Run(x)
		},
	}
}

func (r Reader) Map(f func(Any) Any) Reader {
	return r.Chain(func(x Any) Reader {
		return Reader_.Of(f(x))
	})
}

var (
	Reader_ = reader{}
)

type reader struct{}

func (r reader) Of(a Any) Reader {
	return Reader{
		Run: Constant1(a),
	}
}

func (r reader) Ask() Reader {
	return Reader{
		Run: func(a Any) Any {
			return a
		},
	}
}
