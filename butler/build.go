package butler

import (
	g "github.com/SimonRichardson/butler/generic"
)

type Build interface {
	Build() g.StateT
}
