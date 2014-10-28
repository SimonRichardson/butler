package doc

import (
	"fmt"

	"github.com/SimonRichardson/butler/generic"
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

func (d DocTypes) Run(x generic.Either) generic.Either {
	constant := func(doc Doc) func(generic.Any) generic.Any {
		return func(a generic.Any) generic.Any {
			return doc.Run(a)
		}
	}
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

func (d Doc) Run(a generic.Any) string {
	switch d.doc {
	case InlineText:
		return fmt.Sprintf(d.message, a)
	}
	return ""
}
