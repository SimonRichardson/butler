package output

import (
	"encoding/json"

	g "github.com/SimonRichardson/butler/generic"
)

type JsonEncoder struct{}

func (e JsonEncoder) Encode(a g.Any) g.Either {
	return toEither(json.MarshalIndent(a, "", "\t"))
}

func (e JsonEncoder) Generate(x g.Any) g.Either {
	return generate(e)(x)
}
