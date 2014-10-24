package butler

import "github.com/SimonRichardson/butler/generic"

type Tuple struct {
	_1 generic.Any
	_2 []generic.Any
}

func NewTuple(a generic.Any, b []generic.Any) Tuple {
	return Tuple{
		_1: a,
		_2: b,
	}
}

func (t Tuple) Fst() generic.Any {
	return t._1
}

func (t Tuple) Snd() []generic.Any {
	return t._2
}
