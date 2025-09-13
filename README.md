# Interpolation Parser Plugin for XJS

Package powparser provides a plugin for the **XJS** parser to support the `**` power operator.

## Usage

```go
import (
    "github.com/xjslang/xjs/parser"
    "github.com/xjslang/pow-parser"
)

func main() {
    p := parser.NewParser()
    powparser.InstallPlugin(p)
    // Now the parser supports '**' as a power operator.
}
```