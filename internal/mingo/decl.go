package mingo

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func stringifyDecl(decl ast.Decl) string {
	switch x := decl.(type) {
	case *ast.GenDecl:
		return stringifyGenDecl(x)
	case *ast.FuncDecl:
		return stringifyFuncDecl(x)
	}
	return ""
}

func stringifyGenDecl(n *ast.GenDecl) string {
	switch n.Tok {
	case token.IMPORT:
		imports := []*ast.ImportSpec{}
		for _, spec := range n.Specs {
			imports = append(imports, spec.(*ast.ImportSpec))
		}
		return stringifyImportSpecs(imports)
	case token.CONST:
		// TODO
	case token.VAR:
		// TODO
	case token.TYPE:
		sb := new(strings.Builder)
		for _, spec := range n.Specs {
			sb.WriteString(stringifyTypeSpec(spec.(*ast.TypeSpec)))
			sb.WriteString(";")
		}
		return sb.String()
	}

	return ""
}

func stringifyImportSpecs(ns []*ast.ImportSpec) string {
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

func stringifyTypeSpec(n *ast.TypeSpec) string {
	return fmt.Sprintf("type %s %s", n.Name.Name, stringifyExpr(n.Type))
}
