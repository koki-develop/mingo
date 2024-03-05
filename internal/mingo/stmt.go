package mingo

import (
	"go/ast"
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
	default:
		// fmt.Printf("%#v", x)
	}

	sb.WriteString(";")
	return sb.String()
}
