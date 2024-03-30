package ast

import (
	"fmt"
	"os"
)

type ModifierFunc func(Node) Node

func Modify(node Node, modifier ModifierFunc) Node {
	var ok bool
	switch node := node.(type) {
	case *Program:
		for i, statement := range node.Statements {
			node.Statements[i], ok = Modify(statement, modifier).(Statement)
			if !ok {
				fmt.Fprintf(os.Stderr, "[Modify Error] Statement: %s is not ast.Statement. got=%T", statement.String(), node.Statements[i])
				return nil
			}
		}
	case *ExpressionStatement:
		node.Expression, ok = Modify(node.Expression, modifier).(Expression)
		if !ok {
			fmt.Fprintf(os.Stderr, "[Modify Error] Expression: %s is not ast.Expression. got=%T", node.Expression.String(), node.Expression)
		}
	case *InfixExpression:
		node.Left, ok = Modify(node.Left, modifier).(Expression)
		if !ok {
			fmt.Fprintf(os.Stderr, "[Modify Error] Left part of InfixExpression: %s is not ast.Expression. got=%T", node.Left.String(), node.Left)
			return nil
		}
		node.Right, ok = Modify(node.Right, modifier).(Expression)
		if !ok {
			fmt.Fprintf(os.Stderr, "[Modify Error] Right part of InfixExpression: %s is not ast.Expression. got=%T", node.Right.String(), node.Right)
			return nil
		}
	case *PrefixExpression:
		node.Right, ok = Modify(node.Right, modifier).(Expression)
		if !ok {
			fmt.Fprintf(os.Stderr, "[Modify Error] Right part of PrefixExpression: %s is not ast.Expression. got=%T", node.Right.String(), node.Right)
			return nil
		}
	case *IndexExpression:
		node.Left, ok = Modify(node.Left, modifier).(Expression)
		if !ok {
			fmt.Fprintf(os.Stderr, "[Modify Error] IndexExpression: %s is not ast.Expression. got=%T", node.Left.String(), node.Left)
			return nil
		}
		node.Index, ok = Modify(node.Index, modifier).(Expression)
		if !ok {
			fmt.Fprintf(os.Stderr, "[Modify Error] Index of PrefixExpression: %s is not ast.Expression. got=%T", node.Index.String(), node.Index)
			return nil
		}
	case *IfStatement:
		node.Condition, ok = Modify(node.Condition, modifier).(Expression)
		if !ok {
			fmt.Fprintf(os.Stderr, "[Modify Error] Condition of IfStatement: %s is not ast.Expression. got=%T", node.Condition.String(), node.Condition)
			return nil
		}
		node.Consequence, ok = Modify(node.Consequence, modifier).(*BlockStatement)
		if !ok {
			fmt.Fprintf(os.Stderr, "[Modify Error] Consequence of IfStatement: %s is not ast.BlockStatement. got=%T", node.Consequence.String(), node.Consequence)
			return nil
		}
		if node.Alternative != nil {
			node.Alternative, ok = Modify(node.Alternative, modifier).(*BlockStatement)
			if !ok {
				fmt.Fprintf(os.Stderr, "[Modify Error] Alternative of IfStatement: %s is not ast.BlockStatement. got=%T", node.Alternative.String(), node.Alternative)
				return nil
			}
		}
	case *BlockStatement:
		for i, statement := range node.Statements {
			node.Statements[i], ok = Modify(statement, modifier).(Statement)
			if !ok {
				fmt.Fprintf(os.Stderr, "[Modify Error] BlockStatement[i]: %s is not ast.Statement. got=%T", node.Statements[i].String(), node.Statements[i])
				return nil
			}
		}
	case *ReturnStatement:
		node.ReturnValue, ok = Modify(node.ReturnValue, modifier).(Expression)
		if !ok {
			fmt.Fprintf(os.Stderr, "[Modify Error] ReturnValue: %s is not ast.Expression. got=%T", node.ReturnValue.String(), node.ReturnValue)
			return nil
		}
	case *LetStatement:
		node.Value, ok = Modify(node.Value, modifier).(Expression)
		if !ok {
			fmt.Fprintf(os.Stderr, "[Modify Error] LetStatement: %s is not ast.Expression. got=%T", node.Value.String(), node.Value)
			return nil
		}
	case *FunctionLiteral:
		for i, param := range node.Parameters {
			node.Parameters[i], ok = Modify(param, modifier).(*Identifier)
			if !ok {
				fmt.Fprintf(os.Stderr, "[Modify Error] Parameters[i] of FunctionLiteral: %s is not ast.Identifier. got=%T", node.Parameters[i].String(), node.Parameters[i])
				return nil
			}
		}
		node.Body, _ = Modify(node.Body, modifier).(*BlockStatement)
	case *ArrayLiteral:
		for i, element := range node.Elements {
			node.Elements[i], ok = Modify(element, modifier).(Expression)
			if !ok {
				fmt.Fprintf(os.Stderr, "[Modify Error] Elements[i] of ArrayLiteral: %s is not ast.Expression. got=%T", node.Elements[i].String(), node.Elements[i])
				return nil
			}
		}
	case *HashLiteral:
		newPairs := make(map[Expression]Expression)
		for key, val := range node.Pairs {
			newKey, ok := Modify(key, modifier).(Expression)
			if !ok {
				fmt.Fprintf(os.Stderr, "[Modify Error] newKey of HashLiteral: %s is not ast.Expression. got=%T", newKey.String(), newKey)
				return nil
			}
			newVal, ok := Modify(val, modifier).(Expression)
			if !ok {
				fmt.Fprintf(os.Stderr, "[Modify Error] newVal of HashLiteral: %s is not ast.Expression. got=%T", newVal.String(), newVal)
				return nil
			}
			newPairs[newKey] = newVal
		}
		node.Pairs = newPairs

	}
	return modifier(node)
}
