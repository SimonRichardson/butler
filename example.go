package main

import (
	"fmt"

	. "github.com/SimonRichardson/butler/butler"
)

func main() {
	fmt.Println(ContentType("application/json"))
	/*
		request := GET().
			Path("/name/:id").
			ContentType("application/json").
			QueryInt("limit")

		response := ContentType("application/json").
			Content(JsonMarshall)

		listEmployees := Service(request, response).Then(func(limit int) Result {
			return loadAllEmployees(limit).ToJson()
		})

		server := Compile(listEmployees)
		service := Remotely(listEmployees)("localhost", 80)
		service.Run(server)
	*/
}
