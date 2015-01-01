package generic

func AsEither(x Any) Either {
	return x.(Either)
}

func AsFree(x Any) Free {
	return x.(Free)
}

func AsIO(x Any) IO {
	return x.(IO)
}

func AsList(x Any) List {
	return x.(List)
}

func AsOption(x Any) Option {
	return x.(Option)
}

func AsSet(x Any) Set {
	return x.(Set)
}

func AsStateT(x Any) StateT {
	return x.(StateT)
}

func AsTree(x Any) Tree {
	return x.(Tree)
}

func AsTuple2(x Any) Tuple2 {
	return x.(Tuple2)
}

func AsTuple3(x Any) Tuple3 {
	return x.(Tuple3)
}

func AsWriterTuple(x Any) WriterTuple {
	return x.(WriterTuple)
}

func AsWriter(x Any) Writer {
	return x.(Writer)
}

func AsWriterTTuple(x Any) WriterTTuple {
	return x.(WriterTTuple)
}

func AsWriterT(x Any) WriterT {
	return x.(WriterT)
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
						var (
							exe = AsWriter(c).Run()
							x   = exe.Fst()
							y   = exe.Snd()
						)
						return NewWriter(NewTuple2(z, x), y)
					},
				),
			)
		}
		return NewStateT(AsEither(a.ExecState(Empty{})).Bimap(run, run))
	}
}
