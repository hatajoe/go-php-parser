package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type UnsetStatement struct {
	Token  *token.Token
	Values []Expression
}

func NewUnsetStatement(tok *token.Token, values []Expression) *UnsetStatement {
	return &UnsetStatement{
		Token:  tok,
		Values: values,
	}
}

func (us *UnsetStatement) statementNode()       {}
func (us *UnsetStatement) TokenLiteral() string { return us.Token.Literal }
func (us *UnsetStatement) String() string {
	values := make([]string, 0, len(us.Values))
	for _, value := range us.Values {
		values = append(values, value.String())
	}
	return fmt.Sprintf("%s(%s);", us.TokenLiteral(), strings.Join(values, ", "))
}
