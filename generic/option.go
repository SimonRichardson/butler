package generic

type Option interface {
	Chain(func(Any) Option) Option
	Map(func(Any) Any) Option
	Fold(func(Any) Any, func() Any) Any
	Ap(Any) Option
	Traverse(func(Any) Any) Option
	GetOrElse(func() Any) Any
}

func ToOption(x Any) Option {
	if x == nil {
		return NewNone()
	}
	return NewSome(x)
}

type Some struct {
	x Any
}

func NewSome(x Any) Some {
	return Some{
		x: x,
	}
}

func (x Some) Chain(f func(v Any) Option) Option {
	return f(x.x)
}

func (x Some) Map(f func(v Any) Any) Option {
	return x.Chain(func(v Any) Option {
		return Option_.Of(f(v))
	})
}

func (x Some) Fold(f func(v Any) Any, g func() Any) Any {
	return f(x.x)
}

func (x Some) Ap(v Any) Option {
	return v.(Option).Map(func(y Any) Any {
		return x.x.(func(Option) Cofree)(y.(Option))
	})
}

func (x Some) Traverse(f func(Any) Any) Option {
	return NewSome(f(x.x))
}

func (x Some) GetOrElse(v func() Any) Any {
	return x.x
}

type None struct{}

func NewNone() None {
	return None{}
}

func (x None) Chain(f func(v Any) Option) Option {
	return x
}

func (x None) Map(f func(v Any) Any) Option {
	return x
}

func (x None) Fold(f func(v Any) Any, g func() Any) Any {
	return g()
}

func (x None) Ap(v Any) Option {
	return x
}

func (x None) Traverse(f func(Any) Any) Option {
	return NewSome(NewNone())
}

func (x None) GetOrElse(v func() Any) Any {
	return v()
}

// Static methods

var (
	Option_ = option{}
)

type option struct{}

func (x option) Of(v Any) Option {
	return NewSome(v)
}

func (x option) Empty() Option {
	return NewNone()
}

func (x option) FromBool(b bool, val Any) Option {
	if b {
		return NewSome(val)
	}
	return NewNone()
}
