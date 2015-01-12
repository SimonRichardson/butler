package io

import (
	"encoding/json"

	g "github.com/SimonRichardson/butler/generic"
)

type jsonDecoder struct {
	create func() g.Any
}

func JsonDecoder(create func() g.Any) jsonDecoder {
	return jsonDecoder{
		create: create,
	}
}

func (e jsonDecoder) Keys() g.Either {
	return getAllTagsByName(e.create(), "json")
}

func (e jsonDecoder) Decode(a []byte) g.Either {
	b := e.create()
	if err := json.Unmarshal(a, &b); err != nil {
		return g.NewLeft(err)
	}
	return g.NewRight(b)
}

func (e jsonDecoder) String() string {
	return "JsonDecoder"
}
