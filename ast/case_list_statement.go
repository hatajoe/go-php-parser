package ast

import (
	"bytes"
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type CaseListStatement struct {
	Token        *token.Token
	Expr         Expression
	Stmts        []Statement
	SemiColonSep bool
}

func NewCaseListStatement(tok *token.Token, expr Expression, stmts []Statement, semiColonSep bool) *CaseListStatement {
	return &CaseListStatement{
		Token:        tok,
		Expr:         expr,
		Stmts:        stmts,
		SemiColonSep: semiColonSep,
	}
}

func (cls *CaseListStatement) statementNode()       {}
func (cls *CaseListStatement) TokenLiteral() string { return cls.Token.Literal }
func (cls *CaseListStatement) String() string {

	var out bytes.Buffer

	sep := ":"
	if cls.SemiColonSep {
		sep = ";"
	}

	if cls.Expr != nil {
		out.WriteString(fmt.Sprintf("%s %s%s\n", cls.TokenLiteral(), cls.Expr, sep))
	} else {
		out.WriteString(fmt.Sprintf("%s%s\n", cls.TokenLiteral(), sep))
	}

	for _, stmt := range cls.Stmts {
		out.WriteString("    " + stmt.String() + "\n")
	}

	return out.String()
}
