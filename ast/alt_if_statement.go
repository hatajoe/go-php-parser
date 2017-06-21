package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type AltIfStatement struct {
	Token *token.Token
	Cond  Expression
	Stmts []Statement
	Next  Statement
}

func NewAltIfStatement(tok *token.Token, cond Expression, stmts []Statement, next Statement) *AltIfStatement {
	return &AltIfStatement{
		Token: tok,
		Cond:  cond,
		Stmts: stmts,
		Next:  next,
	}
}

func (is *AltIfStatement) statementNode()       {}
func (is *AltIfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *AltIfStatement) String() string {

	var out bytes.Buffer

	if is.Next != nil {
		out.WriteString(is.Next.String())
	}

	stmts := make([]string, 0, len(is.Stmts))
	if len(is.Stmts) > 0 {
		for _, stmt := range is.Stmts {
			stmts = append(stmts, "    "+stmt.String())
		}
	}

	switch is.TokenLiteral() {
	case "if", "elseif":

		out.WriteString(fmt.Sprintf("%s (%s):\n%s\n", is.TokenLiteral(), is.Cond.String(), strings.Join(stmts, "\n")))
	case "else":
		out.WriteString(fmt.Sprintf("%s:\n%s\n", is.TokenLiteral(), strings.Join(stmts, "\n")))
	case "endif":
		out.WriteString(fmt.Sprintf("%s;", is.TokenLiteral()))
	default:
		panic(fmt.Sprintf("undefined token literal specified %s", is.TokenLiteral()))
	}

	return out.String()
}
