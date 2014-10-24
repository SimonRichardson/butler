package butler

type QueryType string

const (
	QInt    QueryType = "Int"
	QString QueryType = "String"
)

type Query struct {
	Api
	name  QueryType
	value String
}

func NewQuery(name QueryType, value string) Query {
	return Query{
		Api: NewApi(NewDocTypes(
			NewInlineText("Expected query %s"),
			NewInlineText("Unexpected query %s"),
		)),
		name:  name,
		value: NewString(value),
	}
}

func QueryInt(name string) Query {
	return NewQuery(QInt, name)
}

func QueryString(name string) Query {
	return NewQuery(QString, name)
}
