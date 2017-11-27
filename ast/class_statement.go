package ast

import (
	"fmt"
	"strings"
)

type ClassStatement struct {
	VariableModifiers []Expression
	Properties        []Expression
}

func NewClassStatement(variableModifiers []Expression, properties []Expression) *ClassStatement {
	return &ClassStatement{
		VariableModifiers: variableModifiers,
		Properties:        properties,
	}
}

func (cs *ClassStatement) statementNode()       {}
func (cs *ClassStatement) TokenLiteral() string { return "" }
func (cs *ClassStatement) String() string {
	vm := make([]string, 0, len(cs.VariableModifiers))
	for _, v := range cs.VariableModifiers {
		vm = append(vm, v.String())
	}
	props := make([]string, 0, len(cs.Properties))
	for _, prop := range cs.Properties {
		props = append(props, prop.String())
	}
	return fmt.Sprintf("%s %s;", strings.Join(vm, " "), strings.Join(props, ", "))
}
