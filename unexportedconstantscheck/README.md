# Unexported Constant Check

This linter check that the unexported constants starts with `_`.

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

**Exception**: Unexported error values may use the prefix `err` without the underscore.

# How to use it

Install it with:

```bash
go install github.com/manuelarte/presentation-golangci-lint/unexportedconstantscheck/cmd
```

And run it with:

```bash
unexportedconstantscheck [--fix=true|false] ./...
```

- `fix`: To suggest a fix
