package http

import (
	"github.com/SimonRichardson/butler/doc"
	"github.com/SimonRichardson/butler/generic"
)

type String struct {
	doc.Api
	value string
}

func NewString(value string) String {
	return String{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected string %s"),
			doc.NewInlineText("Unexpected string %s"),
		)),
		value: value,
	}
}

func (s String) Build() generic.State {
	return generic.State{}
}
