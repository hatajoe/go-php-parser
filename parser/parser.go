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

//line parser.go.y:1266

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
	-1, 8,
	69, 143,
	139, 143,
	153, 143,
	158, 143,
	-2, 138,
	-1, 10,
	161, 146,
	-2, 155,
	-1, 44,
	69, 145,
	139, 145,
	153, 145,
	158, 145,
	161, 148,
	-2, 133,
	-1, 67,
	153, 106,
	-2, 135,
	-1, 137,
	69, 143,
	139, 143,
	153, 143,
	158, 143,
	-2, 45,
	-1, 140,
	161, 148,
	-2, 145,
	-1, 142,
	69, 143,
	139, 143,
	153, 143,
	158, 143,
	-2, 47,
	-1, 219,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 69,
	-1, 220,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 70,
	-1, 221,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 71,
	-1, 222,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 72,
	-1, 223,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 73,
	-1, 224,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 74,
	-1, 225,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 75,
	-1, 226,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 76,
	-1, 227,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 77,
	-1, 257,
	9, 183,
	162, 183,
	163, 183,
	-2, 143,
	-1, 261,
	69, 144,
	139, 144,
	153, 144,
	158, 144,
	161, 147,
	-2, 79,
	-1, 275,
	161, 169,
	-2, 136,
	-1, 276,
	161, 171,
	-2, 161,
	-1, 280,
	161, 169,
	-2, 137,
	-1, 281,
	161, 171,
	-2, 162,
	-1, 309,
	69, 143,
	139, 143,
	153, 143,
	158, 143,
	-2, 30,
	-1, 321,
	161, 147,
	-2, 144,
	-1, 357,
	9, 182,
	162, 182,
	163, 182,
	-2, 143,
}

const yyPrivate = 57344

const yyLast = 2510

