package main

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
	md "github.com/SimonRichardson/butler/markdown"
)

type User struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Age       uint   `json:"age"`
}

func main() {

	h1 := md.H1()
	ol := md.Ordered()

	doc := md.Empty().Concat(
		h1(md.Content("Header 1")),
	).Concat(
		md.H2()(md.Content("Header 2")),
	).Concat(
		md.HR1(),
	).Concat(
		md.Block()(md.Content("Code")),
	).Concat(
		md.Empty().Concat(
			ol(md.Content("List Item 1")),
		).Concat(
			ol(md.Content("List Item 2")),
		).Concat(
			md.Indent()(
				md.Empty().Concat(
					ol(md.Content("Sub List Item 1")),
				).Concat(
					ol(md.Content("Sub List Item 2")),
				).Concat(
					md.Indent()(
						md.Empty().Concat(
							ol(md.Content("Sub Sub List Item 1")),
						).Concat(
							ol(md.Content("Sub Sub List Item 2")),
						),
					),
				),
			),
		),
	)

	fmt.Println(md.Render(doc))

	ul := md.List(md.Ordered())
	list := ul(
		g.List_.FromArgs(
			md.Content("Sub Sub List Item 1"),
			md.Content("Sub Sub List Item 2"),
		),
	)
	fmt.Println(md.Render(list))

	/*
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
			Body(io.JsonDecoder(create))

		response := b.Response().
			ContentType("application/json").
			Content(io.JsonEncoder{}, g.Constant(hint), accessor)

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
