package butler

import (
	g "github.com/SimonRichardson/butler/generic"
)

type Build interface {
	Build() g.StateT
}

type Builder interface {
	List() g.List
}
