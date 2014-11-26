package butler

import (
	"fmt"
	"net/http"

	"github.com/SimonRichardson/butler/generic"
)

type Remote struct {
	server *http.Server
}

func Remotely(server Server) func(string, int) Remote {
	return func(host string, port int) Remote {
		return Remote{
			server: &http.Server{
				Addr:    fmt.Sprint("%s:%d", host, port),
				Handler: http.HandlerFunc(handle(server)),
			},
		}
	}
}

type RemoteRequest struct {
	request *http.Request
}

func RemotelyRequest(r *http.Request) RemoteRequest {
	return RemoteRequest{
		request: r,
	}
}

func (r Remote) Run() generic.Either {
	if err := r.server.ListenAndServe(); err != nil {
		return generic.NewLeft(err)
	}
	return generic.NewRight(r)
}

func handle(server Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
