package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type ClassDeclarationStatement struct {
	ClassModifiers  []Expression
	Token           *token.Token
	ClassName       *token.Token
	ExtendsFrom     Expression
	ImplementsList  []Expression
	ClassStatements []Statement
}

func NewClassDeclarationStatement(classModifiers []Expression, classLiteral, className *token.Token, extendsFrom Expression, implementsList []Expression, classStatements []Statement) *ClassDeclarationStatement {
	return &ClassDeclarationStatement{
		ClassModifiers:  classModifiers,
		Token:           classLiteral,
		ClassName:       className,
		ExtendsFrom:     extendsFrom,
		ImplementsList:  implementsList,
		ClassStatements: classStatements,
	}
}

func (cds *ClassDeclarationStatement) statementNode()       {}
func (cds *ClassDeclarationStatement) TokenLiteral() string { return cds.Token.Literal }
func (cds *ClassDeclarationStatement) String() string {
	classModifiers := ""
	cm := make([]string, 0, len(cds.ClassModifiers))
	for _, c := range cds.ClassModifiers {
		cm = append(cm, c.String())
	}
	if len(cm) > 0 {
		classModifiers = strings.Join(cm, " ") + " "
	}

	extends := ""
	if cds.ExtendsFrom != nil {
		extends = " extends " + cds.ExtendsFrom.String()
	}

	implements := ""
	is := make([]string, 0, len(cds.ImplementsList))
	for _, i := range cds.ImplementsList {
		is = append(is, i.String())
	}
	if len(is) > 0 {
		implements = strings.Join(is, ", ")
	}

	statements := ""
	st := make([]string, 0, len(cds.ClassStatements))
	for _, s := range cds.ClassStatements {
		st = append(st, "    "+s.String())
	}
	if len(st) > 0 {
		statements = strings.Join(st, "\n")
	}

	return fmt.Sprintf("%s%s %s%s%s\n{\n%s\n}\n", classModifiers, cds.TokenLiteral(), cds.ClassName.Literal, extends, implements, statements)
}
