package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type StaticStatement struct {
	Token  *token.Token
	Values []Expression
}

func NewStaticStatement(tok *token.Token, values []Expression) *StaticStatement {
	return &StaticStatement{
		Token:  tok,
		Values: values,
	}
}

func (ss *StaticStatement) statementNode()       {}
func (ss *StaticStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *StaticStatement) String() string {
	values := make([]string, 0, len(ss.Values))
	for _, value := range ss.Values {
		values = append(values, value.String())
	}
	return fmt.Sprintf("%s %s;", ss.TokenLiteral(), strings.Join(values, ", "))
}
