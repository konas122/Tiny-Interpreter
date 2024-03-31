package evaluator

import (
	"Interpreter/ast"
	"Interpreter/object"
)

func DefineMacros(program *ast.Program, macroEnv *object.Environment) {
	definitions := []int{}

	for i, stmt := range program.Statements {
		if isMacroDefinition(stmt) {
			addMacro(stmt, macroEnv)
			definitions = append(definitions, i)
		}
	}

	for i := len(definitions) - 1; i >= 0; i -= 1 {
		definitionIndex := definitions[i]
		program.Statements = append(program.Statements[:definitionIndex], program.Statements[definitionIndex+1:]...)
	}
}

func isMacroDefinition(node ast.Statement) bool {
	letStatement, ok := node.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStatement.Value.(*ast.MacroLiteral)
	return ok
}

func addMacro(stmt ast.Statement, macroEnv *object.Environment) {
	letStatement, _ := stmt.(*ast.LetStatement)
	macroLiteral, _ := letStatement.Value.(*ast.MacroLiteral)

	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Env:        macroEnv,
		Body:       macroLiteral.Body,
	}
	macroEnv.Set(letStatement.Name.Value, macro)
}

// ================================================

func ExpandMacros(program ast.Node, macroEnv *object.Environment) ast.Node {
	return ast.Modify(program, func(node ast.Node) ast.Node {
		call, ok := node.(*ast.CallStatement)
		if !ok {
			return node
		}

		macro, ok := isMacroCall(call, macroEnv)
		if !ok {
			return node
		}

		args := quoteArgs(call)
		evalEnv := extendMacroEnv(macro, args)
		evaluated := Eval(macro.Body, evalEnv)
		quote, ok := evaluated.(*object.Quote)
		if !ok {
			panic("we only support returning AST-nodes from macro")
		}
		return quote.Node
	})
}

func isMacroCall(exp *ast.CallStatement, macroEnv *object.Environment) (*object.Macro, bool) {
	identifier, ok := exp.Function.(*ast.Identifier)
	if !ok {
		return nil, ok
	}
	obj, ok := macroEnv.Get(identifier.Value)
	if !ok {
		return nil, ok
	}
	macro, ok := obj.(*object.Macro)
	if !ok {
		return nil, ok
	}
	return macro, true
}

func quoteArgs(exp *ast.CallStatement) []*object.Quote {
	args := []*object.Quote{}
	for _, a := range exp.Arguments {
		args = append(args, &object.Quote{Node: a})
	}
	return args
}

func extendMacroEnv(macro *object.Macro, args []*object.Quote) *object.Environment {
	macroEnv := object.NewClosedEnvironment(macro.Env)
	for index, param := range macro.Parameters {
		macroEnv.Set(param.Value, args[index])
	}
	return macroEnv
}
