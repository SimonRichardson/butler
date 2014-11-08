package main

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/http"
	"github.com/SimonRichardson/butler/output"
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

	run := func(a g.Any) g.Any {
		_, y := a.(g.Writer).Run()
		return y
	}

	s := http.Content(output.JsonEncoder{})
	fmt.Println(s.Build().ExecState("").(g.Either).Fold(run, run))

	/*
		listEmployees := Service(request, response).Then(func(args []g.Any) Result {
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
