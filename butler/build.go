package butler

import (
	g "github.com/SimonRichardson/butler/generic"
)

type Build interface {
	Build() g.WriterT
}

type Builder interface {
	List() g.List
}
