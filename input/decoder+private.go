package output

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
		return rec(g.NewCons(t.Field(i).Tag, l), t, i+1)
	}

	elem := reflect.TypeOf(a)
	return g.Either_.FromBool(elem.Kind() == reflect.Struct, elem).Map(func(x g.Any) g.Any {
		return rec(g.List_.Empty(), elem.(reflect.Type), 0)
	})
}
