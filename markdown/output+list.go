package markdown

import "fmt"

type unorderedListType string

var (
	Star   unorderedListType = "*"
	Hyphen unorderedListType = "-"
)

func (t unorderedListType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(t))
}

type orderedListType string

var (
	Hash orderedListType = "#"
)

func (t orderedListType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(t))
}

type list struct {
	Type   marks
	Values []listItem
}

func (l list) String(indent string) string {
	var (
		t   = l.Type.String(indent)
		res = DefaultString
	)
	for _, v := range l.Values {
		res += fmt.Sprintf("%s %s\n", t, v.String(DefaultIndent))
	}
	return res
}

type listItem struct {
	Value raw
}

func (l listItem) String(indent string) string {
	return l.Value.String(indent)
}

func ul(values []listItem) list {
	return list{
		Type:   Hyphen,
		Values: values,
	}
}

func ol(values []listItem) list {
	return list{
		Type:   Hash,
		Values: values,
	}
}

func li(val string) listItem {
	return listItem{
		Value: value(val),
	}
}
