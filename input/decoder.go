package output

import (
	g "github.com/SimonRichardson/butler/generic"
)

type Decoder interface {
	Decode(a g.Any) (g.Any, error)
}
