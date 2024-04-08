package testingdemo

import (
	"antlr-universal-parser/src/models/cpp_ast"
	cpp_parser "antlr-universal-parser/src/parsers/cpp"
	"fmt"
	"path"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getLValueVariableName(stmt cpp_ast.BaseNodeMethods) string {
	node := stmt.(*cpp_ast.StatementNode).
		Child.(*cpp_ast.AssignmentNode)
	return node.LValue.(*cpp_ast.IdentifierNode).Name
}

func getRValue(stmt cpp_ast.BaseNodeMethods) cpp_ast.BaseNodeMethods {
	node := stmt.(*cpp_ast.StatementNode).
		Child.(*cpp_ast.AssignmentNode)
	return node.RValue
}

func parseAssignmentStmt(stmt cpp_ast.BaseNodeMethods) (string, string, []string) {
	node := stmt.(*cpp_ast.StatementNode).
		Child.(*cpp_ast.AssignmentNode)
	return node.LValue.(*cpp_ast.IdentifierNode).Name,
		node.Operator,
		node.RValue.(*cpp_ast.LiteralNode).Values
}

// Basic 1: Basic assignment expressions & call expressions
func TestBasic1(t *testing.T) {
	file := path.Join("./tests/assets/basic1.cpp")
	listener := cpp_parser.ParseFile(file)
	lValue, _, rValue := parseAssignmentStmt(listener.StackPeek(0))

	assert.True(t, lValue == "b")
	assert.True(t, slices.Equal(rValue, []string{"4567"}))

	lValue, _, rValue = parseAssignmentStmt(listener.StackPeek(1))
	assert.True(t, lValue == "a")
	assert.True(t, slices.Equal(rValue, []string{"123"}))

	// call expr and assignment
	assert.True(t, getLValueVariableName(listener.StackPeek(2)) == "e")
	call := getRValue(listener.StackPeek(2)).(*cpp_ast.CallNode)

	assert.True(t, call.Name.(*cpp_ast.IdentifierNode).Name == "call")
	assert.True(t, len(call.Arguments.Expressions) == 2)

	// inplace assignment expr
	lValue, op, rValue := parseAssignmentStmt(listener.StackPeek(3))
	assert.True(t, lValue == "f")
	assert.True(t, op == "+=")
	assert.True(t, slices.Equal(rValue, []string{"1"}))
}

// Basic 2: Array indexing & binary operator
func TestBasic2(t *testing.T) {
	file := path.Join("./tests/assets/basic2.cpp")
	listener := cpp_parser.ParseFile(file)
	fmt.Println("basic2\n", listener.DumpStack())
}
