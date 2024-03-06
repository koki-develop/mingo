package mingo

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func (m *mingo) stringifyStmt(stmt ast.Stmt) string {
	switch x := stmt.(type) {
	case *ast.ReturnStmt:
		return m.stringifyReturnStmt(x)
	case *ast.AssignStmt:
		return m.stringifyAssignStmt(x)
	case *ast.IfStmt:
		return m.stringifyIfStmt(x)
	case *ast.BlockStmt:
		return m.stringifyBlockStmt(x)
	case *ast.ExprStmt:
		return m.stringifyExprStmt(x)
	case *ast.DeclStmt:
		return m.stringifyDeclStmt(x)
	case *ast.DeferStmt:
		return m.stringifyDeferStmt(x)
	case *ast.GoStmt:
		return m.stringifyGoStmt(x)
	case *ast.LabeledStmt:
		return m.stringifyLabeledStmt(x)
	case *ast.SwitchStmt:
		return m.stringifySwitchStmt(x)
	case *ast.SelectStmt:
		return m.stringifySelectStmt(x)
	case *ast.ForStmt:
		return m.stringifyForStmt(x)
	case *ast.RangeStmt:
		return m.stringifyRangeStmt(x)
	case *ast.BranchStmt:
		return m.stringifyBranchStmt(x)
	case *ast.EmptyStmt:
		return m.stringifyEmptyStmt(x)
	case *ast.IncDecStmt:
		return m.stringifyIncDecStmt(x)
	case *ast.SendStmt:
		return m.stringifySendStmt(x)
	case *ast.CaseClause:
		return m.stringifyCaseCaluse(x)
	case *ast.CommClause:
		return m.stringifyCommClause(x)
	case *ast.TypeSwitchStmt:
		return m.stringifyTypeSwitchStmt(x)
	default:
		panic(fmt.Sprintf("unhandled stmt: %#v", x))
	}
}

func (m *mingo) stringifyBlockStmt(stmt *ast.BlockStmt) string {
	sb := new(strings.Builder)

	sb.WriteString("{")
	for i, child := range stmt.List {
		sb.WriteString(m.stringifyStmt(child))

		if _, ok := child.(*ast.DeclStmt); !ok {
			if i < len(stmt.List)-1 {
				sb.WriteString(";")
			}
		}
	}
	sb.WriteString("}")
	return sb.String()
}

func (m *mingo) stringifyReturnStmt(stmt *ast.ReturnStmt) string {
	sb := new(strings.Builder)

	sb.WriteString("return")
	if len(stmt.Results) > 0 {
		sb.WriteString(" ")
	}

	for i, expr := range stmt.Results {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(m.stringifyExpr(expr))
	}

	return sb.String()
}

func (m *mingo) stringifyAssignStmt(stmt *ast.AssignStmt) string {
	sb := new(strings.Builder)

	for i, expr := range stmt.Lhs {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(m.stringifyExpr(expr))
	}
	if stmt.Tok == token.DEFINE {
		sb.WriteString(":=")
	} else {
		sb.WriteString("=")
	}
	for i, expr := range stmt.Rhs {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(m.stringifyExpr(expr))
	}

	return sb.String()
}

func (m *mingo) stringifyIfStmt(stmt *ast.IfStmt) string {
	sb := new(strings.Builder)

	sb.WriteString("if ")
	sb.WriteString(m.stringifyIfStmtBody(stmt))

	return sb.String()
}

func (m *mingo) stringifyElseIfStmt(stmt *ast.IfStmt) string {
	sb := new(strings.Builder)

	sb.WriteString("else if ")
	sb.WriteString(m.stringifyIfStmtBody(stmt))

	return sb.String()
}

func (m *mingo) stringifyIfStmtBody(stmt *ast.IfStmt) string {
	sb := new(strings.Builder)

	if stmt.Init != nil {
		sb.WriteString(m.stringifyStmt(stmt.Init))
		sb.WriteString(";")
	}
	if stmt.Cond != nil {
		sb.WriteString(m.stringifyExpr(stmt.Cond))
	}
	if stmt.Body != nil {
		sb.WriteString(m.stringifyBlockStmt(stmt.Body))
	}
	if stmt.Else != nil {
		if ifstmt, ok := stmt.Else.(*ast.IfStmt); ok {
			sb.WriteString(m.stringifyElseIfStmt(ifstmt))
		} else {
			sb.WriteString("else")
			sb.WriteString(m.stringifyStmt(stmt.Else))
		}
	}

	return sb.String()
}

func (m *mingo) stringifyExprStmt(stmt *ast.ExprStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(m.stringifyExpr(stmt.X))

	return sb.String()
}

func (m *mingo) stringifyDeclStmt(stmt *ast.DeclStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(m.stringifyDecl(stmt.Decl))

	return sb.String()
}

func (m *mingo) stringifyDeferStmt(stmt *ast.DeferStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("defer ")
	sb.WriteString(m.stringifyExpr(stmt.Call))

	return sb.String()
}

