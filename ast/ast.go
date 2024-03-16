package ast

import "Interpreter/token"

type Node interface {
	TokeLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressNode()
}

// ================================================

type Program struct {
	Statements []Statement
}

func (p *Program) TokeLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokeLiteral()
	} else {
		return ""
	}
}

// ================================================

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokeLiteral() string {
	return ls.Token.Literal
}

// ================================================

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokeLiteral() string {
	return rs.Token.Literal
}

// ================================================

type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) expressNode() {}
func (id *Identifier) TokeLiteral() string {
	return id.Token.Literal
}

// ================================================

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokeLiteral() string {
	return es.Token.Literal
}
