package ast

import (
	"bytes"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type NamespaceExpression struct {
	Token         *token.Token
	HeadSeparator *token.Token
	Exprs         []Expression
}

func NewNamespaceExpression(tok, hs *token.Token, exprs ...Expression) *NamespaceExpression {
	return &NamespaceExpression{
		Token:         tok,
		HeadSeparator: hs,
		Exprs:         exprs,
	}
}

func (ne *NamespaceExpression) expressionNode() {}
func (ne *NamespaceExpression) TokenLiteral() string {
	if ne.Token == nil {
		return ""
	}
	return ne.Token.Literal
}
func (ne *NamespaceExpression) String() string {
	var out bytes.Buffer

	exprs := make([]string, 0, len(ne.Exprs))
	for _, expr := range ne.Exprs {
		exprs = append(exprs, expr.String())
	}

	if ne.Token != nil {
		out.WriteString(ne.Token.Literal + " ")
	}
	if ne.HeadSeparator != nil {
		out.WriteString("\\")
	}
	out.WriteString(strings.Join(exprs, "\\"))

	return out.String()
}
