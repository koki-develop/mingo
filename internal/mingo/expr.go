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
	case *ast.FuncLit:
		// TODO
	case *ast.BinaryExpr:
		sb.WriteString(stringifyExpr(x.X))
		sb.WriteString(x.Op.String())
		sb.WriteString(stringifyExpr(x.Y))
	case *ast.SliceExpr:
		sb.WriteString(stringifyExpr(x.X))
		sb.WriteString("[")
		if x.Low != nil {
			sb.WriteString(stringifyExpr(x.Low))
		}
		sb.WriteString(":") // FIXME
		if x.High != nil {
			sb.WriteString(stringifyExpr(x.High))
		}
		if x.Max != nil {
			sb.WriteString(":")
			sb.WriteString(stringifyExpr(x.Max))
		}
		sb.WriteString("]")
	case *ast.UnaryExpr:
		sb.WriteString(x.Op.String())
		sb.WriteString(stringifyExpr(x.X))
	case *ast.CompositeLit:
		sb.WriteString(stringifyExpr(x.Type))
		sb.WriteString("{")
		for i, elt := range x.Elts {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(stringifyExpr(elt))
		}
		sb.WriteString("}")
	case *ast.ParenExpr:
		sb.WriteString("(")
		sb.WriteString(stringifyExpr(x.X))
		sb.WriteString(")")
	case *ast.IndexExpr:
		sb.WriteString(stringifyExpr(x.X))
		sb.WriteString("[")
		sb.WriteString(stringifyExpr(x.Index))
		sb.WriteString("]")
	case *ast.KeyValueExpr:
		sb.WriteString(stringifyExpr(x.Key))
		sb.WriteString(":")
		sb.WriteString(stringifyExpr(x.Value))
	case *ast.TypeAssertExpr:
		sb.WriteString(stringifyExpr(x.X))
		sb.WriteString(".(")
		sb.WriteString(stringifyExpr(x.Type))
		sb.WriteString(")")
	case *ast.ChanType:
		if x.Dir == ast.RECV {
			sb.WriteString("<-")
		}
		sb.WriteString("chan")
		sb.WriteString(stringifyExpr(x.Value))
	case *ast.MapType:
		sb.WriteString("map[")
		sb.WriteString(stringifyExpr(x.Key))
		sb.WriteString("]")
		sb.WriteString(stringifyExpr(x.Value))
	case *ast.InterfaceType:
		sb.WriteString("interface{}")
	case *ast.StructType:
		sb.WriteString("struct{")
		for i, field := range x.Fields.List {
			if i > 0 {
				sb.WriteString(";")
			}
			for j, name := range field.Names {
				if j > 0 {
					sb.WriteString(",")
				}
				sb.WriteString(name.Name)
			}
			sb.WriteString(" ")
			sb.WriteString(stringifyExpr(field.Type))
		}
		sb.WriteString("}")
	case *ast.FuncType:
		sb.WriteString("func(")

		// args
		for i, arg := range x.Params.List {
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
			sb.WriteString(stringifyExpr(arg.Type))
		}
		sb.WriteString(")")

		// result
		if x.Results != nil {
			rb := new(strings.Builder)

			needParens := false
			if len(x.Results.List) > 1 {
				needParens = true
			}

			for i, rslt := range x.Results.List {
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
		}
	default:
		sb.WriteString(x.(*ast.Ident).Name)
	}

	return sb.String()
}
