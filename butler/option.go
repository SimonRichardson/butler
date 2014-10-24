package butler

import "github.com/SimonRichardson/butler/generic"

type Option interface {
	Of(generic.Any) Option
	Chain(func(generic.Any) Option) Option
	Map(func(generic.Any) generic.Any) Option
	Fold(func(generic.Any) Option, func() Option) Option
	Ap(generic.Any) Option
	Traverse(func(generic.Any) generic.Any) Option
	GetOrElse(func() generic.Any) generic.Any
}

func ToOption(x generic.Any) Option {
	if x == nil {
		return NewNone()
	}
	return NewSome(x)
}

type Some struct {
	x generic.Any
}

func NewSome(x generic.Any) Some {
	return Some{
		x: x,
	}
}

func (x Some) Of(v generic.Any) Option {
	return NewSome(v)
}

func (x Some) Chain(f func(v generic.Any) Option) Option {
	return f(x.x)
}

func (x Some) Map(f func(v generic.Any) generic.Any) Option {
	return x.Chain(func(v generic.Any) Option {
		return x.Of(f(v))
	})
}

func (x Some) Fold(f func(v generic.Any) Option, g func() Option) Option {
	return f(x.x)
}

func (x Some) Ap(v generic.Any) Option {
	return v.(Option).Map(func(y generic.Any) generic.Any {
		return x.x.(func(Option) Cofree)(y.(Option))
	})
}

func (x Some) Traverse(f func(generic.Any) generic.Any) Option {
	return NewSome(f(x.x))
}

func (x Some) GetOrElse(v func() generic.Any) generic.Any {
	return x.x
}

type None struct{}

func NewNone() None {
	return None{}
}

func (x None) Of(v generic.Any) Option {
	return NewSome(v)
}

func (x None) Chain(f func(v generic.Any) Option) Option {
	return x
}

func (x None) Map(f func(v generic.Any) generic.Any) Option {
	return x
}

func (x None) Fold(f func(v generic.Any) Option, g func() Option) Option {
	return g()
}

func (x None) Ap(v generic.Any) Option {
	return x
}

func (x None) Traverse(f func(generic.Any) generic.Any) Option {
	return NewSome(NewNone())
}

func (x None) GetOrElse(v func() generic.Any) generic.Any {
	return v()
}
