package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type VarLiteral struct {
	Token *token.Token
	Value string
}

func NewVarLiteral(tok *token.Token, value string) *VarLiteral {
	return &VarLiteral{
		Token: tok,
		Value: value,
	}
}

func (v *VarLiteral) expressionNode()      {}
func (v *VarLiteral) TokenLiteral() string { return v.Token.Literal }
func (v *VarLiteral) String() string       { return v.Value }
