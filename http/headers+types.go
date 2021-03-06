package http

func Accept(value string) Header {
	return NewHeader("Accept", value)
}

func AcceptCharset(value string) Header {
	return NewHeader("Accept-Charset", value)
}

func AcceptEncoding(value string) Header {
	return NewHeader("Accept-Encoding", value)
}

func AcceptLanguage(value string) Header {
	return NewHeader("Accept-Language", value)
}

func AcceptDatetime(value string) Header {
	return NewHeader("Accept-Datetime", value)
}

func Authorization(value string) Header {
	return NewHeader("Authorization", value)
}

func CacheControl(value string) Header {
	return NewHeader("Cache-Control", value)
}

func Connection(value string) Header {
	return NewHeader("Connection", value)
}

func Cookie(value string) Header {
	return NewHeader("Cookie", value)
}

func ContentLength(value string) Header {
	return NewHeader("Content-Length", value)
}

func ContentMD5(value string) Header {
	return NewHeader("Content-MD5", value)
}

func ContentType(value string) Header {
	return NewHeader("Content-Type", value)
}

func Date(value string) Header {
	return NewHeader("Date", value)
}

func Expect(value string) Header {
	return NewHeader("Expect", value)
}

func From(value string) Header {
	return NewHeader("From", value)
}

func Host(value string) Header {
	return NewHeader("Host", value)
}

func IfMatch(value string) Header {
	return NewHeader("If-Match", value)
}

func IfModifiedSince(value string) Header {
	return NewHeader("If-Modified-Since", value)
}

func IfNoneMatch(value string) Header {
	return NewHeader("If-None-Match", value)
}

func IfRange(value string) Header {
	return NewHeader("If-Range", value)
}

func IfUnmodifiedSince(value string) Header {
	return NewHeader("If-Unmodified-Since", value)
}

func MaxForwards(value string) Header {
	return NewHeader("Max-Forwards", value)
}

func Origin(value string) Header {
	return NewHeader("Origin", value)
}

func Pragma(value string) Header {
	return NewHeader("Pragma", value)
}

func ProxyAuthorization(value string) Header {
	return NewHeader("Proxy-Authorization", value)
}

func Range(value string) Header {
	return NewHeader("Range", value)
}

func Referer(value string) Header {
	return NewHeader("Referer", value)
}

func TE(value string) Header {
	return NewHeader("TE", value)
}

func Upgrade(value string) Header {
	return NewHeader("Upgrade", value)
}

func UserAgent(value string) Header {
	return NewHeader("User-Agent", value)
}

func Via(value string) Header {
	return NewHeader("Via", value)
}

func Warning(value string) Header {
	return NewHeader("Warning", value)
}
