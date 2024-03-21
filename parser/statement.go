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

// ================================================

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

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

// ================================================

func (p *Parser) parseIfStatement() ast.Expression {
	expression := &ast.IfStatement{Token: p.curToken}
	if !p.expectedPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectedPeek(token.RPAREN) || !p.expectedPeek(token.LBRACE) {
		return nil
	}
	expression.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		if !p.expectedPeek(token.LBRACE) {
			return nil
		}
		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

// ================================================

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}
