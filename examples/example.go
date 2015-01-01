package main

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
	h "github.com/SimonRichardson/butler/http"
)

type User struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Age       int    `json:"age"`
}

func main() {

	x := g.Free_.Lift(g.Functor_.Either(g.Either_.Of("FUCK")))
	fmt.Println(x.Run())

	fmt.Println()
	str := h.NewString("hello", h.UrlChar())
	fmt.Println(str.Build().Run())
	str.Build().Run().Fst().Fold(
		func(x g.Any) g.Any {
			fmt.Println("FAIL > ", x)
			return x
		},
		func(x g.Any) g.Any {
			fmt.Println("WIN > ", x)
			// run the matcher
			fmt.Println("WAT > ", g.AsStateT(g.AsTuple3(x).Trd()).ExecState("hello"))
			return x
		},
	)
	fmt.Println()

	header := h.NewHeader("Accept", "fuck")
	fmt.Println(header.Build().Run())
	header.Build().Run().Fst().Fold(
		func(x g.Any) g.Any {
			fmt.Println("FAIL > ", x)
			return x
		},
		func(x g.Any) g.Any {
			fmt.Println("WIN > ", x)
			// run the matcher
			fmt.Println("WAT > ", g.AsStateT(g.AsTuple3(x).Trd()).ExecState("Accepter: fucker"))
			return x
		},
	)

	/*
		path := h.NewRoute("/user/name/:id")
		path.Build().ExecState(writer).(g.Either).Fold(
			g.Identity(),
			func(x g.Any) g.Any {
				fmt.Println("Fin > ", run(g.AsWriter(x))("???"))
				return x
			},
		)

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
