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