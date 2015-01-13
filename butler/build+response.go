package butler

import (
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
	"github.com/SimonRichardson/butler/io"
)

func Response() res {
	return res{
		list: g.NewNil(),
	}
}

type res struct {
	list g.List
}

func (b res) List() g.List {
	return b.list
}

func (b req) add(x g.Any) req {
	return b.Extend(func(y req) g.List {
		return g.NewCons(x, y.Extract())
	})
}

func (b res) Extract() g.List {
	return b.list
}

func (b res) Extend(f func(res) g.List) res {
	return res{
		list: f(b.list),
	}
}

// Content

func (b res) Content(encoder io.Encoder, hint func() g.Any) res {
	return b.add(http.Content(encoder, hint))
}

// Headers

func (b res) Accept(value string) res {
	return b.add(http.Accept(value))
}

func (b res) AcceptCharset(value string) res {
	return b.add(http.AcceptCharset(value))
}

func (b res) AcceptEncoding(value string) res {
	return b.add(http.AcceptEncoding(value))
}

func (b res) AcceptLanguage(value string) res {
	return b.add(http.AcceptLanguage(value))
}

func (b res) AcceptDatetime(value string) res {
	return b.add(http.AcceptDatetime(value))
}

func (b res) Authorization(value string) res {
	return b.add(http.Authorization(value))
}

func (b res) CacheControl(value string) res {
	return b.add(http.CacheControl(value))
}

func (b res) Connection(value string) res {
	return b.add(http.Connection(value))
}

func (b res) Cookie(value string) res {
	return b.add(http.Cookie(value))
}

func (b res) ContentLength(value string) res {
	return b.add(http.ContentLength(value))
}

func (b res) ContentMD5(value string) res {
	return b.add(http.ContentMD5(value))
}

func (b res) ContentType(value string) res {
	return b.add(http.ContentType(value))
}

func (b res) Date(value string) res {
	return b.add(http.Date(value))
}

func (b res) Expect(value string) res {
	return b.add(http.Expect(value))
}

func (b res) From(value string) res {
	return b.add(http.From(value))
}

func (b res) Host(value string) res {
	return b.add(http.Host(value))
}

func (b res) IfMatch(value string) res {
	return b.add(http.IfMatch(value))
}

func (b res) IfModifiedSince(value string) res {
	return b.add(http.IfModifiedSince(value))
}

func (b res) IfNoneMatch(value string) res {
	return b.add(http.IfNoneMatch(value))
}

func (b res) IfRange(value string) res {
	return b.add(http.IfRange(value))
}

func (b res) IfUnmodifiedSince(value string) res {
	return b.add(http.IfUnmodifiedSince(value))
}

func (b res) MaxForwards(value string) res {
	return b.add(http.MaxForwards(value))
}

func (b res) Origin(value string) res {
	return b.add(http.Origin(value))
}

func (b res) Pragma(value string) res {
	return b.add(http.Pragma(value))
}

func (b res) ProxyAuthorization(value string) res {
	return b.add(http.ProxyAuthorization(value))
}

func (b res) Range(value string) res {
	return b.add(http.Range(value))
}

func (b res) Referer(value string) res {
	return b.add(http.Referer(value))
}

func (b res) TE(value string) res {
	return b.add(http.TE(value))
}

func (b res) Upgrade(value string) res {
	return b.add(http.Upgrade(value))
}

func (b res) UserAgent(value string) res {
	return b.add(http.UserAgent(value))
}

func (b res) Via(value string) res {
	return b.add(http.Via(value))
}

func (b res) Warning(value string) res {
	return b.add(http.Warning(value))
}
