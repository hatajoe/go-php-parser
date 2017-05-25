package token

import (
	"encoding/json"
)

const (
	EOF     = -1
	UNKNOWN = 0

	// TODO: copy from parser.go to resolve import cycle.
	T_INCLUDE                  = 57346
	T_INCLUDE_ONCE             = 57347
	T_EVAL                     = 57348
	T_REQUIRE                  = 57349
	T_REQUIRE_ONCE             = 57350
	T_LOGICAL_OR               = 57351
	T_LOGICAL_XOR              = 57352
	T_LOGICAL_AND              = 57353
	T_PRINT                    = 57354
	T_YIELD                    = 57355
	T_DOUBLE_ARROW             = 57356
	T_YIELD_FROM               = 57357
	T_PLUS_EQUAL               = 57358
	T_MINUS_EQUAL              = 57359
	T_MUL_EQUAL                = 57360
	T_DIV_EQUAL                = 57361
	T_CONCAT_EQUAL             = 57362
	T_MOD_EQUAL                = 57363
	T_AND_EQUAL                = 57364
	T_OR_EQUAL                 = 57365
	T_XOR_EQUAL                = 57366
	T_SL_EQUAL                 = 57367
	T_SR_EQUAL                 = 57368
	T_POW_EQUAL                = 57369
	T_COALESCE                 = 57370
	T_BOOLEAN_OR               = 57371
	T_BOOLEAN_AND              = 57372
	T_IS_EQUAL                 = 57373
	T_IS_NOT_EQUAL             = 57374
	T_IS_IDENTICAL             = 57375
	T_IS_NOT_IDENTICAL         = 57376
	T_SPACESHIP                = 57377
	T_IS_SMALLER_OR_EQUAL      = 57378
	T_IS_GREATER_OR_EQUAL      = 57379
	T_SL                       = 57380
	T_SR                       = 57381
	T_INSTANCEOF               = 57382
	T_INC                      = 57383
	T_DEC                      = 57384
	T_INT_CAST                 = 57385
	T_DOUBLE_CAST              = 57386
	T_STRING_CAST              = 57387
	T_ARRAY_CAST               = 57388
	T_OBJECT_CAST              = 57389
	T_BOOL_CAST                = 57390
	T_UNSET_CAST               = 57391
	T_POW                      = 57392
	T_NEW                      = 57393
	T_CLONE                    = 57394
	T_NOELSE                   = 57395
	T_ELSEIF                   = 57396
	T_ELSE                     = 57397
	T_ENDIF                    = 57398
	T_STATIC                   = 57399
	T_ABSTRACT                 = 57400
	T_FINAL                    = 57401
	T_PRIVATE                  = 57402
	T_PROTECTED                = 57403
	T_PUBLIC                   = 57404
	T_LNUMBER                  = 57405
	T_DNUMBER                  = 57406
	T_STRING                   = 57407
	T_VARIABLE                 = 57408
	T_INLINE_HTML              = 57409
	T_ENCAPSED_AND_WHITESPACE  = 57410
	T_CONSTANT_ENCAPSED_STRING = 57411
	T_STRING_VARNAME           = 57412
	T_NUM_STRING               = 57413
	T_ECHO                     = 57414
	T_EXIT                     = 57415
	T_IF                       = 57416
	T_DO                       = 57417
	T_WHILE                    = 57418
	T_ENDWHILE                 = 57419
	T_FOR                      = 57420
	T_ENDFOR                   = 57421
	T_FOREACH                  = 57422
	T_ENDFOREACH               = 57423
	T_DECLARE                  = 57424
	T_ENDDECLARE               = 57425
	T_AS                       = 57426
	T_SWITCH                   = 57427
	T_ENDSWITCH                = 57428
	T_CASE                     = 57429
	T_DEFAULT                  = 57430
	T_BREAK                    = 57431
	T_CONTINUE                 = 57432
	T_GOTO                     = 57433
	T_FUNCTION                 = 57434
	T_CONST                    = 57435
	T_RETURN                   = 57436
	T_TRY                      = 57437
	T_CATCH                    = 57438
	T_FINALLY                  = 57439
	T_THROW                    = 57440
	T_USE                      = 57441
	T_INSTEADOF                = 57442
	T_GLOBAL                   = 57443
	T_VAR                      = 57444
	T_UNSET                    = 57445
	T_ISSET                    = 57446
	T_EMPTY                    = 57447
	T_HALT_COMPILER            = 57448
	T_CLASS                    = 57449
	T_TRAIT                    = 57450
	T_INTERFACE                = 57451
	T_EXTENDS                  = 57452
	T_IMPLEMENTS               = 57453
	T_OBJECT_OPERATOR          = 57454
	T_LIST                     = 57455
	T_ARRAY                    = 57456
	T_CALLABLE                 = 57457
	T_LINE                     = 57458
	T_FILE                     = 57459
	T_DIR                      = 57460
	T_CLASS_C                  = 57461
	T_TRAIT_C                  = 57462
	T_METHOD_C                 = 57463
	T_FUNC_C                   = 57464
	T_COMMENT                  = 57465
	T_DOC_COMMENT              = 57466
	T_OPEN_TAG                 = 57467
	T_OPEN_TAG_WITH_ECHO       = 57468
	T_CLOSE_TAG                = 57469
	T_WHITESPACE               = 57470
	T_START_HEREDOC            = 57471
	T_END_HEREDOC              = 57472
	T_DOLLAR_OPEN_CURLY_BRACES = 57473
	T_CURLY_OPEN               = 57474
	T_PAAMAYIM_NEKUDOTAYIM     = 57475
	T_NAMESPACE                = 57476
	T_NS_C                     = 57477
	T_NS_SEPARATOR             = 57478
	T_ELLIPSIS                 = 57479
	T_ERROR                    = 57480
)

