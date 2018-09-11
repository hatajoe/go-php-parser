package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type TraitDeclarationStatement struct {
	Token           *token.Token
	TraitName       *token.Token
	TraitStatements []Statement
}

func NewTraitDeclarationStatement(traitLiteral, traitName *token.Token, traitStatements []Statement) *TraitDeclarationStatement {
	return &TraitDeclarationStatement{
		Token:           traitLiteral,
		TraitName:       traitName,
		TraitStatements: traitStatements,
	}
}

func (tds *TraitDeclarationStatement) statementNode()       {}
func (tds *TraitDeclarationStatement) TokenLiteral() string { return tds.Token.Literal }
func (tds *TraitDeclarationStatement) String() string {
	statements := ""
	st := make([]string, 0, len(tds.TraitStatements))
	for _, s := range tds.TraitStatements {
		st = append(st, "    "+s.String())
	}
	if len(st) > 0 {
		statements = strings.Join(st, "\n")
	}

	return fmt.Sprintf("%s %s\n{\n%s\n}\n", tds.TokenLiteral(), tds.TraitName.Literal, statements)
}
