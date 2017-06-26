package ast

import (
	"bytes"
)

type SwitchCaseListStatement struct {
	Stmts    []Statement
	HasComma bool
}

func NewSwitchCaseListStatement(stmts []Statement, hasComma bool) *SwitchCaseListStatement {
	return &SwitchCaseListStatement{
		Stmts:    stmts,
		HasComma: hasComma,
	}
}

func (scls *SwitchCaseListStatement) statementNode()       {}
func (scls *SwitchCaseListStatement) TokenLiteral() string { return "" }
func (scls *SwitchCaseListStatement) String() string {
	var out bytes.Buffer

	out.WriteString("{")
	if scls.HasComma {
		out.WriteString(";")
	}
	out.WriteString("\n")
	for _, stmt := range scls.Stmts {
		out.WriteString(stmt.String())
	}
	out.WriteString("}")

	return out.String()
}
