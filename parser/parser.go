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
	69, 121,
	139, 121,
	153, 121,
	163, 121,
	-2, 116,
	-1, 11,
	159, 124,
	-2, 133,
	-1, 39,
	69, 123,
	139, 123,
	153, 123,
	159, 126,
	163, 123,
	-2, 111,
	-1, 54,
	153, 90,
	-2, 113,
	-1, 120,
	69, 121,
	139, 121,
	153, 121,
	163, 121,
	-2, 39,
	-1, 123,
	159, 126,
	-2, 123,
	-1, 125,
	69, 121,
	139, 121,
	153, 121,
	163, 121,
	-2, 41,
	-1, 180,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 63,
	-1, 181,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 64,
	-1, 182,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 65,
	-1, 183,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 66,
	-1, 184,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 67,
	-1, 185,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 68,
	-1, 186,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 69,
	-1, 187,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 70,
	-1, 188,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 71,
	-1, 218,
	9, 155,
	160, 155,
	161, 155,
	-2, 121,
	-1, 222,
	69, 122,
	139, 122,
	153, 122,
	159, 125,
	163, 122,
	-2, 73,
	-1, 226,
	159, 141,
	-2, 114,
	-1, 227,
	159, 143,
	-2, 139,
	-1, 231,
	159, 141,
	-2, 115,
	-1, 232,
	159, 143,
	-2, 140,
	-1, 253,
	69, 121,
	139, 121,
	153, 121,
	163, 121,
	-2, 24,
	-1, 265,
	159, 125,
	-2, 122,
	-1, 290,
	9, 154,
	160, 154,
	161, 154,
	-2, 121,
}

const yyPrivate = 57344

const yyLast = 1662

