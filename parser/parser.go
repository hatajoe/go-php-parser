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
	"'{'",
	"'}'",
	"'('",
	"')'",
	"';'",
	"']'",
	"'`'",
	"'\"'",
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
	-1, 15,
	69, 163,
	139, 163,
	153, 163,
	158, 163,
	-2, 158,
	-1, 18,
	160, 166,
	-2, 175,
	-1, 52,
	69, 165,
	139, 165,
	153, 165,
	158, 165,
	160, 168,
	-2, 153,
	-1, 75,
	153, 126,
	-2, 155,
	-1, 154,
	69, 163,
	139, 163,
	153, 163,
	158, 163,
	-2, 65,
	-1, 157,
	160, 168,
	-2, 165,
	-1, 159,
	69, 163,
	139, 163,
	153, 163,
	158, 163,
	-2, 67,
	-1, 241,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 89,
	-1, 242,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 90,
	-1, 243,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 91,
	-1, 244,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 92,
	-1, 245,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 93,
	-1, 246,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 94,
	-1, 247,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 95,
	-1, 248,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 96,
	-1, 249,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 97,
	-1, 285,
	9, 203,
	161, 203,
	163, 203,
	-2, 163,
	-1, 289,
	69, 164,
	139, 164,
	153, 164,
	158, 164,
	160, 167,
	-2, 99,
	-1, 303,
	160, 189,
	-2, 156,
	-1, 304,
	160, 191,
	-2, 181,
	-1, 308,
	160, 189,
	-2, 157,
	-1, 309,
	160, 191,
	-2, 182,
	-1, 344,
	69, 163,
	139, 163,
	153, 163,
	158, 163,
	-2, 50,
	-1, 357,
	160, 167,
	-2, 164,
	-1, 404,
	9, 202,
	161, 202,
	163, 202,
	-2, 163,
}

const yyPrivate = 57344

const yyLast = 3302

