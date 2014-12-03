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

	listEmployees := Service(request, response).Then(func(args map[string]g.Any) g.Any {
		loadAllEmployees := func(x int) g.Any {
			return []g.Any{}
		}
		return loadAllEmployees(args["limit"].(int))
	})

	server := Compile(listEmployees).Run() //.AndThen(listEmployees).Run()

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

	/*
		a := g.NewTreeNode("a",
			g.List_.To(
				g.NewTreeNode("b",
					g.List_.Of(
						g.NewTreeNode("x",
							g.List_.Of(
								g.NewTreeNode("1",
									g.List_.Empty(),
								),
							),
						),
					),
				),
				g.NewTreeNode("c",
					g.List_.To(
						g.NewTreeNode("y",
							g.List_.Empty(),
						),
					),
				),
			),
		)
	*/
	a := g.Tree_.FromList(g.List_.FromString("ab"))
	b := g.Tree_.FromList(g.List_.FromString("ab"))

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a.Merge(b))
}
