package http

import (
	"strconv"

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
	build func(generic.Any) generic.Any
}

func NewQuery(name QueryType, value string, build func(generic.Any) generic.Any) Query {
	return Query{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected query `%s`"),
			doc.NewInlineText("Unexpected query `%s`"),
		)),
		name:  name,
		value: NewString(value, UrlChar()),
		build: build,
	}
}

func QueryInt(name string) Query {
	return NewQuery(QInt, name, func(x generic.Any) generic.Any {
		y, _ := strconv.Atoi(x.(string))
		return y
	})
}

func QueryString(name string) Query {
	return NewQuery(QString, name, generic.Identity())
}
