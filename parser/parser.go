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
	"'`'",
	"'\"'",
	"'{'",
	"'}'",
	"'$'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:1267

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
	69, 134,
	139, 134,
	153, 134,
	164, 134,
	-2, 129,
	-1, 11,
	159, 137,
	-2, 146,
	-1, 43,
	69, 136,
	139, 136,
	153, 136,
	159, 139,
	164, 136,
	-2, 124,
	-1, 61,
	153, 97,
	-2, 126,
	-1, 128,
	69, 134,
	139, 134,
	153, 134,
	164, 134,
	-2, 40,
	-1, 131,
	159, 139,
	-2, 136,
	-1, 133,
	69, 134,
	139, 134,
	153, 134,
	164, 134,
	-2, 42,
	-1, 199,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 64,
	-1, 200,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 65,
	-1, 201,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 66,
	-1, 202,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 67,
	-1, 203,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 68,
	-1, 204,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 69,
	-1, 205,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 70,
	-1, 206,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 71,
	-1, 207,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 72,
	-1, 237,
	9, 174,
	160, 174,
	161, 174,
	-2, 134,
	-1, 241,
	69, 135,
	139, 135,
	153, 135,
	159, 138,
	164, 135,
	-2, 74,
	-1, 254,
	159, 160,
	-2, 127,
	-1, 255,
	159, 162,
	-2, 152,
	-1, 259,
	159, 160,
	-2, 128,
	-1, 260,
	159, 162,
	-2, 153,
	-1, 286,
	69, 134,
	139, 134,
	153, 134,
	164, 134,
	-2, 25,
	-1, 298,
	159, 138,
	-2, 135,
	-1, 332,
	9, 173,
	160, 173,
	161, 173,
	-2, 134,
}

const yyPrivate = 57344

const yyLast = 1998

