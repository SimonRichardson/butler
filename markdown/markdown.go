package markdown

import (
	g "github.com/SimonRichardson/butler/generic"
)

func H1() func(Markup) Markup {
	return Parent("h1", "# ", "", Finish)
}

func H2() func(Markup) Markup {
	return Parent("h2", "## ", "", Finish)
}

func H3() func(Markup) Markup {
	return Parent("h3", "### ", "", Finish)
}

func H4() func(Markup) Markup {
	return Parent("h4", "#### ", "", Finish)
}

func H5() func(Markup) Markup {
	return Parent("h5", "##### ", "", Finish)
}

func H6() func(Markup) Markup {
	return Parent("h6", "###### ", "", Finish)
}

func HR1() Markup {
	return Leaf("hr1", "======", "", Finish)
}

func HR2() Markup {
	return Leaf("hr2", "------", "", Finish)
}

func BlockQuote() func(Markup) Markup {
	return Parent("blockquote", "> ", "", Finish)
}

func Center() func(Markup) Markup {
	return Parent("center", "<- ", " ->", Finish)
}

func Block() func(Markup) Markup {
	return Parent("block", "```", "```", Start|Finish|Before|After)
}

func Inline() func(Markup) Markup {
	return Parent("inline", "`", "`", None)
}

func Ordered() func(Markup) Markup {
	return Parent("ordered", "1. ", "", After)
}

func Unordered() func(Markup) Markup {
	return Parent("unordered", "- ", "", After)
}

func Indent() func(Markup) Markup {
	return Parent("indent", "", "", Pad)
}

func Italic() func(Markup) Markup {
	return Parent("italic", "_", "_", None)
}

func Bold() func(Markup) Markup {
	return Parent("bold", "*", "*", None)
}

func Link(src, content string) Markup {
	return Parent("link-src", "[", "]", None)(Content(src)).
		Concat(Parent("link-content", "(", ")", None)(Content(content)))
}

func Image(src, alt string) Markup {
	return Parent("image-alt", "![", "]", None)(Content(alt)).
		Concat(Parent("image-src", "(", ")", None)(Content(src)))
}

func List(kind func(Markup) Markup) func(g.List) Markup {
	return func(x g.List) Markup {
		var rec func(Markup, g.List) Markup
		rec = func(x Markup, y g.List) Markup {
			return AsMarkup(y.Head().Fold(
				func(z g.Any) g.Any {
					return rec(x.Concat(kind(AsMarkup(z))), y.Tail())
				},
				g.Constant(x),
			))
		}
		return rec(Empty(), x)
	}
}
