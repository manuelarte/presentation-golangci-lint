package simple

const myconstant = "myconstant" // want `unexported constant "myconstant" should be prefixed with _`

const Rate = 0

const errNotFound = "not found"

const (
	group        = "group" // want `unexported constant "group" should be prefixed with _`
	Of           = "Of"
	errConstants = "error"
)

func aFunction(input int) int {
	output := input * 2
	return output
}
