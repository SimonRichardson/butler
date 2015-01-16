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
		f = func(x alphaLowerString) string {
			return fmt.Sprintf("/%s", x.String())
		}
		g = func(x alphaLowerString) string {
			a := makeRouteTest(fmt.Sprintf("/%s", x.String()))
			return g.AsList(a).FoldLeft("", func(a, b g.Any) g.Any {
				return fmt.Sprintf("%s/%s", a, g.AsOption(b).Fold(
					func(a g.Any) g.Any {
						return a.(named).name
					},
					g.Constant("???"),
				))
			}).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_RouteNamedPath_WithTwoLeaves_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x, y alphaLowerString) string {
			return fmt.Sprintf("/%s/%s", x.String(), y.String())
		}
		g = func(x, y alphaLowerString) string {
			a := makeRouteTest(fmt.Sprintf("/%s/%s", x.String(), y.String()))
			return g.AsList(a).FoldRight("", func(a, b g.Any) g.Any {
				return fmt.Sprintf("%s/%s", a, g.AsOption(b).Fold(
					func(a g.Any) g.Any {
						return a.(named).name
					},
					g.Constant("???"),
				))
			}).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_RouteVariablePath_WithOneVariable_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x, y alphaLowerString) string {
			return fmt.Sprintf("/%s/:id", x.String())
		}
		g = func(x, y alphaLowerString) string {
			a := makeRouteTestWithArgs(fmt.Sprintf("/%s/:id", x.String()), fmt.Sprintf("/%s/%s", x.String(), y.String()))
			return g.AsList(a).FoldRight("", func(a, b g.Any) g.Any {
				return fmt.Sprintf("%s/%s", a, g.AsOption(b).Fold(
					func(a g.Any) g.Any {
						switch AsPathNode(a).Type() {
						case Named:
							return a.(named).name
						case Variable:
							return fmt.Sprintf(":%s", a.(variable).name)
						}
						return fail(a)
					},
					g.Constant("???"),
				))
			}).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_RouteVariablePath_WithTwoVariables_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x, y, z alphaLowerString) string {
			return fmt.Sprintf("/%s/:a/:b", x.String())
		}
		g = func(x, y, z alphaLowerString) string {
			a := makeRouteTestWithArgs(fmt.Sprintf("/%s/:a/:b", x.String()), fmt.Sprintf("/%s/%s/%s", x.String(), y.String(), z.String()))
			return g.AsList(a).FoldRight("", func(a, b g.Any) g.Any {
				return fmt.Sprintf("%s/%s", a, g.AsOption(b).Fold(
					func(a g.Any) g.Any {
						switch AsPathNode(a).Type() {
						case Named:
							return a.(named).name
						case Variable:
							return fmt.Sprintf(":%s", a.(variable).name)
						}
						return fail(a)
					},
					g.Constant("???"),
				))
			}).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
