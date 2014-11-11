package markdown

import "fmt"

type blockType string

var (
	HR1        blockType = "======"
	HR2        blockType = "------"
	BlockQuote blockType = ">"
)

func (f blockType) String() string {
	return string(f)
}

type block struct {
	Type  blockType
	Value string
}

func (b block) String() string {
	switch b.Type {
	case HR1, HR2:
		return fmt.Sprintf("%s\n\n", b.Type.String())
	case BlockQuote:
		return fmt.Sprintf("%s %s\n", b.Type.String(), b.Value)
	}
	return ""
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

func blockquote(value string) block {
	return block{
		Type:  BlockQuote,
		Value: value,
	}
}
