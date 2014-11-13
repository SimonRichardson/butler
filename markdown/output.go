package markdown

import (
	"github.com/SimonRichardson/butler/butler"
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
)

const (
	DefaultString string = ""
)

type marks interface {
	IsInline() bool
	Children() g.Option
	String() string
}

type empty struct {
}

func (e empty) IsInline() bool {
	return false
}

func (e empty) Children() g.Option {
	return g.Option_.Empty()
}

func (e empty) String() string {
	return ""
}

func nothing() marks {
	return empty{}
}

type Markdown struct{}

func (m Markdown) Encode(a g.Any) ([]byte, error) {
	return nil, nil
}

func Output(server butler.Server) ([]byte, error) {
	// Build the service and output it as markdown!

	doc := document(
		h1(link("Butler", "http://github.com/simonrichardson/butler")),
		hr1(),
		ul(
			ul(
				link("Nested", "link"),
				ul(
					link("Nested", "link"),
				),
			),
		),
	)

	/*
		request := server.Describe()

		getMethod(request).Map(func(x g.Any) g.Any {
			list = g.NewCons(h2(x.(http.Method).String()), list)
			return x
		})

		getRoute(request).Map(func(x g.Any) g.Any {
			list = g.NewCons(h2(x.(http.Route).String()), list)
			return x
		})
	*/
	return []byte(doc.String()), nil
}

func getMethod(x g.List) g.Option {
	return x.Find(func(a g.Any) bool {
		_, ok := a.(http.Method)
		return ok
	})
}

func getRoute(x g.List) g.Option {
	return x.Find(func(a g.Any) bool {
		_, ok := a.(http.Route)
		return ok
	})
}
