package mingo

import (
	"fmt"
	"go/ast"
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
			fmt.Fprint(sb, minifyGenDecl(x))
		case *ast.FuncDecl:
			fmt.Fprint(sb, minifyFuncDecl(x))
		case *ast.Ident, *ast.BasicLit, *ast.ImportSpec, nil:
			// pass
		default:
			// fmt.Printf("%#v\n", x)
		}
		return true
	})

	return sb.String(), nil
}

func minifyFile(n *ast.File) string {
	return fmt.Sprintf("package %s;", n.Name.Name)
}

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

func minifyFuncDecl(n *ast.FuncDecl) string {
	sb := new(strings.Builder)

	sb.WriteString("func ")

	if n.Recv != nil {
		// TODO
	}

	fmt.Fprintf(sb, "%s(", n.Name.Name)

	// args
	for i, arg := range n.Type.Params.List {
		fmt.Printf("arg: %#v\n", arg)
		if i > 0 {
			sb.WriteString(",")
		}
		for j, name := range arg.Names {
			if j > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(name.Name)
		}

		sb.WriteString(" ")
		sb.WriteString(stringifyArgType(arg.Type))
	}
	sb.WriteString(")")

	// result
	if n.Type.Results != nil {
		rb := new(strings.Builder)

		needParens := false
		if len(n.Type.Results.List) > 1 {
			needParens = true
		}

		for i, rslt := range n.Type.Results.List {
			if i > 0 {
				rb.WriteString(",")
			}

			for j, name := range rslt.Names {
				needParens = true
				if j > 0 {
					rb.WriteString(",")
				}
				rb.WriteString(name.Name)
				rb.WriteString(" ")
			}

			rb.WriteString(stringifyArgType(rslt.Type))
		}

		if needParens {
			sb.WriteString("(")
		}
		sb.WriteString(rb.String())
		if needParens {
			sb.WriteString(")")
		}
	}

	sb.WriteString(";")
	sb.WriteString("\n") // TODO: remove
	return sb.String()
}

func stringifyArgType(t ast.Expr) (rtn string) {
	sb := new(strings.Builder)

	switch x := t.(type) {
	case *ast.StarExpr:
		sb.WriteString("*")
		sb.WriteString(stringifyArgType(x.X))
	case *ast.ArrayType:
		sb.WriteString("[]")
		sb.WriteString(stringifyArgType(x.Elt))
	case *ast.SelectorExpr:
		fmt.Fprintf(sb, "%s.%s", x.X.(*ast.Ident).Name, x.Sel.Name)
	case *ast.Ellipsis:
		sb.WriteString("...")
		sb.WriteString(stringifyArgType(x.Elt))
	default:
		sb.WriteString(x.(*ast.Ident).Name)
	}

	return sb.String()
}
