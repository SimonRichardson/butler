package output

import g "github.com/SimonRichardson/butler/generic"

type Encoder interface {
	Encode(a g.Any) g.Either
}

func toEither(a []byte, b error) g.Either {
	if b != nil {
		return g.NewLeft(b)
	}
	return g.NewRight(a)
}
