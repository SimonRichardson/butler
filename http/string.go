package http

import (
	"github.com/SimonRichardson/butler/doc"
	"github.com/SimonRichardson/butler/generic"
)

type String struct {
	doc.Api
	value     string
	validator func(rune) bool
}

func NewString(value string, validator func(rune) bool) String {
	return String{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected string %s"),
			doc.NewInlineText("Unexpected string %s"),
		)),
		value:     value,
		validator: validator,
	}
}

// Series of predicates, could give more info via a Option or Either
func AnyString() func(rune) bool {
	return func(r rune) bool {
		return true
	}
}

func HeaderString() func(rune) bool {
	return func(r rune) bool {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			fallthrough
		case r >= 32 && r <= 39 || r >= 94 && r <= 96:
			fallthrough
		case r == 42 || r == 43 || r == 45 || r == 46 || r == 124:
			return true
		}
		return false
	}
}

func PathString() func(rune) bool {
	return func(r rune) bool {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			fallthrough
		case r == 47 || r == 58:
			return true
		}
		return false
	}
}

func UrlString() func(rune) bool {
	return func(r rune) bool {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			return true
		}
		return false
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
