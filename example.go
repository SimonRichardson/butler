package main

import (
	"fmt"
	. "github.com/SimonRichardson/butler/butler"
)

func main() {

	request := Butler().Get().
		Path("/name/:id").
		ContentType("application/json").
		QueryInt("limit")

	fmt.Println(request.Run.Run())

	/*
		response := ContentType("application/json").
			Content(HtmlEncoder{})

		listEmployees := Service(request, response).Then(func(limit int) Result {
			return loadAllEmployees(limit)
		})

		server := Compile(listEmployees)
		service := Remotely(listEmployees)("localhost", 80)
		service.Run(server)
	*/
}
