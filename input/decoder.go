package output

import (
	g "github.com/SimonRichardson/butler/generic"
)

type Decoder interface {
	Keys(g.Any) g.Either
	Decode(a g.Any) g.Either
}
