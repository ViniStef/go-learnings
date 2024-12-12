## Loops in Go
<li>The regular loop looks like a loop from Java. Example: you define the INITIAL (i := 0); then the CONDITION (i < 10); then the AFTER (i++);</li>

## Omitting Condition From a For Loop
<li>You can omit the condition from the for loop, making it run forever.</li>

## While Loops?
<li>There is no while loop in go, instead it is possible pass only the CONDITION to the for and that would be the equivalent. Example:</li>

```bash
plantHeight := 1
for plantHeight < 5 {
  fmt.Println("still growing! current height:", plantHeight)
  plantHeight++
}
fmt.Println("plant has grown to ", plantHeight, "inches")
```

## Break and Continue
<li>break stops the current iteration of a loop and exists the loop.</li>
<li>continue stops the current iteration of a loop and continues to the next iteration. It is a powerful way to use the guard clause pattern.</li>