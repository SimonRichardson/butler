package markdown

import "fmt"

type Document struct {
	Type raw
}

func (d Document) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, d.Type.String(DefaultIndent))
}

func document() Document {
	return Document{
		Type: value(""),
	}
}
