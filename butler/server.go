package butler

import "github.com/SimonRichardson/butler/generic"

type Server struct{}

func (s Server) Run(request RemoteRequest) generic.Promise {
	return generic.Promise{}.Of(request)
}
