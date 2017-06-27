package ast

import (
	"bytes"
)

type AltForeachStatement struct {
	Stmts []Statement
}

func NewAltForeachStatement(stmts []Statement) *AltForeachStatement {
	return &AltForeachStatement{
		Stmts: stmts,
	}
}

func (afs *AltForeachStatement) statementNode()       {}
func (afs *AltForeachStatement) TokenLiteral() string { return "" }
func (afs *AltForeachStatement) String() string {
	var out bytes.Buffer

	out.WriteString(":\n")
	for _, stmt := range afs.Stmts {
		out.WriteString("    " + stmt.String() + "\n")
	}
	out.WriteString("endforeach;")

	return out.String()
}
