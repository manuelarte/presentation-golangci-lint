package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/manuelarte/presentation-golangci-lint/unexportedconstantscheck"
)

func main() {
	singlechecker.Main(unexportedconstantscheck.NewAnalyzer())
}
