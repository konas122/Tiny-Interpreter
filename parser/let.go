package parser

import (
	"Interpreter/ast"
	"Interpreter/token"
	"fmt"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectedPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		if p.curToken.Type == token.EOF {
			msg := fmt.Sprintf("expected the last token to be ';', got %s instead", p.peekToken.Type)
			p.errors = append(p.errors, msg)
			return nil
		}
		p.nextToken()
	}

	return stmt
}
