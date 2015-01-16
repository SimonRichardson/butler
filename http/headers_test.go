package http

import (
	"fmt"
	"testing"
	"testing/quick"

	g "github.com/SimonRichardson/butler/generic"
)

func makeHeaderTest(x, y string) []string {
	header := NewHeader(x, y)
	return header.Build().Run().Fst().Fold(
		fail,
		func(z g.Any) g.Any {
			matcher := AsResult(z).Matcher()
			return matcher.ExecState(fmt.Sprintf("%s: %s", x, y)).Fold(
				fail,
				func(z g.Any) g.Any {
					return g.AsTuple3(z).Trd()
				},
			)
		},
	).([]string)
}

func Test_Header_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x, y alphaNumericString) []string {
			return []string{x.String(), y.String()}
		}
		g = func(x, y alphaNumericString) []string {
			return makeHeaderTest(x.String(), y.String())
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
