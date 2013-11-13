package script

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"reflect"
)

// block statement is a list of statements.
type BlockStmt struct {
	children []Expr
	node     []ast.Node
}

// does not enter scope because it does not necessarily needs to (e.g. for, if).
func (w *World) compileBlockStmt_noScope(n *ast.BlockStmt) *BlockStmt {
	b := &BlockStmt{}
	for _, s := range n.List {
		b.append(w.compileStmt(s), s)
	}
	return b
}

func (b *BlockStmt) append(s Expr, n ast.Node) {
	b.children = append(b.children, s)
	b.node = append(b.node, n)
}

func (b *BlockStmt) Eval() interface{} {
	for _, s := range b.children {
		s.Eval()
	}
	return nil
}

func (b *BlockStmt) Type() reflect.Type {
	return nil
}

func (b *BlockStmt) Child() []Expr {
	return b.children
}

func (b *BlockStmt) Format() string {
	var buf bytes.Buffer
	fset := token.NewFileSet()
	for i := range b.children {
		format.Node(&buf, fset, b.node[i])
		fmt.Fprintln(&buf)
	}
	return buf.String()
}
