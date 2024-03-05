package mingo

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func stringifyStmt(stmt ast.Stmt) string {
	switch x := stmt.(type) {
	case *ast.ReturnStmt:
		return stringifyReturnStmt(x)
	case *ast.AssignStmt:
		return stringifyAssignStmt(x)
	case *ast.IfStmt:
		return stringifyIfStmt(x)
	case *ast.BlockStmt:
		return stringifyBlockStmt(x)
	case *ast.ExprStmt:
		return stringifyExprStmt(x)
	case *ast.DeclStmt:
		return stringifyDeclStmt(x)
	case *ast.DeferStmt:
		return stringifyDeferStmt(x)
	case *ast.GoStmt:
		return stringifyGoStmt(x)
	case *ast.LabeledStmt:
		return stringifyLabeledStmt(x)
	case *ast.SwitchStmt:
		return stringifySwitchStmt(x)
	case *ast.SelectStmt:
		return stringifySelectStmt(x)
	case *ast.ForStmt:
		return stringifyForStmt(x)
	case *ast.RangeStmt:
		return stringifyRangeStmt(x)
	case *ast.BranchStmt:
		return stringifyBranchStmt(x)
	case *ast.EmptyStmt:
		return stringifyEmptyStmt(x)
	case *ast.IncDecStmt:
		return stringifyIncDecStmt(x)
	case *ast.SendStmt:
		return stringifySendStmt(x)
	case *ast.CaseClause:
		return stringifyCaseCaluse(x)
	case *ast.CommClause:
		return stringifyCommClause(x)
	case *ast.TypeSwitchStmt:
		return stringifyTypeSwitchStmt(x)
	default:
		panic(fmt.Sprintf("unhandled stmt: %#v", x))
	}
}

func stringifyBlockStmt(stmt *ast.BlockStmt) string {
	sb := new(strings.Builder)

	sb.WriteString("{")
	for i, child := range stmt.List {
		sb.WriteString(stringifyStmt(child))
		if i < len(stmt.List)-1 {
			sb.WriteString(";")
		}
	}
	sb.WriteString("}")
	return sb.String()
}

func stringifyReturnStmt(stmt *ast.ReturnStmt) string {
	sb := new(strings.Builder)

	sb.WriteString("return")
	if len(stmt.Results) > 0 {
		sb.WriteString(" ")
	}

	for i, expr := range stmt.Results {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(stringifyExpr(expr))
	}

	return sb.String()
}

func stringifyAssignStmt(stmt *ast.AssignStmt) string {
	sb := new(strings.Builder)

	for i, expr := range stmt.Lhs {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(stringifyExpr(expr))
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
		sb.WriteString(stringifyExpr(expr))
	}

	return sb.String()
}

func stringifyIfStmt(stmt *ast.IfStmt) string {
	sb := new(strings.Builder)

	sb.WriteString("if ")
	sb.WriteString(stringifyIfStmtBody(stmt))

	return sb.String()
}

func stringifyElseIfStmt(stmt *ast.IfStmt) string {
	sb := new(strings.Builder)

	sb.WriteString("else if ")
	sb.WriteString(stringifyIfStmtBody(stmt))

	return sb.String()
}

func stringifyIfStmtBody(stmt *ast.IfStmt) string {
	sb := new(strings.Builder)

	if stmt.Init != nil {
		sb.WriteString(stringifyStmt(stmt.Init))
		sb.WriteString(";")
	}
	if stmt.Cond != nil {
		sb.WriteString(stringifyExpr(stmt.Cond))
	}
	if stmt.Body != nil {
		sb.WriteString(stringifyBlockStmt(stmt.Body))
	}
	if stmt.Else != nil {
		if ifstmt, ok := stmt.Else.(*ast.IfStmt); ok {
			sb.WriteString(stringifyElseIfStmt(ifstmt))
		} else {
			sb.WriteString("else")
			sb.WriteString(stringifyStmt(stmt.Else))
		}
	}

	return sb.String()
}

func stringifyExprStmt(stmt *ast.ExprStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(stringifyExpr(stmt.X))

	return sb.String()
}

func stringifyDeclStmt(stmt *ast.DeclStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(stringifyDecl(stmt.Decl))

	return sb.String()
}

func stringifyDeferStmt(stmt *ast.DeferStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("defer ")
	sb.WriteString(stringifyExpr(stmt.Call))

	return sb.String()
}

