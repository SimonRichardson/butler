package butler

type QueryType string

const (
	Int    QueryType = "Int"
	String QueryType = "String"
)

func NewQuery(name QueryType, value string) Api {
	return NewApi(
		Query{
			name:  name,
			value: value,
		},
		NewDocTypes(
			NewInlineText("Expected query %s"),
			NewInlineText("Unexpected query %s"),
		),
	)
}

func QueryInt(value string) Api {
	return NewQuery(Int, name)
}

func QueryString(value string) Api {
	return NewQuery(String, name)
}
