package mingo

import (
	"fmt"
	"go/ast"
	"strings"
)

func minifyFuncDecl(n *ast.FuncDecl) string {
	sb := new(strings.Builder)

	sb.WriteString("func")

	if n.Recv != nil {
		sb.WriteString(stringifyFuncParams(n.Recv))
	} else {
		sb.WriteString(" ")
	}

	fmt.Fprintf(sb, "%s", n.Name.Name)

	// args
	sb.WriteString(stringifyFuncParams(n.Type.Params))

	// result
	sb.WriteString(stringifyFuncResults(n.Type.Results))

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

func stringifyFuncParams(params *ast.FieldList) string {
	sb := new(strings.Builder)

	sb.WriteString("(")

	for i, arg := range params.List {
		if i > 0 {
			sb.WriteString(",")
		}
		for j, name := range arg.Names {
			if j > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(name.Name)
		}

		if len(arg.Names) > 0 {
			sb.WriteString(" ")
		}

		sb.WriteString(stringifyExpr(arg.Type))
	}

	sb.WriteString(")")
	return sb.String()
}

func stringifyFuncResults(results *ast.FieldList) string {
	if results == nil {
		return ""
	}

	sb := new(strings.Builder)
	rb := new(strings.Builder)

	needParens := false
	if len(results.List) > 1 {
		needParens = true
	}

	for i, rslt := range results.List {
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

		rb.WriteString(stringifyExpr(rslt.Type))
	}

	if needParens {
		sb.WriteString("(")
	}
	sb.WriteString(rb.String())
	if needParens {
		sb.WriteString(")")
	}

	return sb.String()
}
