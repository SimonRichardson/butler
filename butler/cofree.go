package butler

import "github.com/SimonRichardson/butler/generic"

type Cofree struct {
	value   generic.Any
	functor Option
}

func (c Cofree) Map(f func(generic.Any) generic.Any) Cofree {
	return Cofree{
		value: f(c.value),
		functor: c.functor.Map(func(a generic.Any) generic.Any {
			return a.(Cofree).Map(f)
		}),
	}
}

func (c Cofree) Extract() generic.Any {
	return c.value
}

func (c Cofree) Extend(f func(Cofree) generic.Any) Cofree {
	return Cofree{
		value: f(c),
		functor: c.functor.Map(func(a generic.Any) generic.Any {
			return a.(Cofree).Extend(f)
		}),
	}
}

func (c Cofree) Traverse(g func(generic.Any) Option) Option {
	var do func(generic.Any) generic.Any
	do = func(h generic.Any) generic.Any {
		a := h.(Cofree)
		return g(a.value).Map(func(x generic.Any) generic.Any {
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