var yyAct = [...]int{

	116, 112, 247, 114, 302, 155, 8, 281, 294, 268,
	284, 266, 280, 254, 215, 301, 35, 119, 292, 39,
	126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 286, 140, 121, 121, 267, 123, 123, 52,
	9, 229, 52, 258, 108, 291, 148, 156, 237, 211,
	52, 219, 158, 139, 111, 159, 64, 160, 120, 125,
	143, 157, 142, 37, 54, 8, 163, 164, 165, 166,
	167, 168, 169, 170, 171, 172, 173, 174, 175, 176,
	177, 178, 179, 180, 181, 182, 183, 184, 185, 186,
	187, 188, 147, 192, 193, 195, 196, 197, 198, 199,
	200, 201, 202, 203, 204, 205, 206, 279, 124, 208,
	209, 149, 150, 214, 110, 56, 153, 59, 53, 228,
	148, 53, 243, 220, 62, 52, 221, 212, 58, 53,
	56, 148, 226, 237, 121, 148, 225, 123, 109, 62,
	208, 148, 208, 152, 207, 277, 275, 238, 34, 295,
	240, 62, 251, 5, 244, 190, 191, 248, 218, 141,
	250, 283, 148, 77, 144, 91, 288, 121, 7, 259,
	123, 74, 75, 73, 76, 78, 79, 77, 91, 161,
	57, 273, 276, 216, 233, 149, 150, 274, 146, 113,
	77, 242, 115, 60, 61, 235, 149, 150, 122, 246,
	149, 150, 154, 6, 53, 63, 149, 150, 60, 61,
	2, 121, 272, 257, 123, 12, 256, 239, 261, 11,
	260, 264, 55, 13, 38, 210, 145, 149, 150, 270,
	36, 269, 33, 162, 138, 253, 271, 234, 189, 236,
	151, 10, 4, 3, 236, 234, 1, 0, 0, 0,
	285, 0, 0, 0, 76, 78, 79, 0, 91, 213,
	289, 0, 80, 81, 74, 75, 73, 76, 78, 79,
	77, 91, 0, 0, 0, 0, 231, 0, 252, 121,
	230, 296, 123, 77, 223, 248, 224, 297, 0, 298,
	0, 227, 232, 299, 0, 0, 19, 20, 0, 0,
	0, 0, 21, 290, 22, 17, 18, 24, 25, 26,
	27, 28, 29, 30, 32, 0, 15, 0, 16, 0,
	0, 0, 0, 56, 0, 0, 67, 69, 68, 0,
	40, 41, 62, 52, 0, 0, 58, 0, 0, 42,
	43, 44, 49, 45, 46, 47, 48, 31, 92, 65,
	66, 70, 72, 71, 84, 85, 82, 83, 90, 86,
	87, 88, 89, 80, 81, 74, 75, 73, 76, 78,
	79, 0, 91, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 77, 0, 0, 14, 57, 0,
	0, 0, 0, 0, 0, 0, 50, 0, 117, 0,
	0, 60, 61, 249, 0, 0, 23, 245, 0, 51,
	19, 20, 53, 0, 0, 0, 21, 0, 22, 17,
	18, 24, 25, 26, 27, 28, 29, 30, 32, 0,
	15, 0, 16, 0, 0, 0, 0, 56, 0, 0,
	0, 0, 0, 0, 40, 41, 62, 52, 0, 0,
	58, 0, 0, 42, 43, 44, 49, 45, 46, 47,
	48, 31, 71, 84, 85, 82, 83, 90, 86, 87,
	88, 89, 80, 81, 74, 75, 73, 76, 78, 79,
	293, 91, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 0, 0, 0, 0, 0, 0,
	0, 118, 57, 0, 0, 0, 0, 0, 0, 0,
	50, 0, 0, 0, 0, 60, 61, 0, 0, 0,
	23, 19, 20, 51, 0, 0, 53, 21, 0, 22,
	17, 18, 24, 25, 26, 27, 28, 29, 30, 32,
	0, 15, 0, 16, 0, 0, 0, 0, 56, 0,
	0, 67, 69, 68, 0, 40, 41, 62, 52, 0,
	0, 58, 0, 0, 42, 43, 44, 49, 45, 46,
	47, 48, 31, 92, 65, 66, 70, 72, 71, 84,
	85, 82, 83, 90, 86, 87, 88, 89, 80, 81,
	74, 75, 73, 76, 78, 79, 0, 91, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 77,
	0, 0, 14, 57, 0, 0, 0, 0, 0, 0,
	262, 50, 0, 0, 0, 0, 60, 61, 249, 0,
	0, 23, 19, 20, 51, 0, 0, 53, 21, 0,
	22, 17, 18, 24, 25, 26, 27, 28, 29, 30,
	32, 0, 15, 0, 16, 0, 0, 0, 0, 56,
	0, 0, 0, 0, 0, 0, 40, 41, 62, 52,
	0, 0, 58, 0, 0, 42, 43, 44, 49, 45,
	46, 47, 48, 31, 66, 70, 72, 71, 84, 85,
	82, 83, 90, 86, 87, 88, 89, 80, 81, 74,
	75, 73, 76, 78, 79, 287, 91, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 77, 0,
	0, 0, 0, 263, 57, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 0, 0, 0, 60, 61, 0,
	0, 0, 23, 19, 20, 51, 0, 0, 53, 21,
	0, 22, 17, 18, 24, 25, 26, 27, 28, 29,
	30, 32, 0, 15, 0, 16, 0, 0, 0, 0,
	56, 0, 0, 67, 69, 68, 0, 40, 41, 62,
	52, 0, 0, 58, 241, 0, 42, 43, 44, 49,
	45, 46, 47, 48, 31, 92, 65, 66, 70, 72,
	71, 84, 85, 82, 83, 90, 86, 87, 88, 89,
	80, 81, 74, 75, 73, 76, 78, 79, 0, 91,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 77, 0, 0, 14, 57, 0, 0, 0, 0,
	0, 0, 194, 50, 0, 0, 0, 0, 60, 61,
	0, 0, 0, 23, 19, 20, 51, 0, 0, 53,
	21, 0, 22, 17, 18, 24, 25, 26, 27, 28,
	29, 30, 32, 0, 15, 0, 16, 0, 0, 0,
	0, 56, 0, 0, 0, 0, 0, 0, 40, 41,
	62, 52, 0, 0, 58, 0, 0, 42, 43, 44,
	49, 45, 46, 47, 48, 31, 0, 70, 72, 71,
	84, 85, 82, 83, 90, 86, 87, 88, 89, 80,
	81, 74, 75, 73, 76, 78, 79, 282, 91, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	77, 0, 0, 0, 0, 14, 57, 0, 0, 0,
	0, 0, 0, 0, 50, 0, 0, 0, 0, 60,
	61, 0, 0, 0, 23, 19, 20, 51, 0, 0,
	53, 21, 0, 22, 17, 18, 24, 25, 26, 27,
	28, 29, 30, 32, 0, 15, 0, 16, 0, 0,
	0, 0, 56, 0, 0, 67, 69, 68, 0, 40,
	41, 62, 52, 0, 0, 58, 0, 0, 42, 43,
	44, 49, 45, 46, 47, 48, 31, 92, 65, 66,
	70, 72, 71, 84, 85, 82, 83, 90, 86, 87,
	88, 89, 80, 81, 74, 75, 73, 76, 78, 79,
	0, 91, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 0, 0, 14, 57, 0, 0,
	67, 69, 68, 0, 0, 50, 0, 0, 0, 0,
	60, 61, 0, 0, 0, 23, 0, 0, 51, 0,
	0, 53, 92, 65, 66, 70, 72, 71, 84, 85,
	82, 83, 90, 86, 87, 88, 89, 80, 81, 74,
	75, 73, 76, 78, 79, 0, 91, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 67, 69,
	68, 0, 0, 0, 93, 94, 95, 96, 98, 99,
	100, 101, 102, 103, 104, 105, 97, 0, 0, 278,
	92, 65, 66, 70, 72, 71, 84, 85, 82, 83,
	90, 86, 87, 88, 89, 80, 81, 74, 75, 73,
	76, 78, 79, 0, 91, 106, 107, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 0, 0, 0,
	0, 0, 0, 67, 69, 68, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 255, 92, 65, 66, 70, 72,
	71, 84, 85, 82, 83, 90, 86, 87, 88, 89,
	80, 81, 74, 75, 73, 76, 78, 79, 0, 91,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 77, 0, 0, 0, 0, 0, 0, 67, 69,
	68, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 300,
	92, 65, 66, 70, 72, 71, 84, 85, 82, 83,
	90, 86, 87, 88, 89, 80, 81, 74, 75, 73,
	76, 78, 79, 0, 91, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 0, 0, 0,
	0, 0, 67, 69, 68, 0, 0, 217, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 265, 92, 65, 66, 70, 72, 71,
	84, 85, 82, 83, 90, 86, 87, 88, 89, 80,
	81, 74, 75, 73, 76, 78, 79, 0, 91, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	77, 0, 0, 0, 0, 0, 0, 67, 69, 68,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 222, 92,
	65, 66, 70, 72, 71, 84, 85, 82, 83, 90,
	86, 87, 88, 89, 80, 81, 74, 75, 73, 76,
	78, 79, 0, 91, 69, 68, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 92, 65, 66, 70, 72,
	71, 84, 85, 82, 83, 90, 86, 87, 88, 89,
	80, 81, 74, 75, 73, 76, 78, 79, 68, 91,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 77, 0, 0, 0, 0, 0, 0, 92, 65,
	66, 70, 72, 71, 84, 85, 82, 83, 90, 86,
	87, 88, 89, 80, 81, 74, 75, 73, 76, 78,
	79, 0, 91, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 77, 92, 65, 66, 70, 72,
	71, 84, 85, 82, 83, 90, 86, 87, 88, 89,
	80, 81, 74, 75, 73, 76, 78, 79, 0, 91,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 77, 72, 71, 84, 85, 82, 83, 90, 86,
	87, 88, 89, 80, 81, 74, 75, 73, 76, 78,
	79, 0, 91, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 77, 84, 85, 82, 83, 90,
	86, 87, 88, 89, 80, 81, 74, 75, 73, 76,
	78, 79, 0, 91, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 77, 86, 87, 88, 89,
	80, 81, 74, 75, 73, 76, 78, 79, 0, 91,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 77,
}
var yyPact = [...]int{

	-1000, -1000, 71, -1000, -1000, 916, 47, -1000, 1377, 1117,
	-1000, -1000, -1000, -25, -105, 361, 916, 39, 39, 916,
	916, 916, 916, 916, 916, 916, 916, 916, 916, 916,
	916, -106, 916, -1000, -1000, 90, -1000, -91, -93, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	76, 55, -1000, -47, -112, -112, -1000, -107, -1000, -100,
	-98, 66, -1000, -1000, 916, 916, 916, 916, 916, 916,
	916, 916, 916, 916, 916, 916, 916, 916, 916, 916,
	916, 916, 916, 916, 916, 916, 916, 916, 916, 916,
	916, 54, 916, 805, 916, 916, 916, 916, 916, 916,
	916, 916, 916, 916, 916, 916, -1000, -1000, 916, 916,
	-36, 361, -147, 174, -1000, -1000, 1312, 39, -108, -1000,
	-1000, 90, 916, -1000, 361, -1000, 95, 95, 109, 95,
	1248, 95, 95, 95, 95, 95, 95, 95, -1000, 916,
	95, 916, -44, -44, 34, -1000, 45, -1000, 78, 694,
	39, -40, 49, 916, -1000, -1000, 247, -1000, 361, 67,
	66, -100, -1000, 650, 872, 1423, 1503, 1466, 1536, 1567,
	425, 202, 202, 202, 109, 95, 109, 109, 122, 122,
	1593, 1593, 1593, 1593, 215, 215, 215, 215, 1593, -1000,
	-1000, -1000, 1503, 1503, 39, 1503, 1503, 1503, 1503, 1503,
	1503, 1503, 1503, 1503, 1503, 1503, 1503, -148, 1377, 1050,
	-112, -1000, 916, -1000, -117, 152, 361, 583, -1000, 361,
	1183, -150, -1000, -124, -152, -112, -1000, -1000, 916, -1000,
	-112, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 96, 60,
	985, -57, -157, -1000, 763, -1000, 1, -1000, 1377, 916,
	-128, -1000, -100, -1000, -1000, -1000, -1000, 541, 149, 916,
	-1000, 1377, 39, -114, -142, -1000, -1000, -1000, -1000, -1000,
	316, -1000, -153, -1000, -1000, 58, -1000, -1000, -1000, -1000,
	916, -1000, -1000, -1000, 472, 1377, -1000, -1000, 916, 1503,
	-1000, 361, 149, -1000, -1000, -1000, 1118, -1000, 1503, -145,
	-160, 149, -1000,
}
var yyPgo = [...]int{

	0, 246, 243, 242, 117, 64, 0, 144, 40, 2,
	241, 168, 63, 238, 148, 234, 232, 230, 136, 225,
	224, 19, 16, 223, 222, 219, 215, 92, 212, 210,
	203, 5, 199, 188, 192, 189, 1, 3, 132,
}
var yyR1 = [...]int{

	0, 1, 38, 29, 29, 4, 4, 5, 5, 5,
	2, 3, 31, 31, 32, 32, 9, 9, 30, 30,
	11, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 17, 17, 17, 17, 12,
	12, 13, 15, 15, 21, 21, 21, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 22, 22, 22, 6, 6, 7, 7,
	20, 23, 23, 23, 24, 24, 24, 25, 25, 25,
	25, 25, 25, 8, 8, 8, 14, 14, 14, 26,
	26, 18, 18, 18, 19, 19, 19, 36, 37, 37,
	35, 35, 34, 34, 34, 34, 34, 34, 33, 33,
	33, 33, 27, 27, 27, 27, 27, 27, 27, 28,
	28, 28, 28,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 0, 1, 3, 1, 3, 2,
	1, 3, 2, 3, 1, 3, 1, 2, 3, 1,
	1, 6, 5, 3, 4, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 1, 2, 4, 4, 2, 1,
	1, 1, 0, 3, 4, 3, 1, 1, 1, 1,
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

	-1000, -1, -29, -2, -3, 82, -30, -11, -6, -8,
	-10, -25, -26, -23, 140, 69, 71, 58, 59, 49,
	50, 55, 57, 159, 60, 61, 62, 63, 64, 65,
	66, 100, 67, -16, -14, -22, -17, -12, -20, -21,
	83, 84, 92, 93, 94, 96, 97, 98, 99, 95,
	149, 162, 86, 165, -5, -24, 76, 141, 89, -4,
	154, 155, 85, 158, 9, 33, 34, 10, 12, 11,
	35, 37, 36, 51, 49, 50, 52, 68, 53, 54,
	47, 48, 40, 41, 38, 39, 43, 44, 45, 46,
	42, 56, 32, 17, 18, 19, 20, 29, 21, 22,
	23, 24, 25, 26, 27, 28, 58, 59, 69, 163,
	139, 159, -36, -35, -37, -34, -6, 37, 140, -6,
	-8, -22, 159, -21, 69, -8, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -15, 159,
	-6, 69, 153, 153, 88, 150, -33, -27, 86, 151,
	152, -33, 88, 163, -14, -31, 159, -31, 159, 155,
	155, -4, -11, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -13,
	-12, -5, -6, -6, 37, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -7, -6, -6,
	-19, 85, 163, -14, -36, 161, 9, 15, -8, 159,
	-6, -36, 160, -7, -7, -18, -38, -14, 163, 85,
	-18, -38, -14, 150, -27, 150, -27, 88, 69, 139,
	-6, 90, -8, 162, -6, 160, -32, -9, -6, 156,
	-36, 85, -4, -8, 161, 164, -31, -6, 160, 17,
	-37, -6, 37, 140, -36, 160, 161, 160, 161, -31,
	-6, -31, -28, 85, 91, 50, 86, 85, 164, 164,
	69, 164, 164, 160, 9, -6, 160, 164, 17, -6,
	-8, 159, 160, 164, 161, 91, -6, -9, -6, -36,
	161, 160, 164,
}
var yyDef = [...]int{

	4, -2, 1, 3, 10, 0, 0, 19, 20, -2,
	117, -2, 134, 120, 0, 148, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 92, 0, 84, 127, 112, 132, 0, 0, -2,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	0, 0, 136, 0, -2, 0, 89, 0, 96, 7,
	0, 0, 5, 11, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 38, 40, 118, 0,
	0, 148, 0, 147, 151, 149, 153, 0, 0, 25,
	-2, 0, 0, -2, 148, -2, 59, 60, 61, 62,
	0, 75, 76, 77, 78, 79, 80, 81, 82, 118,
	83, 118, 0, 0, 0, 108, 0, 160, 162, 0,
	0, 0, 0, 0, 138, 85, 0, 88, 148, 0,
	0, 9, 18, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, 72,
	91, 90, 74, 23, 0, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 0, 119, 0,
	135, 144, 0, 146, 0, 95, 148, 0, -2, 148,
	0, 0, -2, 0, 0, 0, -2, -2, 0, 2,
	0, -2, -2, 107, 161, 110, 158, 159, 0, 0,
	0, 0, 121, 109, 0, 12, 0, 14, 16, 0,
	0, 6, 8, -2, 128, 130, 131, 0, 0, 0,
	150, 152, 0, 0, 0, -2, 95, 93, 129, 86,
	0, 87, 0, 169, 170, 0, 172, 164, 165, 166,
	0, 168, 137, 13, 0, 17, 94, 145, 0, 22,
	-2, 148, 157, 142, 163, 171, 0, 15, 21, 0,
	0, 156, 167,
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:870
		{
			yyVAL.expr = ast.NewCloneExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:872
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PlusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:874
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MinusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:876
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MulEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:878
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PowEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:880
		{
			yyVAL.expr = ast.NewAssignExpression(ast.DivEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:882
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ConcatEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:884
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ModEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:886
		{
			yyVAL.expr = ast.NewAssignExpression(ast.AndEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:888
		{
			yyVAL.expr = ast.NewAssignExpression(ast.QrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:890
		{
			yyVAL.expr = ast.NewAssignExpression(ast.XorEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:892
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SlEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:894
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:895
		{
			yyVAL.expr = ast.NewIncrementExpression(ast.PostInc, yyDollar[1].expr)
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:896
		{
			yyVAL.expr = ast.NewIncrementExpression(ast.PreInc, yyDollar[2].expr)
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:897
		{
			yyVAL.expr = ast.NewDecrementExpression(ast.PostDec, yyDollar[1].expr)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:898
		{
			yyVAL.expr = ast.NewDecrementExpression(ast.PreDec, yyDollar[2].expr)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:900
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BooleanOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:902
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BooleanAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:904
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:906
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:908
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalXor, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:909
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:910
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:911
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwXor, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:912
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Concat, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:913
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Add, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:914
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sub, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:915
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Mul, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:916
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Pow, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:917
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Div, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:918
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Mod, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:919
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sl, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:920
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:921
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.UnaryPlus, yyDollar[2].expr)
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:922
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.UnaryMinus, yyDollar[2].expr)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:923
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.BoolNot, yyDollar[2].expr)
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:924
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.BwNot, yyDollar[2].expr)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:926
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsIdentical, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:928
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsNotIdentical, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:930
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:932
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsNotEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:934
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Smaller, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:936
		{
			yyVAL.expr = ast.NewInfixExpression(ast.SmallerOrEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:938
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Greater, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:940
		{
			yyVAL.expr = ast.NewInfixExpression(ast.GreaterOrEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:942
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Spaceship, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:944
		{
			yyVAL.expr = ast.NewInfixExpression(ast.InstanceOf, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:945
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:952
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Coalesce, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:954
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:955
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:956
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:957
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:958
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:959
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:960
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:961
		{
			yyVAL.expr = ast.NewExitExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:962
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.Silence, yyDollar[2].expr)
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:963
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1017
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1019
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1021
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1023
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1027
		{
			yyVAL.expr = ast.NewStaticLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1028
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1032
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 92:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1037
		{
			yyVAL.expr = nil
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1038
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1057
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Long, yyDollar[3].exprs...)
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1058
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...)
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1059
		{
			yyVAL.expr = ast.NewConstantEncapsedStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1063
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1064
		{
			yyVAL.expr = ast.NewDoubleLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1065
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1066
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1067
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1068
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1069
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1070
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1071
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1072
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1073
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1074
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[2].tok)
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1075
		{
			yyVAL.expr = ast.NewEncapsListExpression(yyDollar[2].exprs...)
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1076
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, yyDollar[2].exprs...)
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1077
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1078
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1082
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1084
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1086
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1090
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1091
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 118:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1095
		{
			yyVAL.expr = nil
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1096
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1100
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1104
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1105
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1106
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1110
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1111
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1112
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1117
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 128:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1119
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 129:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1121
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1123
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Curly, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 131:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1125
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Prop, yyDollar[1].expr, []ast.Expression{yyDollar[3].expr, yyDollar[4].expr}...)
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1126
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1131
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1133
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1135
		{
			yyVAL.expr = ast.NewVariableExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1139
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 137:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1140
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.CurlyOpen, yyDollar[3].expr)
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1141
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.Var, yyDollar[2].expr)
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1146
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1148
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1168
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1169
		{
			yyVAL.expr = ast.NewMemberNameExpression(yyDollar[2].expr)
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1170
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1174
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1175
		{
			yyVAL.expr = ast.NewPropertyNameExpression(yyDollar[2].expr)
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1176
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1181
		{ /* allow single trailing comma */
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 148:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1185
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1186
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1191
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].exprs...)
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1193
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1198
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[3].expr, false)}
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1200
		{
			yyVAL.exprs = []ast.Expression{yyDollar[1].expr}
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1202
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[4].expr, true)}
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1204
		{
			yyVAL.exprs = []ast.Expression{ast.NewAmpersandLiteral(yyDollar[2].expr)}
		}
	case 156:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1206
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, ast.NewListExpression(yyDollar[3].tok, yyDollar[5].exprs...), false)}
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1208
		{
			yyVAL.exprs = []ast.Expression{ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...)}
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1213
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1215
		{
			yyVAL.exprs = append(yyDollar[1].exprs, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1217
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1219
		{
			yyVAL.exprs = append(yyVAL.exprs, []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal), yyDollar[2].expr}...)
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1224
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 163:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1226
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Dim, yyDollar[3].expr)
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1228
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Prop, ast.NewStringLiteral(yyDollar[3].tok, yyDollar[3].tok.Literal))
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1230
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, yyDollar[2].expr)
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1232
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 167:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1234
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DimInDollarOpenCurlyBraces, []ast.Expression{ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal), yyDollar[4].expr}...)
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1235
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.CurlyOpen, yyDollar[2].expr)
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1239
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1240
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1241
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[2].tok, "-"+yyDollar[2].tok.Literal)
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1242
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	}
	goto yystack /* stack new state and value */
}
