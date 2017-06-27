package ast

import (
	"bytes"
)

type AltForStatement struct {
	Stmts []Statement
}

func NewAltForStatement(stmts []Statement) *AltForStatement {
	return &AltForStatement{
		Stmts: stmts,
	}
}

func (afs *AltForStatement) statementNode()       {}
func (afs *AltForStatement) TokenLiteral() string { return "" }
func (afs *AltForStatement) String() string {
	var out bytes.Buffer

	out.WriteString(":\n")
	for _, stmt := range afs.Stmts {
		out.WriteString("    " + stmt.String() + "\n")
	}
	out.WriteString("endfor;")

	return out.String()
}
