package http

import (
	"fmt"

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
	program := generic.State{}.Of(s.value)
	return program.
		Map(stringToChars).
		Map(charsToInt).
		Map(validStringValue).
		Map(returnEither)
}

func stringToChars(x generic.Any) generic.Any {
	return generic.FromStringToList(x.(string))
}

func charsToInt(x generic.Any) generic.Any {
	fmt.Println(">>", x)
	return x.(string)
}

func validStringValue(x generic.Any) generic.Any {
	return x
}

func returnEither(x generic.Any) generic.Any {
	return x
}