func (m *mingo) stringifyGoStmt(stmt *ast.GoStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("go ")
	sb.WriteString(m.stringifyExpr(stmt.Call))

	return sb.String()
}

func (m *mingo) stringifyLabeledStmt(stmt *ast.LabeledStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(fmt.Sprintf("%s:", stmt.Label.Name))
	sb.WriteString(m.stringifyStmt(stmt.Stmt))

	return sb.String()
}

func (m *mingo) stringifySwitchStmt(stmt *ast.SwitchStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("switch ")
	if stmt.Init != nil {
		sb.WriteString(m.stringifyStmt(stmt.Init))
		sb.WriteString(";")
	}
	if stmt.Tag != nil {
		sb.WriteString(m.stringifyExpr(stmt.Tag))
	}
	if stmt.Body != nil {
		sb.WriteString(m.stringifyBlockStmt(stmt.Body))
	}

	return sb.String()
}

func (m *mingo) stringifySelectStmt(stmt *ast.SelectStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("select ")
	sb.WriteString(m.stringifyBlockStmt(stmt.Body))

	return sb.String()
}

func (m *mingo) stringifyForStmt(stmt *ast.ForStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("for ")

	if stmt.Init != nil {
		sb.WriteString(m.stringifyStmt(stmt.Init))
	}
	if stmt.Init != nil || stmt.Post != nil {
		sb.WriteString(";")
	}
	if stmt.Cond != nil {
		sb.WriteString(m.stringifyExpr(stmt.Cond))
	}
	if stmt.Init != nil || stmt.Post != nil {
		sb.WriteString(";")
	}
	if stmt.Post != nil {
		sb.WriteString(m.stringifyStmt(stmt.Post))
	}
	sb.WriteString(m.stringifyBlockStmt(stmt.Body))

	return sb.String()
}

func (m *mingo) stringifyRangeStmt(stmt *ast.RangeStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("for ")

	needAssign := false
	if stmt.Key != nil {
		needAssign = true
		sb.WriteString(m.stringifyExpr(stmt.Key))
	}
	if stmt.Value != nil {
		needAssign = true
		sb.WriteString(",")
		sb.WriteString(m.stringifyExpr(stmt.Value))
	}

	if needAssign {
		sb.WriteString(":=")
	}

	sb.WriteString("range ")
	sb.WriteString(m.stringifyExpr(stmt.X))
	sb.WriteString(m.stringifyBlockStmt(stmt.Body))

	return sb.String()
}

func (m *mingo) stringifyBranchStmt(stmt *ast.BranchStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(stmt.Tok.String())
	if stmt.Label != nil {
		sb.WriteString(" ")
		sb.WriteString(stmt.Label.Name)
	}
	return sb.String()
}

func (m *mingo) stringifyEmptyStmt(_ *ast.EmptyStmt) string {
	return ""
}

func (m *mingo) stringifyIncDecStmt(stmt *ast.IncDecStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(m.stringifyExpr(stmt.X))
	sb.WriteString(stmt.Tok.String())

	return sb.String()
}

func (m *mingo) stringifySendStmt(stmt *ast.SendStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(m.stringifyExpr(stmt.Chan))
	sb.WriteString("<-")
	sb.WriteString(m.stringifyExpr(stmt.Value))

	return sb.String()
}

func (m *mingo) stringifyCaseCaluse(stmt *ast.CaseClause) string {
	sb := new(strings.Builder)

	if len(stmt.List) > 0 {
		sb.WriteString("case ")
		for i, expr := range stmt.List {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(m.stringifyExpr(expr))
		}
		sb.WriteString(":")
	} else {
		sb.WriteString("default:")
	}
	for i, child := range stmt.Body {
		sb.WriteString(m.stringifyStmt(child))
		if i < len(stmt.Body)-1 {
			sb.WriteString(";")
		}
	}

	return sb.String()
}

func (m *mingo) stringifyCommClause(stmt *ast.CommClause) string {
	sb := new(strings.Builder)

	if stmt.Comm != nil {
		sb.WriteString("case ")
		sb.WriteString(m.stringifyStmt(stmt.Comm))
		sb.WriteString(":")
	} else {
		sb.WriteString("default:")
	}
	for _, stmt := range stmt.Body {
		sb.WriteString(m.stringifyStmt(stmt))
		sb.WriteString(";")
	}

	return sb.String()
}

func (m *mingo) stringifyTypeSwitchStmt(stmt *ast.TypeSwitchStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("switch ")
	if stmt.Init != nil {
		sb.WriteString(m.stringifyStmt(stmt.Init))
		sb.WriteString(";")
	}
	if stmt.Assign != nil {
		sb.WriteString(m.stringifyStmt(stmt.Assign))
	}
	if stmt.Body != nil {
		sb.WriteString(m.stringifyBlockStmt(stmt.Body))
	}

	return sb.String()
}
