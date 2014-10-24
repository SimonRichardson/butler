package output

import (
	"encoding/xml"

	"github.com/SimonRichardson/butler/generic"
)

type XmlEncoder struct{}

func (e XmlEncoder) Encode(a generic.Any) ([]byte, error) {
	return xml.Marshal(a)
}
