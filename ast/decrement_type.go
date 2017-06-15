package ast

type DecrementType int

const (
	PostDec DecrementType = iota + 1
	PreDec
)
