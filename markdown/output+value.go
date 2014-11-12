package markdown

import "fmt"

type raw struct {
	Value string
}

func (r raw) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, r.Value)
}

func value(str string) raw {
	return raw{
		Value: str,
	}
}
