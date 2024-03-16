package parser

import (
	"Interpreter/ast"
	"Interpreter/token"
	"fmt"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		if p.curToken.Type == token.EOF {
			msg := fmt.Sprintf("expected the last token to be ';', got %s instead", p.peekToken.Type)
			p.errors = append(p.errors, msg)
			break
		}
		p.nextToken()
	}
	return stmt
}
