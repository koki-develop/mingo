package mingo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func MinifyFile(filename string) (string, error) {
	src, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return Minify(filename, src)
}

func Minify(filename string, src []byte) (string, error) {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, filename, string(src), 0)
	if err != nil {
		return "", err
	}

	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		default:
			fmt.Printf("%T\n", x)
		}
		return true
	})

	return "", nil
}
