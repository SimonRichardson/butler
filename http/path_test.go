package http

import (
	"fmt"
	"testing"
	"testing/quick"

	g "github.com/SimonRichardson/butler/generic"
)

func makeRouteTest(x string) g.Any {
	str := NewRoute(x)
	return str.Build().Run().Fst().Fold(
		fail,
		func(z g.Any) g.Any {
			matcher := AsResult(z).Matcher()
			return matcher.ExecState(x).Fold(
				fail,
				func(z g.Any) g.Any {
					return g.AsTuple3(z).Trd()
				},
			)
		},
	)
}

func makeRouteTestWithArgs(x string, y string) g.Any {
	str := NewRoute(x)
	return str.Build().Run().Fst().Fold(
		fail,
		func(z g.Any) g.Any {
			matcher := AsResult(z).Matcher()
			return matcher.ExecState(y).Fold(
				fail,
				func(z g.Any) g.Any {
					return g.AsTuple3(z).Trd()
				},
			)
		},
	)
}

func Test_RouteNamedPath_WithOneLeaf_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x alphaLowerString) g.Set {
			return g.Set_.Empty()
		}
		g = func(x alphaLowerString) g.Set {
			var (
				path = fmt.Sprintf("/%s", x.String())
				a    = makeRouteTestWithArgs(path, path)
			)
			return g.AsSet(a)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_RouteNamedPath_WithTwoLeaves_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x, y alphaLowerString) g.Set {
			return g.Set_.Empty()
		}
		g = func(x, y alphaLowerString) g.Set {
			var (
				path = fmt.Sprintf("/%s/%s", x.String(), y.String())
				a    = makeRouteTestWithArgs(path, path)
			)
			return g.AsSet(a)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_RouteVariablePath_WithOneVariable_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x, y alphaLowerString) g.Set {
			return g.Set_.FromMap(map[g.Any]g.Any{
				"id": y.String(),
			})
		}
		g = func(x, y alphaLowerString) g.Set {
			var (
				route = fmt.Sprintf("/%s/:id", x.String())
				path  = fmt.Sprintf("/%s/%s", x.String(), y.String())
				a     = makeRouteTestWithArgs(route, path)
			)
			return g.AsSet(a)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_RouteVariablePath_WithTwoVariables_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x, y, z alphaLowerString) g.Set {
			return g.Set_.FromMap(map[g.Any]g.Any{
				"a": y.String(),
				"b": z.String(),
			})
		}
		g = func(x, y, z alphaLowerString) g.Set {
			var (
				route = fmt.Sprintf("/%s/:a/:b", x.String())
				path  = fmt.Sprintf("/%s/%s/%s", x.String(), y.String(), z.String())
				a     = makeRouteTestWithArgs(route, path)
			)
			return g.AsSet(a)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
