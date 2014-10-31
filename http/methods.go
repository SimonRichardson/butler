package http

import (
	"github.com/SimonRichardson/butler/doc"
	"github.com/SimonRichardson/butler/generic"
)

type MethodType string

const (
	MDelete  MethodType = "delete"
	MGet     MethodType = "get"
	MHead    MethodType = "head"
	MOptions MethodType = "options"
	MPatch   MethodType = "patch"
	MPost    MethodType = "post"
	MPut     MethodType = "put"
	MTrace   MethodType = "trace"
)

var (
	methodTypes = []MethodType{
		MDelete,
		MGet,
		MHead,
		MOptions,
		MPatch,
		MPost,
		MPut,
		MTrace,
	}
)

type Method struct {
	doc.Api
	method MethodType
}

func NewMethod(method MethodType) Method {
	return Method{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected method `%s`"),
			doc.NewInlineText("Unexpected method `%s`"),
		)),
		method: method,
	}
}

func (m Method) Build() generic.State {
	var (
		extract = func(x generic.Any) func(func(Method, MethodType) generic.Tuple2) generic.Tuple2 {
			return func(f func(Method, MethodType) generic.Tuple2) generic.Tuple2 {
				tuple := x.(generic.Tuple2)
				method := tuple.Fst().(Method)
				methodType := tuple.Snd().(MethodType)

				return f(method, methodType)
			}
		}
		setup = func(x generic.Any) generic.Any {
			return generic.NewTuple2(m, m.method)
		}
		validate = func(types []MethodType) func(generic.Any) generic.Any {
			contains := func(x []MethodType, y MethodType) bool {
				for _, v := range x {
					if v == y {
						return true
					}
				}
				return false
			}
			return func(x generic.Any) generic.Any {
				return extract(x)(func(method Method, methodType MethodType) generic.Tuple2 {
					return generic.NewTuple2(
						method,
						generic.Either_.FromBool(contains(types, methodType), methodType),
					)
				})
			}
		}
		api = func(x generic.Any) generic.Any {
			tuple := x.(generic.Tuple2)
			method := tuple.Fst().(Method)

			sum := func(a generic.Any) generic.Any {
				return []generic.Any{a}
			}
			folded := tuple.Snd().(generic.Either).Bimap(sum, sum)

			return generic.NewTuple2(method, method.Api.Run(folded))
		}
	)
	return generic.State_.Of(m).
		Map(setup).
		Map(validate(methodTypes)).
		Map(api)
}

func Delete() Method {
	return NewMethod(MDelete)
}

func Get() Method {
	return NewMethod(MGet)
}

func Head() Method {
	return NewMethod(MHead)
}

func Options() Method {
	return NewMethod(MOptions)
}

func Patch() Method {
	return NewMethod(MPatch)
}

func Post() Method {
	return NewMethod(MPost)
}

func Put() Method {
	return NewMethod(MPut)
}

func Trace() Method {
	return NewMethod(MTrace)
}
