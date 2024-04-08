package cpp_parser

import (
	cpp_ast "antlr-universal-parser/src/models/cpp_ast"
	"antlr-universal-parser/src/parsers/cpp/parser"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"slices"
	"sort"

	"github.com/antlr4-go/antlr/v4"
)

func formatString(inString string) string {
	indentSize := 3
	addIndent := "   "
	outString := string(inString[0]) // no indent for 1st (
	indent := ""
	for i := 1; i < len(inString); i++ {
		if inString[i] == '(' && (i+1 >= len(inString) || inString[i+1] != ' ') {
			indent += addIndent
			outString += "\n" + indent + "("
		} else if inString[i] == ')' {
			outString += ")"
			if len(indent) > 0 {
				indent = indent[:len(indent)-indentSize]
			}
		} else {
			outString += string(inString[i])
		}
	}
	return outString
}

type evalVisitor struct {
	*parser.BaseCPP14ParserVisitor
	// *CPP14ParserVisitor
}

type CalcListener struct {
	*parser.BaseCPP14ParserListener //继承Listener基类
	*antlr.DefaultErrorListener     //继承错误基类
	stack                           []cpp_ast.BaseNodeMethods
}

func (l *CalcListener) DumpStack() string {
	ret, _ := json.MarshalIndent(l.stack, "", "  ")
	return string(ret)
}

func (l *CalcListener) StackPop() cpp_ast.BaseNodeMethods {
	length := len(l.stack)
	defer func() { l.stack = l.stack[0 : len(l.stack)-1] }()
	return l.stack[length-1]
}

func (l *CalcListener) StackPeek(i int) cpp_ast.BaseNodeMethods {
	length := len(l.stack)
	return l.stack[length-i-1]
}

// 发生错误时，处理错误
func (l *CalcListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {

}

// 发生错误时，处理错误
func (l *CalcListener) pushStack(item cpp_ast.BaseNodeMethods) {
	if reflect.ValueOf(item).Kind() != reflect.Ptr {
		panic("value of item must be a pointer")
	}
	l.stack = append(l.stack, item)
}

func (l *CalcListener) ExitIdExpression(c *parser.IdExpressionContext) {
	fmt.Println("id expr", c.GetText())
	node := &cpp_ast.IdentifierNode{
		Name: c.GetText(),
	}
	node.SetLocationFromToken(
		c.GetStart())
	l.pushStack(node)
}

func (l *CalcListener) ExitPrimaryExpression(c *parser.PrimaryExpressionContext) {
	if c.IdExpression() != nil {
		return
	} else if len(c.AllLiteral()) > 0 {
		node := &cpp_ast.LiteralNode{Values: make([]string, len(c.AllLiteral()))}
		for i, literal := range c.AllLiteral() {
			node.Values[i] = literal.GetText()
		}
		node.SetLocationFromToken(c.GetStart())
		l.pushStack(node)
	} else {
		panic("cannot handle this type!")
	}
}

func (l *CalcListener) ExitPostfixExpression(c *parser.PostfixExpressionContext) {
	if c.PrimaryExpression() != nil {
	} else if c.PostfixExpression() != nil {
		if c.ExpressionList() != nil {
			l.CreateCallNode(c)
		} else if c.LeftBracket() != nil {
			l.CreateArrayReferenceNode(c)
		} else {
			panic("cannot handle this type!")
		}
	} else {
		panic("cannot handle this type!")
	}
}

func (l *CalcListener) ExitUnaryExpression(c *parser.UnaryExpressionContext) {
	if c.PostfixExpression() != nil {
		return // IGNORE
	} else {
		panic("cannot handle this type!")
	}
}

func (l *CalcListener) ExitCastExpression(c *parser.CastExpressionContext) {
	if c.UnaryExpression() != nil {
		return // IGNORE
	} else {
		panic("cannot handle this type!")
	}
}

func (l *CalcListener) ExitPointerMemberExpression(c *parser.PointerMemberExpressionContext) {
	for _, expr := range c.AllCastExpression() {
		fmt.Println("cast-expr", expr)
	}
}

func (l *CalcListener) ExitMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	var operators []antlr.TerminalNode = append(c.AllDiv(), append(c.AllMod(), c.AllStar()...)...)
	sort.Slice(operators, func(i, j int) bool {
		return operators[i].GetSymbol().GetTokenIndex() < operators[j].GetSymbol().GetTokenIndex()
	})

	slices.Reverse(operators)
	pointerMemberExprs := c.AllPointerMemberExpression()
	slices.Reverse(pointerMemberExprs)
	for _, op := range operators {
		l.CreateBinaryOperatorNode(c, op)
	}
}

func (l *CalcListener) ExitAdditiveExpression(c *parser.AdditiveExpressionContext) {
	for _, expr := range c.AllMultiplicativeExpression() {
		fmt.Println("multi-expr", expr)
	}
}

func (l *CalcListener) ExitShiftExpression(c *parser.ShiftExpressionContext) {
	for _, expr := range c.AllAdditiveExpression() {
		fmt.Println("additive-expr", expr)
	}
}

func (l *CalcListener) ExitRelationalExpression(c *parser.RelationalExpressionContext) {
	for _, expr := range c.AllShiftExpression() {
		fmt.Println("shift-expr", expr)
	}
}

