butler
======

Serve content in a monadic style.

```go
request := Butler().
    Get().
    Path("/name/:id").
    ContentType("application/json").
    QueryInt("limit")

response := Butler().
    ContentType("application/json").
    Content(output.HtmlEncoder{})

listEmployees := Service(request, response).Then(func(args map[string]g.Any) g.Any {
    loadAllEmployees := func(x int) g.Any {
        return []g.Any{}
    }
    return loadAllEmployees(args["limit"].(int))
})

server := Compile(listEmployees)

// You can also render the server to markdown, for up to
// date documentation
doc, _ := markdown.Output(server)
fmt.FPrintln(os.Stdout, doc)

service := Remotely(server)("localhost", 80)
service.Run()
```

### TODO

- [ ] Markdown support
- [ ] Query parameters support

#### Notes

- This is not idiomatic go and I don't care!
