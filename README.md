# go-temper

HTML templating in Go

```go
package main

import (
	"fmt"
	. "github.com/gschier/go-temper"
)

func main() {
	root := Div().Children(
		H1().Text("Hello World!"),
		Ul().Children(
			Li().Text("Item 1"),
			Li().Text("Item 2"),
		),
		Button().Class("btn").Text("Click Me!"),
	)

	fmt.Println("HTML:\n\n", root.HTMLPretty())
}

package main

import (
	"fmt"
	. "github.com/gschier/go-temper"
)

func main() {
	root := Div().Children(
		H1().Text("Hello World!"),
		Ul().Children(
			Li().Text("Item 1"),
			Li().Text("Item 2"),
		),
		Button().Class("btn").Text("Click Me!"),
	)

	fmt.Println("HTML:\n\n", root.HTMLPretty())
	
	// <div>
	//  <h1>Hello World!</h1>
	//  <ul>
	//    <li>Item 1</li>
	//    <li>Item 2</li>
	//  </ul>
	//  <button class="btn">Click Me!</button>
	//  </div>
}
```
