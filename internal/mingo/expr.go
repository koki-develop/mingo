package mingo

import (
	"fmt"
	"go/ast"
	"strings"
)

func (m *mingo) stringifyExpr(expr ast.Expr) string {
	switch x := expr.(type) {
	case *ast.BasicLit:
		return m.stringifyBasicLit(x)
	case *ast.CallExpr:
		return m.stringifyCallExpr(x)
	case *ast.SelectorExpr:
		return m.stringifySelectExpr(x)
	case *ast.StarExpr:
		return m.stringifyStarExpr(x)
	case *ast.ArrayType:
		return m.stringifyArrayType(x)
	case *ast.Ellipsis:
		return m.stringifyEllipsis(x)
	case *ast.FuncLit:
		return m.stringifyFuncLit(x)
	case *ast.BinaryExpr:
		return m.stringifyBinaryExpr(x)
	case *ast.SliceExpr:
		return m.stringifySliceExpr(x)
	case *ast.UnaryExpr:
		return m.stringifyUnaryExpr(x)
	case *ast.CompositeLit:
		return m.stringifyCompositeLit(x)
	case *ast.ParenExpr:
		return m.stringifyParenExpr(x)
	case *ast.IndexExpr:
		return m.stringifyIndexExpr(x)
	case *ast.KeyValueExpr:
		return m.stringifyKeyValueExpr(x)
	case *ast.TypeAssertExpr:
		return m.stringifyTypeAssertExpr(x)
	case *ast.ChanType:
		return m.stringifyChanType(x)
	case *ast.MapType:
		return m.stringifyMapType(x)
	case *ast.InterfaceType:
		return m.stringifyInterfaceType(x)
	case *ast.StructType:
		return m.stringifyStructType(x)
	case *ast.FuncType:
		return m.stringifyFuncType(x)
	case *ast.IndexListExpr:
		return m.stringifyIndexListExpr(x)
	case nil:
		return ""
	}

	return expr.(*ast.Ident).Name
}

func (m *mingo) stringifyIndexListExpr(expr *ast.IndexListExpr) string {
	sb := new(strings.Builder)

	sb.WriteString(m.stringifyExpr(expr.X))
	sb.WriteString("[")
	for i, index := range expr.Indices {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(m.stringifyExpr(index))
	}
	sb.WriteString("]")

	return sb.String()
}

func (m *mingo) stringifySelectExpr(expr *ast.SelectorExpr) string {
	switch x := expr.X.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", m.stringifySelectExpr(x), expr.Sel.Name)
	default:
		return fmt.Sprintf("%s.%s", m.stringifyExpr(expr.X), expr.Sel.Name)
	}
}

func (m *mingo) stringifyBasicLit(lit *ast.BasicLit) string {
	return lit.Value
}

func (m *mingo) stringifyCallExpr(expr *ast.CallExpr) string {
	sb := new(strings.Builder)

	sb.WriteString(m.stringifyExpr(expr.Fun))
	sb.WriteString("(")
	for i, arg := range expr.Args {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(m.stringifyExpr(arg))
	}

	if expr.Ellipsis.IsValid() {
		sb.WriteString("...")
	}

	sb.WriteString(")")

	return sb.String()
}

func (m *mingo) stringifyStarExpr(expr *ast.StarExpr) string {
	return fmt.Sprintf("*%s", m.stringifyExpr(expr.X))
}

func (m *mingo) stringifyArrayType(expr *ast.ArrayType) string {
	return fmt.Sprintf("[]%s", m.stringifyExpr(expr.Elt))
}

func (m *mingo) stringifyEllipsis(expr *ast.Ellipsis) string {
	return fmt.Sprintf("...%s", m.stringifyExpr(expr.Elt))
}

func (m *mingo) stringifyFuncLit(expr *ast.FuncLit) string {
	sb := new(strings.Builder)
	sb.WriteString(m.stringifyFuncType(expr.Type))
	sb.WriteString(m.stringifyBlockStmt(expr.Body))

	return sb.String()
}

