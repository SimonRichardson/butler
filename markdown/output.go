package markdown

import "fmt"

func Render(md Markup) string {
	return renderMarkup(md, "")
}

func renderMarkup(md Markup, buf string) string {
	var rec func(Markup, int, string) string
	rec = func(md Markup, indent int, buf string) string {
		if x, ok := md.(parent); ok {
			var (
				y = renderBlock(x.kind)
				z = updateIndent(y.indent, indent)
				p = renderPad(indent)
			)
			return fmt.Sprintf("%s%s%s%s%s%s%s%s%s", buf, y.start, p, x.start, y.before, rec(x.value, z, buf), y.after, x.finish, y.finish)
		} else if x, ok := md.(leaf); ok {
			var (
				y = renderBlock(x.kind)
				p = renderPad(indent)
			)
			return fmt.Sprintf("%s%s%s%s%s%s%s%s", buf, y.start, p, x.start, y.before, y.after, x.finish, y.finish)
		} else if x, ok := md.(content); ok {
			return fmt.Sprintf("%s%s", buf, x.value)
		} else if x, ok := md.(concat); ok {
			return fmt.Sprintf("%s%s", rec(x.left, indent, buf), rec(x.right, indent, buf))
		}
		return buf
	}
	return rec(md, 0, buf)
}

type bloc struct {
	start  string
	finish string
	before string
	after  string
	indent bool
}

func renderBlock(kind BlockKind) bloc {
	b := bloc{}
	if (kind & Start) == Start {
		b.start = "\n"
	}
	if (kind & Finish) == Finish {
		b.finish = "\n"
	}
	if (kind & Before) == Before {
		b.before = "\n"
	}
	if (kind & After) == After {
		b.after = "\n"
	}
	if (kind & Pad) == Pad {
		b.indent = true
		b.before = "\n"
	}
	return b
}

func updateIndent(a bool, b int) int {
	if a {
		return b + 1
	}
	return b
}

func renderPad(a int) string {
	res := ""
	for i := 0; i < a; i++ {
		res += "\t"
	}
	return res
}
