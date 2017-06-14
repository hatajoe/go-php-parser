//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/hatajoe/go-php-parser/ast"
	"github.com/hatajoe/go-php-parser/lexer"
	"github.com/hatajoe/go-php-parser/token"
	"log"
)

//line parser.go.y:13
type yySymType struct {
	yys     int
	program *ast.Program
	stmts   []ast.Statement
	stmt    ast.Statement
	exprs   []ast.Expression
	expr    ast.Expression
	tok     *token.Token
}

const T_INCLUDE = 57346
const T_INCLUDE_ONCE = 57347
const T_EVAL = 57348
const T_REQUIRE = 57349
const T_REQUIRE_ONCE = 57350
const T_LOGICAL_OR = 57351
const T_LOGICAL_XOR = 57352
const T_LOGICAL_AND = 57353
const T_PRINT = 57354
const T_YIELD = 57355
const T_DOUBLE_ARROW = 57356
const T_YIELD_FROM = 57357
const T_PLUS_EQUAL = 57358
const T_MINUS_EQUAL = 57359
const T_MUL_EQUAL = 57360
const T_DIV_EQUAL = 57361
const T_CONCAT_EQUAL = 57362
const T_MOD_EQUAL = 57363
const T_AND_EQUAL = 57364
const T_OR_EQUAL = 57365
const T_XOR_EQUAL = 57366
const T_SL_EQUAL = 57367
const T_SR_EQUAL = 57368
const T_POW_EQUAL = 57369
const T_COALESCE = 57370
const T_BOOLEAN_OR = 57371
const T_BOOLEAN_AND = 57372
const T_IS_EQUAL = 57373
const T_IS_NOT_EQUAL = 57374
const T_IS_IDENTICAL = 57375
const T_IS_NOT_IDENTICAL = 57376
const T_SPACESHIP = 57377
const T_IS_SMALLER_OR_EQUAL = 57378
const T_IS_GREATER_OR_EQUAL = 57379
const T_SL = 57380
const T_SR = 57381
const T_INSTANCEOF = 57382
const T_INC = 57383
const T_DEC = 57384
const T_INT_CAST = 57385
const T_DOUBLE_CAST = 57386
const T_STRING_CAST = 57387
const T_ARRAY_CAST = 57388
const T_OBJECT_CAST = 57389
const T_BOOL_CAST = 57390
const T_UNSET_CAST = 57391
const T_POW = 57392
const T_NEW = 57393
const T_CLONE = 57394
const T_NOELSE = 57395
const T_ELSEIF = 57396
const T_ELSE = 57397
const T_ENDIF = 57398
const T_STATIC = 57399
const T_ABSTRACT = 57400
const T_FINAL = 57401
const T_PRIVATE = 57402
const T_PROTECTED = 57403
const T_PUBLIC = 57404
const T_ECHO = 57405
const T_LNUMBER = 57406
const T_DNUMBER = 57407
const T_STRING = 57408
const T_VARIABLE = 57409
const T_INLINE_HTML = 57410
const T_ENCAPSED_AND_WHITESPACE = 57411
const T_CONSTANT_ENCAPSED_STRING = 57412
const T_STRING_VARNAME = 57413
const T_NUM_STRING = 57414
const T_LINE = 57415
const T_FILE = 57416
const T_DIR = 57417
const T_CLASS_C = 57418
const T_TRAIT_C = 57419
const T_METHOD_C = 57420
const T_FUNC_C = 57421
const T_NS_C = 57422
const T_EXIT = 57423
const T_IF = 57424
const T_DO = 57425
const T_WHILE = 57426
const T_ENDWHILE = 57427
const T_FOR = 57428
const T_ENDFOR = 57429
const T_FOREACH = 57430
const T_ENDFOREACH = 57431
const T_DECLARE = 57432
const T_ENDDECLARE = 57433
const T_AS = 57434
const T_SWITCH = 57435
const T_ENDSWITCH = 57436
const T_CASE = 57437
const T_DEFAULT = 57438
const T_BREAK = 57439
const T_CONTINUE = 57440
const T_GOTO = 57441
const T_FUNCTION = 57442
const T_CONST = 57443
const T_RETURN = 57444
const T_TRY = 57445
const T_CATCH = 57446
const T_FINALLY = 57447
const T_THROW = 57448
const T_USE = 57449
const T_INSTEADOF = 57450
const T_GLOBAL = 57451
const T_VAR = 57452
const T_UNSET = 57453
const T_ISSET = 57454
const T_EMPTY = 57455
const T_HALT_COMPILER = 57456
const T_CLASS = 57457
const T_TRAIT = 57458
const T_INTERFACE = 57459
const T_EXTENDS = 57460
const T_IMPLEMENTS = 57461
const T_OBJECT_OPERATOR = 57462
const T_LIST = 57463
const T_ARRAY = 57464
const T_CALLABLE = 57465
const T_COMMENT = 57466
const T_DOC_COMMENT = 57467
const T_OPEN_TAG = 57468
const T_OPEN_TAG_WITH_ECHO = 57469
const T_CLOSE_TAG = 57470
const T_WHITESPACE = 57471
const T_START_HEREDOC = 57472
const T_END_HEREDOC = 57473
const T_DOLLAR_OPEN_CURLY_BRACES = 57474
const T_CURLY_OPEN = 57475
const T_PAAMAYIM_NEKUDOTAYIM = 57476
const T_NAMESPACE = 57477
const T_NS_SEPARATOR = 57478
const T_ELLIPSIS = 57479
const T_ERROR = 57480

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"T_INCLUDE",
	"T_INCLUDE_ONCE",
	"T_EVAL",
	"T_REQUIRE",
	"T_REQUIRE_ONCE",
	"','",
	"T_LOGICAL_OR",
	"T_LOGICAL_XOR",
	"T_LOGICAL_AND",
	"T_PRINT",
	"T_YIELD",
	"T_DOUBLE_ARROW",
	"T_YIELD_FROM",
	"'='",
	"T_PLUS_EQUAL",
	"T_MINUS_EQUAL",
	"T_MUL_EQUAL",
	"T_DIV_EQUAL",
	"T_CONCAT_EQUAL",
	"T_MOD_EQUAL",
	"T_AND_EQUAL",
	"T_OR_EQUAL",
	"T_XOR_EQUAL",
	"T_SL_EQUAL",
	"T_SR_EQUAL",
	"T_POW_EQUAL",
	"'?'",
	"':'",
	"T_COALESCE",
	"T_BOOLEAN_OR",
	"T_BOOLEAN_AND",
	"'|'",
	"'^'",
	"'&'",
	"T_IS_EQUAL",
	"T_IS_NOT_EQUAL",
	"T_IS_IDENTICAL",
	"T_IS_NOT_IDENTICAL",
	"T_SPACESHIP",
	"'<'",
	"T_IS_SMALLER_OR_EQUAL",
	"'>'",
	"T_IS_GREATER_OR_EQUAL",
	"T_SL",
	"T_SR",
	"'+'",
	"'-'",
	"'.'",
	"'*'",
	"'/'",
	"'%'",
	"'!'",
	"T_INSTANCEOF",
	"'~'",
	"T_INC",
	"T_DEC",
	"T_INT_CAST",
	"T_DOUBLE_CAST",
	"T_STRING_CAST",
	"T_ARRAY_CAST",
	"T_OBJECT_CAST",
	"T_BOOL_CAST",
	"T_UNSET_CAST",
	"'@'",
	"T_POW",
	"'['",
	"T_NEW",
	"T_CLONE",
	"T_NOELSE",
	"T_ELSEIF",
	"T_ELSE",
	"T_ENDIF",
	"T_STATIC",
	"T_ABSTRACT",
	"T_FINAL",
	"T_PRIVATE",
	"T_PROTECTED",
	"T_PUBLIC",
	"T_ECHO",
	"T_LNUMBER",
	"T_DNUMBER",
	"T_STRING",
	"T_VARIABLE",
	"T_INLINE_HTML",
	"T_ENCAPSED_AND_WHITESPACE",
	"T_CONSTANT_ENCAPSED_STRING",
	"T_STRING_VARNAME",
	"T_NUM_STRING",
	"T_LINE",
	"T_FILE",
	"T_DIR",
	"T_CLASS_C",
	"T_TRAIT_C",
	"T_METHOD_C",
	"T_FUNC_C",
	"T_NS_C",
	"T_EXIT",
	"T_IF",
	"T_DO",
	"T_WHILE",
	"T_ENDWHILE",
	"T_FOR",
	"T_ENDFOR",
	"T_FOREACH",
	"T_ENDFOREACH",
	"T_DECLARE",
	"T_ENDDECLARE",
	"T_AS",
	"T_SWITCH",
	"T_ENDSWITCH",
	"T_CASE",
	"T_DEFAULT",
	"T_BREAK",
	"T_CONTINUE",
	"T_GOTO",
	"T_FUNCTION",
	"T_CONST",
	"T_RETURN",
	"T_TRY",
	"T_CATCH",
	"T_FINALLY",
	"T_THROW",
	"T_USE",
	"T_INSTEADOF",
	"T_GLOBAL",
	"T_VAR",
	"T_UNSET",
	"T_ISSET",
	"T_EMPTY",
	"T_HALT_COMPILER",
	"T_CLASS",
	"T_TRAIT",
	"T_INTERFACE",
	"T_EXTENDS",
	"T_IMPLEMENTS",
	"T_OBJECT_OPERATOR",
	"T_LIST",
	"T_ARRAY",
	"T_CALLABLE",
	"T_COMMENT",
	"T_DOC_COMMENT",
	"T_OPEN_TAG",
	"T_OPEN_TAG_WITH_ECHO",
	"T_CLOSE_TAG",
	"T_WHITESPACE",
	"T_START_HEREDOC",
	"T_END_HEREDOC",
	"T_DOLLAR_OPEN_CURLY_BRACES",
	"T_CURLY_OPEN",
	"T_PAAMAYIM_NEKUDOTAYIM",
	"T_NAMESPACE",
	"T_NS_SEPARATOR",
	"T_ELLIPSIS",
	"T_ERROR",
	"';'",
	"'('",
	"')'",
	"']'",
	"'\"'",
	"'{'",
	"'}'",
	"'$'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:1272

