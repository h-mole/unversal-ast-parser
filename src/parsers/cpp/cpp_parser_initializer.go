package cpp_parser

import (
	"antlr-universal-parser/src/models/cpp_ast"
	"antlr-universal-parser/src/parsers/cpp/parser"
	"fmt"
)

func (l *CalcListener) ExitBraceOrEqualInitializer(c *parser.BraceOrEqualInitializerContext) {
	// decl := c.Declarator()
	fmt.Println("boe-initializer", l.DumpStack())
	initializerClause := c.InitializerClause()
	bracedInitList := c.BracedInitList()
	if initializerClause != nil {
		return
	} else if bracedInitList != nil {
		panic("not implemented!")
	} else {
		panic("error condition!")
	}
}

// func (l *CalcListener) ExitInitializerList(c *parser.InitializerListContext) {
// 	initializersCount := len(c.AllInitializerClause())

// }

func (l *CalcListener) ExitInitializer(c *parser.InitializerContext) {
	// decl := c.Declarator()
	fmt.Println("initializer", l.DumpStack())
	boeInitalizer := c.BraceOrEqualInitializer()
	exprList := c.ExpressionList()
	if boeInitalizer != nil {
		return
	} else if exprList != nil {
		length := len(exprList.InitializerList().AllInitializerClause())
		exprListNode := cpp_ast.ExpressionListNode{}
		exprListNode.SetLocationFromToken(c.GetParser().GetCurrentToken())
		for i := 0; i < length; i++ {
			exprListNode.Expressions = append(exprListNode.Expressions, l.StackPop())
		}
		l.pushStack(&exprListNode)
	} else {
		panic("error condition!")
	}
}
