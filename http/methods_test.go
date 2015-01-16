package http

import (
	"testing"
	"testing/quick"

	g "github.com/SimonRichardson/butler/generic"
)

func makeMethodTest(x Method, y string) string {
	return x.Build().Run().Fst().Fold(
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
	).(string)
}

func Test_MethodDelete_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func() string {
			return "delete"
		}
		g = func() string {
			return makeMethodTest(Delete(), "delete")
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_MethodGet_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func() string {
			return "get"
		}
		g = func() string {
			return makeMethodTest(Get(), "get")
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_MethodHead_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func() string {
			return "head"
		}
		g = func() string {
			return makeMethodTest(Head(), "head")
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_MethodOptions_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func() string {
			return "options"
		}
		g = func() string {
			return makeMethodTest(Options(), "options")
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_MethodPatch_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func() string {
			return "patch"
		}
		g = func() string {
			return makeMethodTest(Patch(), "patch")
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_MethodPost_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func() string {
			return "post"
		}
		g = func() string {
			return makeMethodTest(Post(), "post")
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_MethodPut_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func() string {
			return "put"
		}
		g = func() string {
			return makeMethodTest(Put(), "put")
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_MethodTrace_WhenTestingMatchValue(t *testing.T) {
	var (
		f = func() string {
			return "trace"
		}
		g = func() string {
			return makeMethodTest(Trace(), "trace")
		}
	)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
