package butler

type ContentEncoder struct {
	encoder Encoder
}

func Content(encoder Encoder) Api {
	return NewApi(
		ContentEncoder{
			encoder: encoder,
		},
		NewDocTypes(
			NewInlineText("Expected content encoder %s"),
			NewInlineText("Unexpected content encoder %s"),
		),
	)
}
