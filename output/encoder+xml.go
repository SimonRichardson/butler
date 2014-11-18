package output

import (
	"encoding/xml"
	"reflect"

	g "github.com/SimonRichardson/butler/generic"
)

type XmlEncoder struct{}

func (e XmlEncoder) Keys(a g.Any) g.Either {
	return getAllTags(a).Map(func(x g.Any) g.Any {
		return g.AsList(x).Map(func(x g.Any) g.Any {
			return x.(reflect.StructTag).Get("xml")
		}).Filter(func(x g.Any) bool {
			return x.(string) != ""
		})
	})
}

func (e XmlEncoder) Encode(a g.Any) g.Either {
	return toEither(xml.Marshal(a))
}
