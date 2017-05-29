package ast

import (
	"bytes"

	"github.com/hatajoe/go-php-parser/token"
)

type HeredocExpression struct {
	BeginToken *token.Token
	EndToken   *token.Token
	Values     []Expression
}

func NewHeredocExpression(beginTok, endTok *token.Token, values ...Expression) *HeredocExpression {
	return &HeredocExpression{
		BeginToken: beginTok,
		EndToken:   endTok,
		Values:     values,
	}
}

func (he *HeredocExpression) expressionNode()      {}
func (he *HeredocExpression) TokenLiteral() string { return he.BeginToken.Literal }
func (he *HeredocExpression) String() string {
	var out bytes.Buffer

	out.WriteString(he.BeginToken.Literal)
	for _, v := range he.Values {
		out.WriteString(v.String())
	}
	out.WriteString(he.EndToken.Literal)

	return out.String()
}
