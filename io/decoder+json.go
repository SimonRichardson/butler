package io

import (
	"encoding/json"

	g "github.com/SimonRichardson/butler/generic"
)

type jsonDecoder struct {
	create func() g.Any
	fix    func(g.Any, string, g.Any) g.Any
}

func JsonDecoder(create func() g.Any, fix func(g.Any, string, g.Any) g.Any) jsonDecoder {
	return jsonDecoder{
		create: create,
		fix:    fix,
	}
}

func (e jsonDecoder) Keys() g.Either {
	return getAllTagsByName(e.create(), "json")
}

func (e jsonDecoder) Decode(a []byte) g.Either {
	var (
		x = e.create()
		y map[string]interface{}
	)
	if err := json.Unmarshal(a, &y); err != nil {
		return g.NewLeft(err)
	}
	for k, v := range y {
		x = e.fix(x, k, v)
	}
	return g.NewRight(x)
}

func (e jsonDecoder) String() string {
	return "JsonDecoder"
}
