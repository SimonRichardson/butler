package butler

type String struct {
	Api
	value string
}

func NewString(value string) String {
	return String{
		Api: NewApi(NewDocTypes(
			NewInlineText("Expected string %s"),
			NewInlineText("Unexpected string %s"),
		)),
		value: value,
	}
}
