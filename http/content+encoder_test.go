package http

import (
	"testing"
	"testing/quick"

	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/io"
)

func makeContentEncoderTest(x name) string {
	content := Content(io.JsonEncoder{}, g.Constant(name{}))
	return string(content.Build().Run().Fst().Fold(
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
	).([]byte))
}

func Test_ContentEncoding_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x name) string {
			return x.String()
		}
		g = func(x name) string {
			return makeContentEncoderTest(x)
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
