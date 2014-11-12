package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type links struct {
	Name string
	Url  string
}

func (l links) Children() g.Option {
	return g.Option_.Empty()
}

func (l links) String(indent string) string {
	return fmt.Sprintf("[%s](%s)\n", l.Name, l.Url)
}

func link(name, url string) links {
	return links{
		Name: name,
		Url:  url,
	}
}
