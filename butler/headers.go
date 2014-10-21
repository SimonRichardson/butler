package butler

type Header struct {
	name  Api
	value Api
}

func NewHeader(name, value string) Api {
	return NewApi(
		Header{
			name:  NewString(name),
			value: NewString(value),
		},
		NewDocTypes(
			NewInlineText("Expected header %s"),
			NewInlineText("Unexpected header %s"),
		),
	)
}

func Accept(value string) Api {
	return NewHeader("Accept", value)
}

func AcceptCharset(value string) Api {
	return NewHeader("Accept-Charset", value)
}

func AcceptEncoding(value string) Api {
	return NewHeader("Accept-Encoding", value)
}

func AcceptLanguage(value string) Api {
	return NewHeader("Accept-Language", value)
}

func AcceptDatetime(value string) Api {
	return NewHeader("Accept-Datetime", value)
}

func Authorization(value string) Api {
	return NewHeader("Authorization", value)
}

func CacheControl(value string) Api {
	return NewHeader("Cache-Control", value)
}

func Connection(value string) Api {
	return NewHeader("Connection", value)
}

func Cookie(value string) Api {
	return NewHeader("Cookie", value)
}

func ContentLength(value string) Api {
	return NewHeader("Content-Length", value)
}

func ContentMD5(value string) Api {
	return NewHeader("Content-MD5", value)
}

func ContentType(value string) Api {
	return NewHeader("Content-Type", value)
}

func Date(value string) Api {
	return NewHeader("Date", value)
}

func Expect(value string) Api {
	return NewHeader("Expect", value)
}

func From(value string) Api {
	return NewHeader("From", value)
}

func Host(value string) Api {
	return NewHeader("Host", value)
}

func IfMatch(value string) Api {
	return NewHeader("If-Match", value)
}

func IfModifiedSince(value string) Api {
	return NewHeader("If-Modified-Since", value)
}

func IfNoneMatch(value string) Api {
	return NewHeader("If-None-Match", value)
}

func IfRange(value string) Api {
	return NewHeader("If-Range", value)
}

func IfUnmodifiedSince(value string) Api {
	return NewHeader("If-Unmodified-Since", value)
}

func MaxForwards(value string) Api {
	return NewHeader("Max-Forwards", value)
}

func Origin(value string) Api {
	return NewHeader("Origin", value)
}

func Pragma(value string) Api {
	return NewHeader("Pragma", value)
}

func ProxyAuthorization(value string) Api {
	return NewHeader("Proxy-Authorization", value)
}

func Range(value string) Api {
	return NewHeader("Range", value)
}

func Referer(value string) Api {
	return NewHeader("Referer", value)
}

func TE(value string) Api {
	return NewHeader("TE", value)
}

func Upgrade(value string) Api {
	return NewHeader("Upgrade", value)
}

func UserAgent(value string) Api {
	return NewHeader("User-Agent", value)
}

func Via(value string) Api {
	return NewHeader("Via", value)
}

func Warning(value string) Api {
	return NewHeader("Warning", value)
}
