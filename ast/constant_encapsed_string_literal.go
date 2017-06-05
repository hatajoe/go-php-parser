package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type ConstantEncapsedStringLiteral struct {
	Token *token.Token
	Value string
}

func NewConstantEncapsedStringLiteral(tok *token.Token, value string) *ConstantEncapsedStringLiteral {
	return &ConstantEncapsedStringLiteral{
		Token: tok,
		Value: value,
	}
}

func (cesl *ConstantEncapsedStringLiteral) expressionNode()      {}
func (cesl *ConstantEncapsedStringLiteral) TokenLiteral() string { return cesl.Token.Literal }
func (cesl *ConstantEncapsedStringLiteral) String() string       { return cesl.Value }
