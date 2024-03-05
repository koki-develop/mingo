package mingo

import (
	"fmt"
	"go/ast"
	"strings"
)

func stringifySelectExpr(expr *ast.SelectorExpr) string {
	switch x := expr.X.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", stringifySelectExpr(x), expr.Sel.Name)
	default:
		return fmt.Sprintf("%s.%s", x.(*ast.Ident).Name, expr.Sel.Name)
	}
}

func stringifyExpr(expr ast.Expr) string {
	sb := new(strings.Builder)

	switch x := expr.(type) {
	case *ast.BasicLit:
		sb.WriteString(x.Value)
	case *ast.CallExpr:
		sb.WriteString(stringifyExpr(x.Fun))
		sb.WriteString("(")
		for i, arg := range x.Args {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(stringifyExpr(arg))
		}
		sb.WriteString(")")
	case *ast.SelectorExpr:
		sb.WriteString(stringifySelectExpr(x))
	case *ast.StarExpr:
		sb.WriteString("*")
		sb.WriteString(stringifyExpr(x.X))
	case *ast.ArrayType:
		sb.WriteString("[]")
		sb.WriteString(stringifyExpr(x.Elt))
	case *ast.Ellipsis:
		sb.WriteString("...")
		sb.WriteString(stringifyExpr(x.Elt))
	default:
		sb.WriteString(x.(*ast.Ident).Name)
	}

	return sb.String()
}
