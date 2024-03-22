package parser

import (
	"Interpreter/ast"
	"Interpreter/token"
)

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectedPeek(token.RPAREN) {
		return nil
	}
	return identifiers
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	fl := &ast.FunctionLiteral{Token: p.curToken}
	if !p.expectedPeek(token.LPAREN) {
		return nil
	}

	fl.Parameters = p.parseFunctionParameters()

	if !p.expectedPeek(token.LBRACE) {
		return nil
	}

	fl.Body = p.parseBlockStatement()
	return fl
}

// ================================================

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectedPeek(token.RPAREN) {
		return nil
	}
	return args
}

func (p *Parser) parseCallStatement(function ast.Expression) ast.Expression {
	exp := &ast.CallStatement{Token: p.curToken, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}
