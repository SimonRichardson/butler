package io

import (
	"encoding/xml"

	g "github.com/SimonRichardson/butler/generic"
)

type XmlEncoder struct{}

func (e XmlEncoder) Encode(a g.Any) g.Either {
	return toEither(xml.MarshalIndent(a, "", "\t"))
}

func (e XmlEncoder) Keys(a g.Any) g.Either {
	return getAllTagsByName(a, "xml")
}

func (e XmlEncoder) Generate(x g.Any) g.Either {
	return generate(e)(x)
}
