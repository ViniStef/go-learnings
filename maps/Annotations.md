## Maps
<li>Maps in go are similar to JavaScript objects and Python dictionaries. Maps are a data structure base on key value pairs mapping</li>
<li>The zero value of a map is nil</li>
<li>We can create a map by doing the following: </li>

```bash 
cookies := make(map[string]int)
cookies["Jenn"] = 12
cookies["Florensce"] = 19
cookies["Florensce"] = 21 // overwrites 19
```
<p>or</p>

```bash
ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
```
<li>We can use the len() function to get the total number of key value pairs</li>

## Mutations
<li>Insert an element</li>

```bash
m[key] = element
```

<li>Get an element</li>

```bash
element = m[key]
```

<li>Deleting an element</li>

```bash
delete(m, key)
```

<li>Check if a key exists</li>

```bash
element, ok := m[key]
```

<p>If key is in m, then ok is true and elem is the value as expected.

If key is not in the map, then ok is false and elem is the zero value for the map's element type.</p>

<li>It's important to note that just like in slices, when we pass a map as an argument to a function, it is passed as a reference,
meaning that when we write our code inside the function, we will be modifying the original map.
</li>

## Key Types
<li>In Go, a map's key can be of any that is comparable. The comparable types are strings, numerics, booleans, pointers, channels,
interface types and structs or arrays that contain only those types. If they cannot be compared using == such as slices,
maps and functions, they may not be used as map keys
</li>
<li>You can have a a single map with a struct key like this example: </li>

```bash
type Key struct {
    Path, Country string
}
hits := make(map[Key]int)
```

<p>to increment it, we can do this: </p>

```bash
hits[Key{"/", "vn"}]++
```

<p>and to get information: </p>

```bash
n := hits[Key{"/ref/spec", "ch"}]
```

<li>We can combine an if statement with an assignment operation to use the variables inside an if block like this:</li>

```bash
names := map[string]int{}
missingNames := []string{}

if _, ok := names["Denna"]; !ok {
    // if the key doesn't exist yet,
    // append the name to the missingNames slice
    missingNames = append(missingNames, "Denna")
}
```

## Implementing a Set with Map
<li>You can use a map and have the value be a bool. Set the map entry to true to put the value in the set and then test it
by simple indexing. Example:
</li>

```bash
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
```

## Distinguishing a non-existing value from a Zero value
<li>It is important to know sometimes if the value in a map is there because of the zero value or because it's the actual value of that entry.
If you have a numeric value that you might wish to set to 0 occasionally, that can easily be of concern when working with maps. To distinguish them,
we can discriminate with a form of multiple assignment:
</li>

```bash
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

<p>This is called the “comma ok”. If tz is present, seconds will be set appropriately and ok will be true;
if not, seconds will be set to zero and ok will be false. Here's a function that puts it together with a nice error report:</p>

```bash
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

## Nested Maps
<li>maps can contain other maps, creating a nested structure. Example: </li>

```bash
map[string]map[string]int
map[rune]map[string]int
map[int]map[string]map[string]int
```