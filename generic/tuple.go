package generic

import "fmt"

type Tuple2 struct {
	_1 Any
	_2 Any
}

func NewTuple2(a, b Any) Tuple2 {
	return Tuple2{
		_1: a,
		_2: b,
	}
}

func (t Tuple2) Fst() Any {
	return t._1
}

func (t Tuple2) Snd() Any {
	return t._2
}

func (t Tuple2) MapFst(f func(Any) Any) Tuple2 {
	return NewTuple2(f(t._1), t._2)
}

func (t Tuple2) MapSnd(f func(Any) Any) Tuple2 {
	return NewTuple2(t._1, f(t._2))
}

func (t Tuple2) Slice() []Any {
	return []Any{t._1, t._2}
}

func (t Tuple2) String() string {
	return fmt.Sprintf("Tuple2(%s, %s)", t._1, t._2)
}

type Tuple3 struct {
	_1 Any
	_2 Any
	_3 Any
}

func NewTuple3(a, b, c Any) Tuple3 {
	return Tuple3{
		_1: a,
		_2: b,
		_3: c,
	}
}

func (t Tuple3) Fst() Any {
	return t._1
}

func (t Tuple3) Snd() Any {
	return t._2
}

func (t Tuple3) Trd() Any {
	return t._3
}

func (t Tuple3) Slice() []Any {
	return []Any{t._1, t._2, t._3}
}

func (t Tuple3) String() string {
	return fmt.Sprintf("Tuple3(%s, %s, %s)", t._1, t._2, t._3)
}
