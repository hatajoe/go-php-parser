package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type LexicalVariableListExpression struct {
	Token *token.Token
	Exprs []Expression
}

func NewLexicalVariableListExpression(t *token.Token, exprs ...Expression) *LexicalVariableListExpression {
	return &LexicalVariableListExpression{
		Token: t,
		Exprs: exprs,
	}
}

func (lvle *LexicalVariableListExpression) expressionNode()      {}
func (lvle *LexicalVariableListExpression) TokenLiteral() string { return lvle.Token.Literal }
func (lvle *LexicalVariableListExpression) String() string {
	var out bytes.Buffer

	exprs := make([]string, 0, len(lvle.Exprs))
	for _, expr := range lvle.Exprs {
		exprs = append(exprs, expr.String())
	}

	out.WriteString(fmt.Sprintf(" %s (%s)", lvle.TokenLiteral(), strings.Join(exprs, ", ")))

	return out.String()
}
