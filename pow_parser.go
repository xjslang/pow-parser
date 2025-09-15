package powparser

import (
	"strings"

	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
	"github.com/xjslang/xjs/token"
)

// PowExpression represents a 'power' node
type PowExpression struct {
	Token token.Token // the operator token
	Left  ast.Expression
	Right ast.Expression
}

// WriteTo tells the parser how to translate a 'power' node to JavaScript
func (pe *PowExpression) WriteTo(b *strings.Builder) {
	b.WriteString("Math.pow(")
	pe.Left.WriteTo(b)
	b.WriteRune(',')
	pe.Right.WriteTo(b)
	b.WriteRune(')')
}

func Plugin(pb *parser.Builder) {
	// registes a new token type and
	// instructs the parser who to read the new `**` token
	lb := pb.LexerBuilder
	powTokenType := lb.RegisterTokenType("pow")
	lb.UseTokenReader(func(l *lexer.Lexer, next func() token.Token) token.Token {
		if l.CurrentChar == '*' && l.PeekChar() == '*' {
			l.ReadChar() // consume the first '*'
			l.ReadChar() // consume the last '*'
			return token.Token{Type: powTokenType, Literal: "pow", Column: l.Column, Line: l.Line}
		}
		return next()
	})

	// registers the infix `**` operator with a specific precedence
	pb.RegisterInfixOperator(powTokenType, parser.PRODUCT+1, func(token token.Token, left ast.Expression, right func() ast.Expression) ast.Expression {
		return &PowExpression{
			Token: token,
			Left:  left,
			Right: right(),
		}
	})
}
