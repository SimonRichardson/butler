package http

import (
	"github.com/SimonRichardson/butler/doc"
	"github.com/SimonRichardson/butler/output"
)

type ContentEncoder struct {
	doc.Api
	encoder output.Encoder
}

func Content(encoder output.Encoder) ContentEncoder {
	return ContentEncoder{
		Api: doc.NewApi(doc.NewDocTypes(
			doc.NewInlineText("Expected content encoder %s"),
			doc.NewInlineText("Unexpected content encoder %s"),
		)),
		encoder: encoder,
	}
}
