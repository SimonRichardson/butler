package generic

import "fmt"

type List interface {
	Head() Option
	Last() Option
	Tail() List
	Chain(func(Any) List) List
	Map(func(Any) Any) List
	Concat(List) List
	Filter(func(Any) bool) List
	Find(func(Any) bool) Option
	FoldLeft(Any, func(Any, Any) Any) Any
	FoldRight(Any, func(Any, Any) Any) Any
	GroupBy(func(Any) Any) List
	Index(uint) Option
	Partition(func(Any) bool) Tuple2
	ReduceLeft(func(Any, Any) Any) Option
	ReduceRight(func(Any, Any) Any) Option
	Reverse() List
	Size() int
	Zip(List) List
	ZipWithIndex() List
	String() string
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

func (x Cons) Last() Option {
	return x.Reverse().Head()
}

func (x Cons) Tail() List {
	return x.tail
}

func (x Cons) Chain(f func(Any) List) List {
	var rec func(List, List) Free
	rec = func(a List, b List) Free {
		if _, ok := a.(Nil); ok {
			return NewReturn(b)
		}
		return NewSuspend(Functor_.LiftFuncAny(func() Any {
			cons := a.(Cons)
			list := f(cons.head).FoldLeft(b, func(x, y Any) Any {
				return NewCons(y, AsList(x))
			})
			return rec(cons.tail, AsList(list))
		}))
	}
	return AsList(rec(x, NewNil()).Run())
}

func (x Cons) Map(f func(Any) Any) List {
	return x.Chain(func(a Any) List {
		return NewCons(f(a), NewNil())
	})
}

func (x Cons) Concat(y List) List {
	var rec func(List, List) Free
	rec = func(a, b List) Free {
		if _, ok := b.(Nil); ok {
			return NewReturn(a)
		}
		return NewSuspend(Functor_.LiftFuncAny(func() Any {
			cons := b.(Cons)
			return rec(NewCons(cons.head, a), cons.tail)
		}))
	}
	return AsList(rec(x, y).Run())
}

func (x Cons) Filter(f func(Any) bool) List {
	var rec func(List, List) Free
	rec = func(a, b List) Free {
		if _, ok := a.(Nil); ok {
			return NewReturn(b)
		}
		return NewSuspend(Functor_.LiftFuncAny(func() Any {
			var (
				cons = a.(Cons)
				tail = cons.tail
			)
			if f(cons.head) {
				return rec(tail, NewCons(cons.head, b))
			} else {
				return rec(tail, b)
			}
		}))
	}
	return AsList(rec(x, List_.Empty()).Run())
}

func (x Cons) Find(f func(Any) bool) Option {
	var rec func(List, Option) Free
	rec = func(a List, b Option) Free {
		if _, ok := a.(Nil); ok {
			return NewReturn(b)
		}
		return b.Fold(
			func(x Any) Any {
				return NewReturn(Option_.Of(x))
			},
			func() Any {
				var (
					cons = a.(Cons)
					val  = cons.head
					opt  = Option_.FromBool(f(val), val)
				)
				return NewSuspend(Functor_.LiftFuncAny(func() Any {
					return rec(cons.tail, opt)
				}))
			},
		).(Free)
	}
	return AsOption(rec(x, Option_.Empty()).Run())
}

func (x Cons) FoldLeft(v Any, f func(Any, Any) Any) Any {
	var rec func(List, Any) Free
	rec = func(a List, b Any) Free {
		if _, ok := a.(Nil); ok {
			return NewReturn(b)
		}
		return NewSuspend(Functor_.LiftFuncAny(func() Any {
			cons := a.(Cons)
			return rec(cons.tail, f(b, cons.head))
		}))
	}
	return rec(x, v).Run()
}

func (x Cons) FoldRight(v Any, f func(Any, Any) Any) Any {
	var rec func(List, Any) Free
	rec = func(a List, b Any) Free {
		if _, ok := a.(Nil); ok {
			return NewReturn(b)
		}
		cons := a.(Cons)
		return NewSuspend(Functor_.LiftFuncAny(func() Any {
			return rec(cons.tail, v)
		})).Map(func(x Any) Any {
			return f(x, cons.head)
		})
	}
	return rec(x, v).Run()
}

func (x Cons) GroupBy(f func(Any) Any) List {
	var (
		contains = func(a List, b Any) Option {
			return a.Find(func(x Any) bool {
				return AsTuple2(x).Fst() == b
			})
		}
		unique = func(a Any) func(Any) bool {
			return func(b Any) bool {
				return AsTuple2(b).Fst() != a
			}
		}
	)
	return AsList(x.FoldLeft(NewNil(), func(a, b Any) Any {
		var (
			id   = f(b)
			list = AsList(a)
		)
		return contains(list, id).Fold(
			func(x Any) Any {
				merge := AsList(AsTuple2(x).Snd()).Concat(List_.Of(b))
				return NewCons(NewTuple2(id, merge), list.Filter(unique(id)))
			},
			func() Any {
				return NewCons(NewTuple2(id, List_.Of(b)), list)
			},
		)
	}))
}

func (x Cons) Index(index uint) Option {
	var rec func(List, uint) Free
	rec = func(a List, b uint) Free {
		if _, ok := a.(Nil); ok {
			return NewReturn(Option_.Empty())
		}

		cons := a.(Cons)
		if b == 0 {
			return NewReturn(Option_.Of(cons.head))
		}
		return NewSuspend(Functor_.LiftFuncAny(func() Any {
			return rec(cons.tail, b-1)
		}))
	}
	return AsOption(rec(x, index).Run())
}

func (x Cons) Partition(f func(Any) bool) Tuple2 {
	return AsTuple2(x.FoldLeft(NewTuple2(List_.Empty(), List_.Empty()), func(a, b Any) Any {
		x := AsTuple2(a)
		if f(b) {
			return NewTuple2(NewCons(b, AsList(x.Fst())), x.Snd())
		}
		return NewTuple2(x.Fst(), NewCons(b, AsList(x.Snd())))
	}))
}

func (x Cons) ReduceLeft(f func(Any, Any) Any) Option {
	return Option_.Of(x.tail.FoldLeft(x.head, f))
}

func (x Cons) ReduceRight(f func(Any, Any) Any) Option {
	return x.Reverse().ReduceLeft(f)
}

func (x Cons) Reverse() List {
	return AsList(x.FoldLeft(NewNil(), func(a, b Any) Any {
		return NewCons(b, AsList(a))
	}))
}

func (x Cons) Size() int {
	return x.FoldLeft(0, func(a, b Any) Any {
		return a.(int) + 1
	}).(int)
}

func (x Cons) Zip(y List) List {
	var rec func(a, b, c List) Free
	rec = func(a, b, c List) Free {
		_, ok1 := a.(Nil)
		_, ok2 := b.(Nil)

		if ok1 || ok2 {
			return NewReturn(c)
		}
		return NewSuspend(Functor_.LiftFuncAny(func() Any {
			var (
				x = a.(Cons)
				y = b.(Cons)
			)
			return rec(x.tail, y.tail, NewCons(NewTuple2(x.head, y.head), c))
		}))
	}
	return AsList(rec(x, y, NewNil()).Run())
}

func (x Cons) ZipWithIndex() List {
	var rec func(a List, b int, c List) Free
	rec = func(a List, b int, c List) Free {
		if _, ok := a.(Nil); ok {
			return NewReturn(c)
		}
		return NewSuspend(Functor_.LiftFuncAny(func() Any {
			x := a.(Cons)
			return rec(x.tail, b+1, NewCons(NewTuple2(x.head, b), c))
		}))
	}
	return AsList(rec(x, 0, NewNil()).Run())
}

func (x Cons) String() string {
	res := x.ReduceRight(func(a, b Any) Any {
		return fmt.Sprintf("%v, %v", a, b)
	})
	return fmt.Sprintf("List(%s)", res.GetOrElse(Constant("")))
}

type Nil struct{}

func NewNil() Nil {
	return Nil{}
}

func (x Nil) Head() Option {
	return Option_.Empty()
}

func (x Nil) Last() Option {
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

func (x Nil) FoldRight(v Any, f func(Any, Any) Any) Any {
	return v
}

func (x Nil) GroupBy(f func(Any) Any) List {
	return x
}

func (x Nil) Index(index uint) Option {
	return Option_.Empty()
}

func (x Nil) Partition(f func(Any) bool) Tuple2 {
	return NewTuple2(List_.Empty(), List_.Empty())
}

func (x Nil) ReduceLeft(f func(Any, Any) Any) Option {
	return Option_.Empty()
}

func (x Nil) ReduceRight(f func(Any, Any) Any) Option {
	return Option_.Empty()
}

func (x Nil) Reverse() List {
	return x
}

func (x Nil) Size() int {
	return 0
}

func (x Nil) Zip(a List) List {
	return List_.Empty()
}

func (x Nil) ZipWithIndex() List {
	return List_.Empty()
}

func (x Nil) String() string {
	return "List()"
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

func (x list) FromBool(a bool, b Any) List {
	if a {
		return List_.Of(b)
	}
	return List_.Empty()
}

func (x list) FromArgs(s ...Any) List {
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

func (x list) StringSliceToList(s []string) List {
	var rec func(List, []string) Free
	rec = func(l List, v []string) Free {
		num := len(v)
		if num < 1 {
			return NewReturn(l)
		}
		return NewSuspend(Functor_.LiftFuncAny(func() Any {
			return rec(NewCons(v[num-1], l), v[:num-1])
		}))
	}
	return AsList(rec(NewNil(), s).Run())
}

func (x list) Cons(v Any) Any {
	return x.Of(v)
}

func (x list) Nil(v Any) Any {
	return x.Empty()
}
