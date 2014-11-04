package doc

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type DocType string

var (
	InlineText DocType = "inline"
)

type DocTypes struct {
	expected   Doc
	unexpected Doc
}

func NewDocTypes(x, y Doc) DocTypes {
	return DocTypes{
		expected:   x,
		unexpected: y,
	}
}

func (d DocTypes) Run(x g.Either) g.Either {
	return x.Bimap(constant(d.unexpected), constant(d.expected))
}

type Doc struct {
	message string
	doc     DocType
}

func NewInlineText(message string) Doc {
	return Doc{
		message: message,
		doc:     InlineText,
	}
}

func (d Doc) Run(a g.Any) string {
	switch d.doc {
	case InlineText:
		x := a.([]g.Any)
		y := len(x)
		z := make([]interface{}, y, y)
		for k, v := range x {
			z[k] = v
		}
		return fmt.Sprintf(d.message, z...)
	}
	return ""
}
