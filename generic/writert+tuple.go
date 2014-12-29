package generic

import "fmt"

type WriterTTuple struct {
	_1 Either
	_2 []Any
}

func NewWriterTTuple(a Either, b []Any) WriterTTuple {
	return WriterTTuple{
		_1: a,
		_2: b,
	}
}

func (t WriterTTuple) Fst() Either {
	return t._1
}

func (t WriterTTuple) Snd() []Any {
	return t._2
}

func (t WriterTTuple) Slice() []Any {
	return []Any{t._1, t._2}
}

func (t WriterTTuple) String() string {
	return fmt.Sprintf("WriterTTuple(%s, %s)", t._1, t._2)
}
