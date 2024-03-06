package mingo

import (
	"fmt"
	"go/ast"
	"strings"
)

func stringifyExpr(expr ast.Expr) string {
	switch x := expr.(type) {
	case *ast.BasicLit:
		return stringifyBasicLit(x)
	case *ast.CallExpr:
		return stringifyCallExpr(x)
	case *ast.SelectorExpr:
		return stringifySelectExpr(x)
	case *ast.StarExpr:
		return stringifyStarExpr(x)
	case *ast.ArrayType:
		return stringifyArrayType(x)
	case *ast.Ellipsis:
		return stringifyEllipsis(x)
	case *ast.FuncLit:
		return stringifyFuncLit(x)
	case *ast.BinaryExpr:
		return stringifyBinaryExpr(x)
	case *ast.SliceExpr:
		return stringifySliceExpr(x)
	case *ast.UnaryExpr:
		return stringifyUnaryExpr(x)
	case *ast.CompositeLit:
		return stringifyCompositeLit(x)
	case *ast.ParenExpr:
		return stringifyParenExpr(x)
	case *ast.IndexExpr:
		return stringifyIndexExpr(x)
	case *ast.KeyValueExpr:
		return stringifyKeyValueExpr(x)
	case *ast.TypeAssertExpr:
		return stringifyTypeAssertExpr(x)
	case *ast.ChanType:
		return stringifyChanType(x)
	case *ast.MapType:
		return stringifyMapType(x)
	case *ast.InterfaceType:
		return stringifyInterfaceType(x)
	case *ast.StructType:
		return stringifyStructType(x)
	case *ast.FuncType:
		return stringifyFuncType(x)
	case *ast.IndexListExpr:
		return stringifyIndexListExpr(x)
	case nil:
		return ""
	}

	return expr.(*ast.Ident).Name
}

func stringifyIndexListExpr(expr *ast.IndexListExpr) string {
	sb := new(strings.Builder)

	sb.WriteString(stringifyExpr(expr.X))
	sb.WriteString("[")
	for i, index := range expr.Indices {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(stringifyExpr(index))
	}
	sb.WriteString("]")

	return sb.String()
}

func stringifySelectExpr(expr *ast.SelectorExpr) string {
	switch x := expr.X.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", stringifySelectExpr(x), expr.Sel.Name)
	default:
		return fmt.Sprintf("%s.%s", stringifyExpr(expr.X), expr.Sel.Name)
	}
}

func stringifyBasicLit(lit *ast.BasicLit) string {
	return lit.Value
}

func stringifyCallExpr(expr *ast.CallExpr) string {
	sb := new(strings.Builder)

	sb.WriteString(stringifyExpr(expr.Fun))
	sb.WriteString("(")
	for i, arg := range expr.Args {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(stringifyExpr(arg))
	}

	if expr.Ellipsis.IsValid() {
		sb.WriteString("...")
	}

	sb.WriteString(")")

	return sb.String()
}

func stringifyStarExpr(expr *ast.StarExpr) string {
	return fmt.Sprintf("*%s", stringifyExpr(expr.X))
}

func stringifyArrayType(expr *ast.ArrayType) string {
	return fmt.Sprintf("[]%s", stringifyExpr(expr.Elt))
}

func stringifyEllipsis(expr *ast.Ellipsis) string {
	return fmt.Sprintf("...%s", stringifyExpr(expr.Elt))
}

func stringifyFuncLit(expr *ast.FuncLit) string {
	sb := new(strings.Builder)
	sb.WriteString(stringifyFuncType(expr.Type))
	sb.WriteString(stringifyBlockStmt(expr.Body))

	return sb.String()
}

func stringifyBinaryExpr(expr *ast.BinaryExpr) string {
	return fmt.Sprintf("%s%s%s", stringifyExpr(expr.X), expr.Op.String(), stringifyExpr(expr.Y))
}