var yyAct = [...]int{

	133, 129, 300, 7, 290, 249, 131, 79, 263, 78,
	246, 275, 161, 274, 364, 360, 136, 350, 345, 143,
	144, 145, 146, 147, 335, 40, 148, 149, 150, 151,
	152, 153, 154, 72, 157, 322, 310, 165, 166, 167,
	44, 254, 138, 138, 371, 8, 359, 352, 65, 178,
	179, 323, 181, 182, 314, 358, 190, 140, 140, 69,
	250, 65, 137, 142, 65, 191, 278, 65, 75, 65,
	258, 201, 193, 180, 177, 176, 156, 7, 39, 128,
	372, 333, 202, 203, 204, 205, 206, 207, 208, 209,
	210, 211, 212, 213, 214, 215, 216, 217, 218, 219,
	220, 221, 222, 223, 224, 225, 226, 227, 196, 229,
	231, 232, 234, 235, 236, 237, 238, 239, 240, 241,
	242, 243, 244, 245, 175, 192, 247, 248, 66, 332,
	253, 194, 195, 251, 285, 141, 188, 73, 74, 277,
	259, 66, 69, 260, 66, 189, 125, 66, 284, 66,
	170, 75, 65, 169, 361, 71, 267, 247, 355, 162,
	138, 266, 42, 67, 269, 329, 75, 262, 171, 247,
	349, 344, 264, 265, 304, 140, 162, 291, 292, 273,
	257, 293, 280, 168, 279, 162, 175, 266, 162, 297,
	138, 4, 301, 93, 255, 303, 264, 315, 265, 265,
	264, 162, 79, 266, 287, 140, 252, 70, 172, 174,
	271, 306, 130, 132, 282, 162, 127, 187, 288, 331,
	73, 74, 200, 286, 163, 164, 268, 139, 283, 305,
	160, 308, 66, 107, 162, 126, 183, 299, 296, 77,
	294, 163, 164, 158, 162, 93, 159, 327, 276, 281,
	163, 164, 313, 163, 164, 289, 76, 317, 2, 138,
	320, 324, 316, 173, 11, 295, 163, 164, 199, 10,
	172, 174, 68, 334, 140, 12, 228, 43, 337, 309,
	163, 164, 325, 328, 41, 34, 247, 341, 326, 155,
	24, 23, 9, 342, 185, 186, 340, 197, 184, 163,
	164, 3, 1, 351, 0, 0, 312, 0, 353, 163,
	164, 0, 0, 0, 0, 0, 356, 0, 0, 0,
	0, 0, 0, 0, 0, 48, 49, 50, 51, 52,
	0, 336, 0, 362, 36, 37, 338, 38, 90, 91,
	89, 92, 94, 95, 138, 107, 291, 0, 0, 0,
	366, 301, 0, 367, 0, 0, 368, 93, 0, 140,
	369, 0, 0, 339, 357, 0, 252, 343, 0, 0,
	18, 19, 0, 0, 0, 0, 20, 0, 21, 16,
	17, 25, 26, 27, 28, 29, 30, 31, 33, 0,
	14, 45, 15, 0, 0, 0, 0, 69, 0, 0,
	0, 0, 0, 6, 53, 54, 75, 65, 0, 0,
	71, 0, 0, 55, 56, 57, 62, 58, 59, 60,
	61, 32, 0, 0, 102, 103, 104, 105, 96, 97,
	90, 91, 89, 92, 94, 95, 0, 107, 48, 49,
	50, 51, 52, 0, 0, 0, 0, 36, 37, 93,
	38, 0, 46, 47, 0, 92, 94, 95, 0, 107,
	0, 13, 70, 0, 0, 0, 0, 0, 0, 0,
	63, 93, 0, 0, 0, 73, 74, 0, 0, 5,
	198, 0, 22, 18, 19, 35, 64, 66, 0, 20,
	0, 21, 16, 17, 25, 26, 27, 28, 29, 30,
	31, 33, 0, 14, 45, 15, 0, 0, 0, 0,
	69, 0, 0, 0, 0, 0, 0, 53, 54, 75,
	65, 0, 0, 71, 0, 0, 55, 56, 57, 62,
	58, 59, 60, 61, 32, 82, 86, 88, 87, 100,
	101, 98, 99, 106, 102, 103, 104, 105, 96, 97,
	90, 91, 89, 92, 94, 95, 0, 107, 48, 49,
	50, 51, 52, 0, 0, 46, 47, 36, 37, 93,
	38, 0, 0, 0, 13, 70, 0, 0, 0, 0,
	0, 0, 0, 63, 0, 0, 0, 0, 73, 74,
	302, 0, 0, 0, 0, 22, 298, 0, 35, 64,
	66, 0, 0, 18, 19, 0, 0, 0, 0, 20,
	0, 21, 16, 17, 25, 26, 27, 28, 29, 30,
	31, 33, 0, 14, 45, 15, 0, 0, 0, 0,
	69, 0, 0, 0, 0, 0, 6, 53, 54, 75,
	65, 0, 0, 71, 0, 0, 55, 56, 57, 62,
	58, 59, 60, 61, 32, 96, 97, 90, 91, 89,
	92, 94, 95, 0, 107, 0, 0, 0, 0, 0,
	0, 48, 49, 50, 51, 52, 93, 0, 0, 0,
	36, 37, 0, 38, 0, 46, 47, 0, 0, 0,
	0, 0, 0, 0, 13, 70, 0, 0, 0, 0,
	0, 0, 0, 63, 134, 0, 0, 0, 73, 74,
	0, 0, 5, 0, 0, 22, 18, 19, 35, 64,
	66, 0, 20, 0, 21, 16, 17, 25, 26, 27,
	28, 29, 30, 31, 33, 0, 14, 45, 15, 0,
	0, 0, 0, 69, 0, 0, 0, 0, 0, 0,
	53, 54, 75, 65, 0, 0, 71, 0, 0, 55,
	56, 57, 62, 58, 59, 60, 61, 32, 110, 111,
	112, 113, 115, 116, 117, 118, 119, 120, 121, 122,
	114, 0, 0, 0, 48, 49, 50, 51, 52, 0,
	0, 0, 0, 36, 37, 0, 38, 0, 46, 47,
	0, 0, 0, 0, 0, 0, 0, 135, 70, 123,
	124, 0, 0, 0, 0, 0, 63, 0, 0, 0,
	0, 73, 74, 0, 0, 0, 0, 0, 22, 18,
	19, 35, 64, 66, 0, 20, 0, 21, 16, 17,
	25, 26, 27, 28, 29, 30, 31, 33, 0, 14,
	45, 15, 0, 0, 0, 0, 69, 0, 0, 0,
	0, 0, 0, 53, 54, 75, 65, 0, 0, 71,
	0, 0, 55, 56, 57, 62, 58, 59, 60, 61,
	32, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 48, 49, 50,
	51, 52, 0, 0, 0, 0, 36, 37, 0, 38,
	0, 46, 47, 0, 0, 0, 0, 0, 0, 0,
	13, 70, 0, 0, 0, 0, 0, 0, 0, 63,
	318, 0, 0, 0, 73, 74, 302, 0, 0, 0,
	0, 22, 18, 19, 35, 64, 66, 0, 20, 0,
	21, 16, 17, 25, 26, 27, 28, 29, 30, 31,
	33, 0, 14, 45, 15, 0, 0, 0, 0, 69,
	0, 0, 0, 0, 0, 0, 53, 54, 75, 65,
	0, 0, 71, 0, 0, 55, 56, 57, 62, 58,
	59, 60, 61, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	48, 49, 50, 51, 52, 0, 0, 0, 0, 36,
	37, 0, 38, 0, 46, 47, 0, 0, 0, 0,
	0, 0, 0, 319, 70, 0, 0, 0, 0, 0,
	0, 0, 63, 0, 0, 0, 0, 73, 74, 0,
	0, 0, 0, 0, 22, 18, 19, 35, 64, 66,
	0, 20, 0, 21, 16, 17, 25, 26, 27, 28,
	29, 30, 31, 33, 0, 14, 45, 15, 0, 0,
	0, 0, 69, 0, 0, 0, 0, 0, 0, 53,
	54, 75, 65, 0, 0, 71, 270, 0, 55, 56,
	57, 62, 58, 59, 60, 61, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 48, 49, 50, 51, 52, 0, 0,
	0, 0, 36, 37, 0, 38, 0, 46, 47, 0,
	0, 0, 0, 0, 0, 0, 13, 70, 0, 0,
	0, 0, 0, 0, 0, 63, 233, 0, 0, 0,
	73, 74, 0, 0, 0, 0, 0, 22, 18, 19,
	35, 64, 66, 0, 20, 0, 21, 16, 17, 25,
	26, 27, 28, 29, 30, 31, 33, 0, 14, 45,
	15, 0, 0, 0, 0, 69, 0, 0, 0, 0,
	0, 0, 53, 54, 75, 65, 0, 0, 71, 0,
	0, 55, 56, 57, 62, 58, 59, 60, 61, 32,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 48, 49, 50, 51,
	52, 0, 0, 0, 0, 36, 37, 0, 38, 0,
	46, 47, 0, 0, 0, 0, 0, 0, 0, 13,
	70, 0, 0, 230, 0, 0, 0, 0, 63, 0,
	0, 0, 0, 73, 74, 0, 0, 0, 0, 0,
	22, 18, 19, 35, 64, 66, 0, 20, 0, 21,
	16, 17, 25, 26, 27, 28, 29, 30, 31, 33,
	0, 14, 45, 15, 0, 0, 0, 0, 69, 0,
	0, 0, 0, 0, 0, 53, 54, 75, 65, 0,
	0, 71, 0, 0, 55, 56, 57, 62, 58, 59,
	60, 61, 32, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 48,
	49, 50, 51, 52, 0, 0, 0, 0, 36, 37,
	0, 38, 0, 46, 47, 0, 0, 0, 0, 0,
	0, 0, 13, 70, 0, 0, 0, 0, 0, 0,
	0, 63, 0, 0, 0, 0, 73, 74, 0, 0,
	0, 0, 0, 22, 18, 19, 35, 64, 66, 0,
	20, 0, 21, 16, 17, 25, 26, 27, 28, 29,
	30, 31, 33, 0, 14, 45, 15, 0, 0, 0,
	0, 69, 0, 0, 0, 0, 0, 0, 53, 54,
	75, 65, 0, 0, 71, 0, 0, 55, 56, 57,
	62, 58, 59, 60, 61, 32, 108, 0, 109, 81,
	82, 86, 88, 87, 100, 101, 98, 99, 106, 102,
	103, 104, 105, 96, 97, 90, 91, 89, 92, 94,
	95, 0, 107, 0, 0, 0, 46, 47, 0, 0,
	0, 0, 0, 0, 93, 13, 70, 0, 0, 0,
	0, 0, 0, 0, 63, 83, 85, 84, 0, 73,
	74, 0, 0, 0, 0, 0, 22, 0, 0, 35,
	64, 66, 0, 0, 0, 108, 0, 109, 81, 82,
	86, 88, 87, 100, 101, 98, 99, 106, 102, 103,
	104, 105, 96, 97, 90, 91, 89, 92, 94, 95,
	0, 107, 83, 85, 84, 0, 0, 0, 0, 0,
	0, 0, 0, 93, 0, 0, 0, 0, 0, 0,
	0, 0, 108, 0, 109, 81, 82, 86, 88, 87,
	100, 101, 98, 99, 106, 102, 103, 104, 105, 96,
	97, 90, 91, 89, 92, 94, 95, 0, 107, 83,
	85, 84, 0, 0, 0, 0, 0, 0, 0, 0,
	93, 0, 0, 0, 0, 0, 0, 0, 0, 108,
	0, 109, 81, 82, 86, 88, 87, 100, 101, 98,
	99, 106, 102, 103, 104, 105, 96, 97, 90, 91,
	89, 92, 94, 95, 0, 107, 0, 83, 85, 84,
	0, 0, 0, 0, 0, 0, 0, 93, 370, 0,
	0, 0, 0, 0, 0, 0, 0, 108, 0, 109,
	81, 82, 86, 88, 87, 100, 101, 98, 99, 106,
	102, 103, 104, 105, 96, 97, 90, 91, 89, 92,
	94, 95, 0, 107, 83, 85, 84, 0, 0, 0,
	0, 0, 0, 0, 347, 93, 0, 0, 0, 0,
	0, 0, 0, 0, 108, 0, 109, 81, 82, 86,
	88, 87, 100, 101, 98, 99, 106, 102, 103, 104,
	105, 96, 97, 90, 91, 89, 92, 94, 95, 0,
	107, 0, 83, 85, 84, 0, 0, 0, 0, 0,
	0, 346, 93, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 108, 0, 109, 81, 82, 86, 88, 87,
	100, 101, 98, 99, 106, 102, 103, 104, 105, 96,
	97, 90, 91, 89, 92, 94, 95, 0, 107, 83,
	85, 84, 0, 0, 0, 0, 0, 0, 0, 321,
	93, 0, 0, 0, 0, 0, 0, 0, 0, 108,
	0, 109, 81, 82, 86, 88, 87, 100, 101, 98,
	99, 106, 102, 103, 104, 105, 96, 97, 90, 91,
	89, 92, 94, 95, 0, 107, 83, 85, 84, 0,
	0, 0, 0, 0, 0, 0, 261, 93, 0, 0,
	0, 0, 0, 0, 0, 0, 108, 0, 109, 81,
	82, 86, 88, 87, 100, 101, 98, 99, 106, 102,
	103, 104, 105, 96, 97, 90, 91, 89, 92, 94,
	95, 0, 107, 83, 85, 84, 0, 0, 0, 0,
	0, 0, 80, 0, 93, 0, 0, 0, 0, 0,
	0, 0, 0, 108, 0, 109, 81, 82, 86, 88,
	87, 100, 101, 98, 99, 106, 102, 103, 104, 105,
	96, 97, 90, 91, 89, 92, 94, 95, 0, 107,
	83, 85, 84, 0, 0, 0, 0, 0, 365, 0,
	0, 93, 0, 0, 0, 0, 0, 0, 0, 0,
	108, 0, 109, 81, 82, 86, 88, 87, 100, 101,
	98, 99, 106, 102, 103, 104, 105, 96, 97, 90,
	91, 89, 92, 94, 95, 0, 107, 83, 85, 84,
	0, 0, 0, 0, 0, 363, 0, 0, 93, 0,
	0, 0, 0, 0, 0, 0, 0, 108, 0, 109,
	81, 82, 86, 88, 87, 100, 101, 98, 99, 106,
	102, 103, 104, 105, 96, 97, 90, 91, 89, 92,
	94, 95, 0, 107, 83, 85, 84, 0, 0, 0,
	0, 0, 354, 0, 0, 93, 0, 0, 0, 0,
	0, 0, 0, 0, 108, 0, 109, 81, 82, 86,
	88, 87, 100, 101, 98, 99, 106, 102, 103, 104,
	105, 96, 97, 90, 91, 89, 92, 94, 95, 0,
	107, 83, 85, 84, 0, 0, 0, 0, 0, 348,
	0, 0, 93, 0, 0, 0, 0, 0, 0, 0,
	0, 108, 307, 109, 81, 82, 86, 88, 87, 100,
	101, 98, 99, 106, 102, 103, 104, 105, 96, 97,
	90, 91, 89, 92, 94, 95, 0, 107, 83, 85,
	84, 0, 0, 256, 0, 0, 330, 0, 0, 93,
	0, 0, 0, 0, 0, 0, 0, 0, 108, 0,
	109, 81, 82, 86, 88, 87, 100, 101, 98, 99,
	106, 102, 103, 104, 105, 96, 97, 90, 91, 89,
	92, 94, 95, 0, 107, 83, 85, 84, 0, 0,
	0, 0, 0, 311, 0, 0, 93, 0, 0, 0,
	0, 0, 0, 0, 0, 108, 0, 109, 81, 82,
	86, 88, 87, 100, 101, 98, 99, 106, 102, 103,
	104, 105, 96, 97, 90, 91, 89, 92, 94, 95,
	0, 107, 85, 84, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 93, 0, 0, 0, 0, 0, 0,
	0, 108, 0, 109, 81, 82, 86, 88, 87, 100,
	101, 98, 99, 106, 102, 103, 104, 105, 96, 97,
	90, 91, 89, 92, 94, 95, 84, 107, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 93,
	0, 0, 0, 0, 108, 0, 109, 81, 82, 86,
	88, 87, 100, 101, 98, 99, 106, 102, 103, 104,
	105, 96, 97, 90, 91, 89, 92, 94, 95, 272,
	107, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 93, 0, 108, 0, 109, 81, 82, 86,
	88, 87, 100, 101, 98, 99, 106, 102, 103, 104,
	105, 96, 97, 90, 91, 89, 92, 94, 95, 0,
	107, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 93, 109, 81, 82, 86, 88, 87, 100,
	101, 98, 99, 106, 102, 103, 104, 105, 96, 97,
	90, 91, 89, 92, 94, 95, 0, 107, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 93,
	86, 88, 87, 100, 101, 98, 99, 106, 102, 103,
	104, 105, 96, 97, 90, 91, 89, 92, 94, 95,
	0, 107, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 93, 88, 87, 100, 101, 98, 99,
	106, 102, 103, 104, 105, 96, 97, 90, 91, 89,
	92, 94, 95, 0, 107, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 93, 87, 100, 101,
	98, 99, 106, 102, 103, 104, 105, 96, 97, 90,
	91, 89, 92, 94, 95, 0, 107, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 93, 100,
	101, 98, 99, 106, 102, 103, 104, 105, 96, 97,
	90, 91, 89, 92, 94, 95, 0, 107, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 93,
}
var yyPact = [...]int{

	-1000, -1000, 554, -1000, -1000, -1000, 1345, 1722, 751, -1000,
	-1000, -1000, 77, -82, 667, 1345, 66, 66, 1345, 1345,
	1345, 1345, 1345, -1000, -1000, 1345, 1345, 1345, 1345, 1345,
	1345, 1345, -85, 1345, -1000, 158, 1345, 1345, 1345, -1000,
	114, -1000, 0, -3, -1000, -17, -86, -87, 1345, 1345,
	-88, 1345, 1345, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 148, 129, -1000, -22, -96, -96, -1000,
	-89, -1000, -24, -23, 81, -1000, 321, 62, -1000, 2145,
	-1000, 1345, 1345, 1345, 1345, 1345, 1345, 1345, 1345, 1345,
	1345, 1345, 1345, 1345, 1345, 1345, 1345, 1345, 1345, 1345,
	1345, 1345, 1345, 1345, 1345, 1345, 1345, -17, 1232, 1345,
	1119, 1345, 1345, 1345, 1345, 1345, 1345, 1345, 1345, 1345,
	1345, 1345, 1345, -1000, -1000, 1345, 1345, -25, 667, -122,
	185, -1000, -1000, 2098, 66, -91, -1000, -1000, 114, 1345,
	-1000, 667, -1000, 125, 125, 177, 125, 1674, 125, 125,
	125, 125, 125, 125, 125, -1000, 1345, 125, -156, 102,
	99, -1000, 87, 1006, 66, 1416, 2274, 1416, 1345, -19,
	-19, -96, -5, 65, -1000, -1000, 1345, 1345, 2145, 2145,
	1345, 2145, 2145, 90, -1000, 115, 73, 102, 1345, -1000,
	-1000, 434, -1000, 667, 89, 81, -24, -1000, -1000, -1000,
	-1000, 1345, 501, 2345, 2191, 1416, 2234, 2378, 2441, 2410,
	403, 403, 403, 177, 125, 177, 177, 289, 289, 381,
	381, 381, 381, 608, 608, 608, 608, 381, -1000, 2051,
	1345, 2311, 1416, 66, 1416, 1416, 1416, 1416, 1416, 1416,
	1416, 1416, 1416, 1416, 1416, 1416, -127, 2145, 2004, -96,
	-1000, 1345, -1000, -108, 180, 667, 893, -1000, 667, 1627,
	-128, -1000, -111, -1000, -1000, -1000, -1000, 197, 80, 1957,
	60, -78, 1345, -139, -96, -1000, -1000, 1345, -1000, -96,
	-1000, -1000, -1000, -1000, -38, 1345, 1345, -25, -38, 9,
	-1000, 2145, 1579, 1532, -1000, -1000, -1000, 1910, -1000, 8,
	-1000, 2145, 1345, -115, -1000, -24, -1000, 1345, 2311, -1000,
	-1000, -1000, -1000, 1863, 141, 1345, -1000, 2145, 66, -106,
	-116, -1000, -1000, -1000, -148, -1000, -1000, 63, -1000, -1000,
	-1000, -1000, 1345, -1000, 1416, -1000, -1000, 1816, -1000, -1000,
	-149, 1769, -1000, -1000, -1000, 1345, -1000, -1000, -1000, -1000,
	780, 2145, -1000, 2311, -1000, 1345, 1416, -1000, 667, 141,
	-1000, -1000, 1485, -1000, -1000, -1000, -1000, -1000, 1416, -118,
	-79, 141, -1000,
}
var yyPgo = [...]int{

	0, 302, 301, 191, 33, 163, 297, 0, 10, 45,
	2, 292, 9, 291, 162, 168, 78, 290, 289, 285,
	284, 13, 5, 277, 40, 25, 275, 272, 269, 264,
	263, 12, 261, 258, 256, 255, 243, 239, 56, 237,
	214, 230, 213, 212, 1, 6, 4, 11,
}
var yyR1 = [...]int{

	0, 1, 47, 33, 33, 4, 4, 5, 5, 5,
	2, 34, 34, 6, 3, 3, 3, 38, 38, 39,
	39, 10, 10, 37, 37, 12, 13, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 20, 20, 20, 20, 14, 14, 15, 15, 18,
	18, 36, 36, 36, 40, 40, 24, 24, 24, 19,
	19, 19, 19, 19, 19, 19, 19, 19, 19, 19,
	19, 19, 19, 19, 19, 25, 25, 25, 7, 7,
	8, 8, 23, 26, 26, 26, 27, 27, 27, 28,
	28, 28, 28, 28, 28, 9, 9, 9, 16, 16,
	16, 29, 29, 30, 30, 30, 30, 30, 30, 21,
	21, 21, 22, 22, 22, 44, 45, 45, 43, 43,
	42, 42, 42, 42, 42, 42, 41, 41, 41, 41,
	31, 31, 31, 31, 31, 31, 31, 32, 32, 32,
	32, 17, 17, 17, 17, 17, 17, 17, 35, 35,
	46,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 0, 1, 3, 1, 3, 2,
	1, 2, 0, 1, 3, 3, 2, 2, 3, 1,
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

	-1000, -1, -33, -2, -3, 158, 82, -7, -9, -11,
	-28, -29, -26, 140, 69, 71, 58, 59, 49, 50,
	55, 57, 161, -13, -17, 60, 61, 62, 63, 64,
	65, 66, 100, 67, -19, 164, 13, 14, 16, -16,
	-25, -20, -14, -23, -24, 70, 131, 132, 4, 5,
	6, 7, 8, 83, 84, 92, 93, 94, 96, 97,
	98, 99, 95, 149, 165, 86, 166, -5, -27, 76,
	141, 89, -4, 154, 155, 85, -34, -37, -12, -7,
	160, 33, 34, 10, 12, 11, 35, 37, 36, 51,
	49, 50, 52, 68, 53, 54, 47, 48, 40, 41,
	38, 39, 43, 44, 45, 46, 42, 56, 30, 32,
	17, 18, 19, 20, 29, 21, 22, 23, 24, 25,
	26, 27, 28, 58, 59, 69, 158, 139, 161, -44,
	-43, -45, -42, -7, 37, 140, -7, -9, -25, 161,
	-24, 69, -9, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -18, 161, -7, -36, 88,
	-41, -31, 86, 151, 152, -7, -7, -7, 69, 153,
	153, -15, -14, -30, -5, -16, 161, 161, -7, -7,
	161, -7, -7, 88, 150, -41, -41, 88, 158, -16,
	-38, 161, -38, 161, 155, 155, -4, -6, 159, -3,
	160, 9, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -15, -7,
	31, -7, -7, 37, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -8, -7, -7, -22,
	85, 158, -16, -44, 163, 9, 15, -9, 161, -7,
	-44, 162, -8, 164, -31, -31, 88, 69, 139, -7,
	90, -9, 15, -8, -21, -47, -16, 158, 85, -21,
	-47, -16, -40, -38, 153, 69, 158, 139, 153, -35,
	-46, -7, -7, -7, 150, 150, 165, -7, 162, -39,
	-10, -7, 156, -44, 85, -4, -12, 31, -7, -9,
	163, 159, -38, -7, 162, 17, -45, -7, 37, 140,
	-44, 162, 163, 162, -32, 85, 91, 50, 86, 85,
	159, 159, 69, 159, -7, 163, -38, -7, -38, -16,
	-8, -7, -22, -16, 162, 9, 162, 162, 159, 162,
	9, -7, 162, -7, 159, 17, -7, -9, 161, 162,
	163, 91, -7, 159, 163, 159, -46, -10, -7, -44,
	163, 162, 159,
}
var yyDef = [...]int{

	4, -2, 1, 3, 10, 12, 0, 0, -2, 139,
	-2, 156, 142, 0, 176, 0, 0, 0, 0, 0,
	0, 0, 0, 80, 84, 0, 0, 0, 0, 0,
	0, 0, 109, 0, 94, 111, 0, 97, 0, 149,
	134, 154, 0, 0, -2, 0, 0, 0, 0, 0,
	0, 0, 0, 119, 120, 121, 122, 123, 124, 125,
	126, 127, 128, 0, 0, 158, 0, -2, 0, 105,
	0, 118, 7, 0, 0, 5, 0, 0, 24, 25,
	16, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 44, 46, 140, 0, 0, 176, 0,
	175, 179, 177, 181, 0, 0, 31, -2, 0, 0,
	-2, 176, -2, 65, 66, 67, 68, 0, 85, 86,
	87, 88, 89, 90, 91, 92, 140, 93, 0, 112,
	113, 188, 190, 0, 0, 96, 98, 100, 140, 0,
	0, 114, 107, 108, 106, 163, 0, 0, 203, 204,
	0, 206, 207, 0, 130, 0, 0, 0, 0, 160,
	101, 0, 104, 176, 0, 0, 9, 11, 14, 13,
	15, 0, 48, 49, 50, 51, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 61, 62, 63, 64, -2,
	-2, -2, -2, -2, -2, -2, -2, -2, 78, 0,
	0, 83, 29, 0, 32, 33, 34, 35, 36, 37,
	38, 39, 40, 41, 42, 43, 0, 141, 0, 157,
	172, 0, 174, 0, 117, 176, 0, -2, 176, 0,
	0, -2, 0, 95, 189, 186, 187, 0, 0, 0,
	0, 143, 0, 0, 0, -2, -2, 0, 2, 0,
	-2, -2, 26, 115, 0, 140, 0, 0, 0, 0,
	208, 210, 0, 0, 129, 132, 131, 0, 17, 0,
	19, 21, 0, 0, 6, 8, 23, 0, 82, -2,
	150, 152, 153, 0, 0, 0, 178, 180, 0, 0,
	0, -2, 117, 110, 0, 197, 198, 0, 200, 192,
	193, 194, 0, 196, 99, 151, 102, 0, 103, 167,
	0, 0, 166, 168, 201, 0, 202, 205, 159, 18,
	0, 22, 116, 81, 173, 0, 28, -2, 176, 185,
	191, 199, 0, 170, 164, 165, 209, 20, 27, 0,
	0, 184, 195,
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:420
		{
			yyVAL.stmt = ast.NewEchoStatement(yyDollar[1].tok, yyDollar[2].exprs)
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:422
		{
			yyVAL.stmt = ast.NewExpressionStatement(yyDollar[1].expr)
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:655
		{
			yyVAL.expr = ast.NewArgumentListExpression()
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:656
		{
			yyVAL.expr = ast.NewArgumentListExpression(yyDollar[2].exprs...)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:661
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:663
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:667
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:668
		{
			yyVAL.expr = ast.NewArgumentExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:823
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:824
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:828
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:855
		{
			yyVAL.expr = ast.NewNewExpression(yyDollar[1].tok, yyDollar[2].expr, yyDollar[3].expr)
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:862
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...), yyDollar[6].expr, false)
		}
	case 28:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:864
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...), yyDollar[5].expr, false)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:866
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:868
		{
			yyVAL.expr = ast.NewAssignExpression(ast.Equal, yyDollar[1].expr, yyDollar[4].expr, true)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:869
		{
			yyVAL.expr = ast.NewCloneExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:871
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PlusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:873
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MinusEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:875
		{
			yyVAL.expr = ast.NewAssignExpression(ast.MulEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:877
		{
			yyVAL.expr = ast.NewAssignExpression(ast.PowEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:879
		{
			yyVAL.expr = ast.NewAssignExpression(ast.DivEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:881
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ConcatEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:883
		{
			yyVAL.expr = ast.NewAssignExpression(ast.ModEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:885
		{
			yyVAL.expr = ast.NewAssignExpression(ast.AndEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:887
		{
			yyVAL.expr = ast.NewAssignExpression(ast.QrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:889
		{
			yyVAL.expr = ast.NewAssignExpression(ast.XorEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:891
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SlEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:893
		{
			yyVAL.expr = ast.NewAssignExpression(ast.SrEqual, yyDollar[1].expr, yyDollar[3].expr, false)
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:894
		{
			yyVAL.expr = ast.NewIncrementExpression(ast.PostInc, yyDollar[1].expr)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:895
		{
			yyVAL.expr = ast.NewIncrementExpression(ast.PreInc, yyDollar[2].expr)
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:896
		{
			yyVAL.expr = ast.NewDecrementExpression(ast.PostDec, yyDollar[1].expr)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:897
		{
			yyVAL.expr = ast.NewDecrementExpression(ast.PreDec, yyDollar[2].expr)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:899
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BooleanOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:901
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BooleanAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:903
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:905
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:907
		{
			yyVAL.expr = ast.NewInfixExpression(ast.LogicalXor, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:908
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwOr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:909
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwAnd, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:910
		{
			yyVAL.expr = ast.NewInfixExpression(ast.BwXor, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:911
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Concat, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:912
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Add, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:913
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sub, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:914
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Mul, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:915
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Pow, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:916
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Div, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:917
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Mod, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:918
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sl, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:919
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Sr, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:920
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.UnaryPlus, yyDollar[2].expr)
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:921
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.UnaryMinus, yyDollar[2].expr)
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:922
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.BoolNot, yyDollar[2].expr)
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:923
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.BwNot, yyDollar[2].expr)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:925
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsIdentical, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:927
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsNotIdentical, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:929
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:931
		{
			yyVAL.expr = ast.NewInfixExpression(ast.IsNotEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:933
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Smaller, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:935
		{
			yyVAL.expr = ast.NewInfixExpression(ast.SmallerOrEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:937
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Greater, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:939
		{
			yyVAL.expr = ast.NewInfixExpression(ast.GreaterOrEqual, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:941
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Spaceship, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:943
		{
			yyVAL.expr = ast.NewInfixExpression(ast.InstanceOf, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:944
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:945
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:947
		{
			yyVAL.expr = ast.NewTernaryOperatorExpression(yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr)
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:949
		{
			yyVAL.expr = ast.NewTernaryOperatorExpression(yyDollar[1].expr, nil, yyDollar[4].expr)
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:951
		{
			yyVAL.expr = ast.NewInfixExpression(ast.Coalesce, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:952
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:953
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:954
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:955
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:956
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:957
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:958
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:959
		{
			yyVAL.expr = ast.NewCastExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:960
		{
			yyVAL.expr = ast.NewExitExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:961
		{
			yyVAL.expr = ast.NewPrefixExpression(ast.Silence, yyDollar[2].expr)
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:962
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:963
		{
			yyVAL.expr = ast.NewBackticksExpression(yyDollar[2].exprs...)
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:964
		{
			yyVAL.expr = ast.NewPrintExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:965
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, nil)
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:966
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:967
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, ast.NewArrayPairExpression(yyDollar[2].expr, yyDollar[4].expr, false))
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:968
		{
			yyVAL.expr = ast.NewYieldExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1016
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1018
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1020
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.StaticCall, yyDollar[1].expr, yyDollar[3].expr, yyDollar[4].expr)
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1022
		{
			yyVAL.expr = ast.NewFunctionCallExpression(ast.Call, yyDollar[1].expr, nil, yyDollar[2].expr)
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1026
		{
			yyVAL.expr = ast.NewStaticLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1027
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1031
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1032
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1036
		{
			yyVAL.expr = nil
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1037
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1042
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1043
		{
			yyVAL.exprs = []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1044
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1048
		{
			yyVAL.expr = nil
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1049
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1053
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Long, yyDollar[3].exprs...)
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1054
		{
			yyVAL.expr = ast.NewArrayExpression(ast.Short, yyDollar[2].exprs...)
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1055
		{
			yyVAL.expr = ast.NewConstantEncapsedStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1059
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1060
		{
			yyVAL.expr = ast.NewDoubleLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1061
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1062
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1063
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1064
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1065
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1066
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1067
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1068
		{
			yyVAL.expr = ast.NewMagicConstLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1069
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1070
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[2].tok)
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1071
		{
			yyVAL.expr = ast.NewEncapsListExpression(yyDollar[2].exprs...)
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1072
		{
			yyVAL.expr = ast.NewHeredocExpression(yyDollar[1].tok, yyDollar[3].tok, yyDollar[2].exprs...)
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1073
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1074
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1078
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1080
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1082
		{
			yyVAL.expr = ast.NewConstantExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1086
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1087
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 140:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1091
		{
			yyVAL.expr = nil
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1092
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1096
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1100
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1101
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1102
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1106
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1107
		{
			yyVAL.expr = ast.NewDereferencableExpression(ast.Wrapped, yyDollar[2].expr)
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1108
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1113
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1115
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1117
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1119
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Curly, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1121
		{
			yyVAL.expr = ast.NewCallableVariableExpression(ast.Prop, yyDollar[1].expr, []ast.Expression{yyDollar[3].expr, yyDollar[4].expr}...)
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1122
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1127
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1129
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1131
		{
			yyVAL.expr = ast.NewVariableExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1135
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 159:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1136
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.CurlyOpen, yyDollar[3].expr)
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1137
		{
			yyVAL.expr = ast.NewSimpleVariableExpression(ast.Var, yyDollar[2].expr)
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1142
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1144
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1149
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1151
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Dim, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 165:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1153
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Curly, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1155
		{
			yyVAL.expr = ast.NewNVariableExpression(ast.Prop, yyDollar[1].expr, yyDollar[3].expr)
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1157
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1159
		{
			yyVAL.expr = ast.NewStaticMemberExpression(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1163
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1164
		{
			yyVAL.expr = ast.NewMemberNameExpression(yyDollar[2].expr)
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1165
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1169
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1170
		{
			yyVAL.expr = ast.NewPropertyNameExpression(yyDollar[2].expr)
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1171
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1176
		{ /* allow single trailing comma */
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:1180
		{
			yyVAL.exprs = []ast.Expression{}
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1181
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1186
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].exprs...)
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1188
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1193
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[3].expr, false)}
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1195
		{
			yyVAL.exprs = []ast.Expression{yyDollar[1].expr}
		}
	case 182:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1197
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, yyDollar[4].expr, true)}
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1199
		{
			yyVAL.exprs = []ast.Expression{ast.NewAmpersandLiteral(yyDollar[2].expr)}
		}
	case 184:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1201
		{
			yyVAL.exprs = []ast.Expression{ast.NewArrayPairExpression(yyDollar[1].expr, ast.NewListExpression(yyDollar[3].tok, yyDollar[5].exprs...), false)}
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1203
		{
			yyVAL.exprs = []ast.Expression{ast.NewListExpression(yyDollar[1].tok, yyDollar[3].exprs...)}
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1208
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1210
		{
			yyVAL.exprs = append(yyDollar[1].exprs, ast.NewEncapsedAndWhitespaceLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1212
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1214
		{
			yyVAL.exprs = append(yyVAL.exprs, []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal), yyDollar[2].expr}...)
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1219
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 191:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1221
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Dim, yyDollar[3].expr)
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1223
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.Prop, ast.NewStringLiteral(yyDollar[3].tok, yyDollar[3].tok.Literal))
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1225
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, yyDollar[2].expr)
		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1227
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DollarOpenCurlyBraces, ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal))
		}
	case 195:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:1229
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.DimInDollarOpenCurlyBraces, []ast.Expression{ast.NewStringLiteral(yyDollar[2].tok, yyDollar[2].tok.Literal), yyDollar[4].expr}...)
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1230
		{
			yyVAL.expr = ast.NewEncapsVarExpression(yyDollar[1].tok, yyDollar[1].tok.Literal, ast.CurlyOpen, yyDollar[2].expr)
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1234
		{
			yyVAL.expr = ast.NewStringLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1235
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1236
		{
			yyVAL.expr = ast.NewIntegerLiteral(yyDollar[2].tok, "-"+yyDollar[2].tok.Literal)
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1237
		{
			yyVAL.expr = ast.NewVariableLiteral(yyDollar[1].tok, yyDollar[1].tok.Literal)
		}
	case 201:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1242
		{
			yyVAL.expr = ast.NewIssetExpression(yyDollar[1].tok, yyDollar[3].exprs...)
		}
	case 202:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1243
		{
			yyVAL.expr = ast.NewEmptyExpression(yyDollar[1].tok, yyDollar[3].expr)
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1245
		{
			yyVAL.expr = ast.NewIncludeExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1247
		{
			yyVAL.expr = ast.NewIncludeExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 205:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:1249
		{
			yyVAL.expr = ast.NewEvalExpression(yyDollar[1].tok, yyDollar[3].expr)
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1251
		{
			yyVAL.expr = ast.NewRequireExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:1253
		{
			yyVAL.expr = ast.NewRequireExpression(yyDollar[1].tok, yyDollar[2].expr)
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1257
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[1].expr)
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1259
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:1263
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
