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

func either(x g.Any) g.Either {
	return x.(g.Either)
}

func list(x g.Any) g.List {
	return x.(g.List)
}

func tuple2(x g.Any) g.Tuple2 {
	return x.(g.Tuple2)
}

func writer(x g.Any) g.Writer {
	return x.(g.Writer)
}

func merge(a g.StateT) func(g.Any) g.StateT {
	return func(b g.Any) g.StateT {
		run := func(c g.Any) g.Any {
			return g.NewTuple2(
				g.Empty{},
				writer(b).Chain(
					func(z g.Any) g.Writer {
						x, y := writer(c).Run()
						return g.NewWriter(g.NewTuple2(z, x), y)
					},
				),
			)
		}
		return g.NewStateT(either(a.ExecState("")).Bimap(run, run))
	}
}

func modify(a func(g.Any) func(g.Any) g.Any) func(g.Any) g.StateT {
	return compose(g.StateT_.Modify)(a)
}

func get() func(g.Any) g.StateT {
	return func(g.Any) g.StateT {
		return g.StateT_.Get()
	}
}

func singleton(a g.Any) []g.Any {
	return []g.Any{a}
}
