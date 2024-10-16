# strip-literal

Strip comments and string literals from source code (utf-8 encoding).

Supported languages:
* C
* C++
* C#
* CSS (comments)
* Go
* HTML (comments). Not supported: strip in internal JavaScript/CSS.
* JavaScript/TypeScript. Not supported: JSX syntax.
* Java
* Kotlin
* Python
* Swift

[Online demo](https://falconandy.github.io/strip-literal/) (WASM version, compiled by TinyGo 0.33)

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