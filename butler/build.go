package butler

import "github.com/SimonRichardson/butler/generic"

type Build interface {
	Build() generic.State
}
