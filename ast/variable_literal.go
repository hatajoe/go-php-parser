package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type VariableLiteral struct {
	Token *token.Token
	Name  string
}

func NewVariableLiteral(tok *token.Token, name string) *VariableLiteral {
	return &VariableLiteral{
		Token: tok,
		Name:  name,
	}
}

func (vl *VariableLiteral) expressionNode()      {}
func (vl *VariableLiteral) TokenLiteral() string { return vl.Token.Literal }
func (vl *VariableLiteral) String() string {
	return vl.Name
}
