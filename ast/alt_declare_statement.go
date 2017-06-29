package ast

import (
	"bytes"
)

type AltDeclareStatement struct {
	Stmts []Statement
}

func NewAltDeclareStatement(stmts []Statement) *AltDeclareStatement {
	return &AltDeclareStatement{
		Stmts: stmts,
	}
}

func (ads *AltDeclareStatement) statementNode()       {}
func (ads *AltDeclareStatement) TokenLiteral() string { return "" }
func (ads *AltDeclareStatement) String() string {
	var out bytes.Buffer

	out.WriteString(":\n")
	for _, stmt := range ads.Stmts {
		out.WriteString("    " + stmt.String() + "\n")
	}
	out.WriteString("enddeclare;")

	return out.String()
}
