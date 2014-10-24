package main

import (
	"fmt"
	. "github.com/SimonRichardson/butler/butler"
	"github.com/SimonRichardson/butler/generic"
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

	fmt.Println(request.EvalState(generic.Empty{}))
	fmt.Println(response.EvalState(generic.Empty{}))

	/*
		listEmployees := Service(request, response).Then(func(limit int) Result {
			return loadAllEmployees(limit)
		})

		server := Compile(listEmployees)
		service := Remotely(listEmployees)("localhost", 80)
		service.Run(server)
	*/
}
