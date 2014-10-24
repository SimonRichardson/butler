package butler

type ContentEncoder struct {
	Api
	encoder Encoder
}

func Content(encoder Encoder) ContentEncoder {
	return ContentEncoder{
		Api: NewApi(NewDocTypes(
			NewInlineText("Expected content encoder %s"),
			NewInlineText("Unexpected content encoder %s"),
		)),
		encoder: encoder,
	}
}
