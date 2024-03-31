package ast

import (
	"Interpreter/token"
	"bytes"
	"strings"
)

type MacroLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (m *MacroLiteral) expressionNode()      {}
func (m *MacroLiteral) TokenLiteral() string { return m.Token.Literal }
func (m *MacroLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range m.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(m.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(m.Body.String())
	return out.String()
}
