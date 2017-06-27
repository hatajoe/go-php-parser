package ast

import (
	"bytes"
)

type SwitchCaseListStatement struct {
	Stmts        []Statement
	HasSemiColon bool
}

func NewSwitchCaseListStatement(stmts []Statement, hasSemiColon bool) *SwitchCaseListStatement {
	return &SwitchCaseListStatement{
		Stmts:        stmts,
		HasSemiColon: hasSemiColon,
	}
}

func (scls *SwitchCaseListStatement) statementNode()       {}
func (scls *SwitchCaseListStatement) TokenLiteral() string { return "" }
func (scls *SwitchCaseListStatement) String() string {
	var out bytes.Buffer

	out.WriteString("{")
	if scls.HasSemiColon {
		out.WriteString(";")
	}
	out.WriteString("\n")
	for _, stmt := range scls.Stmts {
		out.WriteString(stmt.String())
	}
	out.WriteString("}")

	return out.String()
}
