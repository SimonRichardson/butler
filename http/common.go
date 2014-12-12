package http

import g "github.com/SimonRichardson/butler/generic"

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

// Common aliases

func asContentDecoder(x g.Any) ContentDecoder {
	return x.(ContentDecoder)
}

func asContentEncoder(x g.Any) ContentEncoder {
	return x.(ContentEncoder)
}

func asMethod(x g.Any) Method {
	return x.(Method)
}

func AsPathNode(x g.Any) PathNode {
	return x.(PathNode)
}

func AsRoute(x g.Any) Route {
	return x.(Route)
}

func asString(x g.Any) String {
	return x.(String)
}

func modify(a func(g.Any) func(g.Any) g.Any) func(g.Any) g.StateT {
	return compose(g.StateT_.Modify)(a)
}

func singleton(a g.Any) []g.Any {
	return []g.Any{a}
}
