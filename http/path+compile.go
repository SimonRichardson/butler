package http

import (
	"strings"

	g "github.com/SimonRichardson/butler/generic"
)

type nodeType string

var (
	Named    nodeType = "named"
	Variable nodeType = "variable"
	Wildcard nodeType = "wildcard"
)

type node interface {
	Type() nodeType
}

type named struct {
	name string
}

func (n named) Type() nodeType {
	return Named
}

type variable struct {
	name string
}

func (n variable) Type() nodeType {
	return Variable
}

type wildcard struct{}

func (n wildcard) Type() nodeType {
	return Wildcard
}

func toNode(a string) node {
	return named{
		name: a,
	}
}

func stringToList(s []string) g.List {
	var rec func(g.List, []string) g.List
	rec = func(l g.List, v []string) g.List {
		num := len(v)
		if num < 1 {
			return l
		}
		return rec(g.NewCons(v[num-1], l), v[:num-1])
	}
	return rec(g.NewNil(), s)
}

func compilePath(a String) g.List {
	var (
		x     = stringToList(strings.Split(a.value, "/"))
		nodes = func(a g.Any) g.Any {
			return toNode(a.(string))
		}
	)

	return x.
		Map(nodes)
}
