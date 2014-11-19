package io

import (
	"encoding/xml"

	g "github.com/SimonRichardson/butler/generic"
)

type XmlDecoder struct{}

func (e XmlDecoder) Keys(a g.Any) g.Either {
	return getAllTagsByName(a, "xml")
}

func (e XmlDecoder) Decode(a []byte, b g.Any) (g.Any, error) {
	err := xml.Unmarshal(a, &b)
	return b, err
}