var yyAct = [...]int{

	12, 84, 219, 146, 328, 277, 318, 148, 89, 80,
	274, 291, 90, 411, 207, 407, 303, 302, 48, 371,
	358, 346, 216, 150, 153, 4, 282, 160, 161, 162,
	163, 164, 86, 431, 165, 166, 167, 168, 169, 170,
	171, 52, 174, 155, 155, 182, 183, 184, 426, 15,
	420, 417, 336, 257, 223, 432, 386, 195, 196, 429,
	198, 199, 178, 381, 406, 388, 157, 157, 73, 359,
	350, 313, 405, 77, 154, 159, 73, 278, 73, 208,
	306, 73, 83, 73, 335, 368, 217, 286, 221, 256,
	254, 209, 213, 224, 225, 226, 227, 228, 229, 230,
	231, 232, 233, 234, 235, 236, 237, 238, 239, 240,
	241, 242, 243, 244, 245, 246, 247, 248, 249, 210,
	251, 253, 197, 179, 47, 294, 194, 259, 261, 262,
	263, 264, 265, 266, 267, 268, 269, 270, 271, 272,
	205, 315, 273, 275, 276, 255, 150, 193, 74, 281,
	279, 81, 82, 305, 173, 316, 74, 287, 74, 150,
	314, 74, 288, 74, 145, 142, 141, 87, 85, 430,
	155, 369, 211, 212, 275, 367, 179, 312, 192, 187,
	179, 297, 294, 179, 290, 200, 275, 186, 180, 181,
	188, 295, 218, 157, 319, 320, 301, 50, 321, 206,
	155, 285, 324, 311, 308, 307, 325, 222, 385, 329,
	179, 150, 408, 365, 331, 380, 179, 83, 294, 332,
	121, 122, 333, 157, 90, 123, 125, 124, 179, 75,
	204, 299, 338, 185, 179, 144, 176, 363, 118, 292,
	293, 180, 181, 192, 323, 180, 181, 201, 180, 181,
	104, 189, 104, 340, 143, 341, 419, 342, 258, 402,
	343, 296, 177, 292, 351, 293, 293, 292, 337, 280,
	283, 147, 361, 364, 322, 180, 181, 149, 362, 155,
	349, 180, 181, 191, 150, 353, 310, 150, 220, 327,
	356, 352, 348, 180, 181, 13, 6, 88, 175, 180,
	181, 370, 157, 103, 105, 106, 373, 118, 317, 250,
	344, 304, 309, 14, 275, 377, 189, 372, 7, 104,
	389, 378, 374, 433, 376, 101, 102, 100, 103, 105,
	106, 387, 118, 2, 202, 203, 392, 221, 394, 393,
	395, 360, 190, 19, 104, 18, 76, 20, 191, 51,
	49, 42, 403, 172, 32, 31, 16, 390, 214, 158,
	3, 1, 0, 0, 0, 0, 77, 0, 399, 409,
	0, 0, 0, 155, 0, 83, 73, 0, 0, 79,
	0, 0, 319, 0, 0, 0, 0, 329, 413, 0,
	0, 414, 0, 415, 0, 0, 157, 0, 0, 0,
	0, 0, 421, 422, 404, 0, 150, 0, 0, 423,
	0, 0, 0, 56, 57, 58, 59, 60, 221, 418,
	427, 428, 44, 45, 0, 46, 0, 0, 0, 0,
	0, 78, 0, 0, 0, 0, 0, 375, 0, 0,
	280, 379, 0, 0, 81, 82, 0, 0, 0, 0,
	156, 0, 0, 0, 0, 434, 74, 0, 26, 27,
	0, 0, 0, 0, 28, 0, 29, 24, 25, 33,
	34, 35, 36, 37, 38, 39, 41, 0, 22, 53,
	23, 0, 0, 0, 0, 77, 0, 0, 0, 0,
	0, 11, 61, 62, 83, 73, 0, 0, 79, 0,
	0, 63, 64, 65, 70, 66, 67, 68, 69, 40,
	17, 9, 8, 425, 10, 107, 108, 101, 102, 100,
	103, 105, 106, 0, 118, 56, 57, 58, 59, 60,
	0, 0, 0, 0, 44, 45, 104, 46, 0, 0,
	54, 55, 0, 0, 0, 0, 0, 0, 0, 21,
	78, 0, 400, 0, 0, 0, 0, 0, 71, 0,
	0, 0, 0, 81, 82, 0, 0, 5, 0, 30,
	26, 27, 0, 43, 72, 74, 28, 0, 29, 24,
	25, 33, 34, 35, 36, 37, 38, 39, 41, 0,
	22, 53, 23, 0, 0, 0, 0, 77, 0, 0,
	0, 0, 0, 11, 61, 62, 83, 73, 0, 0,
	79, 0, 0, 63, 64, 65, 70, 66, 67, 68,
	69, 40, 17, 9, 8, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 56, 57, 58,
	59, 60, 0, 0, 0, 0, 44, 45, 0, 46,
	0, 0, 54, 55, 0, 0, 0, 0, 0, 0,
	0, 21, 78, 0, 0, 0, 0, 0, 0, 0,
	71, 0, 0, 0, 0, 81, 82, 0, 0, 5,
	0, 30, 26, 27, 0, 43, 72, 74, 28, 0,
	29, 24, 25, 33, 34, 35, 36, 37, 38, 39,
	41, 0, 22, 53, 23, 0, 0, 0, 398, 77,
	0, 0, 0, 0, 0, 11, 61, 62, 83, 73,
	0, 0, 79, 0, 0, 63, 64, 65, 70, 66,
	67, 68, 69, 40, 17, 9, 8, 0, 10, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 56,
	57, 58, 59, 60, 0, 0, 0, 0, 44, 45,
	0, 46, 0, 0, 54, 55, 0, 0, 0, 0,
	0, 0, 0, 21, 78, 0, 391, 0, 0, 0,
	0, 0, 71, 0, 0, 0, 0, 81, 82, 0,
	0, 5, 0, 30, 26, 27, 0, 43, 72, 74,
	28, 0, 29, 24, 25, 33, 34, 35, 36, 37,
	38, 39, 41, 0, 22, 53, 23, 0, 0, 0,
	0, 77, 0, 0, 0, 0, 0, 11, 61, 62,
	83, 73, 0, 0, 79, 0, 0, 63, 64, 65,
	70, 66, 67, 68, 69, 40, 17, 9, 8, 0,
	10, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 56, 57, 58, 59, 60, 0, 0, 0, 0,
	44, 45, 0, 46, 0, 0, 54, 55, 0, 0,
	0, 0, 0, 0, 0, 21, 78, 0, 0, 0,
	0, 0, 0, 0, 71, 0, 0, 0, 0, 81,
	82, 0, 0, 5, 0, 30, 26, 27, 0, 43,
	72, 74, 28, 0, 29, 24, 25, 33, 34, 35,
	36, 37, 38, 39, 41, 0, 22, 53, 23, 0,
	0, 0, 0, 77, 0, 0, 0, 0, 0, 11,
	61, 62, 83, 73, 0, 0, 79, 0, 0, 63,
	64, 65, 70, 66, 67, 68, 69, 40, 17, 9,
	8, 0, 10, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 56, 57, 58, 59, 60, 0, 0,
	0, 0, 44, 45, 0, 46, 0, 0, 54, 55,
	0, 0, 0, 0, 0, 0, 0, 21, 78, 0,
	0, 0, 0, 0, 0, 0, 71, 0, 0, 0,
	0, 81, 82, 0, 0, 5, 215, 30, 26, 27,
	0, 43, 72, 74, 28, 0, 29, 24, 25, 33,
	34, 35, 36, 37, 38, 39, 41, 0, 22, 53,
	23, 0, 0, 0, 0, 77, 0, 0, 0, 0,
	0, 11, 61, 62, 83, 73, 0, 0, 79, 0,
	0, 63, 64, 65, 70, 66, 67, 68, 69, 40,
	17, 9, 8, 0, 10, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 56, 57, 58, 59, 60,
	0, 0, 0, 0, 44, 45, 0, 46, 0, 0,
	54, 55, 0, 0, 0, 0, 0, 0, 0, 21,
	78, 0, 0, 0, 0, 0, 0, 0, 71, 0,
	0, 0, 0, 81, 82, 0, 0, 5, 0, 30,
	26, 27, 0, 43, 72, 74, 28, 0, 29, 24,
	25, 33, 34, 35, 36, 37, 38, 39, 41, 0,
	22, 53, 23, 0, 0, 0, 0, 77, 0, 0,
	0, 0, 0, 0, 61, 62, 83, 73, 0, 0,
	79, 0, 0, 63, 64, 65, 70, 66, 67, 68,
	69, 40, 0, 0, 113, 114, 115, 116, 107, 108,
	101, 102, 100, 103, 105, 106, 0, 118, 56, 57,
	58, 59, 60, 0, 0, 0, 0, 44, 45, 104,
	46, 0, 54, 55, 0, 0, 0, 0, 0, 0,
	0, 21, 78, 0, 0, 0, 0, 0, 0, 0,
	71, 151, 0, 0, 0, 81, 82, 330, 0, 0,
	0, 30, 326, 26, 27, 43, 72, 74, 0, 28,
	0, 29, 24, 25, 33, 34, 35, 36, 37, 38,
	39, 41, 0, 22, 53, 23, 0, 0, 0, 0,
	77, 0, 0, 0, 0, 0, 0, 61, 62, 83,
	73, 0, 0, 79, 0, 0, 63, 64, 65, 70,
	66, 67, 68, 69, 40, 126, 127, 128, 129, 131,
	132, 133, 134, 135, 136, 137, 138, 130, 0, 0,
	56, 57, 58, 59, 60, 0, 0, 0, 0, 44,
	45, 0, 46, 0, 0, 54, 55, 0, 0, 0,
	0, 0, 0, 0, 152, 78, 139, 140, 0, 0,
	0, 0, 0, 71, 0, 0, 0, 0, 81, 82,
	0, 0, 0, 0, 30, 26, 27, 0, 43, 72,
	74, 28, 0, 29, 24, 25, 33, 34, 35, 36,
	37, 38, 39, 41, 0, 22, 53, 23, 0, 0,
	0, 0, 77, 0, 0, 0, 0, 0, 0, 61,
	62, 83, 73, 0, 0, 79, 0, 0, 63, 64,
	65, 70, 66, 67, 68, 69, 40, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 56, 57, 58, 59, 60, 0, 0,
	0, 0, 44, 45, 0, 46, 0, 54, 55, 0,
	0, 0, 0, 0, 0, 0, 21, 78, 0, 0,
	0, 0, 0, 0, 0, 71, 354, 0, 0, 0,
	81, 82, 330, 0, 0, 0, 30, 0, 26, 27,
	43, 72, 74, 0, 28, 0, 29, 24, 25, 33,
	34, 35, 36, 37, 38, 39, 41, 0, 22, 53,
	23, 0, 0, 0, 0, 77, 0, 0, 0, 0,
	0, 0, 61, 62, 83, 73, 0, 0, 79, 0,
	0, 63, 64, 65, 70, 66, 67, 68, 69, 40,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 56, 57, 58, 59, 60,
	0, 0, 0, 0, 44, 45, 0, 46, 0, 0,
	54, 55, 0, 0, 0, 0, 0, 0, 0, 355,
	78, 0, 0, 0, 0, 0, 0, 0, 71, 0,
	0, 0, 0, 81, 82, 0, 0, 0, 0, 30,
	26, 27, 0, 43, 72, 74, 28, 0, 29, 24,
	25, 33, 34, 35, 36, 37, 38, 39, 41, 0,
	22, 53, 23, 0, 0, 0, 0, 77, 0, 0,
	0, 0, 0, 0, 61, 62, 83, 73, 0, 0,
	79, 298, 0, 63, 64, 65, 70, 66, 67, 68,
	69, 40, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 56, 57,
	58, 59, 60, 0, 0, 0, 0, 44, 45, 0,
	46, 0, 54, 55, 0, 0, 0, 0, 0, 0,
	0, 21, 78, 0, 0, 0, 0, 0, 0, 0,
	71, 260, 0, 0, 0, 81, 82, 0, 0, 0,
	0, 30, 0, 26, 27, 43, 72, 74, 0, 28,
	0, 29, 24, 25, 33, 34, 35, 36, 37, 38,
	39, 41, 0, 22, 53, 23, 0, 0, 0, 0,
	77, 0, 0, 0, 0, 0, 0, 61, 62, 83,
	73, 0, 0, 79, 0, 0, 63, 64, 65, 70,
	66, 67, 68, 69, 40, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	56, 57, 58, 59, 60, 0, 0, 0, 0, 44,
	45, 0, 46, 0, 0, 54, 55, 0, 0, 0,
	0, 0, 0, 0, 21, 78, 0, 252, 0, 0,
	0, 0, 0, 71, 0, 0, 0, 0, 81, 82,
	0, 0, 0, 0, 30, 26, 27, 0, 43, 72,
	74, 28, 0, 29, 24, 25, 33, 34, 35, 36,
	37, 38, 39, 41, 0, 22, 53, 23, 0, 0,
	0, 0, 77, 0, 0, 0, 0, 0, 0, 61,
	62, 83, 73, 0, 0, 79, 0, 0, 63, 64,
	65, 70, 66, 67, 68, 69, 40, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 56, 57, 58, 59, 60, 0, 0, 0,
	0, 44, 45, 0, 46, 0, 0, 54, 55, 0,
	0, 0, 0, 0, 0, 0, 21, 78, 0, 0,
	0, 0, 0, 0, 0, 71, 0, 0, 0, 0,
	81, 82, 0, 0, 0, 0, 30, 26, 27, 0,
	43, 72, 74, 28, 0, 29, 24, 25, 33, 34,
	35, 36, 37, 38, 39, 41, 0, 22, 53, 23,
	0, 0, 0, 0, 77, 0, 0, 0, 0, 0,
	0, 61, 62, 83, 73, 0, 0, 79, 0, 0,
	63, 64, 65, 70, 66, 67, 68, 69, 40, 119,
	0, 120, 92, 93, 97, 99, 98, 111, 112, 109,
	110, 117, 113, 114, 115, 116, 107, 108, 101, 102,
	100, 103, 105, 106, 0, 118, 0, 0, 0, 54,
	55, 0, 0, 0, 0, 0, 0, 104, 21, 78,
	0, 0, 0, 0, 0, 0, 0, 71, 94, 96,
	95, 0, 81, 82, 0, 0, 0, 0, 30, 0,
	0, 0, 43, 72, 74, 0, 0, 0, 119, 0,
	120, 92, 93, 97, 99, 98, 111, 112, 109, 110,
	117, 113, 114, 115, 116, 107, 108, 101, 102, 100,
	103, 105, 106, 0, 118, 94, 96, 95, 0, 0,
	0, 0, 0, 0, 0, 0, 104, 0, 0, 0,
	0, 0, 0, 0, 0, 119, 0, 120, 92, 93,
	97, 99, 98, 111, 112, 109, 110, 117, 113, 114,
	115, 116, 107, 108, 101, 102, 100, 103, 105, 106,
	0, 118, 94, 96, 95, 0, 0, 0, 0, 0,
	0, 0, 0, 104, 0, 0, 0, 0, 0, 0,
	0, 0, 119, 0, 120, 92, 93, 97, 99, 98,
	111, 112, 109, 110, 117, 113, 114, 115, 116, 107,
	108, 101, 102, 100, 103, 105, 106, 0, 118, 0,
	94, 96, 95, 0, 0, 0, 0, 0, 0, 0,
	104, 424, 0, 0, 0, 0, 0, 0, 0, 0,
	119, 0, 120, 92, 93, 97, 99, 98, 111, 112,
	109, 110, 117, 113, 114, 115, 116, 107, 108, 101,
	102, 100, 103, 105, 106, 0, 118, 94, 96, 95,
	0, 0, 0, 0, 0, 0, 0, 91, 104, 0,
	0, 0, 0, 0, 0, 0, 0, 119, 0, 120,
	92, 93, 97, 99, 98, 111, 112, 109, 110, 117,
	113, 114, 115, 116, 107, 108, 101, 102, 100, 103,
	105, 106, 0, 118, 94, 96, 95, 0, 0, 0,
	0, 0, 0, 416, 0, 104, 0, 0, 0, 0,
	0, 0, 0, 0, 119, 0, 120, 92, 93, 97,
	99, 98, 111, 112, 109, 110, 117, 113, 114, 115,
	116, 107, 108, 101, 102, 100, 103, 105, 106, 0,
	118, 94, 96, 95, 0, 0, 0, 0, 0, 0,
	0, 397, 104, 0, 0, 0, 0, 0, 0, 0,
	0, 119, 0, 120, 92, 93, 97, 99, 98, 111,
	112, 109, 110, 117, 113, 114, 115, 116, 107, 108,
	101, 102, 100, 103, 105, 106, 0, 118, 94, 96,
	95, 0, 0, 0, 0, 0, 0, 0, 396, 104,
	0, 0, 0, 0, 0, 0, 0, 0, 119, 0,
	120, 92, 93, 97, 99, 98, 111, 112, 109, 110,
	117, 113, 114, 115, 116, 107, 108, 101, 102, 100,
	103, 105, 106, 0, 118, 94, 96, 95, 0, 0,
	0, 0, 0, 0, 0, 383, 104, 0, 0, 0,
	0, 0, 0, 0, 0, 119, 0, 120, 92, 93,
	97, 99, 98, 111, 112, 109, 110, 117, 113, 114,
	115, 116, 107, 108, 101, 102, 100, 103, 105, 106,
	0, 118, 94, 96, 95, 0, 0, 0, 0, 0,
	0, 0, 382, 104, 0, 0, 0, 0, 0, 0,
	0, 0, 119, 0, 120, 92, 93, 97, 99, 98,
	111, 112, 109, 110, 117, 113, 114, 115, 116, 107,
	108, 101, 102, 100, 103, 105, 106, 0, 118, 94,
	96, 95, 0, 0, 0, 0, 0, 0, 0, 357,
	104, 0, 0, 0, 0, 0, 0, 0, 0, 119,
	0, 120, 92, 93, 97, 99, 98, 111, 112, 109,
	110, 117, 113, 114, 115, 116, 107, 108, 101, 102,
	100, 103, 105, 106, 0, 118, 94, 96, 95, 0,
	0, 0, 0, 0, 0, 0, 345, 104, 0, 0,
	0, 0, 0, 0, 0, 0, 119, 0, 120, 92,
	93, 97, 99, 98, 111, 112, 109, 110, 117, 113,
	114, 115, 116, 107, 108, 101, 102, 100, 103, 105,
	106, 0, 118, 94, 96, 95, 0, 0, 0, 0,
	0, 0, 0, 334, 104, 0, 0, 0, 0, 0,
	0, 0, 0, 119, 0, 120, 92, 93, 97, 99,
	98, 111, 112, 109, 110, 117, 113, 114, 115, 116,
	107, 108, 101, 102, 100, 103, 105, 106, 0, 118,
	94, 96, 95, 0, 0, 0, 0, 0, 0, 0,
	289, 104, 0, 0, 0, 0, 0, 0, 0, 0,
	119, 0, 120, 92, 93, 97, 99, 98, 111, 112,
	109, 110, 117, 113, 114, 115, 116, 107, 108, 101,
	102, 100, 103, 105, 106, 0, 118, 94, 96, 95,
	0, 0, 0, 0, 0, 412, 0, 0, 104, 0,
	0, 0, 0, 0, 0, 0, 0, 119, 0, 120,
	92, 93, 97, 99, 98, 111, 112, 109, 110, 117,
	113, 114, 115, 116, 107, 108, 101, 102, 100, 103,
	105, 106, 0, 118, 94, 96, 95, 0, 0, 0,
	0, 0, 410, 0, 0, 104, 0, 0, 0, 0,
	0, 0, 0, 0, 119, 0, 120, 92, 93, 97,
	99, 98, 111, 112, 109, 110, 117, 113, 114, 115,
	116, 107, 108, 101, 102, 100, 103, 105, 106, 0,
	118, 94, 96, 95, 0, 0, 0, 0, 0, 401,
	0, 0, 104, 0, 0, 0, 0, 0, 0, 0,
	0, 119, 0, 120, 92, 93, 97, 99, 98, 111,
	112, 109, 110, 117, 113, 114, 115, 116, 107, 108,
	101, 102, 100, 103, 105, 106, 0, 118, 94, 96,
	95, 0, 0, 0, 0, 0, 384, 0, 0, 104,
	0, 0, 0, 0, 0, 0, 0, 0, 119, 339,
	120, 92, 93, 97, 99, 98, 111, 112, 109, 110,
	117, 113, 114, 115, 116, 107, 108, 101, 102, 100,
	103, 105, 106, 0, 118, 94, 96, 95, 0, 0,
	284, 0, 0, 366, 0, 0, 104, 0, 0, 0,
	0, 0, 0, 0, 0, 119, 0, 120, 92, 93,
	97, 99, 98, 111, 112, 109, 110, 117, 113, 114,
	115, 116, 107, 108, 101, 102, 100, 103, 105, 106,
	0, 118, 94, 96, 95, 0, 0, 0, 0, 0,
	347, 0, 0, 104, 0, 0, 0, 0, 0, 0,
	0, 0, 119, 0, 120, 92, 93, 97, 99, 98,
	111, 112, 109, 110, 117, 113, 114, 115, 116, 107,
	108, 101, 102, 100, 103, 105, 106, 0, 118, 96,
	95, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	104, 0, 0, 0, 0, 0, 0, 0, 119, 0,
	120, 92, 93, 97, 99, 98, 111, 112, 109, 110,
	117, 113, 114, 115, 116, 107, 108, 101, 102, 100,
	103, 105, 106, 95, 118, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 104, 0, 0, 0,
	0, 119, 0, 120, 92, 93, 97, 99, 98, 111,
	112, 109, 110, 117, 113, 114, 115, 116, 107, 108,
	101, 102, 100, 103, 105, 106, 300, 118, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 104,
	0, 119, 0, 120, 92, 93, 97, 99, 98, 111,
	112, 109, 110, 117, 113, 114, 115, 116, 107, 108,
	101, 102, 100, 103, 105, 106, 0, 118, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 104,
	120, 92, 93, 97, 99, 98, 111, 112, 109, 110,
	117, 113, 114, 115, 116, 107, 108, 101, 102, 100,
	103, 105, 106, 0, 118, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 104, 93, 97, 99,
	98, 111, 112, 109, 110, 117, 113, 114, 115, 116,
	107, 108, 101, 102, 100, 103, 105, 106, 0, 118,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 104, 97, 99, 98, 111, 112, 109, 110, 117,
	113, 114, 115, 116, 107, 108, 101, 102, 100, 103,
	105, 106, 0, 118, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 104, 99, 98, 111, 112,
	109, 110, 117, 113, 114, 115, 116, 107, 108, 101,
	102, 100, 103, 105, 106, 0, 118, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 104, 98,
	111, 112, 109, 110, 117, 113, 114, 115, 116, 107,
	108, 101, 102, 100, 103, 105, 106, 0, 118, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	104, 111, 112, 109, 110, 117, 113, 114, 115, 116,
	107, 108, 101, 102, 100, 103, 105, 106, 0, 118,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 104,
}
var yyPact = [...]int{

	-1000, -1000, 969, -1000, -1000, -1000, -1000, -1000, 8, 969,
	7, 1868, 2055, 147, 152, 1278, -1000, 6, -1000, -1000,
	96, 4, 1194, 1868, 290, 290, 1868, 1868, 1868, 1868,
	1868, -1000, -1000, 1868, 1868, 1868, 1868, 1868, 1868, 1868,
	-6, 1868, -1000, 148, 1868, 1868, 1868, -1000, 164, -1000,
	34, 26, -1000, -3, -13, -34, 1868, 1868, -38, 1868,
	1868, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 97, 142, -1000, -18, -81, -81, -1000, -41, -1000,
	17, 18, 132, -1000, 857, 1868, 89, 1868, 45, -1000,
	2902, -1000, 1868, 1868, 1868, 1868, 1868, 1868, 1868, 1868,
	1868, 1868, 1868, 1868, 1868, 1868, 1868, 1868, 1868, 1868,
	1868, 1868, 1868, 1868, 1868, 1868, 1868, 1868, -3, 1756,
	1868, -70, 969, -71, -109, 227, 1644, 1868, 1868, 1868,
	1868, 1868, 1868, 1868, 1868, 1868, 1868, 1868, 1868, -1000,
	-1000, 1868, 1868, 1868, -8, 1194, -137, 261, -1000, -1000,
	2855, 290, -73, -1000, -1000, 164, 1868, -1000, 1194, -1000,
	184, 184, 182, 184, 2479, 184, 184, 184, 184, 184,
	184, 184, -1000, 1868, 184, -153, 90, 130, -1000, 122,
	1531, 290, 1939, 3031, 1939, 1868, -5, -5, -81, 24,
	2, -1000, -1000, 1868, 1868, 2902, 2902, 1868, 2902, 2902,
	124, -1000, 94, 37, 90, 1868, -1000, -1000, 1081, -1000,
	1194, 134, 132, 17, -1000, -1000, -1000, 2432, -76, -110,
	259, 2902, -1000, 1868, 3103, 3137, 2948, 1939, 2991, 3170,
	3233, 3202, 251, 251, 251, 182, 184, 182, 182, 276,
	276, 1141, 1141, 1141, 1141, 468, 468, 468, 468, 1141,
	-1000, 2808, 1868, 3068, 1868, -1000, 1868, -1000, -1000, 1939,
	290, 1939, 1939, 1939, 1939, 1939, 1939, 1939, 1939, 1939,
	1939, 1939, 1939, 2385, -142, 2902, 2761, -81, -1000, 1868,
	-1000, -91, 247, 1194, 1419, -1000, 1194, 2338, -143, -1000,
	-92, -1000, -1000, -1000, -1000, 187, 128, 2714, 16, 12,
	1868, -144, -81, -1000, -1000, 1868, -1000, -81, -1000, -1000,
	-1000, -1000, -10, 1868, 1868, -8, -10, 54, -1000, 2902,
	2291, 2244, -1000, -1000, -1000, 2667, -1000, 47, -1000, 2902,
	1868, -96, -1000, 17, 745, 1868, 1868, 1868, -1000, 1868,
	3068, 2197, 2150, 633, -1000, 521, -1000, -1000, -1000, 2620,
	242, 1868, -1000, 2902, 290, -88, -97, -1000, -1000, -1000,
	-148, -1000, -1000, 121, -1000, -1000, -1000, -1000, 1868, -1000,
	1939, -1000, -1000, 2573, -1000, -1000, -150, 2526, -1000, -1000,
	-1000, 1868, -1000, -1000, -1000, -1000, 1306, 2902, -1000, -1000,
	-1000, -1000, 2102, -111, 2902, 3068, 969, 225, -112, -1000,
	-1000, -1000, 1868, 1939, -1000, 1194, 242, -1000, -1000, 2008,
	-1000, -1000, -1000, -1000, -1000, 409, -114, 1868, -1000, -1000,
	-1000, 969, 1939, -102, 10, -129, -1000, -106, 969, 242,
	-1000, -1000, 969, -1000, -1000,
}
var yyPgo = [...]int{

	0, 361, 360, 22, 9, 229, 358, 0, 10, 49,
	4, 356, 8, 355, 197, 190, 124, 354, 353, 351,
	350, 17, 5, 349, 41, 18, 347, 346, 345, 343,
	342, 62, 341, 333, 1, 323, 320, 318, 313, 2,
	308, 298, 297, 296, 295, 14, 289, 288, 286, 262,
	277, 271, 3, 7, 6, 16,
}
var yyR1 = [...]int{

	0, 1, 55, 33, 33, 4, 4, 5, 5, 5,
	2, 34, 34, 6, 3, 3, 3, 3, 3, 3,
	3, 3, 35, 36, 36, 44, 44, 43, 43, 38,
	38, 37, 37, 45, 45, 46, 46, 10, 10, 42,
	42, 12, 39, 39, 47, 47, 13, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 20, 20, 20, 20, 14, 14, 15, 15, 18,
	18, 41, 41, 41, 48, 48, 24, 24, 24, 19,
	19, 19, 19, 19, 19, 19, 19, 19, 19, 19,
	19, 19, 19, 19, 19, 25, 25, 25, 7, 7,
	8, 8, 23, 26, 26, 26, 27, 27, 27, 28,
	28, 28, 28, 28, 28, 9, 9, 9, 16, 16,
	16, 29, 29, 30, 30, 30, 30, 30, 30, 21,
	21, 21, 22, 22, 22, 52, 53, 53, 51, 51,
	50, 50, 50, 50, 50, 50, 49, 49, 49, 49,
	31, 31, 31, 31, 31, 31, 31, 32, 32, 32,
	32, 17, 17, 17, 17, 17, 17, 17, 40, 40,
	54,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 0, 1, 3, 1, 3, 2,
	1, 2, 0, 1, 3, 1, 1, 5, 7, 9,
	3, 2, 1, 1, 4, 5, 6, 1, 3, 6,
	7, 3, 6, 2, 3, 1, 3, 1, 2, 3,
	1, 1, 0, 1, 3, 1, 3, 6, 5, 3,
	4, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 5, 4, 3, 1, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 1, 3, 2, 1, 2, 4,
	2, 2, 4, 4, 2, 1, 1, 1, 1, 0,
	3, 0, 1, 1, 0, 1, 4, 3, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	2, 3, 3, 1, 1, 1, 3, 3, 1, 1,
	0, 1, 1, 1, 3, 1, 1, 3, 1, 1,
	4, 4, 4, 4, 1, 1, 1, 3, 1, 4,
	2, 3, 3, 1, 4, 4, 3, 3, 3, 1,
	3, 1, 1, 3, 1, 1, 0, 1, 3, 1,
	3, 1, 4, 2, 6, 4, 2, 2, 1, 2,
	1, 4, 3, 3, 3, 6, 3, 1, 1, 2,
	1, 4, 4, 2, 2, 4, 2, 2, 1, 3,
	1,
}
var yyChk = [...]int{

	-1000, -1, -33, -2, -3, 158, -43, -37, 103, 102,
	105, 82, -7, -44, -38, -9, -11, 101, -28, -29,
	-26, 140, 69, 71, 58, 59, 49, 50, 55, 57,
	160, -13, -17, 60, 61, 62, 63, 64, 65, 66,
	100, 67, -19, 164, 13, 14, 16, -16, -25, -20,
	-14, -23, -24, 70, 131, 132, 4, 5, 6, 7,
	8, 83, 84, 92, 93, 94, 96, 97, 98, 99,
	95, 149, 165, 86, 166, -5, -27, 76, 141, 89,
	-4, 154, 155, 85, -34, 160, -3, 160, -42, -12,
	-7, 162, 33, 34, 10, 12, 11, 35, 37, 36,
	51, 49, 50, 52, 68, 53, 54, 47, 48, 40,
	41, 38, 39, 43, 44, 45, 46, 42, 56, 30,
	32, 73, 74, 73, 75, 74, 17, 18, 19, 20,
	29, 21, 22, 23, 24, 25, 26, 27, 28, 58,
	59, 160, 69, 158, 139, 160, -52, -51, -53, -50,
	-7, 37, 140, -7, -9, -25, 160, -24, 69, -9,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -18, 160, -7, -41, 88, -49, -31, 86,
	151, 152, -7, -7, -7, 69, 153, 153, -15, -14,
	-30, -5, -16, 160, 160, -7, -7, 160, -7, -7,
	88, 150, -49, -49, 88, 158, -16, -45, 160, -45,
	160, 155, 155, -4, -6, 159, -3, -7, 103, -39,
	-47, -7, 162, 9, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-15, -7, 31, -7, 160, -3, 160, 162, 31, -7,
	37, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -8, -7, -7, -22, 85, 158,
	-16, -52, 163, 9, 15, -9, 160, -7, -52, 161,
	-8, 164, -31, -31, 88, 69, 139, -7, 90, -9,
	15, -8, -21, -55, -16, 158, 85, -21, -55, -16,
	-48, -45, 153, 69, 158, 139, 153, -40, -54, -7,
	-7, -7, 150, 150, 165, -7, 161, -46, -10, -7,
	156, -52, 85, -4, 161, 160, 162, 9, -12, 31,
	-7, -7, -7, -34, -9, 161, 163, 159, -45, -7,
	161, 17, -53, -7, 37, 140, -52, 161, 163, 161,
	-32, 85, 91, 50, 86, 85, 159, 159, 69, 159,
	-7, 163, -45, -7, -45, -16, -8, -7, -22, -16,
	161, 9, 161, 161, 159, 161, 9, -7, 161, -36,
	-3, 31, -7, -39, -7, -7, 161, 161, 75, -3,
	31, 159, 17, -7, -9, 160, 161, 163, 91, -7,
	159, 163, 159, -54, -10, -34, 161, 162, -3, 31,
	162, -34, -7, -52, 163, 104, 162, -39, -34, 161,
	159, 162, 161, -35, -3,
}
var yyDef = [...]int{

	4, -2, 1, 3, 10, 12, 15, 16, 0, 0,
	0, 0, 0, 27, 0, -2, 159, 0, -2, 176,
	162, 0, 196, 0, 0, 0, 0, 0, 0, 0,
	0, 100, 104, 0, 0, 0, 0, 0, 0, 0,
	129, 0, 114, 131, 0, 117, 0, 169, 154, 174,
	0, 0, -2, 0, 0, 0, 0, 0, 0, 0,
	0, 139, 140, 141, 142, 143, 144, 145, 146, 147,
	148, 0, 0, 178, 0, -2, 0, 125, 0, 138,
	7, 0, 0, 5, 0, 0, 0, 42, 0, 40,
	41, 21, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	66, 0, 160, 0, 0, 196, 0, 195, 199, 197,
	201, 0, 0, 51, -2, 0, 0, -2, 196, -2,
	85, 86, 87, 88, 0, 105, 106, 107, 108, 109,
	110, 111, 112, 160, 113, 0, 132, 133, 208, 210,
	0, 0, 116, 118, 120, 160, 0, 0, 134, 127,
	128, 126, 183, 0, 0, 223, 224, 0, 226, 227,
	0, 150, 0, 0, 0, 0, 180, 121, 0, 124,
	196, 0, 0, 9, 11, 14, 13, 0, 0, 0,
	43, 45, 20, 0, 68, 69, 70, 71, 72, 73,
	74, 75, 76, 77, 78, 79, 80, 81, 82, 83,
	84, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	98, 0, 0, 103, 0, 28, 0, 31, 12, 49,
	0, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 0, 0, 161, 0, 177, 192, 0,
	194, 0, 137, 196, 0, -2, 196, 0, 0, -2,
	0, 115, 209, 206, 207, 0, 0, 0, 0, 163,
	0, 0, 0, -2, -2, 0, 2, 0, -2, -2,
	46, 135, 0, 160, 0, 0, 0, 0, 228, 230,
	0, 0, 149, 152, 151, 0, 33, 0, 35, 37,
	0, 0, 6, 8, 0, 0, 42, 0, 39, 0,
	102, 0, 0, 0, -2, 0, 170, 172, 173, 0,
	0, 0, 198, 200, 0, 0, 0, -2, 137, 130,
	0, 217, 218, 0, 220, 212, 213, 214, 0, 216,
	119, 171, 122, 0, 123, 187, 0, 0, 186, 188,
	221, 0, 222, 225, 179, 34, 0, 38, 136, 17,
	23, 12, 0, 0, 44, 101, 0, 0, 0, 25,
	12, 193, 0, 48, -2, 196, 205, 211, 219, 0,
	190, 184, 185, 229, 36, 0, 0, 42, 26, 12,
	32, 29, 47, 0, 0, 0, 18, 0, 30, 204,
	215, 24, 0, 19, 22,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 165, 3, 166, 54, 37, 3,
	160, 161, 52, 49, 9, 50, 51, 53, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 31, 162,
	43, 17, 45, 30, 67, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 69, 3, 163, 36, 3, 164, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 158, 35, 159, 57,
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:387
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:389
		{
			yyVAL.stmts = []ast.Statement{}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:393
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:404
		{
			yyVAL.stmt = ast.NewBlockStatement(yyDollar[2].stmts...)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:405
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:406
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:408
		{
			yyVAL.stmt = ast.NewWhileStatement(yyDollar[1].tok, yyDollar[3].expr, yyDollar[5].stmt)
		}
	case 18:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:410
		{
			yyVAL.stmt = ast.NewDoWhileStatement(yyDollar[1].tok, yyDollar[5].expr, yyDollar[2].stmt)
		}
	case 19:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:412
		{
			yyVAL.stmt = ast.NewForStatement(yyDollar[1].tok, yyDollar[3].exprs, yyDollar[5].exprs, yyDollar[7].exprs, yyDollar[9].stmt)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:420
		{
			yyVAL.stmt = ast.NewEchoStatement(yyDollar[1].tok, yyDollar[2].exprs)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:422
		{
			yyVAL.stmt = ast.NewExpressionStatement(yyDollar[1].expr)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:540
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:578
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:579
		{
			yyVAL.stmt = ast.NewAltWhileStatement(yyDollar[3].tok, yyDollar[2].stmts...)
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:584
		{
			yyVAL.stmt = ast.NewIfStatement(yyDollar[1].tok, yyDollar[3].expr, yyDollar[5].stmt, nil)
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:586
		{
			yyVAL.stmt = ast.NewIfStatement(yyDollar[2].tok, yyDollar[4].expr, yyDollar[6].stmt, yyVAL.stmt)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:590
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:592
		{
			yyVAL.stmt = ast.NewIfStatement(yyDollar[2].tok, nil, yyDollar[3].stmt, yyVAL.stmt)
		}
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:598
		{
			yyVAL.stmt = ast.NewAltIfStatement(yyDollar[1].tok, yyDollar[3].expr, yyDollar[6].stmts, nil)
		}
	case 30:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:600
		{
			yyVAL.stmt = ast.NewAltIfStatement(yyDollar[2].tok, yyDollar[4].expr, yyDollar[7].stmts, yyVAL.stmt)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:605
		{
			stmt := yyDollar[1].stmt
			yyVAL.stmt = ast.NewAltIfStatement(yyDollar[2].tok, nil, nil, stmt)
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:610
		{
			stmt := ast.NewAltIfStatement(yyDollar[2].tok, nil, yyDollar[4].stmts, yyVAL.stmt)
			yyVAL.stmt = ast.NewAltIfStatement(yyDollar[5].tok, nil, nil, stmt)
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:661
		{
			yyVAL.expr = ast.NewArgumentListExpression()
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:662
		{
			yyVAL.expr = ast.NewArgumentListExpression(yyDollar[2].exprs...)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:667
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:669
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:673
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:674
		{
			yyVAL.expr = ast.NewArgumentExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:829
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:830
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:834
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:838
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:839
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:843
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:844
		{
			yyVAL.exprs = []ast.Expression{yyDollar[1].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:861
		{
			yyVAL.expr = ast.NewNewExpression(yyDollar[1].tok, yyDollar[2].expr, yyDollar[3].expr)
		}
	case 47:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:868
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...), yyDollar[6].expr, false)
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:870
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...), yyDollar[5].expr, false)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:872
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:874
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, yyDollar[1].expr, yyDollar[4].expr, true)
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:875
		{
			yyVAL.expr = ast.NewCloneExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:877
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PlusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:879
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MinusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:881
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MulEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:883
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PowEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:885
		{
			yyVAL.expr = ast.NewAssignExpression(ast.DivEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:887
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ConcatEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:889
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ModEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:891
		{
			yyVAL.expr = ast.NewAssignExpression(ast.AndEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:893
		{
			yyVAL.expr = ast.NewAssignExpression(ast.QrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:895
		{
			yyVAL.expr = ast.NewAssignExpression(ast.XorEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:897
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SlEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:899
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:900
		{
			yyVAL.expr = ast.NewIncrementExpression(ast.PostInc, yyDollar[1].expr)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:901
		{
			yyVAL.expr = ast.NewIncrementExpression(ast.PreInc, yyDollar[2].expr)
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:902
		{
			yyVAL.expr = ast.NewDecrementExpression(ast.PostDec, yyDollar[1].expr)
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:903
		{
			yyVAL.expr = ast.NewDecrementExpression(ast.PreDec, yyDollar[2].expr)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:905
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BooleanOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:907
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BooleanAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:909
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:911
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:913
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalXor, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:914
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:915
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:916
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwXor, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:917
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Concat, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:918
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Add, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:919
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sub, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:920
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Mul, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:921
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Pow, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:922
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Div, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:923
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Mod, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:924
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sl, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:925
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:926
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.UnaryPlus, yyDollar[2].expr)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:927
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.UnaryMinus, yyDollar[2].expr)
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:928
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.BoolNot, yyDollar[2].expr)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:929
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.BwNot, yyDollar[2].expr)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:931
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsIdentical, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:933
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsNotIdentical, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:935
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:937
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsNotEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:939
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Smaller, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:941
		{
			yyVAL.expr = ast.NewInfixExpression(ast.SmallerOrEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:943
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Greater, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:945
		{
			yyVAL.expr = ast.NewInfixExpression(ast.GreaterOrEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:947
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Spaceship, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:949
		{
			yyVAL.expr = ast.NewInfixExpression(ast.InstanceOf, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:950
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:951
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:953
		{
			yyVAL.expr = ast.NewTernaryOperatorExpression(yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr)
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:955
		{
			yyVAL.expr = ast.NewTernaryOperatorExpression(yyDollar[1].expr, nil, yyDollar[4].expr)
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:957
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Coalesce, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:958
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:959
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:960
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:961
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:962
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:963
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:964
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:965
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:966
		{
			yyVAL.expr = ast.NewExitExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:967
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.Silence, yyDollar[2].expr)
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:968
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:969
		{
			yyVAL.expr = ast.NewBackticksExpression(yyDollar[2].exprs...)
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:970
		{
			yyVAL.expr = ast.NewPrintExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:971
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, nil)
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:972
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:973
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, ast.NewArrayPairExpression(yyDollar[2].expr, yyDollar[4].expr, false))
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:974
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1022
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1024
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1026
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1028
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1032
		{
			yyVAL.expr = ast.NewStaticLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1033
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1037
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1038
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 129:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1042
		{
			yyVAL.expr = nil
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1043
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 131:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1048
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1049
		{
			yyVAL.exprs = []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)}
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1050
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 134:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1054
		{
			yyVAL.expr = nil
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1055
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 136:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1059
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Long, yyDollar[3].exprs...)
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1060
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...)
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1061
		{
			yyVAL.expr = ast.NewConstantEncapsedStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1065
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1066
		{
			yyVAL.expr = ast.NewDoubleLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1067
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1068
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1069
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1070
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1071
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1072
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1073
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1074
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1075
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1076
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[2].tok)
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1077
		{
			yyVAL.expr = ast.NewEncapsListExpression(yyDollar[2].exprs...)
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1078
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, yyDollar[2].exprs...)
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1079
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1080
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1084
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1086
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1088
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1092
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1093
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 160:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1097
		{
			yyVAL.expr = nil
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1098
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1102
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1106
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1107
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1108
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1112
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1113
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1114
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1119
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 170:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1121
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 171:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1123
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 172:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1125
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Curly, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1127
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Prop, yyDollar[1].expr, []ast.Expression{yyDollar[3].expr, yyDollar[4].expr}...)
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1128
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1133
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1135
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1137
		{
			yyVAL.expr = ast.NewVariableExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1141
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1142
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.CurlyOpen, yyDollar[3].expr)
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1143
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.Var, yyDollar[2].expr)
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1148
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1150
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1155
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 184:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1157
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1159
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Curly, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1161
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Prop, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1163
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1165
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1169
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1170
		{
			yyVAL.expr = ast.NewMemberNameExpression(yyDollar[2].expr)
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1171
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1175
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1176
		{
			yyVAL.expr = ast.NewPropertyNameExpression(yyDollar[2].expr)
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1177
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1182
		{ /* allow single trailing comma */
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1186
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1187
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1192
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].exprs...)
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1194
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 200:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1199
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[3].expr, false)}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1201
		{
			yyVAL.exprs = []ast.Expression{yyDollar[1].expr}
		}
	case 202:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1203
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[4].expr, true)}
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1205
		{
			yyVAL.exprs = []ast.Expression{ast.NewAmpersandLiteral(yyDollar[2].expr)}
		}
	case 204:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1207
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, ast.NewListExpression(yyDollar[3].tok, yyDollar[5].exprs...), false)}
		}
	case 205:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1209
		{
			yyVAL.exprs = []ast.Expression{ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...)}
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1214
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1216
		{
			yyVAL.exprs = append(yyDollar[1].exprs, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1218
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1220
		{
			yyVAL.exprs = append(yyVAL.exprs, []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal), yyDollar[2].expr}...)
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1225
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 211:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1227
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Dim, yyDollar[3].expr)
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1229
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Prop, ast.NewStringLiteral(yyDollar[3].tok, yyDollar[3].tok.Literal))
		}
	case 213:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1231
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, yyDollar[2].expr)
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1233
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 215:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1235
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DimInDollarOpenCurlyBraces, []ast.Expression{ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal), yyDollar[4].expr}...)
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1236
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.CurlyOpen, yyDollar[2].expr)
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1240
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1241
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1242
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[2].tok, "-"+yyDollar[2].tok.Literal)
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1243
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 221:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1248
		{
			yyVAL.expr = ast.NewIssetExpression(yyDollar[1].tok, yyDollar[3].exprs...)
		}
	case 222:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1249
		{
			yyVAL.expr = ast.NewEmptyExpression(yyDollar[1].tok, yyDollar[3].expr)
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1251
		{
			yyVAL.expr = ast.NewIncludeExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1253
		{
			yyVAL.expr = ast.NewIncludeExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 225:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1255
		{
			yyVAL.expr = ast.NewEvalExpression(yyDollar[1].tok, yyDollar[3].expr)
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1257
		{
			yyVAL.expr = ast.NewRequireExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 227:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1259
		{
			yyVAL.expr = ast.NewRequireExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1263
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1265
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1269
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
