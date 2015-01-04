package generic

import "fmt"

type Free interface {
	Chain(f func(Any) Free) Free
	Map(f func(Any) Any) Free

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

func (r Return) Map(f func(Any) Any) Free {
	return r.Chain(func(x Any) Free {
		return Free_.Of(f(x))
	})
}

func (r Return) Run() Any {
	return r.val
}

func (r Return) String() string {
	return fmt.Sprintf("Return(%v)", r.val)
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

func (s Suspend) Map(f func(Any) Any) Free {
	return s.Chain(func(x Any) Free {
		return Free_.Of(f(x))
	})
}

func (s Suspend) Run() Any {
	var x Free = s
	for {
		if _, ok := x.(Return); ok {
			break
		}
		x = x.(Suspend).functor.Run()
	}
	return x.(Return).Run()
}

func (s Suspend) String() string {
	return "Suspend"
}

var (
	Free_ = free{}
)

type free struct{}

func (f free) Of(x Any) Free {
	return NewReturn(x)
}

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

func (f functor) LiftEither(x Either) Functor {
	return eitherF{
		val: x,
	}
}

func (f functor) LiftFuncAny(x func() Any) Functor {
	return funcF{
		val: x,
	}
}

func (f functor) LiftFunc(x func() Free) Functor {
	return funcF{
		val: func() Any {
			return x()
		},
	}
}

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

func (x eitherF) String() string {
	return x.val.String()
}

type funcF struct {
	val func() Any
}

func (x funcF) Map(f func(Any) Any) Functor {
	return funcF{
		val: func() Any {
			return f(x.val())
		},
	}
}

func (x funcF) Run() Free {
	return x.val().(Free)
}

func (x funcF) String() string {
	return "FuncF"
}
