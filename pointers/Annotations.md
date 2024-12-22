## Pointers
<li>A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location of where the data is stored NOT the actual data itself.</li>
<li>The * syntax defines a pointer: </li>

```bash
var p *int
```

<li>The & operator generates a pointer to its operand:</li>

```bash
myString := "hello"
myStringPtr := &myString
```