func stringifySliceExpr(expr *ast.SliceExpr) string {
	sb := new(strings.Builder)

	sb.WriteString(stringifyExpr(expr.X))
	sb.WriteString("[")
	if expr.Low != nil {
		sb.WriteString(stringifyExpr(expr.Low))
	}
	sb.WriteString(":")
	if expr.High != nil {
		sb.WriteString(stringifyExpr(expr.High))
	}
	if expr.Max != nil {
		sb.WriteString(":")
		sb.WriteString(stringifyExpr(expr.Max))
	}
	sb.WriteString("]")

	return sb.String()
}

func stringifyUnaryExpr(expr *ast.UnaryExpr) string {
	return fmt.Sprintf("%s%s", expr.Op.String(), stringifyExpr(expr.X))
}

func stringifyCompositeLit(expr *ast.CompositeLit) string {
	sb := new(strings.Builder)

	sb.WriteString(stringifyExpr(expr.Type))
	sb.WriteString("{")
	for i, elt := range expr.Elts {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(stringifyExpr(elt))
	}
	sb.WriteString("}")

	return sb.String()
}

func stringifyParenExpr(expr *ast.ParenExpr) string {
	return fmt.Sprintf("(%s)", stringifyExpr(expr.X))
}

func stringifyIndexExpr(expr *ast.IndexExpr) string {
	return fmt.Sprintf("%s[%s]", stringifyExpr(expr.X), stringifyExpr(expr.Index))
}

func stringifyKeyValueExpr(expr *ast.KeyValueExpr) string {
	return fmt.Sprintf("%s:%s", stringifyExpr(expr.Key), stringifyExpr(expr.Value))
}

func stringifyTypeAssertExpr(expr *ast.TypeAssertExpr) string {
	if expr.Type == nil {
		return fmt.Sprintf("%s.(type)", stringifyExpr(expr.X))
	} else {
		return fmt.Sprintf("%s.(%s)", stringifyExpr(expr.X), stringifyExpr(expr.Type))
	}
}

func stringifyChanType(expr *ast.ChanType) string {
	sb := new(strings.Builder)

	if expr.Dir == ast.RECV {
		sb.WriteString("<-")
	}
	sb.WriteString("chan ")
	sb.WriteString(stringifyExpr(expr.Value))

	return sb.String()
}

func stringifyMapType(expr *ast.MapType) string {
	return fmt.Sprintf("map[%s]%s", stringifyExpr(expr.Key), stringifyExpr(expr.Value))
}

func stringifyInterfaceType(expr *ast.InterfaceType) string {
	sb := new(strings.Builder)

	sb.WriteString("interface{")
	for i, field := range expr.Methods.List {
		if i > 0 {
			sb.WriteString(";")
		}
		for _, name := range field.Names {
			sb.WriteString(name.Name)
		}
		if f, ok := field.Type.(*ast.FuncType); ok {
			sb.WriteString(stringifyFuncTypeParams(f.TypeParams))
			sb.WriteString(stringifyFuncParams(f.Params))
			sb.WriteString(stringifyFuncResults(f.Results))
		} else {
			sb.WriteString(stringifyExpr(field.Type))
		}
	}
	sb.WriteString("}")

	return sb.String()
}

func stringifyStructType(expr *ast.StructType) string {
	sb := new(strings.Builder)

	sb.WriteString("struct{")
	for i, field := range expr.Fields.List {
		if i > 0 {
			sb.WriteString(";")
		}
		for j, name := range field.Names {
			if j > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(name.Name)
		}
		if len(field.Names) > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(stringifyExpr(field.Type))

		if field.Tag != nil {
			sb.WriteString(" ")
			sb.WriteString(field.Tag.Value)
		}
	}
	sb.WriteString("}")

	return sb.String()
}

func stringifyFuncType(expr *ast.FuncType) string {
	sb := new(strings.Builder)

	sb.WriteString("func")
	sb.WriteString(stringifyFuncTypeParams(expr.TypeParams))
	sb.WriteString(stringifyFuncParams(expr.Params))
	sb.WriteString(stringifyFuncResults(expr.Results))

	return sb.String()
}
