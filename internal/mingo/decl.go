package mingo

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func minifyGenDecl(n *ast.GenDecl) string {
	switch n.Tok {
	case token.IMPORT:
		imports := []*ast.ImportSpec{}
		for _, spec := range n.Specs {
			imports = append(imports, spec.(*ast.ImportSpec))
		}
		return minifyImportSpecs(imports)
	case token.TYPE:
		// TODO
	case token.CONST:
		// TODO
	case token.VAR:
		// TODO
	}

	return ""
}

func minifyImportSpecs(ns []*ast.ImportSpec) string {
	sb := new(strings.Builder)
	sb.WriteString("import ")

	if len(ns) > 1 {
		sb.WriteString("(")
	}

	for i, n := range ns {
		if i > 0 {
			sb.WriteString(";")
		}
		if n.Name != nil {
			sb.WriteString(fmt.Sprintf("%s %s", n.Name.Name, n.Path.Value))
		} else {
			sb.WriteString(n.Path.Value)
		}
	}

	if len(ns) > 1 {
		sb.WriteString(")")
	}
	sb.WriteString(";")
	return sb.String()
}
