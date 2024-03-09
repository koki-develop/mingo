package mingo

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

func Minify(filename string, src []byte) ([]byte, error) {
	if s, err := format.Source(src); err != nil {
		return nil, err
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

func (m *mingo) Minify(filename string, src []byte) ([]byte, error) {
	file, err := parser.ParseFile(m.fileSet, filename, string(src), parser.ParseComments)
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	for _, cg := range file.Comments {
		for _, c := range cg.List {
			dirs := []string{"//go:build ", "// +build ", "//go:generate "}
			for _, prefix := range dirs {
				if strings.HasPrefix(c.Text, prefix) {
					fmt.Fprintln(b, c.Text)
				}
			}
		}
	}

	fmt.Fprint(b, m.stringifyFile(file))
	for _, decl := range file.Decls {
		fmt.Fprint(b, m.stringifyDecl(decl))
	}

	return b.Bytes(), nil
}
