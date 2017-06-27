package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type GlobalStatement struct {
	Token  *token.Token
	Values []Expression
}

func NewGlobalStatement(tok *token.Token, values []Expression) *GlobalStatement {
	return &GlobalStatement{
		Token:  tok,
		Values: values,
	}
}

func (gs *GlobalStatement) statementNode()       {}
func (gs *GlobalStatement) TokenLiteral() string { return gs.Token.Literal }
func (gs *GlobalStatement) String() string {
	values := make([]string, 0, len(gs.Values))
	for _, value := range gs.Values {
		values = append(values, value.String())
	}
	return fmt.Sprintf("%s %s;", gs.TokenLiteral(), strings.Join(values, ", "))
}
