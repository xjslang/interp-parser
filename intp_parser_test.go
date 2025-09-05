package intpparser

import (
	"testing"

	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
)

func TestStringInterpolation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic", "let fullName = `$name $surname`", "let fullName=`${name} ${surname}`"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := parser.New(l)
			p.UseExpressionHandler(ParseTemplateExpression)
			ast := p.ParseProgram()
			if ast.String() != tt.expected {
				t.Errorf("Parse(%q) got %q, want %q", tt.input, ast.String(), tt.expected)
				return
			}
		})
	}
}
