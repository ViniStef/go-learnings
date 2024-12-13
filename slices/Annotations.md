## Arrays
<li>Arrays are a fixed size group of variables of the same type.</li>
<li>Example of declaring an array: </li>

```bash
var intsArray [10]int
```

<p>Or declaring an initialized literal: </p>

```bash
primes := [6]int{2, 3, 5, 7, 11, 13}
```

## Slices
<li>Slices are similar to an ArrayList in java, where you have a dynamic-sized flexible array.
</li>
<li>A slice always has an underlying array implicitly, but its also possible to create a slice on top of an array: </li>

```bash
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}
```

<li>This approach of accessing items inside of a slice is similar to how python does it! this is how it works: </li>

```bash
arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]
```

<p>lowIndex is inclusive and highIndex is exclusive. If you omit one side or both, it uses the entire array on that side of the colon.</p>
<li>The zero value of slice is nil</li>

## Make
<li>Slices can also be created by using the make function. Example: </li>

```bash
// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)
```

<p>Slices created with make will be filled with the zero value of the respective type.</p>
<li>To create a slice with a specific set of values, we can use the slice literal: </li>

```bash
mySlice := []string{"I", "love", "go"}
```

<h3>Length of the Slice</h3>

```bash
mySlice := []string{"I", "love", "go"}
fmt.Println(len(mySlice)) // 3
```

<p>It returns the current length of the slice</p>

<h3>Capacity of the Slice</h3>

```bash
mySlice := []string{"I", "love", "go"}
fmt.Println(cap(mySlice)) // 3
```

<p>It returns the maximum length of the slice before reallocation of the array is needed.</p>

<li>If either cap or len is nil, the return is 0</li>

## Variadic and Spread Operator

<li>The variadic syntax is represented by triple dots ... and it allow us to pass an arbitrary amount of final arguments to a function.
A variadic function receives the variadic arguments as a slice. However, it can be convenient to use a variadic becuase you don't need
to create a slice in order to treat it as such. Example:
</li>

```bash
func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}
```

<li>The spread operator allow us to pass a slice into a variadic function, pretty much doing the opposite. It also uses the
triple dots following the the slice in the function call. Example: </li>


```bash
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}
```

## Append 
<li>Just like in other languages, the append is used to add elements dynamically to a slice. If the underlying array is not big enough,
it will create a new one and point the returned slice to it. Its also worth pointing that the append() is variadic!
</li>

## Slices of Slices
<li>A slice can have other slices, creating a matrix. Example: </li>

```bash
rows := [][]int{}
```

<li>When using append, the original slice does not get modified if the capacity is increased. Append creates a new array when
needed and returns a new slice that points to the new array.
</li>

## Range in Go
<li>Go provides a simple helpful way to iterate over elements of a slice, it looks similar to how python does it. This is the syntax for it: </li>

```bash
for INDEX, ELEMENT := range SLICE {
}
```

<p>This allow us to access all the information we generally need out of a slice. Example: </p>

```bash
fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
// 0 apple
// 1 banana
// 2 grape
```

## Currying 
<li>Function currying is a concept from functional programming and involves partial application of functions. It allows a function with multiple arguments to be transformed into a sequence of functions, each taking a single argument.</li>