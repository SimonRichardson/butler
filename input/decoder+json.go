package output

import (
	"encoding/json"
	"reflect"

	g "github.com/SimonRichardson/butler/generic"
)

type JsonDecoder struct{}

func (e JsonDecoder) Keys(a g.Any) g.Either {
	return getAllTags(a).Map(func(x g.Any) g.Any {
		return g.AsList(x).Map(func(x g.Any) g.Any {
			return x.(reflect.StructTag).Get("json")
		}).Filter(func(x g.Any) bool {
			return x.(string) != ""
		})
	})
}

func (e JsonDecoder) Decode(a []byte, b g.Any) (g.Any, error) {
	err := json.Unmarshal(a, &b)
	return b, err
}
