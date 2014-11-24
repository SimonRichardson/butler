package generic

type Result interface {
	IsDone() bool
}

type done struct {
	result Any
}

func Done(result Any) Result {
	return done{
		result: result,
	}
}

func (d done) IsDone() bool {
	return true
}

type cont struct {
	thunk func() Result
}

func Cont(thunk func() Result) cont {
	return cont{
		thunk: thunk,
	}
}

func (d cont) IsDone() bool {
	return false
}

func Trampoline(bounce Result) Any {
	for {
		if bounce.IsDone() {
			break
		} else {
			bounce = bounce.(cont).thunk()
		}
	}
	return bounce.(done).result
}
