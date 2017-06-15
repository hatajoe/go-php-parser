package ast

type IncrementType int

const (
	PostInc IncrementType = iota + 1
	PreInc
)
