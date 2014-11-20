package butler

import (
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
	"github.com/SimonRichardson/butler/io"
)

func Request() req {
	return req{
		list: g.NewNil(),
	}
}

type req struct {
	list g.List
}

func (b req) List() g.List {
	return b.list
}

func (b req) add(x g.Any) req {
	return req{
		list: g.NewCons(x, b.list),
	}
}

// Body

func (b req) Body(encoder io.Decoder) req {
	return b.add(http.Body(encoder))
}

// Headers
func (b req) Accept(value string) req {
	return b.add(http.Accept(value))
}

func (b req) AcceptCharset(value string) req {
	return b.add(http.AcceptCharset(value))
}

func (b req) AcceptEncoding(value string) req {
	return b.add(http.AcceptEncoding(value))
}

func (b req) AcceptLanguage(value string) req {
	return b.add(http.AcceptLanguage(value))
}

func (b req) AcceptDatetime(value string) req {
	return b.add(http.AcceptDatetime(value))
}

func (b req) Authorization(value string) req {
	return b.add(http.Authorization(value))
}

func (b req) CacheControl(value string) req {
	return b.add(http.CacheControl(value))
}

func (b req) Connection(value string) req {
	return b.add(http.Connection(value))
}

func (b req) Cookie(value string) req {
	return b.add(http.Cookie(value))
}

func (b req) ContentLength(value string) req {
	return b.add(http.ContentLength(value))
}

func (b req) ContentMD5(value string) req {
	return b.add(http.ContentMD5(value))
}

func (b req) ContentType(value string) req {
	return b.add(http.ContentType(value))
}

func (b req) Date(value string) req {
	return b.add(http.Date(value))
}

func (b req) Expect(value string) req {
	return b.add(http.Expect(value))
}

func (b req) From(value string) req {
	return b.add(http.From(value))
}

func (b req) Host(value string) req {
	return b.add(http.Host(value))
}

func (b req) IfMatch(value string) req {
	return b.add(http.IfMatch(value))
}

func (b req) IfModifiedSince(value string) req {
	return b.add(http.IfModifiedSince(value))
}

func (b req) IfNoneMatch(value string) req {
	return b.add(http.IfNoneMatch(value))
}

func (b req) IfRange(value string) req {
	return b.add(http.IfRange(value))
}

func (b req) IfUnmodifiedSince(value string) req {
	return b.add(http.IfUnmodifiedSince(value))
}

func (b req) MaxForwards(value string) req {
	return b.add(http.MaxForwards(value))
}

func (b req) Origin(value string) req {
	return b.add(http.Origin(value))
}

func (b req) Pragma(value string) req {
	return b.add(http.Pragma(value))
}

func (b req) ProxyAuthorization(value string) req {
	return b.add(http.ProxyAuthorization(value))
}

func (b req) Range(value string) req {
	return b.add(http.Range(value))
}

func (b req) Referer(value string) req {
	return b.add(http.Referer(value))
}

func (b req) TE(value string) req {
	return b.add(http.TE(value))
}

func (b req) Upgrade(value string) req {
	return b.add(http.Upgrade(value))
}

func (b req) UserAgent(value string) req {
	return b.add(http.UserAgent(value))
}

func (b req) Via(value string) req {
	return b.add(http.Via(value))
}

func (b req) Warning(value string) req {
	return b.add(http.Warning(value))
}

// Methods

func (b req) Delete() req {
	return b.add(http.Delete())
}

func (b req) Get() req {
	return b.add(http.Get())
}

func (b req) Head() req {
	return b.add(http.Head())
}

func (b req) Options() req {
	return b.add(http.Options())
}

func (b req) Patch() req {
	return b.add(http.Patch())
}

func (b req) Post() req {
	return b.add(http.Post())
}

func (b req) Put() req {
	return b.add(http.Put())
}

func (b req) Trace() req {
	return b.add(http.Trace())
}

// Path

func (b req) Path(path string) req {
	return b.add(http.Path(path))
}

// Query

func (b req) Query(q http.RawQuery) req {
	return b.add(q)
}

func (b req) QueryInt(name string) req {
	return b.add(http.QueryInt(name))
}

func (b req) QueryUint(name string) req {
	return b.add(http.QueryUint(name))
}

func (b req) QueryString(name string) req {
	return b.add(http.QueryString(name))
}
