package mingo

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func stringifyStmt(stmt ast.Stmt) string {
	sb := new(strings.Builder)

	switch x := stmt.(type) {
	case *ast.ReturnStmt:
		sb.WriteString("return")
		if len(x.Results) > 0 {
			sb.WriteString(" ")
		}

		for i, expr := range x.Results {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(stringifyExpr(expr))
		}
	case *ast.AssignStmt:
		for i, expr := range x.Lhs {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(stringifyExpr(expr))
		}
		if x.Tok == token.DEFINE {
			sb.WriteString(":=")
		} else {
			sb.WriteString("=")
		}
		for i, expr := range x.Rhs {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(stringifyExpr(expr))
		}
	default:
		fmt.Printf("stmt: %#v\n", x)
	}

	sb.WriteString(";")
	return sb.String()
}
