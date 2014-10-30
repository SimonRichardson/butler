package generic

type Either interface {
	Chain(f func(v Any) Either) Either
	Map(f func(v Any) Any) Either
	Bimap(f func(v Any) Any, g func(v Any) Any) Either
	Fold(f func(v Any) Any, g func(v Any) Any) Any
}

type Right struct {
	x Any
}

func NewRight(x Any) Right {
	return Right{
		x: x,
	}
}

func (x Right) Chain(f func(v Any) Either) Either {
	return f(x.x)
}

func (x Right) Map(f func(v Any) Any) Either {
	return Either_.Of(f(x.x))
}

func (x Right) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return g(x.x)
}

func (x Right) Bimap(f func(v Any) Any, g func(v Any) Any) Either {
	return NewRight(g(x.x))
}

type Left struct {
	x Any
}

func NewLeft(x Any) Left {
	return Left{
		x: x,
	}
}

func (x Left) Chain(f func(v Any) Either) Either {
	return x
}

func (x Left) Map(f func(v Any) Any) Either {
	return NewLeft(x.x)
}

func (x Left) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return f(x.x)
}

func (x Left) Bimap(f func(v Any) Any, g func(v Any) Any) Either {
	return NewLeft(f(x.x))
}

func EitherFromBool(b bool, val Any) Either {
	if b {
		return NewRight(val)
	}
	return NewLeft(val)
}

// Static methods

var (
	Either_ = either{}
)

type either struct{}

func (x either) Of(v Any) Either {
	return NewRight(v)
}
