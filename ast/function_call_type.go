package ast

type FunctionCallType int

const (
	Call FunctionCallType = iota + 1
	StaticCall
)
