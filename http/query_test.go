package http

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	g "github.com/SimonRichardson/butler/generic"
	h "github.com/SimonRichardson/butler/http"
)

var (
	alphaNumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

type alphaNumericString string

func (s alphaNumericString) Generate(rand *rand.Rand, size int) reflect.Value {
	var (
		rnd = rand.Intn(50) + 1
		buf = make([]rune, rnd)
		num = len(buf)
	)
	for k, _ := range buf {
		buf[k] = alphaNumeric[rand.Intn(num)]
	}
	return reflect.ValueOf(alphaNumericString(string(buf)))
}

func (s alphaNumericString) String() string {
	return string(s)
}

func fail(_ g.Any) g.Any {
	return fmt.Errorf("Fail")
}

func Test_QueryInt_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x alphaNumericString, y int) int {
			return y
		}
		g = func(x alphaNumericString, y int) int {
			query := h.QueryInt(x.String())
			return query.Build().Run().Fst().Fold(
				fail,
				func(z g.Any) g.Any {
					var (
						value   = fmt.Sprintf("%s=%v", x, y)
						matcher = h.AsResult(z).Matcher()
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
			query := h.QueryUint(x.String())
			return query.Build().Run().Fst().Fold(
				fail,
				func(z g.Any) g.Any {
					var (
						value   = fmt.Sprintf("%s=%v", x, y)
						matcher = h.AsResult(z).Matcher()
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
			query := h.QueryString(x.String())
			return query.Build().Run().Fst().Fold(
				fail,
				func(z g.Any) g.Any {
					var (
						value   = fmt.Sprintf("%s=%v", x, y)
						matcher = h.AsResult(z).Matcher()
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
