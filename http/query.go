package http

import (
	"github.com/SimonRichardson/butler/doc"
	"github.com/SimonRichardson/butler/generic"
)

type QueryType string

const (
	QInt    QueryType = "Int"
	QString QueryType = "String"
)

type Query struct {
	doc.Api
	name  QueryType
	value String
}

func NewQuery(name QueryType, value string) Query {
	return Query{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected query %s"),
			doc.NewInlineText("Unexpected query %s"),
		)),
		name:  name,
		value: NewString(value, UrlString()),
	}
}

func (q Query) Build() generic.State {
	return generic.State{}
}

func QueryInt(name string) Query {
	return NewQuery(QInt, name)
}

func QueryString(name string) Query {
	return NewQuery(QString, name)
}
