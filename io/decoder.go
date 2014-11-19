package io

import (
	g "github.com/SimonRichardson/butler/generic"
)

type Decoder interface {
	Keys() g.Either
	Decode(a []byte) g.Either
}
