package http

import (
	"github.com/SimonRichardson/butler/doc"
	g "github.com/SimonRichardson/butler/generic"
)

type MethodType string

const (
	DELETE  MethodType = "delete"
	GET     MethodType = "get"
	HEAD    MethodType = "head"
	OPTIONS MethodType = "options"
	PATCH   MethodType = "patch"
	POST    MethodType = "post"
	PUT     MethodType = "put"
	TRACE   MethodType = "trace"
)

var (
	methodTypes = []MethodType{
		DELETE,
		GET,
		HEAD,
		OPTIONS,
		PATCH,
		POST,
		PUT,
		TRACE,
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

func (m Method) Build() g.StateT {
	var (
		contains = func(types []MethodType) func(t MethodType) g.Either {
			return func(t MethodType) g.Either {
				for _, v := range types {
					if t == v {
						return g.NewRight(t)
					}
				}
				return g.NewLeft(t)
			}
		}
		validate = func(f func(t MethodType) g.Either) func(g.Any) func(g.Any) g.Any {
			return func(x g.Any) func(g.Any) g.Any {
				return func(b g.Any) g.Any {
					return f(asMethod(b).method)
				}
			}
		}
		api = func(api doc.Api) func(g.Any) func(g.Any) g.Any {
			return func(g.Any) func(g.Any) g.Any {
				return func(a g.Any) g.Any {
					sum := func(a g.Any) g.Any {
						return singleton(a)
					}
					return api.Run(g.AsEither(a).Bimap(sum, sum))
				}
			}
		}
		finalise = func(m Method) func(g.Any) g.StateT {
			return func(g.Any) g.StateT {
				return g.StateT{
					Run: func(a g.Any) g.Either {
						cast := func(b g.Any) g.Any {
							x := g.NewWriter(m, singleton(a))
							return g.NewTuple2(g.Empty{}, x)
						}
						return g.AsEither(a).Bimap(cast, cast)
					},
				}
			}
		}
	)
	return g.StateT_.Of(m).
		Chain(modify(g.Constant1)).
		Chain(modify(validate(contains(methodTypes)))).
		Chain(modify(api(m.Api))).
		Chain(finalise(m))
}

func Delete() Method {
	return NewMethod(DELETE)
}

func Get() Method {
	return NewMethod(GET)
}

func Head() Method {
	return NewMethod(HEAD)
}

func Options() Method {
	return NewMethod(OPTIONS)
}

func Patch() Method {
	return NewMethod(PATCH)
}

func Post() Method {
	return NewMethod(POST)
}

func Put() Method {
	return NewMethod(PUT)
}

func Trace() Method {
	return NewMethod(TRACE)
}
