package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	unexportedconstantscheck "github.com/manuelarte/presentation-golangci-lint/jsontaglint"
)

func main() {
	singlechecker.Main(unexportedconstantscheck.NewAnalyzer())
}
