package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type ForStatement struct {
	Token       *token.Token
	FirstExprs  []Expression
	SecondExprs []Expression
	ThirdExprs  []Expression
	Stmt        Statement
}

func NewForStatement(tok *token.Token, first, second, third []Expression, stmt Statement) *ForStatement {
	return &ForStatement{
		Token:       tok,
		FirstExprs:  first,
		SecondExprs: second,
		ThirdExprs:  third,
		Stmt:        stmt,
	}
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	firstExprs := make([]string, 0, len(fs.FirstExprs))
	for _, expr := range fs.FirstExprs {
		firstExprs = append(firstExprs, expr.String())
	}
	secondExprs := make([]string, 0, len(fs.SecondExprs))
	for _, expr := range fs.SecondExprs {
		secondExprs = append(secondExprs, expr.String())
	}
	thirdExprs := make([]string, 0, len(fs.ThirdExprs))
	for _, expr := range fs.ThirdExprs {
		thirdExprs = append(thirdExprs, expr.String())
	}
	first := strings.Join(firstExprs, ", ")
	second := strings.Join(secondExprs, ", ")
	third := strings.Join(thirdExprs, ", ")
	if second != "" {
		second = " " + second
	}
	if third != "" {
		third = " " + third
	}
	return fmt.Sprintf("%s (%s;%s;%s) %s", fs.TokenLiteral(), first, second, third, fs.Stmt.String())
}
