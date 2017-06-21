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
	"';'",
	"'('",
	"')'",
	"']'",
	"'`'",
	"'\"'",
	"'$'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:1270

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
	-1, 12,
	69, 153,
	139, 153,
	153, 153,
	158, 153,
	-2, 148,
	-1, 15,
	161, 156,
	-2, 165,
	-1, 49,
	69, 155,
	139, 155,
	153, 155,
	158, 155,
	161, 158,
	-2, 143,
	-1, 72,
	153, 116,
	-2, 145,
	-1, 148,
	69, 153,
	139, 153,
	153, 153,
	158, 153,
	-2, 55,
	-1, 151,
	161, 158,
	-2, 155,
	-1, 153,
	69, 153,
	139, 153,
	153, 153,
	158, 153,
	-2, 57,
	-1, 230,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 79,
	-1, 231,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 80,
	-1, 232,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 81,
	-1, 233,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 82,
	-1, 234,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 83,
	-1, 235,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 84,
	-1, 236,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 85,
	-1, 237,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 86,
	-1, 238,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 87,
	-1, 274,
	9, 193,
	162, 193,
	163, 193,
	-2, 153,
	-1, 278,
	69, 154,
	139, 154,
	153, 154,
	158, 154,
	161, 157,
	-2, 89,
	-1, 292,
	161, 179,
	-2, 146,
	-1, 293,
	161, 181,
	-2, 171,
	-1, 297,
	161, 179,
	-2, 147,
	-1, 298,
	161, 181,
	-2, 172,
	-1, 329,
	69, 153,
	139, 153,
	153, 153,
	158, 153,
	-2, 40,
	-1, 342,
	161, 157,
	-2, 154,
	-1, 383,
	9, 192,
	162, 192,
	163, 192,
	-2, 153,
}

const yyPrivate = 57344

const yyLast = 2943

