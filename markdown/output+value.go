package markdown

import (
	g "github.com/SimonRichardson/butler/generic"
)

type value struct {
	value string
}

func (e value) IsInline() bool {
	return true
}

func (e value) Children() g.Option {
	return g.Option_.Empty()
}

func (e value) String() string {
	return e.value
}

func str(val string) value {
	return value{
		value: val,
	}
}
