package io

import (
	"encoding/json"

	g "github.com/SimonRichardson/butler/generic"
)

type JsonDecoder struct{}

func (e JsonDecoder) Keys(a g.Any) g.Either {
	return getAllTagsByName(a, "json")
}

func (e JsonDecoder) Decode(a []byte, b g.Any) (g.Any, error) {
	err := json.Unmarshal(a, &b)
	return b, err
}
