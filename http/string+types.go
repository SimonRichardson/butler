package http

import g "github.com/SimonRichardson/butler/generic"

func AnyChar() func(byte) g.Either {
	return func(r byte) g.Either {
		return g.NewRight(r)
	}
}

func HeaderNameChar() func(byte) g.Either {
	return func(r byte) g.Either {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			fallthrough
		case r >= 32 && r <= 39 || r >= 94 && r <= 96:
			fallthrough
		case r == 42 || r == 43 || r == 45 || r == 46 || r == 124:
			return g.NewRight(r)
		}
		return g.NewLeft(r)
	}
}

func HeaderValueChar() func(byte) g.Either {
	return func(r byte) g.Either {
		switch {
		case r >= 32 && r <= 126:
			return g.NewRight(r)
		}
		return g.NewLeft(r)
	}
}

func PathChar() func(byte) g.Either {
	return func(r byte) g.Either {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			fallthrough
		case r == 47 || r == 58:
			return g.NewRight(r)
		}
		return g.NewLeft(r)
	}
}

func UrlChar() func(byte) g.Either {
	return func(r byte) g.Either {
		switch {
		case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
			return g.NewRight(r)
		}
		return g.NewLeft(r)
	}
}
