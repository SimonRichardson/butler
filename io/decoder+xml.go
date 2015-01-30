package io

import (
	"encoding/xml"

	g "github.com/SimonRichardson/butler/generic"
)

type xmlDecoder struct {
	create func() g.Any
}

func XmlDecoder(create func() g.Any) xmlDecoder {
	return xmlDecoder{
		create: create,
	}
}

func (e xmlDecoder) Keys() g.Either {
	return getAllTagsByName(e.create(), "xml")
}

func (e xmlDecoder) Decode(a []byte) g.Either {
	b := e.create()
	if err := xml.Unmarshal(a, &b); err != nil {
		return g.NewLeft(err)
	}
	return g.NewRight(b)
}

func (e xmlDecoder) String() string {
	return "XmlDecoder"
}
