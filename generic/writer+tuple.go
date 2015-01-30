package generic

import "fmt"

type WriterTuple struct {
	_1 Any
	_2 []Any
}

func NewWriterTuple(a Any, b []Any) WriterTuple {
	return WriterTuple{
		_1: a,
		_2: b,
	}
}

func (t WriterTuple) Fst() Any {
	return t._1
}

func (t WriterTuple) Snd() []Any {
	return t._2
}

func (t WriterTuple) Slice() []Any {
	return []Any{t._1, t._2}
}

func (t WriterTuple) String() string {
	return fmt.Sprintf("WriterTuple(%s, %s)", t._1, t._2)
}
