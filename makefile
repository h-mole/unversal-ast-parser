mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
project_root := $(dir $(mkfile_path))

cpp-grammar:
	echo $(project_root)
	cd src/parsers/cpp && java -jar $(project_root)antlr-4.13.1-complete.jar -Dlanguage=Cpp -o parser CPP14Lexer.g4
	cd src/parsers/cpp && java -jar $(project_root)antlr-4.13.1-complete.jar -Dlanguage=Cpp -o parser -visitor CPP14Parser.g4

cpp-grammar-go:
	echo $(project_root)
	cd src/parsers/cpp && java -jar $(project_root)antlr-4.13.1-complete.jar -Dlanguage=Go -o parser CPP14Lexer.g4
	cd src/parsers/cpp && java -jar $(project_root)antlr-4.13.1-complete.jar -Dlanguage=Go -o parser -visitor CPP14Parser.g4

build:
	g++ -I ./ -I runtime/Cpp/run/usr/local/include/antlr4-runtime \
	 -I src/parsers/cpp/parser \
	 -Wall -Wextra \
	 -L runtime/Cpp/run/usr/local/lib/ -L /usr/local/lib \
	 -Wl,-rpath=/usr/local/lib/ -lantlr4-runtime \
	   ./src/parsers/cpp/parser/*.cpp main.cpp

build-test-file:
	/usr/local/go/bin/go test -v tests/basic_ast_test.go -c

test: build-test-file
	./testingdemo.test 