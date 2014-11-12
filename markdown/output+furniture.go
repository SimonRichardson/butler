package markdown

import "fmt"

type blockType string

var (
	HR1        blockType = "======"
	HR2        blockType = "------"
	BlockQuote blockType = ">"
)

func (b blockType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(b))
}

type block struct {
	Type  blockType
	Value raw
}

func (b block) String(indent string) string {
	switch b.Type {
	case HR1, HR2:
		return fmt.Sprintf("%s\n\n", b.Type.String(indent))
	case BlockQuote:
		return fmt.Sprintf("%s %s\n", b.Type.String(indent), b.Value.String(DefaultIndent))
	}
	return DefaultString
}

func hr1() block {
	return block{
		Type: HR1,
	}
}

func hr2() block {
	return block{
		Type: HR2,
	}
}

func blockquote(val string) block {
	return block{
		Type:  BlockQuote,
		Value: value(val),
	}
}
