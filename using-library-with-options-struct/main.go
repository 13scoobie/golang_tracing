package main

import (
	"github.com/13scoobie/trackcalltime/withoptions"
)

var Exit, Enter = trackcalltime.New(nil)

// Functions we wish to trace
func bar(i int) bool {
	defer Exit(Enter())
	return foo(i)
}

func foo(i int) bool {
	defer Exit(Enter())
	if i > 0 {
		bar(i - 1)
		return false
	}
	return true
}

func main() {
	defer Exit(Enter())
	foo(1)
}
