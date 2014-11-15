package butler

import (
	g "github.com/SimonRichardson/butler/generic"
)

func AsBuild(x g.Any) Build {
	return x.(Build)
}

func AsServer(x g.Any) Server {
	return x.(Server)
}
