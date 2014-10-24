package butler

import (
	"github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
	"github.com/SimonRichardson/butler/output"
)

func Butler() builder {
	return builder{}.Of(http.Http())
}

type builder struct {
	Run Writer
}

func (x builder) Of(v generic.Any) builder {
	return builder{
		Run: Writer{}.Of(v),
	}
}

func (x builder) Chain(f func(v generic.Any) builder) builder {
	return builder{
		Run: x.Run.Chain(func(y generic.Any) Writer {
			return f(y).Run
		}),
	}
}

func (x builder) Map(f func(v generic.Any) generic.Any) builder {
	return x.Chain(func(y generic.Any) builder {
		return builder{}.Of(f(y))
	})
}

func add(b builder, x generic.Any) builder {
	return b.Map(func(c generic.Any) generic.Any {
		return x
	})
}

// Content

func (b builder) Content(encoder output.Encoder) builder {
	return add(b, http.Content(encoder))
}

// Headers
func (b builder) Accept(value string) builder {
	return add(b, http.Accept(value))
}

func (b builder) AcceptCharset(value string) builder {
	return add(b, http.AcceptCharset(value))
}

func (b builder) AcceptEncoding(value string) builder {
	return add(b, http.AcceptEncoding(value))
}

func (b builder) AcceptLanguage(value string) builder {
	return add(b, http.AcceptLanguage(value))
}

func (b builder) AcceptDatetime(value string) builder {
	return add(b, http.AcceptDatetime(value))
}

func (b builder) Authorization(value string) builder {
	return add(b, http.Authorization(value))
}

func (b builder) CacheControl(value string) builder {
	return add(b, http.CacheControl(value))
}

func (b builder) Connection(value string) builder {
	return add(b, http.Connection(value))
}

func (b builder) Cookie(value string) builder {
	return add(b, http.Cookie(value))
}

func (b builder) ContentLength(value string) builder {
	return add(b, http.ContentLength(value))
}

func (b builder) ContentMD5(value string) builder {
	return add(b, http.ContentMD5(value))
}

func (b builder) ContentType(value string) builder {
	return add(b, http.ContentType(value))
}

func (b builder) Date(value string) builder {
	return add(b, http.Date(value))
}

func (b builder) Expect(value string) builder {
	return add(b, http.Expect(value))
}

func (b builder) From(value string) builder {
	return add(b, http.From(value))
}

func (b builder) Host(value string) builder {
	return add(b, http.Host(value))
}

func (b builder) IfMatch(value string) builder {
	return add(b, http.IfMatch(value))
}

func (b builder) IfModifiedSince(value string) builder {
	return add(b, http.IfModifiedSince(value))
}

func (b builder) IfNoneMatch(value string) builder {
	return add(b, http.IfNoneMatch(value))
}

func (b builder) IfRange(value string) builder {
	return add(b, http.IfRange(value))
}

func (b builder) IfUnmodifiedSince(value string) builder {
	return add(b, http.IfUnmodifiedSince(value))
}

func (b builder) MaxForwards(value string) builder {
	return add(b, http.MaxForwards(value))
}

func (b builder) Origin(value string) builder {
	return add(b, http.Origin(value))
}

func (b builder) Pragma(value string) builder {
	return add(b, http.Pragma(value))
}

func (b builder) ProxyAuthorization(value string) builder {
	return add(b, http.ProxyAuthorization(value))
}

func (b builder) Range(value string) builder {
	return add(b, http.Range(value))
}

func (b builder) Referer(value string) builder {
	return add(b, http.Referer(value))
}

func (b builder) TE(value string) builder {
	return add(b, http.TE(value))
}

func (b builder) Upgrade(value string) builder {
	return add(b, http.Upgrade(value))
}

func (b builder) UserAgent(value string) builder {
	return add(b, http.UserAgent(value))
}

func (b builder) Via(value string) builder {
	return add(b, http.Via(value))
}

func (b builder) Warning(value string) builder {
	return add(b, http.Warning(value))
}

// Methods

func (b builder) Delete() builder {
	return add(b, http.Delete())
}

func (b builder) Get() builder {
	return add(b, http.Get())
}

func (b builder) Head() builder {
	return add(b, http.Head())
}

func (b builder) Options() builder {
	return add(b, http.Options())
}

func (b builder) Patch() builder {
	return add(b, http.Patch())
}

func (b builder) Post() builder {
	return add(b, http.Post())
}

func (b builder) Put() builder {
	return add(b, http.Put())
}

func (b builder) Trace() builder {
	return add(b, http.Trace())
}

// Path

func (b builder) Path(path string) builder {
	return add(b, http.Path(path))
}

// Query

func (b builder) QueryInt(name string) builder {
	return add(b, http.QueryInt(name))
}

func (b builder) QueryString(name string) builder {
	return add(b, http.QueryString(name))
}
