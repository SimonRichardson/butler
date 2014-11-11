package markdown

import "fmt"

type unorderedListType string

var (
	Star   unorderedListType = "*"
	Hyphen unorderedListType = "-"
)

func (t unorderedListType) String() string {
	return string(t)
}

type orderedListType string

var (
	Hash orderedListType = "#"
)

func (t orderedListType) String() string {
	return string(t)
}

type list struct {
	Type   marks
	Values []listItem
}

func (l list) String() string {
	var (
		t   = l.Type.String()
		res = ""
	)
	for _, v := range l.Values {
		res += fmt.Sprintf("%s %s\n", t, v.String())
	}
	return res
}

type listItem struct {
	Value string
}

func (l listItem) String() string {
	return l.Value
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

func li(value string) listItem {
	return listItem{
		Value: value,
	}
}