var yyAct = [...]int{

	9, 81, 140, 210, 317, 44, 4, 307, 266, 84,
	142, 263, 83, 292, 291, 280, 390, 386, 45, 371,
	144, 147, 356, 343, 154, 155, 156, 157, 158, 366,
	49, 159, 160, 161, 162, 163, 164, 165, 12, 168,
	149, 149, 176, 177, 178, 331, 271, 402, 385, 373,
	344, 335, 151, 151, 189, 190, 186, 192, 193, 172,
	148, 153, 152, 77, 384, 201, 202, 275, 245, 74,
	243, 74, 70, 212, 204, 191, 188, 200, 80, 70,
	80, 70, 76, 187, 267, 70, 167, 213, 214, 215,
	216, 217, 218, 219, 220, 221, 222, 223, 224, 225,
	226, 227, 228, 229, 230, 231, 232, 233, 234, 235,
	236, 237, 238, 70, 240, 242, 139, 135, 186, 396,
	244, 248, 250, 251, 252, 253, 254, 255, 256, 257,
	258, 259, 260, 261, 75, 246, 262, 264, 265, 203,
	144, 403, 270, 207, 269, 295, 70, 78, 79, 78,
	79, 276, 71, 144, 150, 277, 173, 268, 283, 71,
	353, 71, 136, 302, 149, 71, 354, 205, 264, 206,
	301, 181, 370, 173, 180, 286, 151, 387, 179, 279,
	264, 182, 365, 350, 274, 199, 293, 298, 308, 309,
	284, 290, 310, 71, 149, 297, 296, 173, 80, 283,
	314, 173, 321, 318, 395, 144, 151, 320, 47, 117,
	119, 118, 98, 84, 288, 115, 116, 247, 294, 381,
	72, 174, 175, 336, 211, 323, 71, 173, 171, 194,
	281, 282, 138, 304, 173, 313, 283, 311, 174, 175,
	141, 272, 325, 173, 326, 198, 327, 305, 300, 328,
	352, 137, 303, 173, 281, 170, 282, 282, 281, 183,
	285, 312, 174, 175, 143, 112, 174, 175, 149, 334,
	322, 185, 299, 144, 338, 316, 144, 98, 341, 10,
	151, 348, 6, 337, 82, 97, 99, 100, 329, 112,
	355, 195, 174, 175, 239, 358, 169, 196, 197, 174,
	175, 98, 306, 264, 362, 11, 7, 360, 174, 175,
	269, 364, 2, 363, 361, 345, 346, 349, 174, 175,
	372, 183, 347, 184, 16, 374, 15, 73, 17, 48,
	46, 39, 333, 185, 378, 166, 29, 382, 107, 108,
	109, 110, 101, 102, 95, 96, 94, 97, 99, 100,
	28, 112, 13, 208, 388, 3, 1, 357, 149, 0,
	0, 0, 359, 98, 0, 0, 0, 308, 0, 0,
	151, 0, 318, 0, 392, 0, 393, 0, 383, 394,
	0, 397, 398, 0, 0, 144, 0, 399, 0, 53,
	54, 55, 56, 57, 0, 0, 0, 401, 41, 42,
	0, 43, 101, 102, 95, 96, 94, 97, 99, 100,
	0, 112, 0, 0, 0, 0, 379, 95, 96, 94,
	97, 99, 100, 98, 112, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 23, 24, 98, 0, 0, 0,
	25, 0, 26, 21, 22, 30, 31, 32, 33, 34,
	35, 36, 38, 0, 19, 50, 20, 0, 0, 0,
	0, 74, 0, 0, 0, 0, 0, 8, 58, 59,
	80, 70, 0, 0, 76, 0, 0, 60, 61, 62,
	67, 63, 64, 65, 66, 37, 14, 120, 121, 122,
	123, 125, 126, 127, 128, 129, 130, 131, 132, 124,
	0, 0, 53, 54, 55, 56, 57, 0, 0, 0,
	0, 41, 42, 0, 43, 0, 51, 52, 0, 0,
	0, 0, 0, 0, 0, 18, 75, 0, 133, 134,
	0, 0, 0, 0, 68, 0, 0, 0, 0, 78,
	79, 0, 0, 5, 0, 0, 27, 23, 24, 40,
	69, 71, 0, 25, 0, 26, 21, 22, 30, 31,
	32, 33, 34, 35, 36, 38, 0, 19, 50, 20,
	0, 0, 0, 377, 74, 0, 0, 0, 0, 0,
	8, 58, 59, 80, 70, 0, 0, 76, 0, 0,
	60, 61, 62, 67, 63, 64, 65, 66, 37, 14,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 53, 54, 55, 56, 57,
	0, 0, 0, 0, 41, 42, 0, 43, 0, 51,
	52, 0, 0, 0, 0, 0, 0, 0, 18, 75,
	0, 0, 0, 0, 0, 0, 0, 68, 0, 0,
	0, 0, 78, 79, 0, 0, 5, 0, 0, 27,
	23, 24, 40, 69, 71, 0, 25, 0, 26, 21,
	22, 30, 31, 32, 33, 34, 35, 36, 38, 0,
	19, 50, 20, 0, 0, 0, 0, 74, 0, 0,
	0, 0, 0, 8, 58, 59, 80, 70, 0, 0,
	76, 0, 0, 60, 61, 62, 67, 63, 64, 65,
	66, 37, 14, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 53, 54,
	55, 56, 57, 0, 0, 0, 0, 41, 42, 0,
	43, 0, 51, 52, 0, 0, 0, 0, 0, 0,
	0, 18, 75, 0, 0, 0, 0, 0, 0, 0,
	68, 0, 0, 0, 0, 78, 79, 0, 0, 5,
	209, 0, 27, 23, 24, 40, 69, 71, 0, 25,
	0, 26, 21, 22, 30, 31, 32, 33, 34, 35,
	36, 38, 0, 19, 50, 20, 0, 0, 0, 0,
	74, 0, 0, 0, 0, 0, 8, 58, 59, 80,
	70, 0, 0, 76, 0, 0, 60, 61, 62, 67,
	63, 64, 65, 66, 37, 14, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 53, 54, 55, 56, 57, 0, 0, 0, 0,
	41, 42, 0, 43, 0, 51, 52, 0, 0, 0,
	0, 0, 0, 0, 18, 75, 0, 0, 0, 0,
	0, 0, 0, 68, 0, 0, 0, 0, 78, 79,
	0, 0, 5, 0, 0, 27, 23, 24, 40, 69,
	71, 0, 25, 0, 26, 21, 22, 30, 31, 32,
	33, 34, 35, 36, 38, 0, 19, 50, 20, 0,
	0, 0, 0, 74, 0, 0, 0, 0, 0, 0,
	58, 59, 80, 70, 0, 0, 76, 0, 0, 60,
	61, 62, 67, 63, 64, 65, 66, 37, 87, 91,
	93, 92, 105, 106, 103, 104, 111, 107, 108, 109,
	110, 101, 102, 95, 96, 94, 97, 99, 100, 0,
	112, 53, 54, 55, 56, 57, 0, 0, 51, 52,
	41, 42, 98, 43, 0, 0, 0, 18, 75, 0,
	0, 0, 0, 0, 0, 0, 68, 0, 0, 0,
	0, 78, 79, 319, 145, 0, 0, 0, 27, 315,
	0, 40, 69, 71, 0, 0, 23, 24, 0, 0,
	0, 0, 25, 0, 26, 21, 22, 30, 31, 32,
	33, 34, 35, 36, 38, 0, 19, 50, 20, 0,
	0, 0, 0, 74, 0, 0, 0, 0, 0, 0,
	58, 59, 80, 70, 0, 0, 76, 0, 0, 60,
	61, 62, 67, 63, 64, 65, 66, 37, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 53, 54, 55, 56, 57, 0,
	0, 0, 0, 41, 42, 0, 43, 0, 51, 52,
	0, 0, 0, 0, 0, 0, 0, 146, 75, 0,
	0, 0, 0, 0, 0, 0, 68, 0, 0, 0,
	0, 78, 79, 0, 0, 0, 0, 0, 27, 23,
	24, 40, 69, 71, 0, 25, 0, 26, 21, 22,
	30, 31, 32, 33, 34, 35, 36, 38, 0, 19,
	50, 20, 0, 0, 0, 0, 74, 0, 0, 0,
	0, 0, 0, 58, 59, 80, 70, 0, 0, 76,
	0, 0, 60, 61, 62, 67, 63, 64, 65, 66,
	37, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 53, 54, 55,
	56, 57, 0, 0, 0, 0, 41, 42, 0, 43,
	0, 51, 52, 0, 0, 0, 0, 0, 0, 0,
	18, 75, 0, 0, 0, 0, 0, 0, 0, 68,
	339, 0, 0, 0, 78, 79, 319, 0, 0, 0,
	0, 27, 23, 24, 40, 69, 71, 0, 25, 0,
	26, 21, 22, 30, 31, 32, 33, 34, 35, 36,
	38, 0, 19, 50, 20, 0, 0, 0, 0, 74,
	0, 0, 0, 0, 0, 0, 58, 59, 80, 70,
	0, 0, 76, 0, 0, 60, 61, 62, 67, 63,
	64, 65, 66, 37, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	53, 54, 55, 56, 57, 0, 0, 0, 0, 41,
	42, 0, 43, 0, 51, 52, 0, 0, 0, 0,
	0, 0, 0, 340, 75, 0, 0, 0, 0, 0,
	0, 0, 68, 0, 0, 0, 0, 78, 79, 0,
	0, 0, 0, 0, 27, 23, 24, 40, 69, 71,
	0, 25, 0, 26, 21, 22, 30, 31, 32, 33,
	34, 35, 36, 38, 0, 19, 50, 20, 0, 0,
	0, 0, 74, 0, 0, 0, 0, 0, 0, 58,
	59, 80, 70, 0, 0, 76, 287, 0, 60, 61,
	62, 67, 63, 64, 65, 66, 37, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 53, 54, 55, 56, 57, 0, 0,
	0, 0, 41, 42, 0, 43, 0, 51, 52, 0,
	0, 0, 0, 0, 0, 0, 18, 75, 0, 0,
	0, 0, 0, 0, 0, 68, 249, 0, 0, 0,
	78, 79, 0, 0, 0, 0, 0, 27, 23, 24,
	40, 69, 71, 0, 25, 0, 26, 21, 22, 30,
	31, 32, 33, 34, 35, 36, 38, 0, 19, 50,
	20, 0, 0, 0, 0, 74, 0, 0, 0, 0,
	0, 0, 58, 59, 80, 70, 0, 0, 76, 0,
	0, 60, 61, 62, 67, 63, 64, 65, 66, 37,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 53, 54, 55, 56,
	57, 0, 0, 0, 0, 41, 42, 0, 43, 0,
	51, 52, 0, 0, 0, 0, 0, 0, 0, 18,
	75, 0, 0, 241, 0, 0, 0, 0, 68, 0,
	0, 0, 0, 78, 79, 0, 0, 0, 0, 0,
	27, 23, 24, 40, 69, 71, 0, 25, 0, 26,
	21, 22, 30, 31, 32, 33, 34, 35, 36, 38,
	0, 19, 50, 20, 0, 0, 0, 0, 74, 0,
	0, 0, 0, 0, 0, 58, 59, 80, 70, 0,
	0, 76, 0, 0, 60, 61, 62, 67, 63, 64,
	65, 66, 37, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 53,
	54, 55, 56, 57, 0, 0, 0, 0, 41, 42,
	0, 43, 0, 51, 52, 0, 0, 0, 0, 0,
	0, 0, 18, 75, 0, 0, 0, 0, 0, 0,
	0, 68, 0, 0, 0, 0, 78, 79, 0, 0,
	0, 0, 0, 27, 23, 24, 40, 69, 71, 0,
	25, 0, 26, 21, 22, 30, 31, 32, 33, 34,
	35, 36, 38, 0, 19, 50, 20, 0, 0, 0,
	0, 74, 0, 0, 0, 0, 0, 0, 58, 59,
	80, 70, 0, 0, 76, 0, 0, 60, 61, 62,
	67, 63, 64, 65, 66, 37, 113, 0, 114, 86,
	87, 91, 93, 92, 105, 106, 103, 104, 111, 107,
	108, 109, 110, 101, 102, 95, 96, 94, 97, 99,
	100, 0, 112, 0, 0, 0, 51, 52, 0, 0,
	0, 0, 0, 0, 98, 18, 75, 0, 0, 0,
	0, 0, 0, 0, 68, 88, 90, 89, 0, 78,
	79, 0, 0, 0, 0, 0, 27, 0, 0, 40,
	69, 71, 0, 0, 0, 113, 0, 114, 86, 87,
	91, 93, 92, 105, 106, 103, 104, 111, 107, 108,
	109, 110, 101, 102, 95, 96, 94, 97, 99, 100,
	0, 112, 88, 90, 89, 0, 0, 0, 0, 0,
	0, 0, 0, 98, 0, 0, 0, 0, 0, 0,
	0, 0, 113, 0, 114, 86, 87, 91, 93, 92,
	105, 106, 103, 104, 111, 107, 108, 109, 110, 101,
	102, 95, 96, 94, 97, 99, 100, 0, 112, 88,
	90, 89, 0, 0, 0, 0, 0, 0, 0, 0,
	98, 0, 0, 0, 0, 0, 0, 0, 0, 113,
	0, 114, 86, 87, 91, 93, 92, 105, 106, 103,
	104, 111, 107, 108, 109, 110, 101, 102, 95, 96,
	94, 97, 99, 100, 0, 112, 0, 88, 90, 89,
	0, 0, 0, 0, 0, 0, 0, 98, 400, 0,
	0, 0, 0, 0, 0, 0, 0, 113, 0, 114,
	86, 87, 91, 93, 92, 105, 106, 103, 104, 111,
	107, 108, 109, 110, 101, 102, 95, 96, 94, 97,
	99, 100, 0, 112, 88, 90, 89, 0, 0, 0,
	0, 0, 0, 0, 376, 98, 0, 0, 0, 0,
	0, 0, 0, 0, 113, 0, 114, 86, 87, 91,
	93, 92, 105, 106, 103, 104, 111, 107, 108, 109,
	110, 101, 102, 95, 96, 94, 97, 99, 100, 0,
	112, 0, 88, 90, 89, 0, 0, 0, 0, 0,
	0, 375, 98, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 113, 0, 114, 86, 87, 91, 93, 92,
	105, 106, 103, 104, 111, 107, 108, 109, 110, 101,
	102, 95, 96, 94, 97, 99, 100, 0, 112, 88,
	90, 89, 0, 0, 0, 0, 0, 0, 0, 368,
	98, 0, 0, 0, 0, 0, 0, 0, 0, 113,
	0, 114, 86, 87, 91, 93, 92, 105, 106, 103,
	104, 111, 107, 108, 109, 110, 101, 102, 95, 96,
	94, 97, 99, 100, 0, 112, 0, 88, 90, 89,
	0, 0, 0, 0, 0, 0, 367, 98, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 113, 0, 114,
	86, 87, 91, 93, 92, 105, 106, 103, 104, 111,
	107, 108, 109, 110, 101, 102, 95, 96, 94, 97,
	99, 100, 0, 112, 88, 90, 89, 0, 0, 0,
	0, 0, 0, 0, 342, 98, 0, 0, 0, 0,
	0, 0, 0, 0, 113, 0, 114, 86, 87, 91,
	93, 92, 105, 106, 103, 104, 111, 107, 108, 109,
	110, 101, 102, 95, 96, 94, 97, 99, 100, 0,
	112, 0, 88, 90, 89, 0, 0, 0, 0, 0,
	0, 330, 98, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 113, 0, 114, 86, 87, 91, 93, 92,
	105, 106, 103, 104, 111, 107, 108, 109, 110, 101,
	102, 95, 96, 94, 97, 99, 100, 0, 112, 88,
	90, 89, 0, 0, 0, 0, 0, 0, 0, 278,
	98, 0, 0, 0, 0, 0, 0, 0, 0, 113,
	0, 114, 86, 87, 91, 93, 92, 105, 106, 103,
	104, 111, 107, 108, 109, 110, 101, 102, 95, 96,
	94, 97, 99, 100, 0, 112, 88, 90, 89, 0,
	0, 0, 0, 0, 85, 0, 0, 98, 0, 0,
	0, 0, 0, 0, 0, 0, 113, 0, 114, 86,
	87, 91, 93, 92, 105, 106, 103, 104, 111, 107,
	108, 109, 110, 101, 102, 95, 96, 94, 97, 99,
	100, 0, 112, 88, 90, 89, 0, 0, 0, 0,
	0, 391, 0, 0, 98, 0, 0, 0, 0, 0,
	0, 0, 0, 113, 0, 114, 86, 87, 91, 93,
	92, 105, 106, 103, 104, 111, 107, 108, 109, 110,
	101, 102, 95, 96, 94, 97, 99, 100, 0, 112,
	88, 90, 89, 0, 0, 0, 0, 0, 389, 0,
	0, 98, 0, 0, 0, 0, 0, 0, 0, 0,
	113, 0, 114, 86, 87, 91, 93, 92, 105, 106,
	103, 104, 111, 107, 108, 109, 110, 101, 102, 95,
	96, 94, 97, 99, 100, 0, 112, 88, 90, 89,
	0, 0, 0, 0, 0, 380, 0, 0, 98, 0,
	0, 0, 0, 0, 0, 0, 0, 113, 0, 114,
	86, 87, 91, 93, 92, 105, 106, 103, 104, 111,
	107, 108, 109, 110, 101, 102, 95, 96, 94, 97,
	99, 100, 0, 112, 88, 90, 89, 0, 0, 0,
	0, 0, 369, 0, 0, 98, 0, 0, 0, 0,
	0, 0, 0, 0, 113, 324, 114, 86, 87, 91,
	93, 92, 105, 106, 103, 104, 111, 107, 108, 109,
	110, 101, 102, 95, 96, 94, 97, 99, 100, 0,
	112, 88, 90, 89, 0, 0, 273, 0, 0, 351,
	0, 0, 98, 0, 0, 0, 0, 0, 0, 0,
	0, 113, 0, 114, 86, 87, 91, 93, 92, 105,
	106, 103, 104, 111, 107, 108, 109, 110, 101, 102,
	95, 96, 94, 97, 99, 100, 0, 112, 88, 90,
	89, 0, 0, 0, 0, 0, 332, 0, 0, 98,
	0, 0, 0, 0, 0, 0, 0, 0, 113, 0,
	114, 86, 87, 91, 93, 92, 105, 106, 103, 104,
	111, 107, 108, 109, 110, 101, 102, 95, 96, 94,
	97, 99, 100, 0, 112, 90, 89, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 98, 0, 0, 0,
	0, 0, 0, 0, 113, 0, 114, 86, 87, 91,
	93, 92, 105, 106, 103, 104, 111, 107, 108, 109,
	110, 101, 102, 95, 96, 94, 97, 99, 100, 89,
	112, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 98, 0, 0, 0, 0, 113, 0, 114,
	86, 87, 91, 93, 92, 105, 106, 103, 104, 111,
	107, 108, 109, 110, 101, 102, 95, 96, 94, 97,
	99, 100, 289, 112, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 98, 0, 113, 0, 114,
	86, 87, 91, 93, 92, 105, 106, 103, 104, 111,
	107, 108, 109, 110, 101, 102, 95, 96, 94, 97,
	99, 100, 0, 112, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 98, 114, 86, 87, 91,
	93, 92, 105, 106, 103, 104, 111, 107, 108, 109,
	110, 101, 102, 95, 96, 94, 97, 99, 100, 0,
	112, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 98, 91, 93, 92, 105, 106, 103, 104,
	111, 107, 108, 109, 110, 101, 102, 95, 96, 94,
	97, 99, 100, 0, 112, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 98, 93, 92, 105,
	106, 103, 104, 111, 107, 108, 109, 110, 101, 102,
	95, 96, 94, 97, 99, 100, 0, 112, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 98,
	92, 105, 106, 103, 104, 111, 107, 108, 109, 110,
	101, 102, 95, 96, 94, 97, 99, 100, 0, 112,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 98, 105, 106, 103, 104, 111, 107, 108, 109,
	110, 101, 102, 95, 96, 94, 97, 99, 100, 0,
	112, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 98,
}
var yyPact = [...]int{

	-1000, -1000, 724, -1000, -1000, -1000, -1000, -1000, 1635, 2154,
	142, 136, 470, -1000, -44, -1000, -1000, 93, -45, 957,
	1635, -7, -7, 1635, 1635, 1635, 1635, 1635, -1000, -1000,
	1635, 1635, 1635, 1635, 1635, 1635, 1635, -75, 1635, -1000,
	167, 1635, 1635, 1635, -1000, 109, -1000, 21, 18, -1000,
	-5, -78, -85, 1635, 1635, -86, 1635, 1635, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 141, 157,
	-1000, 27, -95, -95, -1000, -87, -1000, 12, 14, 113,
	-1000, 611, 64, -1000, 2578, -1000, 1635, 1635, 1635, 1635,
	1635, 1635, 1635, 1635, 1635, 1635, 1635, 1635, 1635, 1635,
	1635, 1635, 1635, 1635, 1635, 1635, 1635, 1635, 1635, 1635,
	1635, 1635, -5, 1522, 1635, -91, 724, -93, -25, 186,
	1409, 1635, 1635, 1635, 1635, 1635, 1635, 1635, 1635, 1635,
	1635, 1635, 1635, -1000, -1000, 1635, 1635, 1635, -1, 957,
	-117, 232, -1000, -1000, 2531, -7, -94, -1000, -1000, 109,
	1635, -1000, 957, -1000, 144, 144, 209, 144, 2107, 144,
	144, 144, 144, 144, 144, 144, -1000, 1635, 144, -149,
	115, 148, -1000, 121, 1296, -7, 1706, 2707, 1706, 1635,
	60, 60, -95, 17, 94, -1000, -1000, 1635, 1635, 2578,
	2578, 1635, 2578, 2578, 87, -1000, 111, 70, 115, 1635,
	-1000, -1000, 837, -1000, 957, 117, 113, 12, -1000, -1000,
	-1000, -1000, 1635, 904, 2778, 2624, 1706, 2667, 2811, 2874,
	2843, 233, 233, 233, 209, 144, 209, 209, 368, 368,
	295, 295, 295, 295, 355, 355, 355, 355, 295, -1000,
	2484, 1635, 2744, 1635, -1000, 1635, -1000, -1000, 1706, -7,
	1706, 1706, 1706, 1706, 1706, 1706, 1706, 1706, 1706, 1706,
	1706, 1706, 2059, -118, 2578, 2437, -95, -1000, 1635, -1000,
	-111, 206, 957, 1183, -1000, 957, 2012, -140, -1000, -112,
	-1000, -1000, -1000, -1000, 231, 98, 2390, 91, 7, 1635,
	-141, -95, -1000, -1000, 1635, -1000, -95, -1000, -1000, -1000,
	-1000, -14, 1635, 1635, -1, -14, 20, -1000, 2578, 1964,
	1917, -1000, -1000, -1000, 2343, -1000, 10, -1000, 2578, 1635,
	-113, -1000, 12, -1000, 1635, 2744, 1869, 1822, 498, -1000,
	385, -1000, -1000, -1000, 2296, 202, 1635, -1000, 2578, -7,
	-97, -114, -1000, -1000, -1000, -146, -1000, -1000, 86, -1000,
	-1000, -1000, -1000, 1635, -1000, 1706, -1000, -1000, 2249, -1000,
	-1000, -147, 2202, -1000, -1000, -1000, 1635, -1000, -1000, -1000,
	-1000, 1070, 2578, -1000, 2744, 724, 173, -41, -1000, -1000,
	-1000, 1635, 1706, -1000, 957, 202, -1000, -1000, 1775, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 724, 1706, -115,
	-18, 724, 202, -1000,
}
var yyPgo = [...]int{

	0, 356, 355, 3, 63, 220, 353, 0, 11, 38,
	4, 352, 12, 350, 208, 181, 5, 336, 335, 331,
	330, 14, 8, 329, 30, 18, 328, 327, 326, 324,
	323, 59, 315, 312, 1, 306, 305, 302, 296, 284,
	282, 279, 65, 275, 272, 228, 264, 240, 2, 10,
	7, 13,
}
var yyR1 = [...]int{

	0, 1, 51, 33, 33, 4, 4, 5, 5, 5,
	2, 34, 34, 6, 3, 3, 3, 3, 3, 41,
	41, 40, 40, 36, 36, 35, 35, 42, 42, 43,
	43, 10, 10, 39, 39, 12, 13, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 20, 20, 20, 20, 14, 14, 15, 15, 18,
	18, 38, 38, 38, 44, 44, 24, 24, 24, 19,
	19, 19, 19, 19, 19, 19, 19, 19, 19, 19,
	19, 19, 19, 19, 19, 25, 25, 25, 7, 7,
	8, 8, 23, 26, 26, 26, 27, 27, 27, 28,
	28, 28, 28, 28, 28, 9, 9, 9, 16, 16,
	16, 29, 29, 30, 30, 30, 30, 30, 30, 21,
	21, 21, 22, 22, 22, 48, 49, 49, 47, 47,
	46, 46, 46, 46, 46, 46, 45, 45, 45, 45,
	31, 31, 31, 31, 31, 31, 31, 32, 32, 32,
	32, 17, 17, 17, 17, 17, 17, 17, 37, 37,
	50,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 0, 1, 3, 1, 3, 2,
	1, 2, 0, 1, 3, 1, 1, 3, 2, 5,
	6, 1, 3, 6, 7, 3, 6, 2, 3, 1,
	3, 1, 2, 3, 1, 1, 3, 6, 5, 3,
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

	-1000, -1, -33, -2, -3, 158, -40, -35, 82, -7,
	-41, -36, -9, -11, 101, -28, -29, -26, 140, 69,
	71, 58, 59, 49, 50, 55, 57, 161, -13, -17,
	60, 61, 62, 63, 64, 65, 66, 100, 67, -19,
	164, 13, 14, 16, -16, -25, -20, -14, -23, -24,
	70, 131, 132, 4, 5, 6, 7, 8, 83, 84,
	92, 93, 94, 96, 97, 98, 99, 95, 149, 165,
	86, 166, -5, -27, 76, 141, 89, -4, 154, 155,
	85, -34, -39, -12, -7, 160, 33, 34, 10, 12,
	11, 35, 37, 36, 51, 49, 50, 52, 68, 53,
	54, 47, 48, 40, 41, 38, 39, 43, 44, 45,
	46, 42, 56, 30, 32, 73, 74, 73, 75, 74,
	17, 18, 19, 20, 29, 21, 22, 23, 24, 25,
	26, 27, 28, 58, 59, 161, 69, 158, 139, 161,
	-48, -47, -49, -46, -7, 37, 140, -7, -9, -25,
	161, -24, 69, -9, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -18, 161, -7, -38,
	88, -45, -31, 86, 151, 152, -7, -7, -7, 69,
	153, 153, -15, -14, -30, -5, -16, 161, 161, -7,
	-7, 161, -7, -7, 88, 150, -45, -45, 88, 158,
	-16, -42, 161, -42, 161, 155, 155, -4, -6, 159,
	-3, 160, 9, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -15,
	-7, 31, -7, 161, -3, 161, 160, 31, -7, 37,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -8, -7, -7, -22, 85, 158, -16,
	-48, 163, 9, 15, -9, 161, -7, -48, 162, -8,
	164, -31, -31, 88, 69, 139, -7, 90, -9, 15,
	-8, -21, -51, -16, 158, 85, -21, -51, -16, -44,
	-42, 153, 69, 158, 139, 153, -37, -50, -7, -7,
	-7, 150, 150, 165, -7, 162, -43, -10, -7, 156,
	-48, 85, -4, -12, 31, -7, -7, -7, -34, -9,
	162, 163, 159, -42, -7, 162, 17, -49, -7, 37,
	140, -48, 162, 163, 162, -32, 85, 91, 50, 86,
	85, 159, 159, 69, 159, -7, 163, -42, -7, -42,
	-16, -8, -7, -22, -16, 162, 9, 162, 162, 159,
	162, 9, -7, 162, -7, 162, 162, 75, -3, 31,
	159, 17, -7, -9, 161, 162, 163, 91, -7, 159,
	163, 159, -50, -10, -3, 31, 160, -34, -7, -48,
	163, -34, 162, 159,
}
var yyDef = [...]int{

	4, -2, 1, 3, 10, 12, 15, 16, 0, 0,
	21, 0, -2, 149, 0, -2, 166, 152, 0, 186,
	0, 0, 0, 0, 0, 0, 0, 0, 90, 94,
	0, 0, 0, 0, 0, 0, 0, 119, 0, 104,
	121, 0, 107, 0, 159, 144, 164, 0, 0, -2,
	0, 0, 0, 0, 0, 0, 0, 0, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 0, 0,
	168, 0, -2, 0, 115, 0, 128, 7, 0, 0,
	5, 0, 0, 34, 35, 18, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 54, 56, 0, 150, 0, 0, 186,
	0, 185, 189, 187, 191, 0, 0, 41, -2, 0,
	0, -2, 186, -2, 75, 76, 77, 78, 0, 95,
	96, 97, 98, 99, 100, 101, 102, 150, 103, 0,
	122, 123, 198, 200, 0, 0, 106, 108, 110, 150,
	0, 0, 124, 117, 118, 116, 173, 0, 0, 213,
	214, 0, 216, 217, 0, 140, 0, 0, 0, 0,
	170, 111, 0, 114, 186, 0, 0, 9, 11, 14,
	13, 17, 0, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 71, 72, 73, 74,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, 88,
	0, 0, 93, 0, 22, 0, 25, 12, 39, 0,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 0, 0, 151, 0, 167, 182, 0, 184,
	0, 127, 186, 0, -2, 186, 0, 0, -2, 0,
	105, 199, 196, 197, 0, 0, 0, 0, 153, 0,
	0, 0, -2, -2, 0, 2, 0, -2, -2, 36,
	125, 0, 150, 0, 0, 0, 0, 218, 220, 0,
	0, 139, 142, 141, 0, 27, 0, 29, 31, 0,
	0, 6, 8, 33, 0, 92, 0, 0, 0, -2,
	0, 160, 162, 163, 0, 0, 0, 188, 190, 0,
	0, 0, -2, 127, 120, 0, 207, 208, 0, 210,
	202, 203, 204, 0, 206, 109, 161, 112, 0, 113,
	177, 0, 0, 176, 178, 211, 0, 212, 215, 169,
	28, 0, 32, 126, 91, 0, 0, 0, 19, 12,
	183, 0, 38, -2, 186, 195, 201, 209, 0, 180,
	174, 175, 219, 30, 20, 12, 26, 23, 37, 0,
	0, 24, 194, 205,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 165, 3, 166, 54, 37, 3,
	161, 162, 52, 49, 9, 50, 51, 53, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 31, 160,
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:420
		{
			yyVAL.stmt = ast.NewEchoStatement(yyDollar[1].tok, yyDollar[2].exprs)
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:422
		{
			yyVAL.stmt = ast.NewExpressionStatement(yyDollar[1].expr)
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:582
		{
			yyVAL.stmt = ast.NewIfStatement(yyDollar[1].tok, yyDollar[3].expr, yyDollar[5].stmt, nil)
		}
	case 20:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:584
		{
			yyVAL.stmt = ast.NewIfStatement(yyDollar[2].tok, yyDollar[4].expr, yyDollar[6].stmt, yyVAL.stmt)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:588
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:590
		{
			yyVAL.stmt = ast.NewIfStatement(yyDollar[2].tok, nil, yyDollar[3].stmt, yyVAL.stmt)
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:596
		{
			yyVAL.stmt = ast.NewAltIfStatement(yyDollar[1].tok, yyDollar[3].expr, yyDollar[6].stmts, nil)
		}
	case 24:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:598
		{
			yyVAL.stmt = ast.NewAltIfStatement(yyDollar[2].tok, yyDollar[4].expr, yyDollar[7].stmts, yyVAL.stmt)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:603
		{
			stmt := yyDollar[1].stmt
			yyVAL.stmt = ast.NewAltIfStatement(yyDollar[2].tok, nil, nil, stmt)
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:608
		{
			stmt := ast.NewAltIfStatement(yyDollar[2].tok, nil, yyDollar[4].stmts, yyVAL.stmt)
			yyVAL.stmt = ast.NewAltIfStatement(yyDollar[5].tok, nil, nil, stmt)
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:659
		{
			yyVAL.expr = ast.NewArgumentListExpression()
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:660
		{
			yyVAL.expr = ast.NewArgumentListExpression(yyDollar[2].exprs...)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:665
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:667
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:671
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:672
		{
			yyVAL.expr = ast.NewArgumentExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:827
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:828
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:832
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:859
		{
			yyVAL.expr = ast.NewNewExpression(yyDollar[1].tok, yyDollar[2].expr, yyDollar[3].expr)
		}
	case 37:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:866
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...), yyDollar[6].expr, false)
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:868
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...), yyDollar[5].expr, false)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:870
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:872
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, yyDollar[1].expr, yyDollar[4].expr, true)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:873
		{
			yyVAL.expr = ast.NewCloneExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:875
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PlusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:877
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MinusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:879
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MulEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:881
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PowEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:883
		{
			yyVAL.expr = ast.NewAssignExpression(ast.DivEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:885
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ConcatEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:887
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ModEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:889
		{
			yyVAL.expr = ast.NewAssignExpression(ast.AndEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:891
		{
			yyVAL.expr = ast.NewAssignExpression(ast.QrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:893
		{
			yyVAL.expr = ast.NewAssignExpression(ast.XorEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:895
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SlEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:897
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:898
		{
			yyVAL.expr = ast.NewIncrementExpression(ast.PostInc, yyDollar[1].expr)
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:899
		{
			yyVAL.expr = ast.NewIncrementExpression(ast.PreInc, yyDollar[2].expr)
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:900
		{
			yyVAL.expr = ast.NewDecrementExpression(ast.PostDec, yyDollar[1].expr)
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:901
		{
			yyVAL.expr = ast.NewDecrementExpression(ast.PreDec, yyDollar[2].expr)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:903
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BooleanOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:905
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BooleanAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:907
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:909
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:911
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalXor, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:912
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:913
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:914
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwXor, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:915
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Concat, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:916
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Add, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:917
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sub, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:918
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Mul, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:919
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Pow, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:920
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Div, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:921
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Mod, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:922
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sl, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:923
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:924
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.UnaryPlus, yyDollar[2].expr)
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:925
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.UnaryMinus, yyDollar[2].expr)
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:926
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.BoolNot, yyDollar[2].expr)
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:927
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.BwNot, yyDollar[2].expr)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:929
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsIdentical, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:931
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsNotIdentical, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:933
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:935
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsNotEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:937
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Smaller, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:939
		{
			yyVAL.expr = ast.NewInfixExpression(ast.SmallerOrEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:941
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Greater, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:943
		{
			yyVAL.expr = ast.NewInfixExpression(ast.GreaterOrEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:945
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Spaceship, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:947
		{
			yyVAL.expr = ast.NewInfixExpression(ast.InstanceOf, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:948
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:949
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:951
		{
			yyVAL.expr = ast.NewTernaryOperatorExpression(yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr)
		}
	case 92:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:953
		{
			yyVAL.expr = ast.NewTernaryOperatorExpression(yyDollar[1].expr, nil, yyDollar[4].expr)
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:955
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Coalesce, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:956
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:957
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:958
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:959
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:960
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:961
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:962
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:963
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:964
		{
			yyVAL.expr = ast.NewExitExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:965
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.Silence, yyDollar[2].expr)
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:966
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:967
		{
			yyVAL.expr = ast.NewBackticksExpression(yyDollar[2].exprs...)
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:968
		{
			yyVAL.expr = ast.NewPrintExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:969
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, nil)
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:970
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:971
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, ast.NewArrayPairExpression(yyDollar[2].expr, yyDollar[4].expr, false))
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:972
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1020
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1022
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1024
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1026
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1030
		{
			yyVAL.expr = ast.NewStaticLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1031
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1035
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1036
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 119:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1040
		{
			yyVAL.expr = nil
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1041
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 121:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1046
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1047
		{
			yyVAL.exprs = []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)}
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1048
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 124:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1052
		{
			yyVAL.expr = nil
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1053
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1057
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Long, yyDollar[3].exprs...)
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1058
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...)
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1059
		{
			yyVAL.expr = ast.NewConstantEncapsedStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1063
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1064
		{
			yyVAL.expr = ast.NewDoubleLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1065
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1066
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1067
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1068
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1069
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1070
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1071
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1072
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1073
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1074
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[2].tok)
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1075
		{
			yyVAL.expr = ast.NewEncapsListExpression(yyDollar[2].exprs...)
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1076
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, yyDollar[2].exprs...)
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1077
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1078
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1082
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1084
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1086
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1090
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1091
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 150:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1095
		{
			yyVAL.expr = nil
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1096
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1100
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1104
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1105
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1106
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1110
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1111
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1112
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1117
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 160:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1119
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 161:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1121
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 162:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1123
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Curly, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 163:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1125
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Prop, yyDollar[1].expr, []ast.Expression{yyDollar[3].expr, yyDollar[4].expr}...)
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1126
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1131
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1133
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1135
		{
			yyVAL.expr = ast.NewVariableExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1139
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 169:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1140
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.CurlyOpen, yyDollar[3].expr)
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1141
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.Var, yyDollar[2].expr)
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1146
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1148
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1153
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 174:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1155
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 175:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1157
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Curly, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1159
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Prop, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1161
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1163
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1167
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1168
		{
			yyVAL.expr = ast.NewMemberNameExpression(yyDollar[2].expr)
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1169
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1173
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1174
		{
			yyVAL.expr = ast.NewPropertyNameExpression(yyDollar[2].expr)
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1175
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1180
		{ /* allow single trailing comma */
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1184
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1185
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1190
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].exprs...)
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1192
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1197
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[3].expr, false)}
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1199
		{
			yyVAL.exprs = []ast.Expression{yyDollar[1].expr}
		}
	case 192:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1201
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[4].expr, true)}
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1203
		{
			yyVAL.exprs = []ast.Expression{ast.NewAmpersandLiteral(yyDollar[2].expr)}
		}
	case 194:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1205
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, ast.NewListExpression(yyDollar[3].tok, yyDollar[5].exprs...), false)}
		}
	case 195:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1207
		{
			yyVAL.exprs = []ast.Expression{ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...)}
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1212
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1214
		{
			yyVAL.exprs = append(yyDollar[1].exprs, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1216
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1218
		{
			yyVAL.exprs = append(yyVAL.exprs, []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal), yyDollar[2].expr}...)
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1223
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 201:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1225
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Dim, yyDollar[3].expr)
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1227
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Prop, ast.NewStringLiteral(yyDollar[3].tok, yyDollar[3].tok.Literal))
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1229
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, yyDollar[2].expr)
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1231
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 205:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1233
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DimInDollarOpenCurlyBraces, []ast.Expression{ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal), yyDollar[4].expr}...)
		}
	case 206:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1234
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.CurlyOpen, yyDollar[2].expr)
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1238
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1239
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1240
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[2].tok, "-"+yyDollar[2].tok.Literal)
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1241
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 211:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1246
		{
			yyVAL.expr = ast.NewIssetExpression(yyDollar[1].tok, yyDollar[3].exprs...)
		}
	case 212:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1247
		{
			yyVAL.expr = ast.NewEmptyExpression(yyDollar[1].tok, yyDollar[3].expr)
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1249
		{
			yyVAL.expr = ast.NewIncludeExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1251
		{
			yyVAL.expr = ast.NewIncludeExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 215:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1253
		{
			yyVAL.expr = ast.NewEvalExpression(yyDollar[1].tok, yyDollar[3].expr)
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1255
		{
			yyVAL.expr = ast.NewRequireExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1257
		{
			yyVAL.expr = ast.NewRequireExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1261
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 219:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1263
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1267
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
