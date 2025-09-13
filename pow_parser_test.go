package powparser

import (
	"fmt"
	"testing"

	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
)

func TestMain(m *testing.T) {
	input := "let area = side**2"
	l := lexer.New(input)
	p := parser.New(l)
	InstallPlugin(p)
	ast, err := p.ParseProgram()
	if err != nil {
		panic(fmt.Sprintf("ParseProgram() error: %v", err))
	}
	fmt.Println(ast.String())
}
