package main

import (

	//. "github.com/SimonRichardson/butler/butler"

	"fmt"
	"net/http"

	g "github.com/SimonRichardson/butler/generic"
	h "github.com/SimonRichardson/butler/http"
)

type User struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Age       int    `json:"age"`
}

func main() {

	writer := g.Writer_.Of(g.Empty{})

	header := h.NewHeader("Accept", "fuck")
	header.Build().ExecState(writer).(g.Either).Fold(
		func(x g.Any) g.Any {
			return x
		},
		func(x g.Any) g.Any {

			header := make(http.Header)
			header.Add("Accept", "fuck")

			set := g.Set_.HttpHeaderToSet(header)

			fmt.Println("Fin > ", x.(g.Writer).Run().Fst().(g.StateT).ExecState(set))
			return x
		},
	)
	/*
		var (
			hint   = User{}
			create = func() g.Any {
				return User{}
			}
		)

		request := Request().
			Post().
			Path("/name/:id").
			ContentType("application/json").
			AcceptLanguage("en").
			QueryUint("offset").
			QueryUint("limit").
			Body(io.JsonDecoder(create))

		response := Response().
			ContentType("application/json").
			Content(io.JsonEncoder{}, g.Constant(hint))

		listEmployees := Service(request, response).Then(func() g.Any {
			loadAllEmployees := func() g.Any {
				return []g.Any{}
			}
			return loadAllEmployees()
		})

		server := Compile(listEmployees).AndThen(listEmployees).Run()

		// You can also render the server to markdown, for up to
		// date documentation
		markdown.Output(server).Fold(
			func(err g.Any) g.Any {
				fmt.Println(err)
				return err
			},
			func(doc g.Any) g.Any {
				// fmt.Println(doc)
				return doc
			},
		)

		// Run the documentation
		Remotely(server)("localhost", "8080")
	*/
}
