package generic

type Free interface {
	Chain(f func(Any) Free) Free
	Run() Any
}

type Return struct {
	val Any
}

func NewReturn(x Any) Return {
	return Return{
		val: x,
	}
}

func (r Return) Chain(f func(Any) Free) Free {
	return f(r.val)
}

func (r Return) Run() Any {
	return r.val
}

type Suspend struct {
	functor Functor
}

func NewSuspend(f Functor) Suspend {
	return Suspend{
		functor: f,
	}
}

func (s Suspend) Chain(f func(Any) Free) Free {
	return Suspend{
		functor: s.functor.Map(func(x Any) Any {
			return AsFree(x).Chain(f)
		}),
	}
}

func (s Suspend) Run() Any {
	var rec func(x Free) Result
	rec = func(x Free) Result {
		if r, ok := x.(Return); ok {
			return Done(r.val)
		}
		return Cont(func() Result {
			s := x.(Suspend)
			return rec(s.functor.Run())
		})
	}
	return Trampoline(rec(s))
}

var (
	Free_ = free{}
)

type free struct{}

func (f free) Lift(x Functor) Free {
	return NewSuspend(x.Map(Free_.Return))
}

func (f free) Return(x Any) Any {
	return NewReturn(x)
}

func (f free) Suspend(x Any) Any {
	return f.Lift(x.(Functor))
}

var (
	Functor_ = functor{}
)

type Functor interface {
	Map(f func(Any) Any) Functor
	Run() Free
}

type functor struct{}

type eitherF struct {
	val Either
}

func (x eitherF) Map(f func(Any) Any) Functor {
	return eitherF{
		val: x.val.Map(f),
	}
}

func (x eitherF) Run() Free {
	return NewReturn(x.val)
}

func (f functor) Either(x Either) Functor {
	return eitherF{
		val: x,
	}
}
