## Functions
<li>In Go, the variable type comes after the variable name: Here, func sub(x int, y int) int is known as the "function signature".</li>

## Type Declarations Examples
<li>x int</li>
<li>p *int</li>
<li>a [3]int</li>

## Passing Variables as Values
<li>when a variable is passed into a function, that function receives a copy of the variable. The function is unable to mutate the caller's original data.</li>

## Blank Identifier: _
<li>If you don't care about a value that's being returned from a function, you can use the underscore to completely ignore it. Unlike other programming languages like Python
where the underscore is a convention, in Go it is a real language feature.
</li>
<li>The Go compiler <strong>will throw an error</strong> if you have any unused variable declarations in your code, so you need to ignore anything you don't intend to use.</li>

## Named Return Values
<li>A return statement without arguments returns the named return values. This is known as a "naked" return.
Naked return statements should be used only in short functions. They can harm readability in longer functions.</li>

## Explicit Returns
<li>You can have named return values but still explicitly return values if wanted. Example:</li>
func getCoords() (x, y int){
  return x, y // this is explicit
}
<li>We also don't need to return the named return values, we can simply sort of override it. Example: </li>
func getCoords() (x, y int){
  return 5, 6 // this is explicit, x and y are NOT returned
}

## Guard Clause
<li>It is an early return from a function if a given condition is met</li>

## Closures
<li>A closure is a function that captures and retains access to variables from its surrounding scope, allowing it to access and modify those variables even after the scope in which they were defined has ended.</li>
<p>Example: </p>
func main() {
    // Creates a closure that captures the variable `x`
    x := 10
    increment := func() int {
        x++
        return x
    }

    fmt.Println(increment()) // prints 11
    fmt.Println(increment()) // prints 12
}

## Defer
<li>Defer is a feature in Go that allows you to guarantee the execution of a process. The deferred call's arguments
are evaluated immediately but the actual call is not executed until the surrounding function returns. It is useful when
you need to assure that something happens inside a function regardless of the outcome, specially when there's a lot of 
returns. It is important to note that the defer call happens as soon as the return happens, but the surrounding function
returns the actual value after it. Think of it like a finally on a try-catch-finally in other programming languages.
</li>

## Block Scope
<li>Go is block-scoped, unlike some other languages such as python which is function scoped. Variables defined inside
a block can only be accessed within that block (and its nested blocks).Example:</li>
package main

// scoped to the entire "main" package (basically global)
var age = 19

func sendEmail() {
// scoped to the "sendEmail" function
name := "Jon Snow"

    for i := 0; i < 5; i++ {
        // scoped to the "for" body
        email := "snow@winterfell.net"
    }
}