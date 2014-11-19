package output

import (
	"encoding/xml"
	"reflect"

	g "github.com/SimonRichardson/butler/generic"
)

type XmlDecoder struct{}

func (e XmlDecoder) Keys(a g.Any) g.Either {
	return getAllTags(a).Map(func(x g.Any) g.Any {
		return g.AsList(x).Map(func(x g.Any) g.Any {
			return x.(reflect.StructTag).Get("xml")
		}).Filter(func(x g.Any) bool {
			return x.(string) != ""
		})
	})
}

func (e XmlDecoder) Decode(a []byte, b g.Any) (g.Any, error) {
	err := xml.Unmarshal(a, &b)
	return b, err
}
