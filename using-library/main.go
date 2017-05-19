package main

import (
	"github.com/13scoobie/trackcalltime/withoutoptions"
)

// Functions we wish to trace
func bar(i int) bool {
	defer trackcalltime.Exit(trackcalltime.Enter())
	return foo(i)
}

func foo(i int) bool {
	defer trackcalltime.Exit(trackcalltime.Enter())
	if i > 0 {
		bar(i - 1)
		return false
	}
	return true
}

func main() {
	defer trackcalltime.Exit(trackcalltime.Enter())
	foo(1)
}
