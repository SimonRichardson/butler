package http

import (
	"regexp"

	g "github.com/SimonRichardson/butler/generic"
)

var (
	headerNameChar  *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9\-!"#$%&'^*+| ]$`)
	headerValueChar *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9\-!"#$%&'^*+| /]$`)
	pathChar        *regexp.Regexp = regexp.MustCompile(`^[a-z0-9/:*]$`)
	urlChar         *regexp.Regexp = regexp.MustCompile(`^[a-z0-9]$`)
	queryChar       *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9_]$`)
)

func AnyChar() func(byte) g.Either {
	return func(b byte) g.Either {
		return g.NewRight(b)
	}
}

func HeaderNameChar() func(byte) g.Either {
	return func(b byte) g.Either {
		return g.Either_.FromBool(headerNameChar.MatchString(string(b)), b)
	}
}

func HeaderValueChar() func(byte) g.Either {
	return func(b byte) g.Either {
		return g.Either_.FromBool(headerValueChar.MatchString(string(b)), b)
	}
}

func PathChar() func(byte) g.Either {
	return func(b byte) g.Either {
		return g.Either_.FromBool(pathChar.MatchString(string(b)), b)
	}
}

func UrlChar() func(byte) g.Either {
	return func(b byte) g.Either {
		return g.Either_.FromBool(urlChar.MatchString(string(b)), b)
	}
}

func QueryChar() func(byte) g.Either {
	return func(b byte) g.Either {
		return g.Either_.FromBool(queryChar.MatchString(string(b)), b)
	}
}
