package ast

import (
	"bytes"
	"strings"
)

type AltSwitchCaseListStatement struct {
	Stmts        []Statement
	HasSemiColon bool
}

func NewAltSwitchCaseListStatement(stmts []Statement, hasSemiColon bool) *AltSwitchCaseListStatement {
	return &AltSwitchCaseListStatement{
		Stmts:        stmts,
		HasSemiColon: hasSemiColon,
	}
}

func (scls *AltSwitchCaseListStatement) statementNode()       {}
func (scls *AltSwitchCaseListStatement) TokenLiteral() string { return "" }
func (scls *AltSwitchCaseListStatement) String() string {
	var out bytes.Buffer

	stmts := make([]string, 0, len(scls.Stmts))
	for _, stmt := range scls.Stmts {
		stmts = append(stmts, stmt.String())
	}

	out.WriteString(":")
	if scls.HasSemiColon {
		out.WriteString(";")
	}
	out.WriteString("\n")
	out.WriteString(strings.Join(stmts, "\n"))
	out.WriteString("endswitch;")

	return out.String()
}
