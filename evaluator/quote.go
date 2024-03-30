package evaluator

import (
	"Interpreter/ast"
	"Interpreter/object"
	"Interpreter/token"
	"fmt"
)

func quote(node ast.Node, env *object.Environment) object.Object {
	node = evalUnquote(node, env)
	return &object.Quote{Node: node}
}

func evalUnquote(quoted ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(quoted, func(node ast.Node) ast.Node {
		if !isUnquote(node) {
			return node
		}
		call, ok := node.(*ast.CallStatement)
		if !ok {
			return node
		}
		if len(call.Arguments) != 1 {
			return node
		}
		unquote := Eval(call.Arguments[0], env)
		return Object2ASTNode(unquote)
	})
}

func isUnquote(node ast.Node) bool {
	call, ok := node.(*ast.CallStatement)
	if !ok {
		return false
	}
	return call.Function.TokenLiteral() == "unquote"
}

func Object2ASTNode(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Integer:
		t := token.Token{
			Type:    token.INT,
			Literal: fmt.Sprintf("%d", obj.Value),
		}
		return &ast.IntegerLiteral{Token: t, Value: obj.Value}
	case *object.Boolean:
		var t token.Token
		if obj.Value {
			t = token.Token{Type: token.TRUE, Literal: "true"}
		} else {
			t = token.Token{Type: token.FALSE, Literal: "false"}
		}
		return &ast.Boolean{Token: t, Value: obj.Value}
	case *object.String:
		t := token.Token{
			Type:    token.STRING,
			Literal: obj.Value,
		}
		return &ast.StringLiteral{Token: t, Value: obj.Value}
	case *object.Quote:
		return obj.Node
	default:
		return nil
	}
}
