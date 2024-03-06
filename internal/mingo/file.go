package mingo

import (
	"fmt"
	"go/ast"
)

func (m *mingo) stringifyFile(n *ast.File) string {
	return fmt.Sprintf("package %s;", n.Name.Name)
}
