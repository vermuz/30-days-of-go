// https://golang.org/ref/spec#Function_declarations
// https://golang.org/doc/effective_go.html#functions

/*
* Go functions supports higher order functions, closures, recursions and multiple return statements
* Nested function declarations are not supported unless the nested fucntion is
* assigned to a varible
*
* the syntax of a function declaration in go looks like this:
*
*     optional    function name   supports multiple returns
*         /             /                  /
* func recevier() identifier(param) and (return) {
*     ...block
* }
*
* - receiver is used for attaching methods to a type. We will look into how methods works
* when learning about structs
*
* - if the first letter of the identifier keyword is uppercase
* then your function is exported
*
 */

package functions

import "fmt"

// function with multiple return statements
func MultipleReturns(a string) (string, bool) {
	return fmt.Sprintf("Learning %v", a), true
}

// recursion
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// closures
func Curry(a string) func(string) string {
	return func(b string) string {
		return a + b
	}
}
