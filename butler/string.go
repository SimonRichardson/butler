package butler

type String struct {
	value string
}

func NewString(value string) Api {
	return NewApi(
		String{
			value: value,
		},
		NewDocTypes(
			NewInlineText("Expected string %s"),
			NewInlineText("Unexpected string %s"),
		),
	)
}
