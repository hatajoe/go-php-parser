package ast

type AssignType int

const (
	Equal AssignType = iota + 1
	PlusEqual
	MinusEqual
	MulEqual
	PowEqual
	DivEqual
	ConcatEqual
	ModEqual
	AndEqual
	QrEqual
	XorEqual
	SlEqual
	SrEqual
)
