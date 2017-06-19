package ast

import (
	"bytes"
	"strings"
)

type BlockStatement struct {
	Stmts []Statement
}

func NewBlockStatement(stmts ...Statement) *BlockStatement {
	return &BlockStatement{
		Stmts: stmts,
	}
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return "" }
func (bs *BlockStatement) String() string {

	if len(bs.Stmts) <= 0 {
		return "{}"
	}

	var out bytes.Buffer

	stmts := make([]string, 0, len(bs.Stmts))
	for _, v := range bs.Stmts {
		stmts = append(stmts, "    "+v.String())
	}

	out.WriteString("{\n")
	out.WriteString(strings.Join(stmts, "\n"))
	out.WriteString("\n}")

	return out.String()
}
