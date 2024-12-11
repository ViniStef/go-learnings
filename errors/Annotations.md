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