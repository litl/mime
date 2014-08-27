# mime #

mime is a simple Go package that extends the `mime` package from the
standard library with:

 * A `mime.Type` type, that represents a MIME Type
 * A `mime.DefaultExtension` function, which returns an extension given a `mime.Type`

## Example ##

```go
package main

import (
	"fmt"
	"github.com/litl/mime"
)

func main() {
	t := mime.Type("video/x-matroska")
	fmt.Printf("Type %s has default extension %s\n", t, t.DefaultExtension())
}
```
