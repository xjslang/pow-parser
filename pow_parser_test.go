package powparser

import (
	"fmt"
	"testing"

	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
)

func TestMain(m *testing.T) {
	input := "let x = a ** 2"
	lb := lexer.NewBuilder()
	p := parser.NewBuilder(lb).Install(Plugin).Build(input)
	ast, err := p.ParseProgram()
	if err != nil {
		panic(fmt.Sprintf("ParseProgram() error: %v", err))
	}
	fmt.Println(ast.String())
}
