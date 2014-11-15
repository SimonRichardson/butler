package output

import (
	"encoding/xml"

	g "github.com/SimonRichardson/butler/generic"
)

type XmlEncoder struct{}

func (e XmlEncoder) Encode(a g.Any) g.Either {
	return toEither(xml.Marshal(a))
}
