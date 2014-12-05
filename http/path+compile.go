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
	Empty    nodeType = "empty"
)

type PathNode interface {
	Type() nodeType
}

type named struct {
	name string
}

func newNamed(name string) named {
	return named{
		name: name,
	}
}

func (n named) Type() nodeType {
	return Named
}

type variable struct {
	name string
}

func newVariable(name string) variable {
	return variable{
		name: name,
	}
}

func (n variable) Type() nodeType {
	return Variable
}

type wildcard struct{}

func newWildcard() wildcard {
	return wildcard{}
}

func (n wildcard) Type() nodeType {
	return Wildcard
}

type empty struct{}

func newEmpty() empty {
	return empty{}
}

func (n empty) Type() nodeType {
	return Empty
}

func toNode(a string) g.Either {
	switch {
	case string(a[0]) == ":":
		if len(a) > 1 {
			str := string(a[1:])
			return g.NewRight(g.NewSome(newVariable(str)))
		}
		return g.NewLeft(g.NewSome(a))
	case a == "*":
		return g.NewRight(g.NewSome(newWildcard()))
	default:
		return g.NewRight(g.NewSome(newNamed(a)))
	}
}

func stringToList(s []string) g.List {
	var rec func(g.List, []string) g.Result
	rec = func(l g.List, v []string) g.Result {
		num := len(v)
		if num < 1 {
			return g.Done(l)
		}
		return g.Cont(func() g.Result {
			return rec(g.NewCons(v[num-1], l), v[:num-1])
		})

	}
	return g.AsList(g.Trampoline(rec(g.NewNil(), s)))
}

func compilePath(a string) g.List {
	var (
		x      = stringToList(strings.Split(a, "/")).Reverse()
		option = func(a g.Any) g.Any {
			return g.Option_.FromBool(strings.TrimSpace(a.(string)) != "", a)
		}
		nodes = func(a g.Any) g.Any {
			return g.AsOption(a).Fold(
				func(a g.Any) g.Any {
					return toNode(a.(string))
				},
				func() g.Any {
					return g.NewLeft(g.NewNone())
				},
			)
		}
		nones = func(a g.Any) bool {
			return g.AsEither(a).Fold(
				g.Constant1(false),
				func(a g.Any) g.Any {
					return g.Option_.ToBool(g.AsOption(a))
				},
			).(bool)
		}
		extract = func(a g.Any) g.Any {
			return g.AsEither(a).Fold(
				g.Constant1(newEmpty()),
				func(a g.Any) g.Any {
					return g.AsOption(a).GetOrElse(g.Constant(newEmpty()))
				},
			).(PathNode)
		}
	)

	return x.
		Map(option).
		Map(nodes).
		Filter(nones).
		Map(extract)
}
