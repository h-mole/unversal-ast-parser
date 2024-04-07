package testingdemo

import (
	cpp_parser "antlr-universal-parser/src/parsers/cpp"
	"path"
	"testing"
)

func TestBasic1(t *testing.T) {
	// _, filename, _, _ := runtime.Caller(1)
	// dirname := (path.Dir(filename))
	file := path.Join("./tests/assets/basic1.cpp")
	stack := cpp_parser.ParseFile(file)
	// stack
}
