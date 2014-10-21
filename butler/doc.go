package butler

type DocType string

var (
	InlineText DocType = "inline"
)

type DocTypes struct {
	expected   Doc
	unexpected Doc
}

func NewDocTypes(x, y Doc) DocTypes {
	return DocTypes{
		expected:   x,
		unexpected: y,
	}
}

type Doc struct {
	message string
	doc     DocType
}

func NewInlineText(message string) Doc {
	return Doc{
		message: message,
		doc:     InlineText,
	}
}
