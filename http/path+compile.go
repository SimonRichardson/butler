package http

import (
	"fmt"
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
	Match(string) bool
	Type() nodeType
	String() string
}

type named struct {
	name string
}

func newNamed(name string) named {
	return named{
		name: name,
	}
}

func (n named) Match(x string) bool {
	return n.name == x
}

func (n named) Type() nodeType {
	return Named
}

func (n named) String() string {
	return fmt.Sprintf("Named(%v)", n.name)
}

type variable struct {
	name string
}

func newVariable(name string) variable {
	return variable{
		name: name,
	}
}

func (n variable) Match(x string) bool {
	return true
}

func (n variable) Type() nodeType {
	return Variable
}

func (n variable) String() string {
	return fmt.Sprintf("Variable(%v)", n.name)
}

type wildcard struct{}

func newWildcard() wildcard {
	return wildcard{}
}

func (n wildcard) Match(x string) bool {
	return true
}

func (n wildcard) Type() nodeType {
	return Wildcard
}

func (n wildcard) String() string {
	return "Wildcard"
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

func compilePath(a string) g.Either {
	var (
		x      = g.List_.StringSliceToList(strings.Split(a, "/")).Reverse()
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
		traverse = func(a g.List) g.Either {
			// TODO : If list & either implemented traverse, this would be easy!
			var (
				x = g.NewTuple2(g.Either_.Right, g.List_.Empty())
				y = a.FoldLeft(x, func(a, b g.Any) g.Any {
					return g.AsEither(b).Fold(
						func(c g.Any) g.Any {
							var (
								x = g.AsTuple2(a)
								y = g.AsList(x.Snd())
							)
							return g.NewTuple2(g.Either_.Left, g.NewCons(c, y))
						},
						func(c g.Any) g.Any {
							var (
								x = g.AsTuple2(a)
								y = g.AsList(x.Snd())
							)
							return g.NewTuple2(x.Fst(), g.NewCons(c, y))
						},
					)
				})
				z = g.AsTuple2(y)
				f = z.Fst().(func(g.Any) g.Any)
			)
			return g.AsEither(f(z.Snd()))
		}
		program = x.
			Map(option).
			Map(nodes).
			Filter(nones)
	)

	return traverse(program)
}
