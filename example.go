package main

import (
	"fmt"
	. "github.com/SimonRichardson/butler/butler"
	"github.com/SimonRichardson/butler/http"
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
	service.Build()

	s := http.Path("/users/:id")
	fmt.Println(">>", s.Build().EvalState(""))

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
