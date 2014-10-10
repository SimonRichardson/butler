package butler

type List interface {
	Head() Option
	FoldLeft(Any, func(Any, Any) Any) Any
}

type Cons struct {
	head Any
	tail List
}

func (x Cons) Head() Option {
	return NewSome(x.head)
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

func (x Nil) Head() Option {
	return NewNone()
}

func (x Nil) FoldLeft(v Any, f func(Any, Any) Any) Any {
	return v
}

func FromStringToList(s string, f func(string) Any) List {
	num := len(s)
	res := make([]Any, num, num)
	for i := 0; i < num; i++ {
		res[i] = f(string(s[i]))
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
