package markdown

import g "github.com/SimonRichardson/butler/generic"

type inlineBlock struct {
	values g.List
}

func (b inlineBlock) IsBlock() bool {
	return false
}

func (b inlineBlock) Children() g.Option {
	return g.Option_.Of(b.values)
}

func (b inlineBlock) String() string {
	return ""
}

func group(val ...mark) inlineBlock {
	return inlineBlock{
		values: fromMarks(val),
	}
}
