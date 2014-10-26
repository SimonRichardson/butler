package generic

type Cofree struct {
	value   Any
	functor Option
}

func (c Cofree) Map(f func(Any) Any) Cofree {
	return Cofree{
		value: f(c.value),
		functor: c.functor.Map(func(a Any) Any {
			return a.(Cofree).Map(f)
		}),
	}
}

func (c Cofree) Extract() Any {
	return c.value
}

func (c Cofree) Extend(f func(Cofree) Any) Cofree {
	return Cofree{
		value: f(c),
		functor: c.functor.Map(func(a Any) Any {
			return a.(Cofree).Extend(f)
		}),
	}
}

func (c Cofree) Traverse(g func(Any) Option) Option {
	var do func(Any) Any
	do = func(h Any) Any {
		a := h.(Cofree)
		return g(a.value).Map(func(x Any) Any {
			return func(i Option) Cofree {
				return Cofree{
					value:   x,
					functor: i,
				}
			}
		}).Ap(a.functor.Traverse(do))
	}
	return do(c).(Option)
}
