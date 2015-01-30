package http

import (
	"testing"
	"testing/quick"

	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/io"
)

func makeContentDecoderTest(x string) name {
	content := Body(io.JsonDecoder(
		g.Constant(name{}),
		// Turns out this is quicker than reflection!
		func(x g.Any, y string, z g.Any) g.Any {
			n := x.(name)
			switch y {
			case "name":
				n.Name = z.(string)
			}
			return n
		},
	))
	return content.Build().Run().Fst().Fold(
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
	).(name)
}

func Test_ContentDecoding_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func(x name) name {
			return x
		}
		g = func(x name) name {
			return makeContentDecoderTest(x.String())
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
