package butler

import (
	g "github.com/SimonRichardson/butler/generic"
)

func asBuild(x g.Any) Build {
	return x.(Build)
}
