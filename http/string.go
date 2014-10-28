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

// Build up the state, so it runs this when required.
// 1) Convert string to list of chars
// 2) Convert character to number
// 3) Check number is in range
// 4) Return either (expected/unexpected)
// State<Writer<Either<String>, []Doc>>
func (s String) Build() generic.State {
	return generic.State{}
}
