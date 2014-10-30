package generic

type List interface {
	Head() Option
	Chain(func(Any) List) List
	Map(func(Any) Any) List
	FoldLeft(Any, func(Any, Any) Any) Any
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

func (x Cons) Chain(f func(Any) List) List {
	var rec func(List, List) List
	rec = func(a List, b List) List {
		if _, ok := a.(Nil); ok {
			return b
		}
		cons := a.(Cons)
		list := b.FoldLeft(f(cons.head), func(x Any, y Any) Any {
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

type Nil struct{}

func NewNil() Nil {
	return Nil{}
}

func (x Nil) Head() Option {
	return NewNone()
}

func (x Nil) Chain(f func(Any) List) List {
	return x
}

func (x Nil) Map(f func(Any) Any) List {
	return x
}

func (x Nil) FoldLeft(v Any, f func(Any, Any) Any) Any {
	return v
}

func FromStringToList(s string) List {
	num := len(s)
	res := make([]Any, num, num)
	for i := 0; i < num; i++ {
		res[i] = string(s[i])
	}
	return SliceToList(res)
}

func SliceToList(s []Any) List {
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

func ToList(args ...Any) List {
	return SliceToList(args)
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
