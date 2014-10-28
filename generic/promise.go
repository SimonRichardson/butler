package generic

type Promise struct {
	Fork func(func(Any) Any) Any
}

func NewPromise(f func(func(Any) Any) Any) Promise {
	return Promise{
		Fork: f,
	}
}

func (x Promise) Of(v Any) Promise {
	return Promise{
		Fork: func(resolve func(Any) Any) Any {
			return resolve(v)
		},
	}
}

func (x Promise) Chain(f func(v Any) Promise) Promise {
	return Promise{
		Fork: func(resolve func(x Any) Any) Any {
			return x.Fork(func(a Any) Any {
				return f(a).Fork(resolve)
			})
		},
	}
}

func (x Promise) Map(f func(Any) Any) Promise {
	return Promise{func(resolve func(Any) Any) Any {
		return x.Fork(func(a Any) Any {
			return resolve(f(a))
		})
	}}
}

func (x Promise) Extract() Any {
	return x.Fork(Identity())
}

func (x Promise) Extend(f func(Promise) Any) Promise {
	return x.Map(func(y Any) Any {
		return f(x.Of(y))
	})
}
