package butler

import (
	"fmt"
	"net/http"

	g "github.com/SimonRichardson/butler/generic"
)

func Remotely(server g.Either) func(string, string) {
	return func(host, port string) {

		var (
			getIO = func(x g.Any) g.Any {
				return AsServerWithIO(x).IO()
			}
			perform = func(x g.Any) g.Any {
				return g.AsIO(x).UnsafePerform()
			}
			run = func(x g.Any) g.Any {
				addr := fmt.Sprintf("%s:%s", host, port)
				http.ListenAndServe(addr, x.(*http.ServeMux))
				return x
			}
		)

		server.
			Map(getIO).
			Map(perform).
			Map(run)
	}
}
