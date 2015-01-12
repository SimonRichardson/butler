package http

import (
	"strings"

	g "github.com/SimonRichardson/butler/generic"
)

func compose(a func(func(g.Any) g.Any) g.StateT) func(func(g.Any) func(g.Any) g.Any) func(g.Any) g.StateT {
	return func(b func(g.Any) func(g.Any) g.Any) func(g.Any) g.StateT {
		return func(c g.Any) g.StateT {
			return a(b(c))
		}
	}
}

func constant(a g.StateT) func(g.Any) g.StateT {
	return func(g.Any) g.StateT {
		return a
	}
}

func join(a g.WriterT, f func(g.Either) g.WriterT, h func(g.Any) []g.Any) g.WriterT {
	var (
		x   = a.Run()
		y   = x.Snd()
		run = func(x func(g.Any) g.Either) func(g.Any) g.Any {
			return func(a g.Any) g.Any {
				// We could use compose, if we had the right signatures!
				b := f(x(h(a))).Run()
				return g.NewWriterT(b.Fst(), append(y, b.Snd()...)).
					Tell("Join")
			}
		}
	)
	return g.AsWriterT(x.Fst().Fold(
		run(func(x g.Any) g.Either {
			return g.NewLeft(x)
		}),
		run(func(x g.Any) g.Either {
			return g.NewRight(x)
		}),
	))
}

func matchSplit(s string) func(a g.Any) func(g.Any) g.Any {
	return func(a g.Any) func(g.Any) g.Any {
		return func(b g.Any) g.Any {
			var (
				src = strings.Split(b.(string), s)
				dst = make([]string, len(src))
			)
			for k, v := range src {
				dst[k] = strings.TrimSpace(v)
			}
			return g.NewTuple2(a, dst)
		}
	}
}

func matchGet(a g.Any) func(g.Any) g.Any {
	return func(b g.Any) g.Any {
		return g.NewTuple2(a, b)
	}
}

func matchPut(a g.Any) func(g.Any) g.Any {
	return func(b g.Any) g.Any {
		var (
			x = g.AsTuple2(a)
			y = g.AsTuple2(x.Fst())
			z = y.Append(b)
		)
		return g.NewTuple2(z, z)
	}
}

func matchFlatten(a g.Any) g.StateT {
	return g.NewStateT(g.AsEither(a))
}

// Common aliases

func AsContentDecoder(x g.Any) ContentDecoder {
	return x.(ContentDecoder)
}

func AsContentEncoder(x g.Any) ContentEncoder {
	return x.(ContentEncoder)
}

func AsHeader(x g.Any) Header {
	return x.(Header)
}

func AsMethod(x g.Any) Method {
	return x.(Method)
}

func AsPathNode(x g.Any) PathNode {
	return x.(PathNode)
}

func AsRoute(x g.Any) Route {
	return x.(Route)
}

func AsString(x g.Any) String {
	return x.(String)
}

func modify(a func(g.Any) func(g.Any) g.Any) func(g.Any) g.StateT {
	return compose(g.StateT_.Modify)(a)
}

func singleton(a g.Any) []g.Any {
	return []g.Any{a}
}
