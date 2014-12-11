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

	/*
		toList := func(a g.Any) g.Any {
			return g.List_.Of(a)
		}

		a := g.Tree_.FromList(g.List_.FromString("abcd")).Map(toList)
		b := g.Tree_.FromList(g.List_.FromString("ab12")).Map(toList)
		fmt.Println(g.Walker_.Combine(a, b, func(a, b g.Any) g.Option {
			x := g.AsList(a)
			y := g.AsList(b)
			if g.Option_.Equals(x.Head(), y.Head()) {
				return g.Option_.Of(x.Concat(y))
			} else {
				return g.Option_.Empty()
			}
		}))*/
}
