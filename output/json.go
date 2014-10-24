package output

import (
	"encoding/json"

	"github.com/SimonRichardson/butler/generic"
)

type JsonEncoder struct{}

func (e JsonEncoder) Encode(a generic.Any) ([]byte, error) {
	return json.Marshal(a)
}
