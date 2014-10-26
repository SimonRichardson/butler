package butler

import "github.com/SimonRichardson/butler/generic"

type List interface {
	Head() Option
	Of(generic.Any) List
	Empty() List
	Map(func(generic.Any) generic.Any) List
	FoldLeft(generic.Any, func(generic.Any, generic.Any) generic.Any) generic.Any
}

type Cons struct {
	head generic.Any
	tail List
}

func NewCons(x generic.Any, y List) Cons {
	return Cons{
		head: x,
		tail: y,
	}
}

func (x Cons) Head() Option {
	return NewSome(x.head)
}

func (x Cons) Of(v generic.Any) List {
	return NewCons(v, NewNil())
}

func (x Cons) Empty() List {
	return NewNil()
}

func (x Cons) Map(f func(generic.Any) generic.Any) List {
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

func (x Cons) FoldLeft(v generic.Any, f func(generic.Any, generic.Any) generic.Any) generic.Any {
	var rec func(List, generic.Any) generic.Any
	rec = func(a List, b generic.Any) generic.Any {
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

func (x Nil) Of(v generic.Any) List {
	return NewCons(v, NewNil())
}

func (x Nil) Empty() List {
	return NewNil()
}

func (x Nil) Map(f func(generic.Any) generic.Any) List {
	return x
}

func (x Nil) FoldLeft(v generic.Any, f func(generic.Any, generic.Any) generic.Any) generic.Any {
	return v
}

func FromStringToList(s string, f func(string) generic.Any) List {
	num := len(s)
	res := make([]generic.Any, num, num)
	for i := 0; i < num; i++ {
		res[i] = f(string(s[i]))
	}
	return SliceToList(res)
}

func SliceToList(s []generic.Any) List {
	var rec func(List, []generic.Any) List
	rec = func(l List, v []generic.Any) List {
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
