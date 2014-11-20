package http

import (
	"strconv"

	g "github.com/SimonRichardson/butler/generic"
)

func QueryInt(name string) Query {
	return NewQuery(name, QInt, func(x g.Any) g.Any {
		if y, err := strconv.Atoi(x.(string)); err != nil {
			return y
		} else {
			return -1
		}
	})
}

func QueryUint(name string) Query {
	atou := func(s string) (i uint, err error) {
		u64, err := strconv.ParseUint(s, 10, 0)
		return uint(u64), err
	}
	return NewQuery(name, QUint, func(x g.Any) g.Any {
		if y, err := atou(x.(string)); err != nil {
			return y
		} else {
			return 0
		}
	})
}

func QueryString(name string) Query {
	return NewQuery(name, QString, g.Identity())
}
