package generic

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

func (t Tuple2) Slice() []Any {
	return []Any{t._1, t._2}
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
