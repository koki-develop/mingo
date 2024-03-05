package mingo

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func MinifyFile(filename string) (string, error) {
	src, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	fmted, err := format.Source(src)
	if err != nil {
		return "", err
	}

	return Minify(filename, fmted)
}

func Minify(filename string, src []byte) (string, error) {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, filename, string(src), 0)
	if err != nil {
		return "", err
	}

	sb := new(strings.Builder)
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.File:
			fmt.Fprint(sb, minifyFile(x))
		case *ast.GenDecl:
			fmt.Fprint(sb, stringifyGenDecl(x))
		case *ast.FuncDecl:
			fmt.Fprint(sb, stringifyFuncDecl(x))
		}
		return true
	})

	return sb.String(), nil
}

func minifyFile(n *ast.File) string {
	return fmt.Sprintf("package %s;", n.Name.Name)
}
