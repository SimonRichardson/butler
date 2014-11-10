package butler

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type Server struct{}

func (s Server) Run(request RemoteRequest) g.Promise {
	return g.Promise_.Of(request)
}

func Compile(x service) Server {
	run := func(a g.Any) g.Any {
		_, y := a.(g.Writer).Run()
		return y
	}
	fmt.Println(x.Build().ExecState(g.Writer_.Of(g.Empty{})).(g.Either).Fold(run, run))
	return Server{}
}
