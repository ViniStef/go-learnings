## Structs in Go
<li>It is used to represent structured data, similar to how we use dictionary or objects in other languages</li>

## Embedded Structs
<li>Embedded structs allow you to have kind of a data-only inheritance. Go doesn't support classes or inheritance in
the complete sense, but embedded structs are a way to elevate and share fields between struct definitions.Example:</li>

```bash
type car struct {
  brand string
  model string
}

type truck struct {
// "car" is embedded, so the definition of a
// "truck" now also additionally contains all
// of the fields of the car struct
car
bedSize int
}
```

```bash
lanesTruck := truck{
bedSize: 10,
car: car{
brand: "toyota",
model: "camry",
},
}

fmt.Println(lanesTruck.bedSize)

// embedded fields promoted to the top-level
// instead of lanesTruck.car.brand
fmt.Println(lanesTruck.brand)
fmt.Println(lanesTruck.model)
```

## Struct Methods
<li>Struct methods are basically functions that have a receiver. The receiver is a a special parameter that goes before
the function's name.Example:
</li>

```bash
type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
return r.width * r.height
}

var r = rect{
width: 5,
height: 10,
}

fmt.Println(r.area())
// prints 50
```

## Memory Layout
<li>Structs sit in memory in a contiguous block, with fields placed one after another as defined in the struct. The order
can have a big impact on memory usage. Example of a good one first and a bad one after:</li>

```bash
type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}
```

```bash
type stats struct {
	NumPosts uint8
	Reach    uint16
	NumLikes uint8
}
```