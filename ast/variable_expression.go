package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type VariableExpression struct {
	Token *token.Token
	Name  string
}

func NewVariableExpression(tok *token.Token, name string) *VariableExpression {
	return &VariableExpression{
		Token: tok,
		Name:  name,
	}
}

func (ve *VariableExpression) expressionNode()      {}
func (ve *VariableExpression) TokenLiteral() string { return ve.Token.Literal }
func (ve *VariableExpression) String() string       { return ve.Name }
