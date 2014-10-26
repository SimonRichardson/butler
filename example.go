package main

import (
	"fmt"
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

	fmt.Println(service)

	/*
		listEmployees := Service(request, response).Then(func(limit int) Result {
			return loadAllEmployees(limit)
		})

		server := Compile(listEmployees)
		service := Remotely(listEmployees)("localhost", 80)
		service.Run(server)
	*/
}
