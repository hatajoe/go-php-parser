package ast

import (
	"bytes"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type EchoStatement struct {
	Token  *token.Token
	Values []Expression
}

func NewEchoStatement(tok *token.Token, values []Expression) *EchoStatement {
	return &EchoStatement{
		Token:  tok,
		Values: values,
	}
}

func (es *EchoStatement) statementNode()       {}
func (es *EchoStatement) TokenLiteral() string { return es.Token.Literal }
func (es *EchoStatement) String() string {
	var out bytes.Buffer

	values := make([]string, 0, len(es.Values))
	for _, v := range es.Values {
		values = append(values, v.String())
	}

	out.WriteString(es.TokenLiteral() + " ")
	out.WriteString(strings.Join(values, ", "))
	out.WriteString(";")

	return out.String()
}
