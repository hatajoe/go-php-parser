package ast

type EmptyStatement struct {
}

func NewEmptyStatement() *EmptyStatement {
	return &EmptyStatement{}
}

func (es *EmptyStatement) statementNode()       {}
func (es *EmptyStatement) TokenLiteral() string { return ";" }
func (es *EmptyStatement) String() string {
	return es.TokenLiteral()
}
