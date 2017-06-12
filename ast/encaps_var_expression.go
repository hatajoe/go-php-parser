package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type EncapsVarExpression struct {
	Token *token.Token
	Name  string
	Type  VariableType
	Exprs []Expression
}

func NewEncapsVarExpression(tok *token.Token, name string, t VariableType, exprs ...Expression) *EncapsVarExpression {
	return &EncapsVarExpression{
		Token: tok,
		Name:  name,
		Type:  t,
		Exprs: exprs,
	}
}

func (eve *EncapsVarExpression) expressionNode()      {}
func (eve *EncapsVarExpression) TokenLiteral() string { return "" }
func (eve *EncapsVarExpression) String() string {
	switch eve.Type {
	case Dim:
		return fmt.Sprintf("%s[%s]", eve.Name, eve.Exprs[0].String())
	case Prop:
		return fmt.Sprintf("%s->%s", eve.Name, eve.Exprs[0].String())
	case DollarOpenCurlyBraces:
		return fmt.Sprintf("%s%s}", eve.Name, eve.Exprs[0].String())
	case DimInDollarOpenCurlyBraces:
		return fmt.Sprintf("%s%s[%s]}", eve.Name, eve.Exprs[0].String(), eve.Exprs[1].String())
	case CurlyOpen:
		return fmt.Sprintf("%s%s}", eve.Name, eve.Exprs[0].String())
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", eve.Type))
	}
}
