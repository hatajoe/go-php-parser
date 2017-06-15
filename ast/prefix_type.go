package ast

type PrefixType int

const (
	UnaryPlus PrefixType = iota + 1
	UnaryMinus
	BoolNot
	BwNot
	Silence
)
