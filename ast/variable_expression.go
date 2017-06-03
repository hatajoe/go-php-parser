package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type VariableType int

const (
	Var VariableType = iota + 1
	Dim
	Prop
	DollarOpenCurlyBraces
	DimInDollarOpenCurlyBraces
)

type VariableExpression struct {
	Token *token.Token
	Type  VariableType
	Name  string
	Exprs []Expression
}

func NewVariableExpression(tok *token.Token, t VariableType, name string, exprs ...Expression) *VariableExpression {
	return &VariableExpression{
		Token: tok,
		Type:  t,
		Name:  name,
		Exprs: exprs,
	}
}

func (ve *VariableExpression) expressionNode()      {}
func (ve *VariableExpression) TokenLiteral() string { return ve.Token.Literal }
func (ve *VariableExpression) String() string {
	switch ve.Type {
	case Var:
		return ve.Name
	case Dim:
		return fmt.Sprintf("%s[%s]", ve.Name, ve.Exprs[0].String())
	case Prop:
		return fmt.Sprintf("%s->%s", ve.Name, ve.Exprs[0].String())
	case DollarOpenCurlyBraces:
		return fmt.Sprintf("%s%s}", ve.Name, ve.Exprs[0].String())
	case DimInDollarOpenCurlyBraces:
		return fmt.Sprintf("%s%s[%s]}", ve.Name, ve.Exprs[0].String(), ve.Exprs[1].String())
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", ve.Type))
	}
}
