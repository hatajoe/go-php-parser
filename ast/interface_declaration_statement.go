package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type InterfaceDeclarationStatement struct {
	Token           *token.Token
	InterfaceName   *token.Token
	ExtendsFrom     []Expression
	ClassStatements []Statement
}

func NewInterfaceDeclarationStatement(classLiteral, className *token.Token, extendsFrom []Expression, classStatements []Statement) *InterfaceDeclarationStatement {
	return &InterfaceDeclarationStatement{
		Token:           classLiteral,
		InterfaceName:   className,
		ExtendsFrom:     extendsFrom,
		ClassStatements: classStatements,
	}
}

func (ids *InterfaceDeclarationStatement) statementNode()       {}
func (ids *InterfaceDeclarationStatement) TokenLiteral() string { return ids.Token.Literal }
func (ids *InterfaceDeclarationStatement) String() string {
	extends := ""
	ex := make([]string, 0, len(ids.ExtendsFrom))
	for _, e := range ids.ExtendsFrom {
		ex = append(ex, e.String())
	}
	if len(ex) > 0 {
		extends = fmt.Sprintf(" extends %s", strings.Join(ex, ", "))
	}

	statements := ""
	st := make([]string, 0, len(ids.ClassStatements))
	for _, s := range ids.ClassStatements {
		st = append(st, "    "+s.String())
	}
	if len(st) > 0 {
		statements = strings.Join(st, "\n")
	}

	return fmt.Sprintf("%s %s%s\n{\n%s\n}\n", ids.TokenLiteral(), ids.InterfaceName.Literal, extends, statements)
}
