package http

import (
	"fmt"
	"testing"
	"testing/quick"

	g "github.com/SimonRichardson/butler/generic"
)

func makeStringTest(x string, f func(byte) g.Either) g.Any {
	str := NewString(x, f)
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

func Test_StringAnyChar_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x string) string {
			return x
		}
		g = func(x string) string {
			return makeStringTest(x, AnyChar()).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StringHeaderNameChar_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x alphaString) string {
			return x.String()
		}
		g = func(x alphaString) string {
			return makeStringTest(x.String(), HeaderNameChar()).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StringHeaderValueChar_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x alphaString) string {
			return x.String()
		}
		g = func(x alphaString) string {
			return makeStringTest(x.String(), HeaderValueChar()).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StringPathChar_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x alphaLowerString) string {
			return x.String()
		}
		g = func(x alphaLowerString) string {
			return makeStringTest(x.String(), PathChar()).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StringUrlChar_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x alphaLowerString) string {
			return x.String()
		}
		g = func(x alphaLowerString) string {
			fmt.Println(x)
			return makeStringTest(x.String(), UrlChar()).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StringMethodChar_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x alphaLowerString) string {
			return x.String()
		}
		g = func(x alphaLowerString) string {
			fmt.Println(x)
			return makeStringTest(x.String(), MethodChar()).(string)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