func stringifyGoStmt(stmt *ast.GoStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("go ")
	sb.WriteString(stringifyExpr(stmt.Call))

	return sb.String()
}

func stringifyLabeledStmt(stmt *ast.LabeledStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(fmt.Sprintf("%s:", stmt.Label.Name))
	sb.WriteString(stringifyStmt(stmt.Stmt))

	return sb.String()
}

func stringifySwitchStmt(stmt *ast.SwitchStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("switch ")
	if stmt.Init != nil {
		sb.WriteString(stringifyStmt(stmt.Init))
	}
	if stmt.Tag != nil {
		sb.WriteString(stringifyExpr(stmt.Tag))
	}
	if stmt.Body != nil {
		sb.WriteString(stringifyBlockStmt(stmt.Body))
	}

	return sb.String()
}

func stringifySelectStmt(stmt *ast.SelectStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("select ")
	sb.WriteString(stringifyBlockStmt(stmt.Body))

	return sb.String()
}

func stringifyForStmt(stmt *ast.ForStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("for ")

	if stmt.Init != nil {
		sb.WriteString(stringifyStmt(stmt.Init))
		if stmt.Cond != nil || stmt.Post != nil {
			sb.WriteString(";")
		}
	}
	if stmt.Cond != nil {
		sb.WriteString(stringifyExpr(stmt.Cond))
		if stmt.Post != nil {
			sb.WriteString(";")
		}
	}
	if stmt.Post != nil {
		sb.WriteString(stringifyStmt(stmt.Post))
	}
	sb.WriteString(stringifyBlockStmt(stmt.Body))

	return sb.String()
}

func stringifyRangeStmt(stmt *ast.RangeStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("for ")

	needAssign := false
	if stmt.Key != nil {
		needAssign = true
		sb.WriteString(stringifyExpr(stmt.Key))
		sb.WriteString(",")
	}
	if stmt.Value != nil {
		needAssign = true
		sb.WriteString(stringifyExpr(stmt.Value))
	}

	if needAssign {
		sb.WriteString(":=")
	}

	sb.WriteString("range ")
	sb.WriteString(stringifyExpr(stmt.X))
	sb.WriteString(stringifyBlockStmt(stmt.Body))

	return sb.String()
}

func stringifyBranchStmt(stmt *ast.BranchStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(stmt.Tok.String())
	return sb.String()
}

func stringifyEmptyStmt(_ *ast.EmptyStmt) string {
	return ""
}

func stringifyIncDecStmt(stmt *ast.IncDecStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(stringifyExpr(stmt.X))
	sb.WriteString(stmt.Tok.String())

	return sb.String()
}

func stringifySendStmt(stmt *ast.SendStmt) string {
	sb := new(strings.Builder)
	sb.WriteString(stringifyExpr(stmt.Chan))
	sb.WriteString("<-")
	sb.WriteString(stringifyExpr(stmt.Value))

	return sb.String()
}

func stringifyCaseCaluse(stmt *ast.CaseClause) string {
	sb := new(strings.Builder)

	if len(stmt.List) > 0 {
		sb.WriteString("case ")
		for i, expr := range stmt.List {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(stringifyExpr(expr))
		}
		sb.WriteString(":")
	} else {
		sb.WriteString("default:")
	}
	for i, child := range stmt.Body {
		sb.WriteString(stringifyStmt(child))
		if i < len(stmt.Body)-1 {
			sb.WriteString(";")
		}
	}

	return sb.String()
}

func stringifyCommClause(stmt *ast.CommClause) string {
	sb := new(strings.Builder)

	if stmt.Comm != nil {
		sb.WriteString("case ")
		sb.WriteString(stringifyStmt(stmt.Comm))
		sb.WriteString(":")
	} else {
		sb.WriteString("default:")
	}
	for _, stmt := range stmt.Body {
		sb.WriteString(stringifyStmt(stmt))
		sb.WriteString(";")
	}

	return sb.String()
}

func stringifyTypeSwitchStmt(stmt *ast.TypeSwitchStmt) string {
	sb := new(strings.Builder)
	sb.WriteString("switch ")
	if stmt.Init != nil {
		sb.WriteString(stringifyStmt(stmt.Init))
	}
	if stmt.Assign != nil {
		sb.WriteString(stringifyStmt(stmt.Assign))
	}
	if stmt.Body != nil {
		sb.WriteString(stringifyBlockStmt(stmt.Body))
	}

	return sb.String()
}
