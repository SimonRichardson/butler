package markdown

import (
	g "github.com/SimonRichardson/butler/generic"
)

func AsMarkup(x g.Any) Markup {
	return x.(Markup)
}
