package main

import (
	"fmt"

	"github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
)

func main() {

	/*request := Butler().
		Get().
		Path("/name/:id").
		ContentType("application/json").
		QueryInt("limit")

	response := Butler().
		ContentType("application/json").
		Content(output.HtmlEncoder{})

	service := Service(request, response)
	service.Build()*/

	s := http.NewString("/na£me/:id", http.UrlChar())
	s.Build().ExecState("").(generic.Either).Fold(generic.Identity(), func(a generic.Any) generic.Any {
		_, y := a.(generic.Writer).Run()
		fmt.Println(y)
		return y
	})

	/*
		listEmployees := Service(request, response).Then(func(args []generic.Any) Result {
			return loadAllEmployees(args[0].(int))
		})

		server := Compile(listEmployees)

		// You can also render the server to markdown, for up to
		// date documentation
		fmt.Println(markdown.Output(server))

		// Run the documentation
		service := Remotely(server)("localhost", 80)
		service.Run()
	*/
}
