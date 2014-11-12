package markdown

import "fmt"

type headerType string

var (
	H1 headerType = "#"
	H2 headerType = "##"
	H3 headerType = "###"
	H4 headerType = "####"
	H5 headerType = "#####"
	H6 headerType = "######"
)

func (h headerType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(h))
}

type header struct {
	Type  headerType
	Value raw
}

func (h header) String(indent string) string {
	return fmt.Sprintf("%s %s\n", h.Type.String(indent), h.Value.String(DefaultIndent))
}

func h1(val string) header {
	return header{
		Type:  H1,
		Value: value(val),
	}
}

func h2(val string) header {
	return header{
		Type:  H2,
		Value: value(val),
	}
}

func h3(val string) header {
	return header{
		Type:  H3,
		Value: value(val),
	}
}

func h4(val string) header {
	return header{
		Type:  H4,
		Value: value(val),
	}
}

func h5(val string) header {
	return header{
		Type:  H5,
		Value: value(val),
	}
}

func h6(val string) header {
	return header{
		Type:  H6,
		Value: value(val),
	}
}
