# Presentation `golangci-lint`

Here I have the content to give the presentation about [golangci-lint][1]

🔗[Link to the presentation][5]

## Tools

- I am using [reveal.js][2] to create the slides.
- I am using [Jujutsu][3] as a version control.

To start the slides, inside the [./slides/](./slides/) folder, run:

```bash
npm start
```

## Go Projects

### Ast Example

Example on how the AST looks like for a small .go file

```bash
cd astexample && make t
```

### Unexported Constants Check

Exercise to implement uber style guideline [Prefix Unerxported Globals with _](https://github.com/uber-go/guide/blob/master/style.md#prefix-unexported-globals-with-_).

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
// foo.go

const (
  defaultPort = 8080
  defaultUser = "user"
)

// bar.go

func Bar() {
  defaultPort := 9090
  ...
  fmt.Println("Default port", defaultPort)

  // We will not see a compile error if the first line of
  // Bar() is deleted.
}
```

</td><td>

```go
// foo.go

const (
  _defaultPort = 8080
  _defaultUser = "user"
)
```

</td></tr>
</tbody></table>

To implement: [analyzer.go](./unexportedconstantscheck/analyzer.go)

### Custom Module

Small web app to run [custom plugin linters][4].

## TODO

- Improve slides

[1]: https://golangci-lint.run/
[2]: https://revealjs.com/
[3]: https://github.com/jj-vcs/jj
[4]: https://golangci-lint.run/plugins/module-plugins/
[5]: https://manuelarte.github.io/presentation-golangci-lint/
