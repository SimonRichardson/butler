package generic

func Done(x Any) Free {
	return NewReturn(x)
}

func Cont(f func() Free) Free {
	return NewSuspend(Functor_.LiftFunc(f))
}

func Trampoline(x Free) Any {
	return x.Run()
}
