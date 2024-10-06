# strip-literal

Strip comments and string literals from code.

Supported languages:
* Go

## Usage

```go
package main

import (
	"github.com/falconandy/strip-literal"
)

func main() {
	code := []byte("GO CODE")
	count := strip.StripLiterals(code, strip.LangGo, strip.Options{
		Comments: strip.Remove,
		Strings:  strip.RuneToSpace,
	})
	println(string(code[:count]))
}
```