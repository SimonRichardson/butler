package http

import (
	"fmt"
	"testing"
	"testing/quick"

	g "github.com/SimonRichardson/butler/generic"
)

func Test_QueryInt_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x alphaNumericString, y int) int {
			return y
		}
		g = func(x alphaNumericString, y int) int {
			query := QueryInt(x.String())
			return query.Build().Run().Fst().Fold(
				fail,
				func(z g.Any) g.Any {
					var (
						value   = fmt.Sprintf("%s=%v", x, y)
						matcher = AsResult(z).Matcher()
					)
					return matcher.ExecState(value).Fold(
						fail,
						func(z g.Any) g.Any {
							return g.AsTuple3(z).Trd()
						},
					)
				},
			).(int)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_QueryUint_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x alphaNumericString, y uint) uint {
			return y
		}
		g = func(x alphaNumericString, y uint) uint {
			query := QueryUint(x.String())
			return query.Build().Run().Fst().Fold(
				fail,
				func(z g.Any) g.Any {
					var (
						value   = fmt.Sprintf("%s=%v", x, y)
						matcher = AsResult(z).Matcher()
					)
					return matcher.ExecState(value).Fold(
						fail,
						func(z g.Any) g.Any {
							return g.AsTuple3(z).Trd()
						},
					)
				},
			).(uint)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_QueryString_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x alphaNumericString, y string) string {
			return y
		}
		g = func(x alphaNumericString, y string) string {
			query := QueryString(x.String())
			return query.Build().Run().Fst().Fold(
				fail,
				func(z g.Any) g.Any {
					var (
						value   = fmt.Sprintf("%s=%v", x, y)
						matcher = AsResult(z).Matcher()
					)
					return matcher.ExecState(value).Fold(
						fail,
						func(z g.Any) g.Any {
							return g.AsTuple3(z).Trd()
						},
					)
				},
			).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
