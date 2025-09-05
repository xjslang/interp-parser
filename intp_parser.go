package intpparser

import (
	"regexp"

	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/parser"
	"github.com/xjslang/xjs/token"
)

func ParseTemplateExpression(p *parser.Parser, next func() ast.Expression) ast.Expression {
	if p.CurrentToken.Type == token.RAW_STRING {
		r := regexp.MustCompile(`\$(\w+)`)
		p.CurrentToken.Literal = r.ReplaceAllString(p.CurrentToken.Literal, "${$1}")
	}
	return next()
}
