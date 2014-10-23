package main

import . "github.com/SimonRichardson/butler/butler"

func main() {

	request := Get().
		Path("/name/:id").
		ContentType("application/json").
		QueryInt("limit")

	/*
		response := ContentType("application/json").
			Content(JsonEncoder{})

		listEmployees := Service(request, response).Then(func(limit int) Result {
			return loadAllEmployees(limit)
		})

		server := Compile(listEmployees)
		service := Remotely(listEmployees)("localhost", 80)
		service.Run(server)
	*/
}
