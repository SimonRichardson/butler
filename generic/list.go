package generic

type List interface {
	Head() Option
	Tail() List
	Chain(func(Any) List) List
	Map(func(Any) Any) List
	Concat(List) List
	Filter(func(Any) bool) List
	Find(func(Any) bool) Option
	FoldLeft(Any, func(Any, Any) Any) Any
	ReduceLeft(func(Any, Any) Any) Option
}

type Cons struct {
	head Any
	tail List
}

func NewCons(x Any, y List) Cons {
	return Cons{
		head: x,
		tail: y,
	}
}

func (x Cons) Head() Option {
	return NewSome(x.head)
}

func (x Cons) Tail() List {
	return x.tail
}

func (x Cons) Chain(f func(Any) List) List {
	var rec func(List, List) List
	rec = func(a List, b List) List {
		if _, ok := a.(Nil); ok {
			return b
		}
		cons := a.(Cons)
		list := f(cons.head).FoldLeft(b, func(x, y Any) Any {
			return NewCons(y, x.(List))
		})
		return rec(cons.tail, list.(List))
	}
	return rec(x, NewNil())
}

func (x Cons) Map(f func(Any) Any) List {
	return x.Chain(func(a Any) List {
		return NewCons(f(a), NewNil())
	})
}

func (x Cons) Concat(y List) List {
	var rec func(List, List) Result
	rec = func(a, b List) Result {
		if _, ok := b.(Nil); ok {
			return Done(a)
		}
		return Cont(func() Result {
			cons := b.(Cons)
			return rec(NewCons(cons.head, a), cons.tail)
		})
	}
	return Trampoline(rec(x, y)).(List)
}

func (x Cons) Filter(f func(Any) bool) List {
	var rec func(List, List) List
	rec = func(a, b List) List {
		if _, ok := a.(Nil); ok {
			return b
		}
		cons := a.(Cons)
		if f(cons.head) {
			return rec(cons.tail, NewCons(cons.head, b))
		} else {
			return rec(cons.tail, b)
		}
	}
	return rec(x, List_.Empty())
}

func (x Cons) Find(f func(Any) bool) Option {
	var rec func(List, Option) Option
	rec = func(a List, b Option) Option {
		if _, ok := a.(Nil); ok {
			return b
		}
		return b.Fold(
			func(x Any) Any {
				return Option_.Of(x)
			},
			func() Any {
				var (
					cons = a.(Cons)
					val  = cons.head
					opt  = Option_.FromBool(f(val), val)
				)
				return rec(cons.tail, opt)
			},
		).(Option)
	}
	return rec(x, Option_.Empty())
}

func (x Cons) FoldLeft(v Any, f func(Any, Any) Any) Any {
	var rec func(List, Any) Any
	rec = func(a List, b Any) Any {
		if _, ok := a.(Nil); ok {
			return b
		}
		cons := a.(Cons)
		return rec(cons.tail, f(b, cons.head))
	}
	return rec(x, v)
}

func (x Cons) ReduceLeft(f func(Any, Any) Any) Option {
	return Option_.Of(x.tail.FoldLeft(x.head, f))
}

type Nil struct{}

func NewNil() Nil {
	return Nil{}
}

func (x Nil) Head() Option {
	return Option_.Empty()
}

func (x Nil) Tail() List {
	return x
}

func (x Nil) Chain(f func(Any) List) List {
	return x
}

func (x Nil) Map(f func(Any) Any) List {
	return x
}

func (x Nil) Concat(y List) List {
	return y
}

func (x Nil) Filter(func(Any) bool) List {
	return x
}

func (x Nil) Find(f func(Any) bool) Option {
	return Option_.Empty()
}

func (x Nil) FoldLeft(v Any, f func(Any, Any) Any) Any {
	return v
}

func (x Nil) ReduceLeft(f func(Any, Any) Any) Option {
	return Option_.Empty()
}

// Static methods

var (
	List_ = list{}
)

type list struct{}

func (x list) Of(v Any) List {
	return NewCons(v, NewNil())
}

func (x list) Empty() List {
	return NewNil()
}

func (x list) To(args ...Any) List {
	return x.FromSlice(args)
}

func (x list) ToSlice(l List) []Any {
	return l.FoldLeft(make([]Any, 0, 0), func(a, b Any) Any {
		return append(a.([]Any), b)
	}).([]Any)
}

func (x list) FromAmount(s int) List {
	var rec func(List, int) List
	rec = func(l List, v int) List {
		if v < 1 {
			return l
		}
		return rec(Cons{
			head: v,
			tail: l,
		}, v-1)
	}
	return rec(Nil{}, s)
}

func (x list) FromSlice(s []Any) List {
	var rec func(List, []Any) List
	rec = func(l List, v []Any) List {
		num := len(v)
		if num < 1 {
			return l
		}
		return rec(Cons{
			head: v[num-1],
			tail: l,
		}, v[:num-1])
	}
	return rec(Nil{}, s)
}

func (x list) FromString(s string) List {
	num := len(s)
	res := make([]Any, num, num)
	for i := 0; i < num; i++ {
		res[i] = string(s[i])
	}
	return x.FromSlice(res)
}
