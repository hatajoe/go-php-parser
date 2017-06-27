package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type InlineHTMLStatement struct {
	Token *token.Token
}

func NewInlineHTMLStatement(tok *token.Token) *InlineHTMLStatement {
	return &InlineHTMLStatement{
		Token: tok,
	}
}

func (ihs *InlineHTMLStatement) statementNode()       {}
func (ihs *InlineHTMLStatement) TokenLiteral() string { return ihs.Token.Literal }
func (ihs *InlineHTMLStatement) String() string {
	return ihs.TokenLiteral()
}
