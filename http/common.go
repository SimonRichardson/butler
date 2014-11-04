package http

import (
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

// Common aliases

func modify(a func(g.Any) func(g.Any) g.Any) func(g.Any) g.StateT {
	return compose(g.StateT_.Modify)(a)
}

func get() func(g.Any) g.StateT {
	return func(g.Any) g.StateT {
		return g.StateT_.Get()
	}
}
