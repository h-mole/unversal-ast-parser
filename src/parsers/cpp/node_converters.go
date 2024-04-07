package cpp_parser

import "antlr-universal-parser/src/parsers/cpp/parser"

func convertAssignmentOperator(c parser.IAssignmentOperatorContext) string {
	if c.PlusAssign() != nil {
		return "+="
	} else if c.MinusAssign() != nil {
		return "-="
	} else if c.StarAssign() != nil {
		return "*="
	} else if c.DivAssign() != nil {
		return "/="
	} else if c.ModAssign() != nil {
		return "%="
	} else if c.LeftShiftAssign() != nil {
		return "<<="
	} else if c.RightShiftAssign() != nil {
		return ">>="
	} else if c.AndAssign() != nil {
		return "&="
	} else if c.XorAssign() != nil {
		return "^="
	} else if c.OrAssign() != nil {
		return "|="
	} else {
		panic("not implemented")
	}
}
