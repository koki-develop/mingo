package mingo

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

func Minify(filename string, src []byte) (string, error) {
	if s, err := format.Source(src); err != nil {
		return "", err
	} else {
		src = s
	}

	fset := token.NewFileSet()
	m := &mingo{fileSet: fset}
	return m.Minify(filename, src)
}

type mingo struct {
	fileSet *token.FileSet
}

func (m *mingo) Minify(filename string, src []byte) (string, error) {
	file, err := parser.ParseFile(m.fileSet, filename, string(src), 0)
	if err != nil {
		return "", err
	}

	sb := new(strings.Builder)
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.File:
			fmt.Fprint(sb, stringifyFile(x))
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
