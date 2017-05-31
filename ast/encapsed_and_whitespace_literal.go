package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type EncapsedAndWhitespaceLiteral struct {
	Token *token.Token
	Value string
}

func NewEncapsedAndWhitespaceLiteral(tok *token.Token, value string) *EncapsedAndWhitespaceLiteral {
	return &EncapsedAndWhitespaceLiteral{
		Token: tok,
		Value: value,
	}
}

func (eawl *EncapsedAndWhitespaceLiteral) expressionNode()      {}
func (eawl *EncapsedAndWhitespaceLiteral) TokenLiteral() string { return eawl.Token.Literal }
func (eawl *EncapsedAndWhitespaceLiteral) String() string       { return eawl.Value }
