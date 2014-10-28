package butler

type Result interface {
	Headers() map[string]string
	StatusCode() int
	Body() []byte
}
