package generic

type List interface {
	Head() Option
	Of(Any) List
	Empty() List
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

func (x Cons) Of(v Any) List {
	return NewCons(v, NewNil())
}

func (x Cons) Empty() List {
	return NewNil()
}

func (x Cons) Map(f func(Any) Any) List {
	var rec func(List, List) List
	rec = func(a List, b List) List {
		if _, ok := a.(Nil); ok {
			return b
		}
		cons := a.(Cons)
		return rec(cons.tail, NewCons(f(cons.head), b))
	}
	return rec(x, NewNil())
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

func (x Nil) Of(v Any) List {
	return NewCons(v, NewNil())
}

func (x Nil) Empty() List {
	return NewNil()
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
		if len(v) < 1 {
			return l
		}
		return rec(Cons{
			head: v[0],
			tail: l,
		}, v[1:])
	}
	return rec(Nil{}, s)
}
