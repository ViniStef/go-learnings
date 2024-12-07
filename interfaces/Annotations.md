## Interfaces in Go
<li>An interface is a collection of method signatures. A type "implements" the interface if it has methods that
match the interface's method signatures. When a type implements and interface, it can then be used as that interface type
</li>
<li>An empty interface doesn't require a type to have any methods at all, meaning every type already implements the
empty interface.
</li>

## Interface Implementation
<li>An interface implementation in go is done implicitly, the type never declares that it implements a given interface.
If an interface exists and the type has the methods defined, then the type automatically fulfills that interface.
</li>