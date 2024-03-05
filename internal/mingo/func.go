package mingo

import (
	"fmt"
	"go/ast"
	"strings"
)

func minifyFuncDecl(n *ast.FuncDecl) string {
	sb := new(strings.Builder)

	sb.WriteString("func ")

	if n.Recv != nil {
		// TODO
	}

	fmt.Fprintf(sb, "%s(", n.Name.Name)
	sb.WriteString(stringifyExpr(n.Type))

	// body
	sb.WriteString("{")
	for _, stmt := range n.Body.List {
		sb.WriteString(stringifyStmt(stmt))
	}
	sb.WriteString("}")

	sb.WriteString(";")
	sb.WriteString("\n") // TODO: remove
	return sb.String()
}