var yyAct = [...]int{

	124, 120, 278, 269, 229, 122, 8, 39, 174, 347,
	310, 243, 339, 335, 325, 311, 299, 127, 287, 59,
	134, 135, 136, 137, 138, 129, 129, 139, 140, 141,
	142, 143, 144, 145, 43, 148, 9, 226, 156, 234,
	321, 309, 152, 346, 230, 59, 66, 334, 327, 63,
	132, 300, 131, 131, 128, 133, 291, 63, 69, 59,
	254, 333, 257, 59, 59, 253, 69, 59, 175, 238,
	65, 176, 8, 182, 183, 184, 185, 186, 187, 188,
	189, 190, 191, 192, 193, 194, 195, 196, 197, 198,
	199, 200, 201, 202, 203, 204, 205, 206, 207, 60,
	209, 211, 212, 214, 215, 216, 217, 218, 219, 220,
	221, 222, 223, 224, 225, 180, 38, 227, 228, 177,
	166, 233, 64, 231, 165, 60, 147, 67, 68, 116,
	119, 239, 153, 129, 240, 67, 68, 308, 264, 60,
	130, 256, 172, 60, 60, 71, 178, 60, 227, 179,
	263, 153, 159, 246, 158, 249, 41, 247, 227, 61,
	131, 164, 237, 129, 153, 324, 270, 271, 336, 262,
	306, 160, 69, 275, 282, 5, 279, 173, 153, 281,
	246, 153, 157, 167, 7, 242, 153, 84, 246, 330,
	131, 320, 251, 244, 245, 252, 272, 154, 155, 118,
	292, 161, 151, 153, 163, 171, 235, 153, 266, 150,
	244, 285, 245, 245, 244, 164, 154, 155, 304, 98,
	259, 129, 267, 121, 117, 258, 283, 248, 274, 154,
	155, 84, 290, 265, 123, 232, 261, 294, 289, 277,
	297, 293, 273, 154, 155, 168, 154, 155, 131, 6,
	286, 154, 155, 302, 305, 161, 181, 313, 163, 303,
	169, 170, 312, 149, 268, 227, 317, 314, 154, 155,
	208, 318, 154, 155, 2, 255, 260, 301, 162, 12,
	11, 326, 83, 85, 86, 328, 98, 62, 13, 42,
	40, 35, 146, 331, 70, 25, 24, 10, 84, 4,
	3, 1, 316, 129, 0, 0, 37, 0, 0, 0,
	337, 0, 87, 88, 81, 82, 80, 83, 85, 86,
	0, 98, 270, 0, 0, 341, 279, 0, 342, 0,
	131, 343, 332, 84, 0, 344, 0, 0, 0, 0,
	0, 0, 19, 20, 0, 0, 0, 0, 21, 0,
	22, 17, 18, 26, 27, 28, 29, 30, 31, 32,
	34, 0, 15, 44, 16, 0, 0, 0, 0, 63,
	0, 0, 0, 0, 0, 0, 47, 48, 69, 59,
	315, 0, 65, 232, 319, 49, 50, 51, 56, 52,
	53, 54, 55, 33, 99, 0, 100, 72, 73, 77,
	79, 78, 91, 92, 89, 90, 97, 93, 94, 95,
	96, 87, 88, 81, 82, 80, 83, 85, 86, 0,
	98, 37, 0, 0, 45, 46, 0, 0, 0, 0,
	0, 0, 84, 14, 64, 0, 0, 0, 0, 0,
	0, 0, 57, 0, 0, 125, 0, 67, 68, 280,
	0, 0, 23, 276, 0, 36, 58, 19, 20, 60,
	0, 0, 0, 21, 0, 22, 17, 18, 26, 27,
	28, 29, 30, 31, 32, 34, 0, 15, 44, 16,
	0, 0, 0, 0, 63, 0, 0, 0, 0, 0,
	0, 47, 48, 69, 59, 0, 0, 65, 0, 0,
	49, 50, 51, 56, 52, 53, 54, 55, 33, 100,
	72, 73, 77, 79, 78, 91, 92, 89, 90, 97,
	93, 94, 95, 96, 87, 88, 81, 82, 80, 83,
	85, 86, 0, 98, 0, 0, 37, 0, 0, 45,
	46, 0, 0, 0, 0, 84, 0, 0, 126, 64,
	0, 0, 0, 0, 0, 0, 0, 57, 0, 0,
	0, 0, 67, 68, 0, 0, 0, 23, 0, 0,
	36, 58, 19, 20, 60, 0, 0, 0, 21, 0,
	22, 17, 18, 26, 27, 28, 29, 30, 31, 32,
	34, 0, 15, 44, 16, 0, 0, 0, 0, 63,
	0, 0, 0, 0, 0, 0, 47, 48, 69, 59,
	0, 0, 65, 0, 0, 49, 50, 51, 56, 52,
	53, 54, 55, 33, 73, 77, 79, 78, 91, 92,
	89, 90, 97, 93, 94, 95, 96, 87, 88, 81,
	82, 80, 83, 85, 86, 0, 98, 0, 0, 0,
	0, 37, 0, 0, 45, 46, 0, 0, 84, 0,
	0, 0, 0, 14, 64, 0, 0, 0, 0, 0,
	0, 0, 57, 0, 0, 295, 0, 67, 68, 280,
	0, 0, 23, 0, 0, 36, 58, 19, 20, 60,
	0, 0, 0, 21, 0, 22, 17, 18, 26, 27,
	28, 29, 30, 31, 32, 34, 0, 15, 44, 16,
	0, 0, 0, 0, 63, 0, 0, 0, 0, 0,
	0, 47, 48, 69, 59, 0, 0, 65, 0, 0,
	49, 50, 51, 56, 52, 53, 54, 55, 33, 77,
	79, 78, 91, 92, 89, 90, 97, 93, 94, 95,
	96, 87, 88, 81, 82, 80, 83, 85, 86, 0,
	98, 0, 0, 0, 0, 0, 37, 0, 0, 45,
	46, 0, 84, 0, 0, 0, 0, 0, 296, 64,
	0, 0, 0, 0, 0, 0, 0, 57, 0, 0,
	0, 0, 67, 68, 0, 0, 0, 23, 0, 0,
	36, 58, 19, 20, 60, 0, 0, 0, 21, 0,
	22, 17, 18, 26, 27, 28, 29, 30, 31, 32,
	34, 0, 15, 44, 16, 0, 0, 0, 0, 63,
	0, 0, 0, 0, 0, 0, 47, 48, 69, 59,
	0, 0, 65, 250, 0, 49, 50, 51, 56, 52,
	53, 54, 55, 33, 79, 78, 91, 92, 89, 90,
	97, 93, 94, 95, 96, 87, 88, 81, 82, 80,
	83, 85, 86, 0, 98, 0, 0, 0, 0, 0,
	0, 37, 0, 0, 45, 46, 84, 0, 0, 0,
	0, 0, 0, 14, 64, 0, 0, 0, 0, 0,
	0, 0, 57, 0, 0, 213, 0, 67, 68, 0,
	0, 0, 23, 0, 0, 36, 58, 19, 20, 60,
	0, 0, 0, 21, 0, 22, 17, 18, 26, 27,
	28, 29, 30, 31, 32, 34, 0, 15, 44, 16,
	0, 0, 0, 0, 63, 0, 0, 0, 0, 0,
	0, 47, 48, 69, 59, 0, 0, 65, 0, 0,
	49, 50, 51, 56, 52, 53, 54, 55, 33, 0,
	78, 91, 92, 89, 90, 97, 93, 94, 95, 96,
	87, 88, 81, 82, 80, 83, 85, 86, 0, 98,
	0, 0, 0, 0, 0, 0, 37, 0, 0, 45,
	46, 84, 0, 0, 0, 0, 0, 0, 14, 64,
	0, 0, 0, 0, 210, 0, 0, 57, 0, 0,
	0, 0, 67, 68, 0, 0, 0, 23, 0, 0,
	36, 58, 19, 20, 60, 0, 0, 0, 21, 0,
	22, 17, 18, 26, 27, 28, 29, 30, 31, 32,
	34, 0, 15, 44, 16, 0, 0, 0, 0, 63,
	0, 0, 0, 0, 0, 0, 47, 48, 69, 59,
	0, 0, 65, 0, 0, 49, 50, 51, 56, 52,
	53, 54, 55, 33, 0, 0, 91, 92, 89, 90,
	97, 93, 94, 95, 96, 87, 88, 81, 82, 80,
	83, 85, 86, 0, 98, 0, 0, 0, 0, 0,
	0, 37, 0, 0, 45, 46, 84, 0, 0, 0,
	0, 0, 0, 14, 64, 0, 0, 0, 0, 0,
	0, 0, 57, 0, 0, 0, 0, 67, 68, 0,
	0, 0, 23, 0, 0, 36, 58, 19, 20, 60,
	0, 0, 0, 21, 0, 22, 17, 18, 26, 27,
	28, 29, 30, 31, 32, 34, 0, 15, 44, 16,
	0, 0, 0, 0, 63, 0, 0, 0, 0, 0,
	0, 47, 48, 69, 59, 0, 0, 65, 0, 0,
	49, 50, 51, 56, 52, 53, 54, 55, 33, 93,
	94, 95, 96, 87, 88, 81, 82, 80, 83, 85,
	86, 0, 98, 81, 82, 80, 83, 85, 86, 0,
	98, 0, 0, 0, 84, 0, 0, 0, 0, 45,
	46, 0, 84, 0, 0, 0, 0, 0, 14, 64,
	0, 0, 0, 74, 76, 75, 0, 57, 0, 0,
	0, 0, 67, 68, 0, 0, 0, 23, 0, 0,
	36, 58, 0, 99, 60, 100, 72, 73, 77, 79,
	78, 91, 92, 89, 90, 97, 93, 94, 95, 96,
	87, 88, 81, 82, 80, 83, 85, 86, 0, 98,
	74, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 0, 0, 0, 0, 0,
	99, 0, 100, 72, 73, 77, 79, 78, 91, 92,
	89, 90, 97, 93, 94, 95, 96, 87, 88, 81,
	82, 80, 83, 85, 86, 0, 98, 74, 76, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 84, 0,
	0, 0, 0, 0, 0, 0, 0, 99, 0, 100,
	72, 73, 77, 79, 78, 91, 92, 89, 90, 97,
	93, 94, 95, 96, 87, 88, 81, 82, 80, 83,
	85, 86, 0, 98, 74, 76, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 340, 0,
	0, 0, 0, 0, 99, 0, 100, 72, 73, 77,
	79, 78, 91, 92, 89, 90, 97, 93, 94, 95,
	96, 87, 88, 81, 82, 80, 83, 85, 86, 0,
	98, 74, 76, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 84, 0, 0, 338, 0, 0, 0, 0,
	0, 99, 0, 100, 72, 73, 77, 79, 78, 91,
	92, 89, 90, 97, 93, 94, 95, 96, 87, 88,
	81, 82, 80, 83, 85, 86, 0, 98, 74, 76,
	75, 0, 0, 0, 0, 0, 0, 0, 0, 84,
	0, 0, 329, 0, 0, 0, 0, 0, 99, 0,
	100, 72, 73, 77, 79, 78, 91, 92, 89, 90,
	97, 93, 94, 95, 96, 87, 88, 81, 82, 80,
	83, 85, 86, 0, 98, 74, 76, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 84, 0, 0, 323,
	0, 0, 0, 0, 0, 99, 0, 100, 72, 73,
	77, 79, 78, 91, 92, 89, 90, 97, 93, 94,
	95, 96, 87, 88, 81, 82, 80, 83, 85, 86,
	0, 98, 74, 76, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 84, 0, 0, 307, 0, 0, 0,
	0, 0, 99, 0, 100, 72, 73, 77, 79, 78,
	91, 92, 89, 90, 97, 93, 94, 95, 96, 87,
	88, 81, 82, 80, 83, 85, 86, 0, 98, 74,
	76, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	84, 0, 0, 288, 0, 0, 0, 0, 0, 99,
	0, 100, 72, 73, 77, 79, 78, 91, 92, 89,
	90, 97, 93, 94, 95, 96, 87, 88, 81, 82,
	80, 83, 85, 86, 0, 98, 74, 76, 75, 0,
	0, 0, 0, 0, 0, 0, 345, 84, 0, 0,
	0, 0, 0, 0, 0, 0, 99, 0, 100, 72,
	73, 77, 79, 78, 91, 92, 89, 90, 97, 93,
	94, 95, 96, 87, 88, 81, 82, 80, 83, 85,
	86, 0, 98, 74, 76, 75, 0, 0, 0, 0,
	0, 0, 322, 0, 84, 0, 0, 0, 0, 0,
	0, 0, 0, 99, 284, 100, 72, 73, 77, 79,
	78, 91, 92, 89, 90, 97, 93, 94, 95, 96,
	87, 88, 81, 82, 80, 83, 85, 86, 0, 98,
	74, 76, 75, 0, 0, 236, 0, 0, 0, 298,
	0, 84, 0, 0, 0, 0, 0, 0, 0, 0,
	99, 0, 100, 72, 73, 77, 79, 78, 91, 92,
	89, 90, 97, 93, 94, 95, 96, 87, 88, 81,
	82, 80, 83, 85, 86, 0, 98, 74, 76, 75,
	0, 0, 0, 0, 0, 0, 241, 0, 84, 0,
	0, 0, 0, 0, 0, 0, 0, 99, 0, 100,
	72, 73, 77, 79, 78, 91, 92, 89, 90, 97,
	93, 94, 95, 96, 87, 88, 81, 82, 80, 83,
	85, 86, 0, 98, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	0, 0, 0, 99, 0, 100, 72, 73, 77, 79,
	78, 91, 92, 89, 90, 97, 93, 94, 95, 96,
	87, 88, 81, 82, 80, 83, 85, 86, 75, 98,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 0, 99, 0, 100, 72,
	73, 77, 79, 78, 91, 92, 89, 90, 97, 93,
	94, 95, 96, 87, 88, 81, 82, 80, 83, 85,
	86, 0, 98, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 84, 101, 102, 103, 104, 106,
	107, 108, 109, 110, 111, 112, 113, 105, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 114, 115,
}
var yyPact = [...]int{

	-1000, -1000, 93, -1000, -1000, 1098, 136, -1000, 1797, 1938,
	-1000, -1000, -1000, 60, -29, 408, 1098, -19, -19, 1098,
	1098, 1098, 1098, 1098, -1000, -1000, 1098, 1098, 1098, 1098,
	1098, 1098, 1098, -33, 1098, -1000, 121, 1098, -1000, 113,
	-1000, 1, -1, -1000, -27, -35, -39, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 95, 117, -1000,
	-22, -91, -91, -1000, -40, -1000, -9, -6, 87, -1000,
	-1000, 1098, 1098, 1098, 1098, 1098, 1098, 1098, 1098, 1098,
	1098, 1098, 1098, 1098, 1098, 1098, 1098, 1098, 1098, 1098,
	1098, 1098, 1098, 1098, 1098, 1098, 1098, 1098, -27, 983,
	1098, 868, 1098, 1098, 1098, 1098, 1098, 1098, 1098, 1098,
	1098, 1098, 1098, 1098, -1000, -1000, 1098, 1098, -41, 408,
	-122, 197, -1000, -1000, 1750, -19, -90, -1000, -1000, 113,
	1098, -1000, 408, -1000, 119, 119, 163, 119, 1656, 119,
	119, 119, 119, 119, 119, 119, -1000, 1098, 119, -151,
	78, 100, -1000, 88, 753, -19, 364, 1098, -23, -23,
	-91, -3, 69, -1000, -1000, 1098, 1098, 46, -1000, 92,
	65, 78, 1098, -1000, -1000, 293, -1000, 408, 89, 87,
	-9, -1000, 590, 704, 1843, 364, 1886, 818, 1048, 933,
	230, 230, 230, 163, 119, 163, 163, 1164, 1164, 1156,
	1156, 1156, 1156, 265, 265, 265, 265, 1156, -1000, 1703,
	1098, 477, 364, -19, 364, 364, 364, 364, 364, 364,
	364, 364, 364, 364, 364, 364, -143, 1797, 1468, -91,
	-1000, 1098, -1000, -104, 183, 408, 638, -1000, 408, 1609,
	-145, -1000, -109, -1000, -1000, -1000, -1000, 168, 85, 1421,
	-28, -155, -146, -91, -1000, -1000, 1098, -1000, -91, -1000,
	-1000, -1000, -1000, -67, 1098, 1098, -41, -67, 31, -1000,
	1797, 1562, -1000, -1000, -1000, 1374, -1000, 5, -1000, 1797,
	1098, -112, -1000, -9, 1098, 477, -1000, -1000, -1000, -1000,
	1327, 172, 1098, -1000, 1797, -19, -98, -113, -1000, -1000,
	-1000, -148, -1000, -1000, 77, -1000, -1000, -1000, -1000, 1098,
	-1000, -1000, -1000, 1280, -1000, -1000, -149, 1233, -1000, -1000,
	-1000, 1098, -1000, -1000, -1000, 523, 1797, -1000, 477, -1000,
	1098, 364, -1000, 408, 172, -1000, -1000, 1515, -1000, -1000,
	-1000, -1000, -1000, 364, -117, -156, 172, -1000,
}
var yyPgo = [...]int{

	0, 301, 300, 299, 46, 159, 0, 37, 36, 2,
	297, 184, 296, 156, 171, 116, 295, 292, 291, 290,
	65, 4, 289, 34, 7, 288, 287, 280, 279, 278,
	42, 277, 274, 264, 263, 249, 8, 239, 236, 202,
	234, 223, 1, 5, 3, 60,
}
var yyR1 = [...]int{

	0, 1, 45, 32, 32, 4, 4, 5, 5, 5,
	2, 3, 36, 36, 37, 37, 9, 9, 35, 35,
	11, 12, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 19, 19, 19, 19, 13, 13, 14, 14,
	17, 17, 34, 34, 34, 38, 38, 23, 23, 23,
	18, 18, 18, 18, 18, 18, 18, 18, 18, 18,
	18, 18, 18, 18, 18, 18, 24, 24, 24, 6,
	6, 7, 7, 22, 25, 25, 25, 26, 26, 26,
	27, 27, 27, 27, 27, 27, 8, 8, 8, 15,
	15, 15, 28, 28, 29, 29, 29, 29, 29, 29,
	20, 20, 20, 21, 21, 21, 42, 43, 43, 41,
	41, 40, 40, 40, 40, 40, 40, 39, 39, 39,
	39, 30, 30, 30, 30, 30, 30, 30, 31, 31,
	31, 31, 16, 16, 33, 33, 44,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 0, 1, 3, 1, 3, 2,
	1, 3, 2, 3, 1, 3, 1, 2, 3, 1,
	1, 3, 6, 5, 3, 4, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 1, 5, 4, 3, 1,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 1,
	3, 2, 2, 4, 4, 2, 1, 1, 1, 1,
	0, 3, 0, 1, 1, 0, 1, 4, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 2, 3, 3, 1, 1, 1, 3, 3, 1,
	1, 0, 1, 1, 1, 3, 1, 1, 3, 1,
	1, 4, 4, 4, 4, 1, 1, 1, 3, 1,
	4, 2, 3, 3, 1, 4, 4, 3, 3, 3,
	1, 3, 1, 1, 3, 1, 1, 0, 1, 3,
	1, 3, 1, 4, 2, 6, 4, 2, 2, 1,
	2, 1, 4, 3, 3, 3, 6, 3, 1, 1,
	2, 1, 4, 4, 1, 3, 1,
}
var yyChk = [...]int{

	-1000, -1, -32, -2, -3, 82, -35, -11, -6, -8,
	-10, -27, -28, -25, 140, 69, 71, 58, 59, 49,
	50, 55, 57, 159, -12, -16, 60, 61, 62, 63,
	64, 65, 66, 100, 67, -18, 162, 13, -15, -24,
	-19, -13, -22, -23, 70, 131, 132, 83, 84, 92,
	93, 94, 96, 97, 98, 99, 95, 149, 163, 86,
	166, -5, -26, 76, 141, 89, -4, 154, 155, 85,
	158, 9, 33, 34, 10, 12, 11, 35, 37, 36,
	51, 49, 50, 52, 68, 53, 54, 47, 48, 40,
	41, 38, 39, 43, 44, 45, 46, 42, 56, 30,
	32, 17, 18, 19, 20, 29, 21, 22, 23, 24,
	25, 26, 27, 28, 58, 59, 69, 164, 139, 159,
	-42, -41, -43, -40, -6, 37, 140, -6, -8, -24,
	159, -23, 69, -8, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -17, 159, -6, -34,
	88, -39, -30, 86, 151, 152, -6, 69, 153, 153,
	-14, -13, -29, -5, -15, 159, 159, 88, 150, -39,
	-39, 88, 164, -15, -36, 159, -36, 159, 155, 155,
	-4, -11, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -14, -6,
	31, -6, -6, 37, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -7, -6, -6, -21,
	85, 164, -15, -42, 161, 9, 15, -8, 159, -6,
	-42, 160, -7, 162, -30, -30, 88, 69, 139, -6,
	90, -8, -7, -20, -45, -15, 164, 85, -20, -45,
	-15, -38, -36, 153, 69, 164, 139, 153, -33, -44,
	-6, -6, 150, 150, 163, -6, 160, -37, -9, -6,
	156, -42, 85, -4, 31, -6, -8, 161, 165, -36,
	-6, 160, 17, -43, -6, 37, 140, -42, 160, 161,
	160, -31, 85, 91, 50, 86, 85, 165, 165, 69,
	165, 161, -36, -6, -36, -15, -7, -6, -21, -15,
	160, 9, 160, 165, 160, 9, -6, 160, -6, 165,
	17, -6, -8, 159, 160, 161, 91, -6, 165, 161,
	165, -44, -9, -6, -42, 161, 160, 165,
}
var yyDef = [...]int{

	4, -2, 1, 3, 10, 0, 0, 19, 20, -2,
	130, -2, 147, 133, 0, 167, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 79, 0, 0, 0, 0,
	0, 0, 0, 100, 0, 89, 102, 0, 140, 125,
	145, 0, 0, -2, 0, 0, 0, 110, 111, 112,
	113, 114, 115, 116, 117, 118, 119, 0, 0, 149,
	0, -2, 0, 96, 0, 109, 7, 0, 0, 5,
	11, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 39, 41, 131, 0, 0, 167,
	0, 166, 170, 168, 172, 0, 0, 26, -2, 0,
	0, -2, 167, -2, 60, 61, 62, 63, 0, 80,
	81, 82, 83, 84, 85, 86, 87, 131, 88, 0,
	103, 104, 179, 181, 0, 0, 91, 131, 0, 0,
	105, 98, 99, 97, 154, 0, 0, 0, 121, 0,
	0, 0, 0, 151, 92, 0, 95, 167, 0, 0,
	9, 18, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 55, 56, 57, 58, 59, -2,
	-2, -2, -2, -2, -2, -2, -2, -2, 73, 0,
	0, 78, 24, 0, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 0, 132, 0, 148,
	163, 0, 165, 0, 108, 167, 0, -2, 167, 0,
	0, -2, 0, 90, 180, 177, 178, 0, 0, 0,
	0, 134, 0, 0, -2, -2, 0, 2, 0, -2,
	-2, 21, 106, 0, 131, 0, 0, 0, 0, 194,
	196, 0, 120, 123, 122, 0, 12, 0, 14, 16,
	0, 0, 6, 8, 0, 77, -2, 141, 143, 144,
	0, 0, 0, 169, 171, 0, 0, 0, -2, 108,
	101, 0, 188, 189, 0, 191, 183, 184, 185, 0,
	187, 142, 93, 0, 94, 158, 0, 0, 157, 159,
	192, 0, 193, 150, 13, 0, 17, 107, 76, 164,
	0, 23, -2, 167, 176, 182, 190, 0, 161, 155,
	156, 195, 15, 22, 0, 0, 175, 186,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 163, 3, 166, 54, 37, 3,
	159, 160, 52, 49, 9, 50, 51, 53, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 31, 158,
	43, 17, 45, 30, 67, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 69, 3, 161, 36, 3, 162, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 164, 35, 165, 57,
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:856
		{
			yyVAL.expr = ast.NewNewExpression(yyDollar[1].tok, yyDollar[2].expr, yyDollar[3].expr)
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:863
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...), yyDollar[6].expr, false)
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:865
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...), yyDollar[5].expr, false)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:867
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:869
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, yyDollar[1].expr, yyDollar[4].expr, true)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:870
		{
			yyVAL.expr = ast.NewCloneExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:872
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PlusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:874
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MinusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:876
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MulEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:878
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PowEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:880
		{
			yyVAL.expr = ast.NewAssignExpression(ast.DivEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:882
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ConcatEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:884
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ModEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:886
		{
			yyVAL.expr = ast.NewAssignExpression(ast.AndEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:888
		{
			yyVAL.expr = ast.NewAssignExpression(ast.QrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:890
		{
			yyVAL.expr = ast.NewAssignExpression(ast.XorEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:892
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SlEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:894
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:895
		{
			yyVAL.expr = ast.NewIncrementExpression(ast.PostInc, yyDollar[1].expr)
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:896
		{
			yyVAL.expr = ast.NewIncrementExpression(ast.PreInc, yyDollar[2].expr)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:897
		{
			yyVAL.expr = ast.NewDecrementExpression(ast.PostDec, yyDollar[1].expr)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:898
		{
			yyVAL.expr = ast.NewDecrementExpression(ast.PreDec, yyDollar[2].expr)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:900
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BooleanOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:902
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BooleanAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:904
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:906
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:908
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalXor, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:909
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:910
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:911
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwXor, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:912
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Concat, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:913
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Add, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:914
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sub, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:915
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Mul, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:916
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Pow, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:917
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Div, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:918
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Mod, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:919
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sl, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:920
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:921
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.UnaryPlus, yyDollar[2].expr)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:922
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.UnaryMinus, yyDollar[2].expr)
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:923
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.BoolNot, yyDollar[2].expr)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:924
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.BwNot, yyDollar[2].expr)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:926
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsIdentical, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:928
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsNotIdentical, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:930
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:932
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsNotEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:934
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Smaller, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:936
		{
			yyVAL.expr = ast.NewInfixExpression(ast.SmallerOrEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:938
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Greater, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:940
		{
			yyVAL.expr = ast.NewInfixExpression(ast.GreaterOrEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:942
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Spaceship, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:944
		{
			yyVAL.expr = ast.NewInfixExpression(ast.InstanceOf, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:945
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:946
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:948
		{
			yyVAL.expr = ast.NewTernaryOperatorExpression(yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr)
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:950
		{
			yyVAL.expr = ast.NewTernaryOperatorExpression(yyDollar[1].expr, nil, yyDollar[4].expr)
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:952
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Coalesce, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:953
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:954
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:955
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:956
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:957
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:958
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:959
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:960
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:961
		{
			yyVAL.expr = ast.NewExitExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:962
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.Silence, yyDollar[2].expr)
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:963
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:964
		{
			yyVAL.expr = ast.NewBackticksExpression(yyDollar[2].exprs...)
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:965
		{
			yyVAL.expr = ast.NewPrintExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1017
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1019
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1021
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1023
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1027
		{
			yyVAL.expr = ast.NewStaticLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1028
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1032
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1033
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1037
		{
			yyVAL.expr = nil
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1038
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1043
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1044
		{
			yyVAL.exprs = []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1045
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1049
		{
			yyVAL.expr = nil
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1050
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1054
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Long, yyDollar[3].exprs...)
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1055
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...)
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1056
		{
			yyVAL.expr = ast.NewConstantEncapsedStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1060
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1061
		{
			yyVAL.expr = ast.NewDoubleLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1062
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1063
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1064
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1065
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1066
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1067
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1068
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1069
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1070
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1071
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[2].tok)
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1072
		{
			yyVAL.expr = ast.NewEncapsListExpression(yyDollar[2].exprs...)
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1073
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, yyDollar[2].exprs...)
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1074
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1075
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1079
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1081
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1083
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1087
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1088
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 131:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1092
		{
			yyVAL.expr = nil
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1093
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1097
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1101
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1102
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1103
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1107
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1108
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1109
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1114
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 141:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1116
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 142:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1118
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 143:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1120
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Curly, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 144:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1122
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Prop, yyDollar[1].expr, []ast.Expression{yyDollar[3].expr, yyDollar[4].expr}...)
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1123
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1128
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1130
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1132
		{
			yyVAL.expr = ast.NewVariableExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1136
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1137
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.CurlyOpen, yyDollar[3].expr)
		}
	case 151:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1138
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.Var, yyDollar[2].expr)
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1143
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1145
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1150
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1152
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1154
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Curly, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1156
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Prop, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1158
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1160
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1164
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1165
		{
			yyVAL.expr = ast.NewMemberNameExpression(yyDollar[2].expr)
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1166
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1170
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1171
		{
			yyVAL.expr = ast.NewPropertyNameExpression(yyDollar[2].expr)
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1172
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1177
		{ /* allow single trailing comma */
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1181
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1182
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1187
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].exprs...)
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1189
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1194
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[3].expr, false)}
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1196
		{
			yyVAL.exprs = []ast.Expression{yyDollar[1].expr}
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1198
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[4].expr, true)}
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1200
		{
			yyVAL.exprs = []ast.Expression{ast.NewAmpersandLiteral(yyDollar[2].expr)}
		}
	case 175:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1202
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, ast.NewListExpression(yyDollar[3].tok, yyDollar[5].exprs...), false)}
		}
	case 176:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1204
		{
			yyVAL.exprs = []ast.Expression{ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...)}
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1209
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1211
		{
			yyVAL.exprs = append(yyDollar[1].exprs, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1213
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1215
		{
			yyVAL.exprs = append(yyVAL.exprs, []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal), yyDollar[2].expr}...)
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1220
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 182:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1222
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Dim, yyDollar[3].expr)
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1224
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Prop, ast.NewStringLiteral(yyDollar[3].tok, yyDollar[3].tok.Literal))
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1226
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, yyDollar[2].expr)
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1228
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 186:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1230
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DimInDollarOpenCurlyBraces, []ast.Expression{ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal), yyDollar[4].expr}...)
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1231
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.CurlyOpen, yyDollar[2].expr)
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1235
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1236
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1237
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[2].tok, "-"+yyDollar[2].tok.Literal)
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1238
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 192:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1243
		{
			yyVAL.expr = ast.NewIssetExpression(yyDollar[1].tok, yyDollar[3].exprs...)
		}
	case 193:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1244
		{
			yyVAL.expr = ast.NewEmptyExpression(yyDollar[1].tok, yyDollar[3].expr)
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1258
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1260
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1264
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
