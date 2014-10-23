package butler

type QueryType string

const (
	QInt    QueryType = "Int"
	QString QueryType = "String"
)

type Query struct {
	name  QueryType
	value Api
}

func NewQuery(name QueryType, value string) Api {
	return NewApi(
		Query{
			name:  name,
			value: NewString(value),
		},
		NewDocTypes(
			NewInlineText("Expected query %s"),
			NewInlineText("Unexpected query %s"),
		),
	)
}

func QueryInt(name string) Api {
	return NewQuery(QInt, name)
}

func QueryString(name string) Api {
	return NewQuery(QString, name)
}
