package cpp_ast

import (
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

type SourceLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type BaseNodeInfo struct {
	Location SourceLocation `json:"location"`
}

func (m *BaseNodeInfo) SetLocation(line, column int) {
	m.Location.Line = line
	m.Location.Column = column
}

func (m *BaseNodeInfo) SetLocationFromToken(token antlr.Token) {
	m.Location.Line = token.GetLine()
	m.Location.Column = token.GetColumn()
}

type BaseNodeMethods interface {
}

func GetNodeLocation(node BaseNodeMethods) SourceLocation {
	switch v := node.(type) { //v表示b1 接口转换成Bag对象的值
	case *IdentifierNode:
		return v.Location
	case *LiteralNode:
		return v.Location
	case *AssignmentNode:
		return v.Location
	case *StatementNode:
		return v.Location
	case *CallNode:
		return v.Location
	case *ExpressionListNode:
		return v.Location
	default:
		panic("not implemented type: " + fmt.Sprintf("%+v", reflect.TypeOf(node)))
	}

}

type IdentifierNode struct {
	BaseNodeInfo
	BaseNodeMethods `json:"-"`
	Name            string `json:"name"`
}

type LiteralNode struct {
	BaseNodeInfo
	BaseNodeMethods `json:"-"`
	Values          []string `json:"value"`
}

type InitializerNode struct {
	BaseNodeInfo
	BaseNodeMethods `json:"-"`
	Values          []string `json:"value"`
}

type AssignmentNode struct {
	BaseNodeInfo
	BaseNodeMethods `json:"-"`
	Operator        string          `json:"operator"`
	LValue          BaseNodeMethods `json:"l_value"`
	RValue          BaseNodeMethods `json:"r_value"`
}

type CallNode struct {
	BaseNodeInfo
	BaseNodeMethods `json:"-"`
	Name            BaseNodeMethods     `json:"name"`
	Arguments       *ExpressionListNode `json:"arguments"`
}

type StatementNode struct {
	BaseNodeInfo
	BaseNodeMethods `json:"-"`
	Child           BaseNodeMethods `json:"child"`
}

type ExpressionListNode struct {
	BaseNodeInfo
	BaseNodeMethods `json:"-"`
	Expressions     []BaseNodeMethods `json:"expressions"`
}
