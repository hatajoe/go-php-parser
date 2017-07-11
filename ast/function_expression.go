package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type FunctionExpression struct {
	Token      *token.Token
	ReturnRef  int
	Parameters []Expression
	LexicalVar Expression
	ReturnType Expression
	Stmts      []Statement
}

func NewFunctionExpression(tok *token.Token, returnRef int, params []Expression, lexicalVar, returnType Expression, stmts []Statement) *FunctionExpression {
	return &FunctionExpression{
		Token:      tok,
		ReturnRef:  returnRef,
		Parameters: params,
		LexicalVar: lexicalVar,
		ReturnType: returnType,
		Stmts:      stmts,
	}
}

func (fe *FunctionExpression) expressionNode()      {}
func (fe *FunctionExpression) TokenLiteral() string { return fe.Token.Literal }
func (fe *FunctionExpression) String() string {
	ref := ""
	if fe.ReturnRef != 0 {
		ref = "&"
	}
	lexical := ""
	if fe.LexicalVar != nil {
		lexical = fe.LexicalVar.String()
	}
	ret := ""
	if fe.ReturnType != nil {
		ret = fe.ReturnType.String() + " "
	} else {
		lexical = lexical + " "
	}
	params := make([]string, 0, len(fe.Parameters))
	for _, p := range fe.Parameters {
		params = append(params, p.String())
	}
	stmts := make([]string, 0, len(fe.Stmts))
	for _, stmt := range fe.Stmts {
		stmts = append(stmts, "    "+stmt.String())
	}
	return fmt.Sprintf("%s %s(%s)%s%s{\n%s\n}", fe.TokenLiteral(), ref, strings.Join(params, ", "), lexical, ret, strings.Join(stmts, "\n"))
}
