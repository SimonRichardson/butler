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
