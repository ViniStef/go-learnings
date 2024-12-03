## Fast and Compiled

<li>Go compiles the code directly to machine code, it is generally faster than interpreted languages</li>
<li>Go programs don't run quite as fast as other compiled languages such as Rust, Zig or C. However, Go does compile much
faster than they do.
</li>
<br>

<hr>

## Type Sizes

<h3>The size represents how many bits in memory will be used to store the variable
</h3>

*<strong>Whole Numbers (no decimal)</strong>*
<li>int int8 int16 int32 int64</li>
<br>

*<strong>Positive Whole Numbers (no decimal)</strong>*
<br>
<em>"uint" stands for "unsigned integer".</em>
<li>uint uint8 uint16 uint32 uint64 uintptr</li>
<br>

*<strong>Signed Decimal Numbers</strong>*
<li>float32 float64</li>
<br>

*<strong>Imaginary Numbers (rarely used)</strong>*
<li>complex64 complex128</li>
<br>
<hr>

## Computed Constants
<li>Constants cannot be created using the walrus operator</li>
<li>Constants must be known at runtime, requiring an initial value</li>
<li>However, constants can be computed as long as the computation can happen at compile time</li>
<li>Cannot declare a constant that can only be computed at run-time like in javascript. breaking example:<br>
<strong>the current time can only be known when the program is running</strong><br> 
<p style="color: red">const currentTime = time.Now()</p>
</li>
<br>
<hr>

## Go's Speed
<li>Go is generally faster than interpreted languages or VM powered languages, but slighly slower than other compiled 
languages, that is mostly due to its automated memory management, also known as "Go runtime".</li>
<br>
<hr>

## Formatting Strings
<li>fmt.Printf - Prints a formatted string to standard out</li>
<li>fmt.Sprintf() - Returns the formatted string</li>
<strong>Default Representation</strong>
<li>The <span style="color: teal">%v</span> variant represents the Go syntax representation of a value</li>
To be more specific:<br>
<strong>String</strong>
<li>s := fmt.Sprintf("I am <span style="color: teal">%s</span> years old", "way too many")</li>
<strong>Integer</strong>
<li>s := fmt.Sprintf("I am <span style="color: teal">%d</span> years old", 10)</li>
<strong>Float</strong>
<li>s := fmt.Sprintf("I am <span style="color: teal">%f</span> years old", 10.523)</li>
<li>s := fmt.Sprintf("I am <span style="color: teal">%.2f</span> years old", 10.523)</li>
<em>The ".2" rounds the number to 2 decimal places</em>
<li>You can use <span style="color: teal">%t</span> for boolean values</li>
<br>
<hr>

## Runes and String encoding
<li>Rune is an alias for int32, meaning a rune is a 32-bit integer, which is large enough to hold any Unicode code point</li>
<li>If I need to work with individual characters in a string, I should use the rune type. It breaks strings up
into their individual characters, which can be more than one byte long.</li>
<li>We can use zany characters like emojis and Chinese characters in our strings, and Go would handle them just fine</li>


