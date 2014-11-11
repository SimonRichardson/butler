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

func (h headerType) String() string {
	return string(h)
}

type header struct {
	Type  headerType
	Value string
}

func (h header) String() string {
	return fmt.Sprintf("%s %s\n", h.Type.String(), h.Value)
}

func h1(value string) header {
	return header{
		Type:  H1,
		Value: value,
	}
}

func h2(value string) header {
	return header{
		Type:  H2,
		Value: value,
	}
}

func h3(value string) header {
	return header{
		Type:  H3,
		Value: value,
	}
}

func h4(value string) header {
	return header{
		Type:  H4,
		Value: value,
	}
}

func h5(value string) header {
	return header{
		Type:  H5,
		Value: value,
	}
}

func h6(value string) header {
	return header{
		Type:  H6,
		Value: value,
	}
}
