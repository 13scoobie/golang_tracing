package main

import (
	"fmt"
	"regexp"
	"runtime"
)

// Trace Functions
func enter() {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	fnName := extractFnName.ReplaceAllString(functionObject.Name(), "$1")
	fmt.Printf("Entering %s\n", fnName)
}

func exit() {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	fnName := extractFnName.ReplaceAllString(functionObject.Name(), "$1")
	fmt.Printf("Exiting  %s\n", fnName)
}

func bar(i int) bool {
	enter()
	defer exit()
	return foo(i)
}

func foo(i int) bool {
	enter()
	defer exit()
	if i > 0 {
		bar(i - 1)
		return false
	}
	return true
}

func main() {
	enter()
	defer exit()
	foo(1)
}
