package http

import "github.com/SimonRichardson/butler/doc"

type http struct {
	doc.Api
}

func Http() http {
	return http{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected http %s"),
			doc.NewInlineText("Unexpected http %s"),
		)),
	}
}
