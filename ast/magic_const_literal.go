package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type MagicConstLiteral struct {
	Token *token.Token
	Value string
}

func NewMagicConstLiteral(tok *token.Token, value string) *MagicConstLiteral {
	return &MagicConstLiteral{
		Token: tok,
		Value: value,
	}
}

func (mcl *MagicConstLiteral) expressionNode()      {}
func (mcl *MagicConstLiteral) TokenLiteral() string { return mcl.Token.Literal }
func (mcl *MagicConstLiteral) String() string       { return mcl.Token.Literal }