func (m *mingo) stringifyBinaryExpr(expr *ast.BinaryExpr) string {
	return fmt.Sprintf("%s%s%s", m.stringifyExpr(expr.X), expr.Op.String(), m.stringifyExpr(expr.Y))
}

func (m *mingo) stringifySliceExpr(expr *ast.SliceExpr) string {
	sb := new(strings.Builder)

	sb.WriteString(m.stringifyExpr(expr.X))
	sb.WriteString("[")
	if expr.Low != nil {
		sb.WriteString(m.stringifyExpr(expr.Low))
	}
	sb.WriteString(":")
	if expr.High != nil {
		sb.WriteString(m.stringifyExpr(expr.High))
	}
	if expr.Max != nil {
		sb.WriteString(":")
		sb.WriteString(m.stringifyExpr(expr.Max))
	}
	sb.WriteString("]")

	return sb.String()
}

func (m *mingo) stringifyUnaryExpr(expr *ast.UnaryExpr) string {
	return fmt.Sprintf("%s%s", expr.Op.String(), m.stringifyExpr(expr.X))
}

func (m *mingo) stringifyCompositeLit(expr *ast.CompositeLit) string {
	sb := new(strings.Builder)

	sb.WriteString(m.stringifyExpr(expr.Type))
	sb.WriteString("{")
	for i, elt := range expr.Elts {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(m.stringifyExpr(elt))
	}
	sb.WriteString("}")

	return sb.String()
}

func (m *mingo) stringifyParenExpr(expr *ast.ParenExpr) string {
	return fmt.Sprintf("(%s)", m.stringifyExpr(expr.X))
}

func (m *mingo) stringifyIndexExpr(expr *ast.IndexExpr) string {
	return fmt.Sprintf("%s[%s]", m.stringifyExpr(expr.X), m.stringifyExpr(expr.Index))
}

func (m *mingo) stringifyKeyValueExpr(expr *ast.KeyValueExpr) string {
	return fmt.Sprintf("%s:%s", m.stringifyExpr(expr.Key), m.stringifyExpr(expr.Value))
}

func (m *mingo) stringifyTypeAssertExpr(expr *ast.TypeAssertExpr) string {
	if expr.Type == nil {
		return fmt.Sprintf("%s.(type)", m.stringifyExpr(expr.X))
	} else {
		return fmt.Sprintf("%s.(%s)", m.stringifyExpr(expr.X), m.stringifyExpr(expr.Type))
	}
}

func (m *mingo) stringifyChanType(expr *ast.ChanType) string {
	sb := new(strings.Builder)

	if expr.Dir == ast.RECV {
		sb.WriteString("<-")
	}
	sb.WriteString("chan ")
	sb.WriteString(m.stringifyExpr(expr.Value))

	return sb.String()
}

func (m *mingo) stringifyMapType(expr *ast.MapType) string {
	return fmt.Sprintf("map[%s]%s", m.stringifyExpr(expr.Key), m.stringifyExpr(expr.Value))
}

func (m *mingo) stringifyInterfaceType(expr *ast.InterfaceType) string {
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
			sb.WriteString(m.stringifyFuncTypeParams(f.TypeParams))
			sb.WriteString(m.stringifyFuncParams(f.Params))
			sb.WriteString(m.stringifyFuncResults(f.Results))
		} else {
			sb.WriteString(m.stringifyExpr(field.Type))
		}
	}
	sb.WriteString("}")

	return sb.String()
}

func (m *mingo) stringifyStructType(expr *ast.StructType) string {
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
		sb.WriteString(m.stringifyExpr(field.Type))

		if field.Tag != nil {
			sb.WriteString(" ")
			sb.WriteString(field.Tag.Value)
		}
	}
	sb.WriteString("}")

	return sb.String()
}

func (m *mingo) stringifyFuncType(expr *ast.FuncType) string {
	sb := new(strings.Builder)

	sb.WriteString("func")
	sb.WriteString(m.stringifyFuncTypeParams(expr.TypeParams))
	sb.WriteString(m.stringifyFuncParams(expr.Params))
	sb.WriteString(m.stringifyFuncResults(expr.Results))

	return sb.String()
}
