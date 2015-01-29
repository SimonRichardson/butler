package main

import (
	"fmt"

	b "github.com/SimonRichardson/butler/butler"
	g "github.com/SimonRichardson/butler/generic"
	"github.com/SimonRichardson/butler/io"
)

type User struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Age       uint   `json:"age"`
}

func main() {

	var (
		hint   = User{}
		create = func() g.Any {
			return User{}
		}
		accessor = func(x g.Any, prop string, val g.Any) g.Any {
			u := x.(User)
			switch prop {
			case "first-name":
				u.FirstName = val.(string)
			case "last-name":
				u.LastName = val.(string)
			case "age":
				u.Age = val.(uint)
			}
			return u
		}
	)

	request := b.Request().
		Post().
		Path("/name/:id").
		ContentType("application/json").
		AcceptLanguage("en").
		QueryUint("offset").
		QueryUint("limit").
		Body(io.JsonDecoder(create, accessor))

	response := b.Response().
		ContentType("application/json").
		Content(io.JsonEncoder{}, g.Constant(hint))

	fmt.Println(request, response)

	/*
		listEmployees := Service(request, response).Then(func() g.Any {
			loadAllEmployees := func() g.Any {
				return []g.Any{}
			}
			return loadAllEmployees()
		})

		fmt.Println(listEmployees)

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