var keywords = map[string]int{
	"T_INCLUDE":                  T_INCLUDE,
	"T_INCLUDE_ONCE":             T_INCLUDE_ONCE,
	"T_EVAL":                     T_EVAL,
	"T_REQUIRE":                  T_REQUIRE,
	"T_REQUIRE_ONCE":             T_REQUIRE_ONCE,
	"T_LOGICAL_OR":               T_LOGICAL_OR,
	"T_LOGICAL_XOR":              T_LOGICAL_XOR,
	"T_LOGICAL_AND":              T_LOGICAL_AND,
	"T_PRINT":                    T_PRINT,
	"T_YIELD":                    T_YIELD,
	"T_DOUBLE_ARROW":             T_DOUBLE_ARROW,
	"T_YIELD_FROM":               T_YIELD_FROM,
	"T_PLUS_EQUAL":               T_PLUS_EQUAL,
	"T_MINUS_EQUAL":              T_MINUS_EQUAL,
	"T_MUL_EQUAL":                T_MUL_EQUAL,
	"T_DIV_EQUAL":                T_DIV_EQUAL,
	"T_CONCAT_EQUAL":             T_CONCAT_EQUAL,
	"T_MOD_EQUAL":                T_MOD_EQUAL,
	"T_AND_EQUAL":                T_AND_EQUAL,
	"T_OR_EQUAL":                 T_OR_EQUAL,
	"T_XOR_EQUAL":                T_XOR_EQUAL,
	"T_SL_EQUAL":                 T_SL_EQUAL,
	"T_SR_EQUAL":                 T_SR_EQUAL,
	"T_POW_EQUAL":                T_POW_EQUAL,
	"T_COALESCE":                 T_COALESCE,
	"T_BOOLEAN_OR":               T_BOOLEAN_OR,
	"T_BOOLEAN_AND":              T_BOOLEAN_AND,
	"T_IS_EQUAL":                 T_IS_EQUAL,
	"T_IS_NOT_EQUAL":             T_IS_NOT_EQUAL,
	"T_IS_IDENTICAL":             T_IS_IDENTICAL,
	"T_IS_NOT_IDENTICAL":         T_IS_NOT_IDENTICAL,
	"T_SPACESHIP":                T_SPACESHIP,
	"T_IS_SMALLER_OR_EQUAL":      T_IS_SMALLER_OR_EQUAL,
	"T_IS_GREATER_OR_EQUAL":      T_IS_GREATER_OR_EQUAL,
	"T_SL":                       T_SL,
	"T_SR":                       T_SR,
	"T_INSTANCEOF":               T_INSTANCEOF,
	"T_INC":                      T_INC,
	"T_DEC":                      T_DEC,
	"T_INT_CAST":                 T_INT_CAST,
	"T_DOUBLE_CAST":              T_DOUBLE_CAST,
	"T_STRING_CAST":              T_STRING_CAST,
	"T_ARRAY_CAST":               T_ARRAY_CAST,
	"T_OBJECT_CAST":              T_OBJECT_CAST,
	"T_BOOL_CAST":                T_BOOL_CAST,
	"T_UNSET_CAST":               T_UNSET_CAST,
	"T_POW":                      T_POW,
	"T_NEW":                      T_NEW,
	"T_CLONE":                    T_CLONE,
	"T_NOELSE":                   T_NOELSE,
	"T_ELSEIF":                   T_ELSEIF,
	"T_ELSE":                     T_ELSE,
	"T_ENDIF":                    T_ENDIF,
	"T_STATIC":                   T_STATIC,
	"T_ABSTRACT":                 T_ABSTRACT,
	"T_FINAL":                    T_FINAL,
	"T_PRIVATE":                  T_PRIVATE,
	"T_PROTECTED":                T_PROTECTED,
	"T_PUBLIC":                   T_PUBLIC,
	"T_LNUMBER":                  T_LNUMBER,
	"T_DNUMBER":                  T_DNUMBER,
	"T_STRING":                   T_STRING,
	"T_VARIABLE":                 T_VARIABLE,
	"T_INLINE_HTML":              T_INLINE_HTML,
	"T_ENCAPSED_AND_WHITESPACE":  T_ENCAPSED_AND_WHITESPACE,
	"T_CONSTANT_ENCAPSED_STRING": T_CONSTANT_ENCAPSED_STRING,
	"T_STRING_VARNAME":           T_STRING_VARNAME,
	"T_NUM_STRING":               T_NUM_STRING,
	"T_EXIT":                     T_EXIT,
	"T_IF":                       T_IF,
	"T_ECHO":                     T_ECHO,
	"T_DO":                       T_DO,
	"T_WHILE":                    T_WHILE,
	"T_ENDWHILE":                 T_ENDWHILE,
	"T_FOR":                      T_FOR,
	"T_ENDFOR":                   T_ENDFOR,
	"T_FOREACH":                  T_FOREACH,
	"T_ENDFOREACH":               T_ENDFOREACH,
	"T_DECLARE":                  T_DECLARE,
	"T_ENDDECLARE":               T_ENDDECLARE,
	"T_AS":                       T_AS,
	"T_SWITCH":                   T_SWITCH,
	"T_ENDSWITCH":                T_ENDSWITCH,
	"T_CASE":                     T_CASE,
	"T_DEFAULT":                  T_DEFAULT,
	"T_BREAK":                    T_BREAK,
	"T_CONTINUE":                 T_CONTINUE,
	"T_GOTO":                     T_GOTO,
	"T_FUNCTION":                 T_FUNCTION,
	"T_CONST":                    T_CONST,
	"T_RETURN":                   T_RETURN,
	"T_TRY":                      T_TRY,
	"T_CATCH":                    T_CATCH,
	"T_FINALLY":                  T_FINALLY,
	"T_THROW":                    T_THROW,
	"T_USE":                      T_USE,
	"T_INSTEADOF":                T_INSTEADOF,
	"T_GLOBAL":                   T_GLOBAL,
	"T_VAR":                      T_VAR,
	"T_UNSET":                    T_UNSET,
	"T_ISSET":                    T_ISSET,
	"T_EMPTY":                    T_EMPTY,
	"T_HALT_COMPILER":            T_HALT_COMPILER,
	"T_CLASS":                    T_CLASS,
	"T_TRAIT":                    T_TRAIT,
	"T_INTERFACE":                T_INTERFACE,
	"T_EXTENDS":                  T_EXTENDS,
	"T_IMPLEMENTS":               T_IMPLEMENTS,
	"T_OBJECT_OPERATOR":          T_OBJECT_OPERATOR,
	"T_LIST":                     T_LIST,
	"T_ARRAY":                    T_ARRAY,
	"T_CALLABLE":                 T_CALLABLE,
	"T_LINE":                     T_LINE,
	"T_FILE":                     T_FILE,
	"T_DIR":                      T_DIR,
	"T_CLASS_C":                  T_CLASS_C,
	"T_TRAIT_C":                  T_TRAIT_C,
	"T_METHOD_C":                 T_METHOD_C,
	"T_FUNC_C":                   T_FUNC_C,
	"T_COMMENT":                  T_COMMENT,
	"T_DOC_COMMENT":              T_DOC_COMMENT,
	"T_OPEN_TAG":                 T_OPEN_TAG,
	"T_OPEN_TAG_WITH_ECHO":       T_OPEN_TAG_WITH_ECHO,
	"T_CLOSE_TAG":                T_CLOSE_TAG,
	"T_WHITESPACE":               T_WHITESPACE,
	"T_START_HEREDOC":            T_START_HEREDOC,
	"T_END_HEREDOC":              T_END_HEREDOC,
	"T_DOLLAR_OPEN_CURLY_BRACES": T_DOLLAR_OPEN_CURLY_BRACES,
	"T_CURLY_OPEN":               T_CURLY_OPEN,
	"T_PAAMAYIM_NEKUDOTAYIM":     T_PAAMAYIM_NEKUDOTAYIM,
	"T_NAMESPACE":                T_NAMESPACE,
	"T_NS_C":                     T_NS_C,
	"T_NS_SEPARATOR":             T_NS_SEPARATOR,
	"T_ELLIPSIS":                 T_ELLIPSIS,
	"T_ERROR":                    T_ERROR,
}

type (
	Tokens []*Token

	TokenType int

	Token struct {
		Type     TokenType
		Literal  string
		Position Position
	}

	Position struct {
		Line   int
		Column int
	}
)

func (t *Token) UnmarshalJSON(data []byte) error {
	aux := struct {
		Tok string `json:"tok"`
		Lit string `json:"lit"`
		Pos int    `json:"pos"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	tok, ok := keywords[aux.Tok]
	if !ok {
		tok = int([]rune(aux.Tok)[0])
	}
	t.Type = TokenType(tok)
	t.Literal = aux.Lit
	t.Position = Position{
		Line:   aux.Pos,
		Column: 0, // TODO
	}
	return nil
}
