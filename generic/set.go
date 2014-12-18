package generic

import "net/http"

type Set struct {
	set map[Any]Any
}

func (s Set) Get(x Any) Option {
	if val, ok := s.set[x]; ok {
		return Option_.Of(val)
	}
	return Option_.Empty()
}

func (s Set) Set(x, y Any) Set {
	add := func() Any {
		c := s.copy()
		c[x] = y
		return c
	}
	return AsSet(s.Get(x).Fold(
		func(a Any) Any {
			if a == y {
				return s
			}
			return add()
		},
		add,
	))
}

func (s Set) copy() map[Any]Any {
	r := make(map[Any]Any)
	for k, v := range s.set {
		r[k] = v
	}
	return r
}

var (
	Set_ = set{}
)

type set struct{}

func (s set) FromMap(m map[Any]Any) Set {
	return Set{
		set: m,
	}
}

func (s set) HttpHeaderToSet(m http.Header) Set {
	x := make(map[Any]Any)
	for k, v := range m {
		x[k] = List_.StringSliceToList(v)
	}
	return Set{
		set: x,
	}
}
