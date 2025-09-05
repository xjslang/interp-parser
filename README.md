# Interpolation Parser Plugin for XJS

This is a plugin for the XJS transpiler that simplifies template literal usage by providing a more concise syntax for variable interpolation, eliminating the need for curly braces around variables.

## Simplified Variable Interpolation
```js
let greetings = `Hello, $name $surname!`
// → `Hello, ${name} ${surname}!`
```

## Usage

```go
package main

import (
    "github.com/xjslang/xjs/lexer"
    "github.com/xjslang/xjs/parser"
    intpparser "github.com/xjslang/intp-parser"
)

func main() {
    input := `let message = \`Hello, $name!\``
    
    // Create lexer and parser
    l := lexer.New(input)
    p := parser.New(l)
    
    // Register the interpolation middleware and parse the program
    p.UseExpressionHandler(intpparser.ParseInterpolationExpression)
    ast := p.ParseProgram()
    output := ast.String()
    // → "let message = `Hello, ${name}!`"
}
```