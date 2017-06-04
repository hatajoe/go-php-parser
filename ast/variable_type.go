package ast

type VariableType int

const (
	Var                        VariableType = iota + 1 // $foo
	Dim                                                // $foo[
	Prop                                               // $foo->
	DollarOpenCurlyBraces                              // ${
	DimInDollarOpenCurlyBraces                         // ${foo[
	CurlyOpen                                          // {
	MultipleDollars                                    // $$foo
	Wrapped                                            // ($foo)
)
