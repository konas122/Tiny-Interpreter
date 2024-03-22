package ast

import (
	"Interpreter/token"
	"bytes"
	"strings"
)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString(" (")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// ================================================

type CallStatement struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (cs *CallStatement) expressionNode() {}
func (cs *CallStatement) TokenLiteral() string {
	return cs.Token.Literal
}

func (cs *CallStatement) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range cs.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(cs.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
