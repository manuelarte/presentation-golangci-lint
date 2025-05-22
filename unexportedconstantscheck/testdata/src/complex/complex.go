package complex

// rate the case
const rate = 0 // want `unexported constant "rate" should be prefixed with _`

const one, two = "ONE", "TWO" // want `unexported constants "ONE, TWO" should be prefixed with _`
