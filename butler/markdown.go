package butler

import (
	g "github.com/SimonRichardson/butler/generic"
	md "github.com/SimonRichardson/butler/markdown"
)

var (
	h1     = md.H1()
	h2     = md.H2()
	h3     = md.H3()
	h4     = md.H4()
	ul     = md.List(md.Unordered())
	center = md.Center()
	doc    = md.Empty()
)

func buildTemplateHeader() md.Markup {
	return doc.Concat(
		h1(md.Link("https://github.com/simonrichardson/butler", "Butler")),
	).Concat(
		h3(md.Content("Serving you content in a monadic style.")),
	).Concat(
		md.HR1(),
	)
}

func buildTemplateFooter() md.Markup {
	return doc.Concat(
		md.HR2(),
	).Concat(
		center(md.Link("https://github.com/simonrichardson/butler", "Served by Butler")),
	)
}

func buildTemplateContent() md.Markup {
	return doc.Concat(
		ul(
			g.List_.FromArgs(
				md.Link("#routes", "Route definitions"),
			),
		),
	).Concat(
		h4(md.Content("Routes")),
	).Concat(
		md.Content("The route definitions for your service are as follows:"),
	)
}

func buildTemplateError(x g.List) md.Markup {
	return doc.Concat(
		h2(md.Content("Error:")),
	).Concat(
		md.Content("Failed to build the service. The errors are as follows:"),
	)
}

func RenderTemplate(server g.Either) g.Either {
	var (
		header = buildTemplateHeader()
		footer = buildTemplateFooter()
		render = func(x g.Any) g.Any {
			return md.Render(md.AsMarkup(x))
		}
	)
	return server.Bimap(
		func(x g.Any) g.Any {
			// Implement errors:
			return doc.Concat(header).
				Concat(footer)
		},
		func(x g.Any) g.Any {
			// Implement routes
			var (
				content = buildTemplateContent()
			)
			return doc.Concat(header).
				Concat(content).
				Concat(footer)
		},
	).Bimap(render, render)
}
