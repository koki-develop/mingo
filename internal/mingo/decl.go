package mingo

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func (m *mingo) stringifyDecl(decl ast.Decl) string {
	switch x := decl.(type) {
	case *ast.GenDecl:
		return m.stringifyGenDecl(x)
	case *ast.FuncDecl:
		return m.stringifyFuncDecl(x)
	}
	return ""
}

func (m *mingo) stringifyGenDecl(n *ast.GenDecl) string {
	switch n.Tok {
	case token.IMPORT:
		return m.stringifyImportDecl(n)
	case token.CONST:
		return m.stringifyConstDecl(n)
	case token.VAR:
		return m.stringifyVarDecl(n)
	case token.TYPE:
		return m.stringifyTypeSpec(n)
	}

	return ""
}

func (m *mingo) stringifyImportDecl(decl *ast.GenDecl) string {
	sb := new(strings.Builder)
	sb.WriteString("import")

	if len(decl.Specs) > 1 {
		sb.WriteString("(")
	} else {
		sb.WriteString(" ")
	}

	for i, n := range decl.Specs {
		n := n.(*ast.ImportSpec)

		if i > 0 {
			sb.WriteString(";")
		}
		if n.Name != nil {
			sb.WriteString(fmt.Sprintf("%s %s", n.Name.Name, n.Path.Value))
		} else {
			sb.WriteString(n.Path.Value)
		}
	}

	if len(decl.Specs) > 1 {
		sb.WriteString(")")
	}
	sb.WriteString(";")
	return sb.String()
}

func (m *mingo) stringifyConstDecl(decl *ast.GenDecl) string {
	sb := new(strings.Builder)
	sb.WriteString("const")

	if len(decl.Specs) > 1 {
		sb.WriteString("(")
	} else {
		sb.WriteString(" ")
	}

	for i, spec := range decl.Specs {
		spec := spec.(*ast.ValueSpec)

		if i > 0 {
			sb.WriteString(";")
		}
		for j, name := range spec.Names {
			if j > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(name.Name)
		}

		if spec.Type != nil {
			sb.WriteString(" ")
			sb.WriteString(m.stringifyExpr(spec.Type))
		}

		if spec.Values != nil {
			sb.WriteString("=")
			for k, value := range spec.Values {
				if k > 0 {
					sb.WriteString(",")
				}
				sb.WriteString(m.stringifyExpr(value))
			}
		}
	}

	if len(decl.Specs) > 1 {
		sb.WriteString(")")
	}
	sb.WriteString(";")
	return sb.String()
}

func (m *mingo) stringifyVarDecl(decl *ast.GenDecl) string {
	sb := new(strings.Builder)

	if decl.Doc != nil {
		for _, cmt := range decl.Doc.List {
			if strings.HasPrefix(cmt.Text, "//go:embed ") {
				fmt.Fprintf(sb, "\n%s\n", cmt.Text)
			}
		}
	}

	sb.WriteString("var")

	if len(decl.Specs) > 1 {
		sb.WriteString("(")
	} else {
		sb.WriteString(" ")
	}

	for i, spec := range decl.Specs {
		spec := spec.(*ast.ValueSpec)

		if i > 0 {
			sb.WriteString(";")
		}
		for j, name := range spec.Names {
			if j > 0 {
				sb.WriteString(",")
			}

			if spec.Doc != nil {
				for _, cmt := range spec.Doc.List {
					if strings.HasPrefix(cmt.Text, "//go:embed ") {
						fmt.Fprintf(sb, "\n%s\n", cmt.Text)
					}
				}
			}
			sb.WriteString(name.Name)
		}

		if spec.Type != nil {
			sb.WriteString(" ")
			sb.WriteString(m.stringifyExpr(spec.Type))
		}

		if spec.Values != nil {
			sb.WriteString("=")
			for k, value := range spec.Values {
				if k > 0 {
					sb.WriteString(",")
				}
				sb.WriteString(m.stringifyExpr(value))
			}
		}
	}

	if len(decl.Specs) > 1 {
		sb.WriteString(")")
	}
	sb.WriteString(";")
	return sb.String()
}

func (m *mingo) stringifyTypeSpec(decl *ast.GenDecl) string {
	sb := new(strings.Builder)

	for _, n := range decl.Specs {
		n := n.(*ast.TypeSpec)

		sb.WriteString(fmt.Sprintf("type %s", n.Name.Name))
		if n.TypeParams != nil {
			sb.WriteString(m.stringifyFuncTypeParams(n.TypeParams))
		}
		if n.Assign != 0 {
			sb.WriteString("=")
		} else {
			sb.WriteString(" ")
		}
		sb.WriteString(m.stringifyExpr(n.Type))
		sb.WriteString(";")
	}

	return sb.String()
}
