package main

import (
	. "github.com/SimonRichardson/butler/butler"
	"github.com/SimonRichardson/butler/output"
)

func main() {

	request := Butler().
		Get().
		Path("/name/:id").
		ContentType("application/json").
		QueryInt("limit")

	response := Butler().
		ContentType("application/json").
		Content(output.HtmlEncoder{})

	service := Service(request, response)

	/*
		listEmployees := Service(request, response).Then(func(args []g.Any) Result {
			return loadAllEmployees(args[0].(int))
		})
	*/

	/*
		server := Compile(listEmployees)

		// You can also render the server to markdown, for up to
		// date documentation
		fmt.Println(markdown.Output(server))

		// Run the documentation
		service := Remotely(server)("localhost", 80)
		service.Run()
	*/
}
