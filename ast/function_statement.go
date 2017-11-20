package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type FunctionStatement struct {
	Token      *token.Token
	ReturnRef  int
	Name       *token.Token
	Parameters []Expression
	ReturnType Expression
	Stmts      []Statement
}

func NewFunctionStatement(tok *token.Token, returnRef int, name *token.Token, params []Expression, returnType Expression, stmts []Statement) *FunctionStatement {
	return &FunctionStatement{
		Token:      tok,
		ReturnRef:  returnRef,
		Name:       name,
		Parameters: params,
		ReturnType: returnType,
		Stmts:      stmts,
	}
}

func (fs *FunctionStatement) statementNode()       {}
func (fs *FunctionStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *FunctionStatement) String() string {
	ref := ""
	if fs.ReturnRef != 0 {
		ref = "&"
	}
	ret := ""
	if fs.ReturnType != nil {
		ret = fs.ReturnType.String() + " "
	}
	params := make([]string, 0, len(fs.Parameters))
	for _, p := range fs.Parameters {
		params = append(params, p.String())
	}
	stmts := make([]string, 0, len(fs.Stmts))
	for _, stmt := range fs.Stmts {
		stmts = append(stmts, "    "+stmt.String())
	}
	return fmt.Sprintf("%s %s%s(%s)%s{\n%s\n}", fs.TokenLiteral(), ref, fs.Name.Literal, strings.Join(params, ", "), ret, strings.Join(stmts, "\n"))
}
