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

	return Minify(filename, src)
}

func Minify(filename string, src []byte) (string, error) {
	if s, err := format.Source(src); err != nil {
		return "", err
	} else {
		src = s
	}

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
			return true
		case *ast.GenDecl:
			fmt.Fprint(sb, stringifyGenDecl(x))
			return true
		case *ast.FuncDecl:
			fmt.Fprint(sb, stringifyFuncDecl(x))
			return true
		}
		return false
	})

	return sb.String(), nil
}

func minifyFile(n *ast.File) string {
	return fmt.Sprintf("package %s;", n.Name.Name)
}
