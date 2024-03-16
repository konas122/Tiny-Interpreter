package parser

import (
	"Interpreter/ast"
	"Interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 83823;
`
	l := lexer.New(input)
	p := New(l)

	program := p.PaserProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatal("PaserProgram() return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokeLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%s", s.TokeLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("s not %s. got=%s", name, letStmt.Name.TokeLiteral())
		return false
	}

	return true
}

// ================================================

func TestReturnStatement(t *testing.T) {
	input := `
return 5;
return 10;
return 993 332;
`
	l := lexer.New(input)
	p := New(l)

	program := p.PaserProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ResturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokeLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokeLiteral())
		}
	}
}

// ================================================

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
