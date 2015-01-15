package doc

import (
	g "github.com/SimonRichardson/butler/generic"
)

func constant(doc Doc) func(g.Any) g.Any {
	return func(a g.Any) g.Any {
		return doc.Run(a)
	}
}

func AsApi(x g.Any) Api {
	return x.(Api)
}
