package parser

import (
	"Interpreter/ast"
	"Interpreter/token"
)

func (p *Parser) parseMacroLiteral() ast.Expression {
	m := &ast.MacroLiteral{Token: p.curToken}
	if !p.expectedPeek(token.LPAREN) {
		return nil
	}
	m.Parameters = p.parseFunctionParameters()
	if !p.expectedPeek(token.LBRACE) {
		return nil
	}
	m.Body = p.parseBlockStatement()
	return m
}
