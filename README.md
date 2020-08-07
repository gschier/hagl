# HAGL (HTML Abstraction Go Library)

HAGL (HTML Abstraction Go Library) is a Go library for rendering HTML, inspired 
by [Elm](https://elm-lang.org) and [HAML](https://haml.info).

## Example

```go
package main

import (
    . "github.com/gschier/hagl"
)

func main() {
    loggedIn := true

    header := func() Node {
        return H1().Text("Hello World!")
    }

    root := Div().Children(
        // Comments
        Comment("This is a simple example"),

        // Composable elements
        header().Class("hero"),

        // Looping
        Ul().Range(3, func(i int) Node {
            return Li().Textf("Item %d", i)
        }),

        // Conditional rendering
        Switch(loggedIn).
            Case(true, func() Node {
                return A().Attr("href", "/logout").Text("Logout")
            }).
            Default(func() Node {
                return A().Attr("href", "/login").Text("Log In")
            }),
    )

    println(root.HTMLPretty())
}
```

```html
<div>
  <!-- This is a simple example -->
  <h1 class="hero">Hello World!</h1>
  <ul>
    <li>Item 0</li>
    <li>Item 1</li>
    <li>Item 2</li>
  </ul>
  <a href="/logout">Logout</a>
</div>
```
