package output

import "github.com/SimonRichardson/butler/generic"

type Encoder interface {
	Encode(a generic.Any) ([]byte, error)
}
