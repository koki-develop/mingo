package mingo

import (
	"fmt"
	"go/ast"
	"strings"
)

func (m *mingo) stringifyFuncDecl(n *ast.FuncDecl) string {
	sb := new(strings.Builder)

	sb.WriteString("func")

	if n.Recv != nil {
		sb.WriteString(m.stringifyFuncParams(n.Recv))
	} else {
		sb.WriteString(" ")
	}

	fmt.Fprintf(sb, "%s", n.Name.Name)

	sb.WriteString(m.stringifyFuncTypeParams(n.Type.TypeParams))
	sb.WriteString(m.stringifyFuncParams(n.Type.Params))
	sb.WriteString(m.stringifyFuncResults(n.Type.Results))
	sb.WriteString(m.stringifyBlockStmt(n.Body))

	sb.WriteString(";")
	return sb.String()
}

func (m *mingo) stringifyFuncTypeParams(params *ast.FieldList) string {
	if params == nil {
		return ""
	}

	sb := new(strings.Builder)

	sb.WriteString("[")
	for i, param := range params.List {
		if i > 0 {
			sb.WriteString(",")
		}
		for j, name := range param.Names {
			if j > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(name.Name)
		}
		if len(param.Names) > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(m.stringifyExpr(param.Type))
	}
	sb.WriteString("]")

	return sb.String()
}

func (m *mingo) stringifyFuncParams(params *ast.FieldList) string {
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

		sb.WriteString(m.stringifyExpr(arg.Type))
	}

	sb.WriteString(")")
	return sb.String()
}

func (m *mingo) stringifyFuncResults(results *ast.FieldList) string {
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

		rb.WriteString(m.stringifyExpr(rslt.Type))
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
