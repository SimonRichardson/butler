package markdown

import (
	g "github.com/SimonRichardson/butler/generic"
)

type empty struct {
}

func (e empty) IsBlock() bool {
	return false
}

func (e empty) Children() g.Option {
	return g.Option_.Empty()
}

func (e empty) String() string {
	return ""
}

func nothing() mark {
	return empty{}
}
