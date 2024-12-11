## Error interface
<li>Go programs express errors with error values. An Error is any type that implements the simple built-in error interface:

```bash
type error interface {
  Error() string
}
```

<li>When something can go wrong in a function, that function should return an error as its last return value. Any code that calls a function that can return an error should handle errors by testing whether the error is nil.</li>

<li>Because errors are just interfaces, it is possible to build custom types that implement the error interface. Example:</li>

```bash
type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}

func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}
```

## Error Package

```bash
var err error = errors.New("something went wrong")
```

## Panic
<li>When a function calls panic, the program crashes and prints a stack trace.</li>
<li>The panic function goes out of the flow of the current function and up to the call stack until it reaches a function that
defers a recover. If no function calls recover, the goroutine (often the entire program) crashes.
</li>

```bash
func enrichUser(userID string) User {
    user, err := getUser(userID)
    if err != nil {
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    // this panics, but the defer/recover block catches it
    // a truly astonishingly bad way to handle errors
    enrichUser("123")
}
```

<li>log.Fatal() is a nice alternative to exit your program from a unrecoverable state.</li>