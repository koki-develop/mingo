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
		imports := []*ast.ImportSpec{}
		for _, spec := range n.Specs {
			imports = append(imports, spec.(*ast.ImportSpec))
		}
		return m.stringifyImportSpecs(imports)
	case token.CONST:
		consts := []*ast.ValueSpec{}
		for _, spec := range n.Specs {
			consts = append(consts, spec.(*ast.ValueSpec))
		}
		return m.stringifyConstSpecs(consts)
	case token.VAR:
		vars := []*ast.ValueSpec{}
		for _, spec := range n.Specs {
			vars = append(vars, spec.(*ast.ValueSpec))
		}
		return m.stringifyVarSpecs(vars)
	case token.TYPE:
		sb := new(strings.Builder)
		for _, spec := range n.Specs {
			sb.WriteString(m.stringifyTypeSpec(spec.(*ast.TypeSpec)))
			sb.WriteString(";")
		}
		return sb.String()
	}

	return ""
}

func (m *mingo) stringifyImportSpecs(specs []*ast.ImportSpec) string {
	sb := new(strings.Builder)
	sb.WriteString("import")

	if len(specs) > 1 {
		sb.WriteString("(")
	} else {
		sb.WriteString(" ")
	}

	for i, n := range specs {
		if i > 0 {
			sb.WriteString(";")
		}
		if n.Name != nil {
			sb.WriteString(fmt.Sprintf("%s %s", n.Name.Name, n.Path.Value))
		} else {
			sb.WriteString(n.Path.Value)
		}
	}

	if len(specs) > 1 {
		sb.WriteString(")")
	}
	sb.WriteString(";")
	return sb.String()
}

func (m *mingo) stringifyConstSpecs(specs []*ast.ValueSpec) string {
	sb := new(strings.Builder)
	sb.WriteString("const")

	if len(specs) > 1 {
		sb.WriteString("(")
	} else {
		sb.WriteString(" ")
	}

	for i, spec := range specs {
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

	if len(specs) > 1 {
		sb.WriteString(")")
	}
	sb.WriteString(";")
	return sb.String()
}

func (m *mingo) stringifyVarSpecs(specs []*ast.ValueSpec) string {
	sb := new(strings.Builder)
	sb.WriteString("var")

	if len(specs) > 1 {
		sb.WriteString("(")
	} else {
		sb.WriteString(" ")
	}

	for i, spec := range specs {
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

	if len(specs) > 1 {
		sb.WriteString(")")
	}
	sb.WriteString(";")
	return sb.String()
}

func (m *mingo) stringifyTypeSpec(n *ast.TypeSpec) string {
	sb := new(strings.Builder)

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

	return sb.String()
}