type LexerWrapper struct {
	l         *lexer.Lexer
	recentLit string
	recentPos token.Position
	program   *ast.Program
}

func (l *LexerWrapper) Lex(lval *yySymType) int {
	tok := l.l.Scan()
	if tok.Type == token.EOF {
		return 0
	}
	lval.tok = tok
	l.recentLit = tok.Literal
	l.recentPos = tok.Position
	return int(tok.Type)
}

func (l *LexerWrapper) Error(e string) {
	log.Fatalf("Line %d, Column %d: %q %s", l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func Parse(l *lexer.Lexer) *ast.Program {
	w := LexerWrapper{l: l}
	if yyParse(&w) != 0 {
		panic("Parse error")
	}
	return w.program
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 9,
	69, 71,
	139, 71,
	153, 71,
	163, 71,
	-2, 66,
	-1, 11,
	159, 74,
	-2, 83,
	-1, 23,
	69, 73,
	139, 73,
	153, 73,
	159, 76,
	163, 73,
	-2, 61,
	-1, 38,
	153, 43,
	-2, 63,
	-1, 121,
	9, 105,
	160, 105,
	161, 105,
	-2, 71,
	-1, 123,
	159, 76,
	-2, 73,
	-1, 128,
	159, 91,
	-2, 64,
	-1, 129,
	159, 93,
	-2, 89,
	-1, 133,
	159, 91,
	-2, 65,
	-1, 134,
	159, 93,
	-2, 90,
	-1, 135,
	159, 75,
	-2, 72,
	-1, 156,
	69, 71,
	139, 71,
	153, 71,
	163, 71,
	-2, 24,
	-1, 191,
	9, 104,
	160, 104,
	161, 104,
	-2, 71,
}

const yyPrivate = 57344

const yyLast = 521

var yyAct = [...]int{

	70, 66, 36, 150, 131, 36, 8, 181, 18, 204,
	195, 188, 183, 182, 15, 179, 158, 114, 36, 23,
	202, 40, 196, 76, 193, 185, 169, 157, 24, 25,
	46, 36, 118, 203, 42, 194, 9, 26, 27, 28,
	33, 29, 30, 31, 32, 62, 81, 187, 140, 8,
	96, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 111, 112, 161, 135, 117, 192, 89,
	125, 91, 65, 48, 111, 92, 93, 75, 81, 86,
	122, 37, 130, 143, 37, 14, 41, 147, 74, 141,
	151, 123, 122, 153, 34, 115, 197, 37, 178, 44,
	45, 152, 180, 123, 22, 148, 122, 35, 121, 71,
	37, 82, 83, 81, 46, 64, 160, 123, 68, 154,
	145, 164, 146, 5, 128, 110, 167, 168, 73, 120,
	189, 171, 162, 119, 156, 67, 79, 69, 43, 63,
	149, 15, 136, 82, 83, 6, 2, 127, 40, 81,
	80, 85, 81, 186, 140, 24, 25, 46, 36, 142,
	173, 42, 12, 190, 26, 27, 28, 33, 29, 30,
	31, 32, 84, 11, 122, 81, 184, 77, 82, 83,
	39, 15, 198, 13, 94, 123, 151, 21, 40, 199,
	200, 113, 19, 16, 201, 24, 25, 46, 36, 126,
	133, 42, 191, 20, 26, 27, 28, 33, 29, 30,
	31, 32, 72, 41, 82, 83, 138, 82, 83, 176,
	10, 34, 47, 132, 7, 38, 44, 45, 137, 88,
	139, 22, 155, 165, 35, 139, 137, 37, 163, 78,
	82, 83, 124, 4, 3, 1, 0, 0, 0, 40,
	0, 0, 14, 41, 174, 177, 0, 0, 46, 36,
	175, 34, 42, 0, 17, 15, 44, 45, 152, 90,
	0, 22, 40, 95, 35, 0, 0, 37, 0, 24,
	25, 46, 36, 0, 0, 42, 0, 0, 26, 27,
	28, 33, 29, 30, 31, 32, 15, 0, 0, 0,
	0, 0, 87, 40, 0, 0, 0, 0, 0, 0,
	24, 25, 46, 36, 41, 0, 42, 144, 0, 26,
	27, 28, 33, 29, 30, 31, 32, 44, 45, 116,
	0, 0, 22, 0, 0, 0, 166, 41, 37, 129,
	134, 0, 0, 159, 0, 34, 0, 0, 97, 0,
	44, 45, 0, 0, 0, 22, 0, 170, 35, 0,
	0, 37, 172, 0, 0, 0, 0, 14, 41, 0,
	0, 0, 0, 0, 0, 0, 34, 0, 0, 0,
	15, 44, 45, 0, 0, 0, 22, 40, 0, 35,
	0, 0, 37, 0, 24, 25, 46, 36, 0, 0,
	42, 0, 0, 26, 27, 28, 33, 29, 30, 31,
	32, 15, 0, 0, 0, 0, 0, 0, 40, 0,
	0, 0, 0, 0, 0, 24, 25, 46, 36, 0,
	0, 42, 0, 0, 26, 27, 28, 33, 29, 30,
	31, 32, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 14, 41, 0, 0, 0, 0, 0, 0, 0,
	34, 0, 0, 0, 0, 44, 45, 0, 0, 0,
	22, 0, 0, 35, 0, 0, 37, 0, 0, 0,
	0, 0, 14, 41, 0, 0, 0, 0, 0, 0,
	0, 34, 0, 0, 0, 0, 44, 45, 0, 0,
	0, 22, 0, 0, 35, 0, 0, 37, 49, 50,
	51, 52, 54, 55, 56, 57, 58, 59, 60, 61,
	53,
}
var yyPact = [...]int{

	-1000, -1000, 41, -1000, -1000, 342, 64, -1000, -1000, 491,
	-1000, -1000, -1000, -24, -87, 72, -1000, -1000, 59, -1000,
	-65, -76, 342, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 89, 63, -1000, -84, -90, -90,
	-1000, -88, -1000, -80, -79, 29, -1000, -1000, 342, 311,
	342, 342, 342, 342, 342, 342, 342, 342, 342, 342,
	342, 342, 342, 342, -68, 72, -129, 124, -1000, -1000,
	114, 173, -89, 342, -81, -81, -94, -8, -1000, 66,
	-1000, 20, 227, 173, -40, 27, 342, -1000, -1000, -55,
	-1000, 72, 34, 29, -80, -1000, -1000, 173, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-134, -1000, -148, -90, -1000, 342, -1000, -95, 115, 72,
	196, -1000, 59, -1000, 72, 72, -135, -90, -1000, -1000,
	342, -1000, -90, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 169, 13, -149, -62, -151, -1000, -152, -1000, 16,
	-1000, -1000, 342, -113, -1000, -80, -1000, -1000, -1000, -1000,
	-153, 113, 342, -1000, -1000, 173, -91, -137, -125, -1000,
	-1000, -154, -1000, -139, -1000, -1000, 5, -1000, -1000, -1000,
	-1000, 342, -1000, -1000, -1000, 112, -1000, -1000, -1000, 342,
	-1000, -1000, 72, -1000, 113, -1000, -1000, -1000, -141, -1000,
	-1000, -127, -155, 113, -1000,
}
var yyPgo = [...]int{

	0, 245, 244, 243, 138, 225, 0, 125, 36, 3,
	220, 224, 203, 264, 193, 192, 147, 191, 187, 19,
	8, 183, 180, 173, 162, 150, 160, 146, 145, 229,
	140, 136, 137, 135, 1, 118, 124,
}
var yyR1 = [...]int{

	0, 1, 36, 27, 27, 4, 4, 5, 5, 5,
	2, 3, 29, 29, 30, 30, 9, 9, 28, 28,
	11, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 15, 15,
	15, 15, 12, 12, 19, 19, 19, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 20, 20, 20, 6, 6, 7, 7,
	18, 21, 21, 21, 22, 22, 22, 23, 23, 23,
	23, 23, 23, 8, 8, 8, 13, 13, 13, 24,
	24, 16, 16, 16, 17, 17, 17, 34, 35, 35,
	33, 33, 32, 32, 32, 32, 32, 32, 31, 31,
	31, 31, 25, 25, 25, 25, 25, 25, 25, 26,
	26, 26, 26,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 0, 1, 3, 1, 3, 2,
	1, 3, 2, 3, 1, 3, 1, 2, 3, 1,
	1, 6, 5, 3, 4, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 1, 2, 4,
	4, 2, 1, 1, 4, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 2, 3,
	3, 1, 1, 1, 3, 3, 1, 1, 0, 1,
	1, 1, 3, 1, 1, 3, 1, 1, 4, 4,
	4, 4, 1, 1, 1, 3, 1, 4, 2, 3,
	3, 1, 3, 1, 1, 3, 1, 1, 0, 1,
	3, 1, 3, 1, 4, 2, 6, 4, 2, 2,
	1, 2, 1, 4, 3, 3, 3, 6, 3, 1,
	1, 2, 1,
}
var yyChk = [...]int{

	-1000, -1, -27, -2, -3, 82, -28, -11, -6, -8,
	-10, -23, -24, -21, 140, 69, -14, -13, -20, -15,
	-12, -18, 159, -19, 83, 84, 92, 93, 94, 96,
	97, 98, 99, 95, 149, 162, 86, 165, -5, -22,
	76, 141, 89, -4, 154, 155, 85, 158, 9, 17,
	18, 19, 20, 29, 21, 22, 23, 24, 25, 26,
	27, 28, 69, 163, 139, 159, -34, -33, -35, -32,
	-6, 37, 140, 69, 153, 153, -6, 88, 150, -31,
	-25, 86, 151, 152, -31, 88, 163, -13, -29, 159,
	-29, 159, 155, 155, -4, -11, -6, 37, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-7, -6, -6, -17, 85, 163, -13, -34, 161, 9,
	15, -8, -20, -19, 69, 159, -7, -16, -36, -13,
	163, 85, -16, -36, -13, 160, 150, -25, 150, -25,
	88, 69, 139, -6, 90, -8, 162, -6, 160, -30,
	-9, -6, 156, -34, 85, -4, -8, 161, 164, -29,
	-6, 160, 17, -35, -6, 37, 140, -34, -34, 161,
	-29, -6, -29, -26, 85, 91, 50, 86, 85, 164,
	164, 69, 164, 164, 160, 9, -6, 160, 164, 17,
	-6, -8, 159, 161, 160, 164, 161, 91, -6, -9,
	-6, -34, 161, 160, 164,
}
var yyDef = [...]int{

	4, -2, 1, 3, 10, 0, 0, 19, 20, -2,
	67, -2, 84, 70, 0, 98, 37, 77, 62, 82,
	0, 0, 0, -2, 47, 48, 49, 50, 51, 52,
	53, 54, 55, 56, 0, 0, 86, 0, -2, 0,
	42, 0, 46, 7, 0, 0, 5, 11, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 68, 0, 0, 98, 0, 97, 101, 99,
	103, 0, 0, 68, 0, 0, 0, 0, 58, 0,
	110, 112, 0, 0, 0, 0, 0, 88, 38, 0,
	41, 98, 0, 0, 9, 18, 23, 0, 25, 26,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36,
	0, 69, 0, 85, 94, 0, 96, 0, 45, 98,
	0, -2, 0, -2, 98, 98, 0, 0, -2, -2,
	0, 2, 0, -2, -2, -2, 57, 111, 60, 108,
	109, 0, 0, 0, 0, 71, 59, 0, 12, 0,
	14, 16, 0, 0, 6, 8, -2, 78, 80, 81,
	0, 0, 0, 100, 102, 0, 0, 0, 0, 79,
	39, 0, 40, 0, 119, 120, 0, 122, 114, 115,
	116, 0, 118, 87, 13, 0, 17, 44, 95, 0,
	22, -2, 98, 45, 107, 92, 113, 121, 0, 15,
	21, 0, 0, 106, 117,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 162, 3, 165, 54, 37, 3,
	159, 160, 52, 49, 9, 50, 51, 53, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 31, 158,
	43, 17, 45, 30, 67, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 69, 3, 161, 36, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 163, 35, 164, 57,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 10, 11, 12,
	13, 14, 15, 16, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 32, 33, 34, 38,
	39, 40, 41, 42, 44, 46, 47, 48, 56, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 68, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
	121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
	151, 152, 153, 154, 155, 156, 157,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:239
		{
			yyVAL.program = &ast.Program{Statements: yyDollar[1].stmts}
			if l, ok := yylex.(*LexerWrapper); ok {
				l.program = yyVAL.program
			}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:265
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:275
		{
			yyVAL.stmts = []ast.Statement{}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:280
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:281
		{
			yyVAL.expr = ast.NewNamespaceExpression(nil, nil, yyDollar[1].expr, ast.NewStringLiteral(yyDollar[3].tok, yyDollar[3].tok.Literal))
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:285
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.expr = ast.NewNamespaceExpression(yyDollar[1].tok, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:287
		{
			yyVAL.expr = ast.NewNamespaceExpression(nil, yyDollar[1].tok, yyDollar[2].expr)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:291
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:421
		{
			yyVAL.stmt = ast.NewEchoStatement(yyDollar[1].tok, yyDollar[2].exprs)
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:656
		{
			yyVAL.expr = ast.NewArgumentListExpression()
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:657
		{
			yyVAL.expr = ast.NewArgumentListExpression(yyDollar[2].exprs...)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:662
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:664
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:668
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:669
		{
			yyVAL.expr = ast.NewArgumentExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:824
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:825
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:829
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 21:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:863
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...), yyDollar[6].expr, false)
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:865
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...), yyDollar[5].expr, false)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:867
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:869
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, yyDollar[1].expr, yyDollar[4].expr, true)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:872
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PlusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:874
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MinusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:876
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MulEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:878
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PowEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:880
		{
			yyVAL.expr = ast.NewAssignExpression(ast.DivEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:882
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ConcatEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:884
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ModEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:886
		{
			yyVAL.expr = ast.NewAssignExpression(ast.AndEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:888
		{
			yyVAL.expr = ast.NewAssignExpression(ast.QrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:890
		{
			yyVAL.expr = ast.NewAssignExpression(ast.XorEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:892
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SlEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:894
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:963
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1017
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1019
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1021
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1023
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1027
		{
			yyVAL.expr = ast.NewStaticLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1028
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1057
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Long, yyDollar[3].exprs...)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1058
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1059
		{
			yyVAL.expr = ast.NewConstantEncapsedStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1063
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1064
		{
			yyVAL.expr = ast.NewDoubleLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1065
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1066
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1067
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1068
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1069
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1070
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1071
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1072
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1073
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1074
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[2].tok)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1075
		{
			yyVAL.expr = ast.NewEncapsListExpression(yyDollar[2].exprs...)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1076
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, yyDollar[2].exprs...)
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1077
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1078
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1082
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1084
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1086
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1090
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1091
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1095
		{
			yyVAL.expr = nil
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1096
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1100
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1104
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1105
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1106
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1110
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1111
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1112
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1117
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1119
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1121
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1123
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Curly, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1125
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Prop, yyDollar[1].expr, []ast.Expression{yyDollar[3].expr, yyDollar[4].expr}...)
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1126
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1131
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1133
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1135
		{
			yyVAL.expr = ast.NewVariableExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1139
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1140
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.CurlyOpen, yyDollar[3].expr)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1141
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.Var, yyDollar[2].expr)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1146
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1148
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1168
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1169
		{
			yyVAL.expr = ast.NewMemberNameExpression(yyDollar[2].expr)
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1170
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1174
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1175
		{
			yyVAL.expr = ast.NewPropertyNameExpression(yyDollar[2].expr)
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1176
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1181
		{ /* allow single trailing comma */
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1185
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1186
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1191
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].exprs...)
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1193
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1198
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[3].expr, false)}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1200
		{
			yyVAL.exprs = []ast.Expression{yyDollar[1].expr}
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1202
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[4].expr, true)}
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1204
		{
			yyVAL.exprs = []ast.Expression{ast.NewAmpersandLiteral(yyDollar[2].expr)}
		}
	case 106:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1206
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, ast.NewListExpression(yyDollar[3].tok, yyDollar[5].exprs...), false)}
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1208
		{
			yyVAL.exprs = []ast.Expression{ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...)}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1213
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1215
		{
			yyVAL.exprs = append(yyDollar[1].exprs, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1217
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1219
		{
			yyVAL.exprs = append(yyVAL.exprs, []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal), yyDollar[2].expr}...)
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1224
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1226
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Dim, yyDollar[3].expr)
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1228
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Prop, ast.NewStringLiteral(yyDollar[3].tok, yyDollar[3].tok.Literal))
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1230
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, yyDollar[2].expr)
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1232
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 117:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1234
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DimInDollarOpenCurlyBraces, []ast.Expression{ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal), yyDollar[4].expr}...)
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1235
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.CurlyOpen, yyDollar[2].expr)
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1239
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1240
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1241
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[2].tok, "-"+yyDollar[2].tok.Literal)
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1242
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	}
	goto yystack /* stack new state and value */
}
