package butler

type Error struct {
	Description string
	Code        string
}

var (
	error404 = func() Error {
		return Error{
			Description: "Not Found",
			Code:        "404",
		}
	}
)
