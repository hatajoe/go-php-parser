package ast

import (
	"bytes"
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type IfStatement struct {
	Token *token.Token
	Cond  Expression
	Stmt  Statement
	Next  Statement
}

func NewIfStatement(tok *token.Token, cond Expression, stmt Statement, next Statement) *IfStatement {
	return &IfStatement{
		Token: tok,
		Cond:  cond,
		Stmt:  stmt,
		Next:  next,
	}
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IfStatement) String() string {

	var out bytes.Buffer

	if is.Next != nil {
		out.WriteString(is.Next.String())
	}

	switch is.TokenLiteral() {
	case "if":
		out.WriteString(fmt.Sprintf("%s (%s) %s", is.TokenLiteral(), is.Cond.String(), is.Stmt.String()))
	case "elseif":
		out.WriteString(fmt.Sprintf(" %s (%s) %s", is.TokenLiteral(), is.Cond.String(), is.Stmt.String()))
	case "else":
		out.WriteString(fmt.Sprintf(" %s %s", is.TokenLiteral(), is.Stmt.String()))
	default:
		panic(fmt.Sprintf("undefined token literal specified %s", is.TokenLiteral()))
	}

	return out.String()
}
