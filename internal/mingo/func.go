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

	// args
	for i, arg := range n.Type.Params.List {
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