func (l *CalcListener) ExitEqualityExpression(c *parser.EqualityExpressionContext) {
	for _, expr := range c.AllRelationalExpression() {
		fmt.Println("relational-expr", expr)
	}
}

func (l *CalcListener) ExitAndExpression(c *parser.AndExpressionContext) {
	for _, expr := range c.AllEqualityExpression() {
		fmt.Println("eq-expr", expr)
	}
}

func (l *CalcListener) ExitExclusiveOrExpression(c *parser.ExclusiveOrExpressionContext) {
	for _, expr := range c.AllAndExpression() {
		fmt.Println("andexpr", expr)
	}
}

func (l *CalcListener) ExitInclusiveOrExpression(c *parser.InclusiveOrExpressionContext) {
	for _, exclusiveOrExpr := range c.AllExclusiveOrExpression() {
		fmt.Println("excl-or", exclusiveOrExpr)
	}
}

func (l *CalcListener) ExitLogicalAndExpression(c *parser.LogicalAndExpressionContext) {
	for _, inclusiveExpr := range c.AllInclusiveOrExpression() {
		fmt.Println("inclusive", inclusiveExpr)
	}
}

func (l *CalcListener) ExitLogicalOrExpression(c *parser.LogicalOrExpressionContext) {
	for _, logicalAnd := range c.AllLogicalAndExpression() {
		fmt.Println("logicalAnd", logicalAnd)
	}
}

func (l *CalcListener) ExitConditionalExpression(c *parser.ConditionalExpressionContext) {
	logicalOr := c.LogicalOrExpression()
	fmt.Println("logicalOr", logicalOr)
}

func (l *CalcListener) ExitDeclaratorid(c *parser.DeclaratoridContext) {
	if c.Ellipsis() != nil {
		panic("not implemented!")
	}
	// node:=Identif
	// l.pushStack()
}

func (l *CalcListener) ExitNoPointerDeclarator(c *parser.NoPointerDeclaratorContext) {

}

func (l *CalcListener) ExitPointerDeclarator(c *parser.PointerDeclaratorContext) {
	fmt.Println("pointerDeclarator", l.DumpStack())

}

func (l *CalcListener) ExitInitDeclarator(c *parser.InitDeclaratorContext) {
	fmt.Println("init-declarator", l.DumpStack())
	if c.Initializer() != nil {
		if c.Initializer().BraceOrEqualInitializer() != nil {
			fmt.Println("initializer", l.DumpStack())
			rValue, lValue := l.StackPop(), l.StackPop()
			assignment := &cpp_ast.AssignmentNode{Operator: "=", RValue: rValue, LValue: lValue}
			l_location := cpp_ast.GetNodeLocation(lValue)
			assignment.SetLocation(l_location.Line, l_location.Column)
			l.pushStack(assignment)
		} else {
			l.CreateCallNode(c)
		}
	} else {
		panic("not implemented such method!")
	}
}

func (l *CalcListener) ExitAssignmentExpression(c *parser.AssignmentExpressionContext) {
	strings := make([]string, 0)
	ret2 := c.ToStringTree(strings, c.GetParser())
	fmt.Printf("%+v\n", ret2)
	op := c.AssignmentOperator()
	fmt.Println("op", op, "cond", c.ConditionalExpression())
	if op == nil {
		return // IGNORE
	} else {
		initClause, logicalOrExpr := l.StackPop(), l.StackPop()
		out := convertAssignmentOperator(op)
		l.pushStack(&cpp_ast.AssignmentNode{
			BaseNodeInfo: cpp_ast.BaseNodeInfo{
				Location: cpp_ast.GetNodeLocation(logicalOrExpr),
			},
			Operator: out,
			LValue:   logicalOrExpr,
			RValue:   initClause,
		})

	}
	assign := op.Assign()
	fmt.Println("assign", assign)
	// if assign != nil {
	// 	fmt.Println(assign.GetText())
	// }
}

func (l *CalcListener) ExitStatement(c *parser.StatementContext) {
	child := l.StackPop()
	stmt := &cpp_ast.StatementNode{Child: child}
	childLoc := cpp_ast.GetNodeLocation(child)
	stmt.SetLocation(childLoc.Line, childLoc.Column)
	l.pushStack(stmt)
}

func ParseFile(file string) *CalcListener {
	input, err := antlr.NewFileStream(file)
	if err != nil {
		panic(err)
	}
	lexer := parser.NewCPP14Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCPP14Parser(stream)
	tree := p.TranslationUnit()

	strings := make([]string, 0)
	ret2 := tree.ToStringTree(strings, p)
	fmt.Printf("%+v\n", formatString(ret2))
	var listener CalcListener
	listener.stack = make([]cpp_ast.BaseNodeMethods, 0)
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)

	return &listener
}

func Main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewCPP14Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCPP14Parser(stream)
	tree := p.TranslationUnit()

	strings := make([]string, 0)
	ret2 := tree.ToStringTree(strings, p)
	fmt.Printf("%+v\n", formatString(ret2))
	var listener CalcListener
	listener.stack = make([]cpp_ast.BaseNodeMethods, 0)
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)

	ret, _ := json.MarshalIndent(listener.stack, "", "  ")
	ioutil.WriteFile("test.json", ret, 666)
	fmt.Println("写入成功")
}
