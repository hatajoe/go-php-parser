package ast

type VariableType int

const (
	Var                        VariableType = iota + 1 // $foo
	Dim                                                // $foo[
	Curly                                              // $foo{
	Prop                                               // $foo->
	DollarOpenCurlyBraces                              // ${
	DimInDollarOpenCurlyBraces                         // ${foo[
	CurlyOpen                                          // {
	Wrapped                                            // ($foo)
)
