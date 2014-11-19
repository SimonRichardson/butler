package output

import g "github.com/SimonRichardson/butler/generic"

type Encoder interface {
	Encode(g.Any) g.Either
	Generate(g.Any) g.Either
}
