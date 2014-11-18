package output

import (
	"encoding/json"
	"reflect"

	g "github.com/SimonRichardson/butler/generic"
)

type JsonEncoder struct{}

func (e JsonEncoder) Keys(a g.Any) g.Either {
	return getAllTags(a).Map(func(x g.Any) g.Any {
		return g.AsList(x).Map(func(x g.Any) g.Any {
			return x.(reflect.StructTag).Get("json")
		}).Filter(func(x g.Any) bool {
			return x.(string) != ""
		})
	})
}

func (e JsonEncoder) Encode(a g.Any) g.Either {
	return toEither(json.Marshal(a))
}
