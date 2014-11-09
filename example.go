package main

import (
	"fmt"
	. "github.com/SimonRichardson/butler/butler"
	g "github.com/SimonRichardson/butler/generic"
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

	run := func(a g.Any) g.Any {
		_, y := a.(g.Writer).Run()
		return y
	}

	fmt.Println(service.Build().ExecState(g.Empty{}).(g.Either).Fold(run, run))

	/*
		listEmployees := Service(request, response).Then(func(args map[string]g.Any) g.Any {
			loadAllEmployees := func(x int) g.Any {
				return []g.Any{}
			}
			return loadAllEmployees(args["limit"].(int))
		})

		server := Compile(listEmployees)
		fmt.Println(server)

			// You can also render the server to markdown, for up to
			// date documentation
			fmt.Println(markdown.Output(server))

			// Run the documentation
			service := Remotely(server)("localhost", 80)
			service.Run()
	*/
}
