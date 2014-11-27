package generic

type IO struct {
	UnsafePerform func() Any
}

func NewIO(unsafe func() Any) IO {
	return IO{
		UnsafePerform: unsafe,
	}
}

func (x IO) Chain(f func(x Any) IO) IO {
	return NewIO(func() Any {
		io := f(x.UnsafePerform())
		return io.UnsafePerform()
	})
}

func (x IO) Map(f func(x Any) Any) IO {
	return x.Chain(func(x Any) IO {
		return IO{func() Any {
			return f(x)
		}}
	})
}

// Static methods

var (
	IO_ = io{}
)

type io struct{}

func (x io) Of(v Any) IO {
	return NewIO(func() Any {
		return v
	})
}
