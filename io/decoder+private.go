package io

import (
	"reflect"

	g "github.com/SimonRichardson/butler/generic"
)

func getAllTags(a g.Any) g.Either {
	var rec func(g.List, reflect.Type, int) g.List
	rec = func(l g.List, t reflect.Type, i int) g.List {
		if i >= t.NumField() {
			return l
		}
		var (
			field = t.Field(i)
			tuple = g.NewTuple2(field.Type, field.Tag)
		)
		return rec(g.NewCons(tuple, l), t, i+1)
	}

	elem := reflect.TypeOf(a)
	return g.Either_.FromBool(elem.Kind() == reflect.Struct, elem).Map(func(x g.Any) g.Any {
		return rec(g.List_.Empty(), elem.(reflect.Type), 0)
	})
}

func getAllTagsByName(a g.Any, b string) g.Either {
	var (
		get = func(x g.Any) g.Any {
			var (
				tuple = g.AsTuple2(x)
				tag   = tuple.Snd().(reflect.StructTag)
			)
			return g.NewTuple2(tuple.Fst(), tag.Get(b))
		}
		isEmpty = func(x g.Any) bool {
			var (
				tuple = g.AsTuple2(x)
				str   = tuple.Snd().(string)
			)
			return str != ""
		}
		stringify = func(x g.Any) g.Any {
			var (
				tuple = g.AsTuple2(x)
				a     = tuple.Fst().(reflect.Type)
				b     = tuple.Snd().(string)
			)
			return g.NewTuple2(a.String(), b)
		}
		filter = func(x g.Any) g.Any {
			return g.AsList(x).
				Map(get).
				Filter(isEmpty).
				Map(stringify)
		}
	)
	return getAllTags(a).
		Map(filter)
}
