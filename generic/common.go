package generic

func AsEither(x Any) Either {
	return x.(Either)
}

func AsList(x Any) List {
	return x.(List)
}

func AsStateT(x Any) StateT {
	return x.(StateT)
}

func AsTuple2(x Any) Tuple2 {
	return x.(Tuple2)
}

func AsWriter(x Any) Writer {
	return x.(Writer)
}

func Get() func(Any) StateT {
	return func(Any) StateT {
		return StateT_.Get()
	}
}

func Merge(a StateT) func(Any) StateT {
	return func(b Any) StateT {
		run := func(c Any) Any {
			return NewTuple2(
				Empty{},
				AsWriter(b).Chain(
					func(z Any) Writer {
						x, y := AsWriter(c).Run()
						return NewWriter(NewTuple2(z, x), y)
					},
				),
			)
		}
		return NewStateT(AsEither(a.ExecState("")).Bimap(run, run))
	}
}
