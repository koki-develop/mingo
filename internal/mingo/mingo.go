package mingo

import (
	"fmt"
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
	fmt.Fprint(sb, m.stringifyFile(file))
	for _, decl := range file.Decls {
		fmt.Fprint(sb, m.stringifyDecl(decl))
	}

	return sb.String(), nil
}
