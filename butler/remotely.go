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
				return AsServer(x).Run()
			}
			perform = func(x g.Any) g.Any {
				return g.AsIO(x).
					Map(func(f g.Any) g.Any {
					return &http.Server{
						Addr:    fmt.Sprintf("%s:%s", host, port),
						Handler: http.HandlerFunc(f.(func(w http.ResponseWriter, r *http.Request))),
					}
				})
			}
			run = func(x g.Any) g.Any {
				return g.AsIO(x).
					Map(func(a g.Any) g.Any {
					return a.(*http.Server).ListenAndServe()
				})
			}
			unsafe = func(x g.Any) g.Any {
				return g.AsIO(x).UnsafePerform()
			}
		)

		server.
			Map(getIO).
			Map(perform).
			Map(run).
			Map(unsafe)
	}
}
