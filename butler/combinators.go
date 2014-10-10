package butler

func Identity() func(Any) Any {
	return func(x Any) Any {
		return x
	}
}

func Constant(x Any) func() Any {
	return func() Any {
		return x
	}
}

func Constant1(x Any) func(Any) Any {
	return func(y Any) Any {
		return x
	}
}

func Compose(f func(x Any) Any) func(func(Any) Any) func(Any) Any {
	return func(g func(Any) Any) func(Any) Any {
		return func(a Any) Any {
			return f(g(a))
		}
	}
}
