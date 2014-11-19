package main

import (
	"fmt"
	. "github.com/SimonRichardson/butler/butler"
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/io"
	"github.com/SimonRichardson/butler/markdown"
)

type User struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Age       int    `json:"age"`
}

func main() {

	hint := User{}

	request := Request().
		Post().
		Path("/name/:id").
		ContentType("application/json").
		AcceptLanguage("en").
		QueryInt("limit")

	response := Response().
		ContentType("application/json").
		Content(io.JsonEncoder{}, g.Constant(hint))

	listEmployees := Service(request, response).Then(func(args map[string]g.Any) g.Any {
		loadAllEmployees := func(x int) g.Any {
			return []g.Any{}
		}
		return loadAllEmployees(args["limit"].(int))
	})

	server := Compile(listEmployees)

	// You can also render the server to markdown, for up to
	// date documentation
	markdown.Output(server).Fold(
		func(err g.Any) g.Any {
			fmt.Println(err)
			return err
		},
		func(doc g.Any) g.Any {
			fmt.Println(doc)
			return doc
		},
	)

	/*
		// Run the documentation
		service := Remotely(server)("localhost", 80)
		service.Run()
	*/
}
