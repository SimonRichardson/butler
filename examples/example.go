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

	x := g.Free_.Of(g.Functor_.LiftEither(g.Either_.Of("FUCK")))
	fmt.Println(x.Run())

	y := g.List_.FromString("1234").FoldRight("", func(a, b g.Any) g.Any {
		return a.(string) + b.(string)
	})
	fmt.Println(">>", g.List_.FromString("1234").Zip(g.List_.FromString("abcd")), y)

	var rec func(g.Any) g.Free
	rec = func(x g.Any) g.Free {
		y := x.(int)
		if y < 10 {
			return g.Cont(func() g.Free {
				return rec(y + 1)
			})
		}
		return g.Done(y)
	}
	fmt.Println(g.Trampoline(rec(1)))

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
			fmt.Println("WAT > ", g.AsStateT(g.AsTuple3(x).Trd()).ExecState("Accept: fuck"))
			return x
		},
	)
	fmt.Println()

	path := h.NewRoute("/user/name/*/:id")
	fmt.Println(path.Build().Run())
	path.Build().Run().Fst().Fold(
		func(x g.Any) g.Any {
			fmt.Println("FAIL > ", x)
			return x
		},
		func(x g.Any) g.Any {
			fmt.Println("WIN > ", x)
			// run the matcher
			fmt.Println("WAT > ", g.AsStateT(g.AsTuple3(x).Trd()).ExecState("/user/name/_/1"))
			return x
		},
	)
	fmt.Println()

	query := h.QueryString("user_name")
	fmt.Println(query.Build().Run())
	query.Build().Run().Fst().Fold(
		func(x g.Any) g.Any {
			fmt.Println("FAIL > ", x)
			return x
		},
		func(x g.Any) g.Any {
			fmt.Println("WIN > ", x)
			// run the matcher
			fmt.Println("WAT > ", g.AsStateT(g.AsTuple3(x).Trd()).ExecState("user_name=fred"))
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
