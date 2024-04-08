package cpp_parser

import (
	cpp_ast "antlr-universal-parser/src/models/cpp_ast"
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

// Create a call node by the first two items on stack, then push
// the created call node back
func (l *CalcListener) CreateCallNode(ctx interface{ GetParser() antlr.Parser }) {
	arguments, callee := l.StackPop(), l.StackPop()
	exprList, ok := arguments.(*cpp_ast.ExpressionListNode)
	if !ok {
		fmt.Println(reflect.TypeOf(callee), reflect.TypeOf(arguments))
		panic("stack item does not match")
	}
	callNode := &cpp_ast.CallNode{Name: callee, Arguments: exprList}
	callNode.SetLocationFromToken(ctx.GetParser().GetCurrentToken())
	l.pushStack(callNode)
}

// Create a array reference by the first two items on stack, then push
// the created call node back
func (l *CalcListener) CreateArrayReferenceNode(ctx interface{ GetParser() antlr.Parser }) {
	subscript, callee := l.StackPop(), l.StackPop()
	// exprList, _ := arguments//.(*cpp_ast.ExpressionListNode)
	// if !ok {
	// 	fmt.Println(reflect.TypeOf(callee), reflect.TypeOf(arguments))
	// 	panic("stack item does not match")
	// }
	callNode := &cpp_ast.ArrayReferenceNode{Name: callee, Subscript: subscript}
	callNode.SetLocationFromToken(ctx.GetParser().GetCurrentToken())
	l.pushStack(callNode)
}

// Create a binary operator by the first two items on stack, then push
// the created call node back
func (l *CalcListener) CreateBinaryOperatorNode(ctx interface{ GetParser() antlr.Parser }, operator antlr.TerminalNode) {
	rOp, lOp := l.StackPop(), l.StackPop()
	binOpNode := &cpp_ast.BinaryOperatorNode{Operator: operator.GetText(), LOperand: lOp, ROperand: rOp}
	binOpNode.SetLocationFromToken(ctx.GetParser().GetCurrentToken())
	l.pushStack(binOpNode)
}
