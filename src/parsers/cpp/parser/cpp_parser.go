package parser

import (
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

// type evalVisitor struct {
// 	*BaseCPP14ParserVisitor
// 	// *CPP14ParserVisitor
// }

// type CalcListener struct {
// 	*BaseCPP14ParserListener    //继承Listener基类
// 	*antlr.DefaultErrorListener //继承错误基类
// 	stack                       interface{}
// }

// // 发生错误时，处理错误
// func (l *CalcListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {

// }

// func (l *CalcListener) ExitPrimaryExpression(c *PrimaryExpressionContext) {
// 	if c.IdExpression() != nil {
// 		fmt.Println("id expr", c.IdExpression().GetText())
// 	} else if len(c.AllLiteral()) > 0 {
// 		for _, literal := range c.AllLiteral() {
// 			fmt.Println("literal", literal.GetText())
// 		}
// 	} else {
// 		panic("cannot handle this type!")
// 	}
// }

// func (l *CalcListener) ExitPostfixExpression(c *PostfixExpressionContext) {
// 	if c.PrimaryExpression() != nil {
// 		fmt.Println("primary expr")
// 	} else {
// 		panic("cannot handle this type!")
// 	}
// }

// func (l *CalcListener) ExitUnaryExpression(c *UnaryExpressionContext) {
// 	if c.PostfixExpression() != nil {
// 		fmt.Println("postfix expr")
// 	} else {
// 		panic("cannot handle this type!")
// 	}
// }

// func (l *CalcListener) ExitCastExpression(c *CastExpressionContext) {
// 	if c.UnaryExpression() != nil {
// 		fmt.Println("cast_expr-unary expr")
// 	} else {
// 		panic("cannot handle this type!")
// 	}
// }

// func (l *CalcListener) ExitPointerMemberExpression(c *PointerMemberExpressionContext) {
// 	for _, expr := range c.AllCastExpression() {
// 		fmt.Println("cast-expr", expr)
// 	}
// }

// func (l *CalcListener) ExitMultiplicativeExpression(c *MultiplicativeExpressionContext) {
// 	for _, expr := range c.AllPointerMemberExpression() {
// 		fmt.Println("pointer-member-expr", expr)
// 	}
// }

// func (l *CalcListener) ExitAdditiveExpression(c *AdditiveExpressionContext) {
// 	for _, expr := range c.AllMultiplicativeExpression() {
// 		fmt.Println("multi-expr", expr)
// 	}
// }

// func (l *CalcListener) ExitShiftExpression(c *ShiftExpressionContext) {
// 	for _, expr := range c.AllAdditiveExpression() {
// 		fmt.Println("additive-expr", expr)
// 	}
// }

// func (l *CalcListener) ExitRelationalExpression(c *RelationalExpressionContext) {
// 	for _, expr := range c.AllShiftExpression() {
// 		fmt.Println("shift-expr", expr)
// 	}
// }

// func (l *CalcListener) ExitEqualityExpression(c *EqualityExpressionContext) {
// 	for _, expr := range c.AllRelationalExpression() {
// 		fmt.Println("relational-expr", expr)
// 	}
// }

// func (l *CalcListener) ExitAndExpression(c *AndExpressionContext) {
// 	for _, expr := range c.AllEqualityExpression() {
// 		fmt.Println("eq-expr", expr)
// 	}
// }

// func (l *CalcListener) ExitExclusiveOrExpression(c *ExclusiveOrExpressionContext) {
// 	for _, expr := range c.AllAndExpression() {
// 		fmt.Println("andexpr", expr)
// 	}
// }

// func (l *CalcListener) ExitInclusiveOrExpression(c *InclusiveOrExpressionContext) {
// 	for _, exclusiveOrExpr := range c.AllExclusiveOrExpression() {
// 		fmt.Println("excl-or", exclusiveOrExpr)
// 	}
// }

// func (l *CalcListener) ExitLogicalAndExpression(c *LogicalAndExpressionContext) {
// 	for _, inclusiveExpr := range c.AllInclusiveOrExpression() {
// 		fmt.Println("inclusive", inclusiveExpr)
// 	}
// }

// func (l *CalcListener) ExitLogicalOrExpression(c *LogicalOrExpressionContext) {
// 	for _, logicalAnd := range c.AllLogicalAndExpression() {
// 		fmt.Println("logicalAnd", logicalAnd)
// 	}
// }

// func (l *CalcListener) ExitConditionalExpression(c *ConditionalExpressionContext) {
// 	logicalOr := c.LogicalOrExpression()
// 	fmt.Println("logicalOr", logicalOr)
// }

// func (l *CalcListener) ExitAssignmentExpression(c *AssignmentExpressionContext) {
// 	op := c.AssignmentOperator()
// 	fmt.Println("op", op, "cond", c.ConditionalExpression())
// 	if op == nil {
// 		return
// 	}
// 	assign := op.Assign()
// 	fmt.Println("assign", assign)
// 	// if assign != nil {
// 	// 	fmt.Println(assign.GetText())
// 	// }
// }

// ParserBase implementation.
type CPP14ParserBase struct {
	*antlr.BaseParser
}

// Returns true if the current Token is a closing bracket (")" or "}")
func (p *CPP14ParserBase) IsPureSpecifierAllowed() bool {
	x := p.GetParserRuleContext()
	y := x.GetChild(0)
	c := y.GetChild(0)
	c2 := c.GetChild(0)
	if c2.GetChildCount() <= 1 {
		return false
	}
	c3 := c2.GetChild(1)
	if c3 == nil {
		return false
	}
	ce1 := reflect.TypeOf(c3).Elem()
	ce2 := reflect.TypeOf(ParametersAndQualifiersContext{})
	ce := ce1 == ce2
	return ce
}

// func Main() {
// 	input, _ := antlr.NewFileStream(os.Args[1])
// 	lexer := NewCPP14Lexer(input)
// 	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
// 	p := NewCPP14Parser(stream)
// 	// p.BuildParseTrees = true
// 	tree := p.TranslationUnit()
// 	// println(tree)
// 	// var visitor evalVisitor
// 	// visitor.Visit(tree)
// 	var listener CalcListener
// 	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
// 	//
// 	return
// }
