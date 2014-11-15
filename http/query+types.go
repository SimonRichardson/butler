package http

import (
	"strconv"

	g "github.com/SimonRichardson/butler/generic"
)

func QueryInt(name string) Query {
	return NewQuery(name, func(x g.Any) g.Any {
		if y, err := strconv.Atoi(x.(string)); err != nil {
			return y
		} else {
			return -1
		}
	})
}

func QueryString(name string) Query {
	return NewQuery(name, g.Identity())
}
