package ast

type InfixType int

const (
	BooleanOr InfixType = iota + 1
	BooleanAnd
	LogicalOr
	LogicalAnd
	LogicalXor
	BwOr
	BwAnd
	BwXor
	Concat
	Add
	Sub
	Mul
	Pow
	Div
	Mod
	Sl
	Sr
	IsIdentical
	IsNotIdentical
	IsEqual
	IsNotEqual
	Smaller
	SmallerOrEqual
	Greater
	GreaterOrEqual
	Spaceship
	InstanceOf
	Coalesce
)
