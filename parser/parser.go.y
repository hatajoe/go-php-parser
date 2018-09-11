%{
package parser

import (
    "log"
    "github.com/hatajoe/go-php-parser/ast"
    "github.com/hatajoe/go-php-parser/token"
    "github.com/hatajoe/go-php-parser/lexer"
)

%}

%union {
    program *ast.Program
    stmts []ast.Statement
    stmt ast.Statement
    exprs []ast.Expression
    expr ast.Expression
    num int
    tok *token.Token
}

%left T_INCLUDE T_INCLUDE_ONCE T_EVAL T_REQUIRE T_REQUIRE_ONCE
%left ','
%left T_LOGICAL_OR
%left T_LOGICAL_XOR
%left T_LOGICAL_AND
%right T_PRINT
%right T_YIELD
%right T_DOUBLE_ARROW
%right T_YIELD_FROM
%left '=' T_PLUS_EQUAL T_MINUS_EQUAL T_MUL_EQUAL T_DIV_EQUAL T_CONCAT_EQUAL T_MOD_EQUAL T_AND_EQUAL T_OR_EQUAL T_XOR_EQUAL T_SL_EQUAL T_SR_EQUAL T_POW_EQUAL
%left '?' ':'
%right T_COALESCE
%left T_BOOLEAN_OR
%left T_BOOLEAN_AND
%left '|'
%left '^'
%left '&'
%nonassoc T_IS_EQUAL T_IS_NOT_EQUAL T_IS_IDENTICAL T_IS_NOT_IDENTICAL T_SPACESHIP
%nonassoc '<' T_IS_SMALLER_OR_EQUAL '>' T_IS_GREATER_OR_EQUAL
%left T_SL T_SR
%left '+' '-' '.'
%left '*' '/' '%'
%right '!'
%nonassoc T_INSTANCEOF
%right '~' T_INC T_DEC T_INT_CAST T_DOUBLE_CAST T_STRING_CAST T_ARRAY_CAST T_OBJECT_CAST T_BOOL_CAST T_UNSET_CAST '@'
%right T_POW
%right '['
%nonassoc T_NEW T_CLONE
%left T_NOELSE
%left T_ELSEIF
%left T_ELSE
%left T_ENDIF
%right T_STATIC T_ABSTRACT T_FINAL T_PRIVATE T_PROTECTED T_PUBLIC

%token <tok> T_ECHO
%token <tok> T_LNUMBER
%token <tok> T_DNUMBER
%token <tok> T_STRING
%token <tok> T_VARIABLE
%token <tok> T_INLINE_HTML
%token <tok> T_ENCAPSED_AND_WHITESPACE
%token <tok> T_CONSTANT_ENCAPSED_STRING
%token <tok> T_STRING_VARNAME
%token <tok> T_NUM_STRING
%token <tok> T_LINE
%token <tok> T_FILE
%token <tok> T_DIR
%token <tok> T_CLASS_C
%token <tok> T_TRAIT_C
%token <tok> T_METHOD_C
%token <tok> T_FUNC_C
%token <tok> T_NS_C

%token <tok> T_INCLUDE
%token <tok> T_INCLUDE_ONCE
%token <tok> T_EVAL
%token <tok> T_REQUIRE
%token <tok> T_REQUIRE_ONCE
%token <tok> T_LOGICAL_OR
%token <tok> T_LOGICAL_XOR
%token <tok> T_LOGICAL_AND
%token <tok> T_PRINT
%token <tok> T_YIELD
%token <tok> T_YIELD_FROM
%token T_PLUS_EQUAL
%token T_MINUS_EQUAL
%token T_MUL_EQUAL
%token T_DIV_EQUAL
%token T_CONCAT_EQUAL
%token T_MOD_EQUAL
%token T_AND_EQUAL
%token T_OR_EQUAL
%token T_XOR_EQUAL
%token T_SL_EQUAL
%token T_SR_EQUAL
%token T_BOOLEAN_OR
%token T_BOOLEAN_AND
%token T_IS_EQUAL
%token T_IS_NOT_EQUAL
%token T_IS_IDENTICAL
%token T_IS_NOT_IDENTICAL
%token T_IS_SMALLER_OR_EQUAL
%token T_IS_GREATER_OR_EQUAL
%token T_SPACESHIP
%token T_SL
%token T_SR
%token <tok> T_INSTANCEOF
%token T_INC
%token T_DEC
%token <tok> T_INT_CAST
%token <tok> T_DOUBLE_CAST
%token <tok> T_STRING_CAST
%token <tok> T_ARRAY_CAST
%token <tok> T_OBJECT_CAST
%token <tok> T_BOOL_CAST
%token <tok> T_UNSET_CAST
%token <tok> T_NEW
%token <tok> T_CLONE
%token <tok> T_EXIT
%token <tok> T_IF
%token <tok> T_ELSEIF
%token <tok> T_ELSE
%token <tok> T_ENDIF
%token <tok> T_DO
%token <tok> T_WHILE
%token <tok> T_ENDWHILE
%token <tok> T_FOR
%token <tok> T_ENDFOR
%token <tok> T_FOREACH
%token <tok> T_ENDFOREACH
%token <tok> T_DECLARE
%token <tok> T_ENDDECLARE
%token <tok> T_AS
%token <tok> T_SWITCH
%token <tok> T_ENDSWITCH
%token <tok> T_CASE
%token <tok> T_DEFAULT
%token <tok> T_BREAK
%token <tok> T_CONTINUE
%token <tok> T_GOTO
%token <tok> T_FUNCTION
%token <tok> T_CONST
%token <tok> T_RETURN
%token <tok> T_TRY
%token <tok> T_CATCH
%token <tok> T_FINALLY
%token <tok> T_THROW
%token <tok> T_USE
%token <tok> T_INSTEADOF
%token <tok> T_GLOBAL
%token <tok> T_STATIC
%token <tok> T_ABSTRACT
%token <tok> T_FINAL
%token <tok> T_PRIVATE
%token <tok> T_PROTECTED
%token <tok> T_PUBLIC
%token <tok> T_VAR
%token <tok> T_UNSET
%token <tok> T_ISSET
%token <tok> T_EMPTY
%token T_HALT_COMPILER
%token <tok> T_CLASS
%token <tok>T_TRAIT
%token <tok>T_INTERFACE
%token <tok> T_EXTENDS
%token <tok> T_IMPLEMENTS
%token T_OBJECT_OPERATOR
%token <tok> T_DOUBLE_ARROW
%token <tok> T_LIST
%token <tok> T_ARRAY
%token <tok> T_CALLABLE
%token T_COMMENT
%token T_DOC_COMMENT
%token T_OPEN_TAG
%token T_OPEN_TAG_WITH_ECHO
%token T_CLOSE_TAG
%token T_WHITESPACE
%token <tok> T_START_HEREDOC
%token <tok> T_END_HEREDOC
%token <tok> T_DOLLAR_OPEN_CURLY_BRACES
%token <tok> T_CURLY_OPEN
%token <tok> T_PAAMAYIM_NEKUDOTAYIM
%token <tok> T_NAMESPACE
%token <tok> T_NS_SEPARATOR
%token <tok> T_ELLIPSIS
%token T_COALESCE
%token T_POW
%token T_POW_EQUAL

/* Token used to force a parse error from the lexer */
%token T_ERROR

%type <program> start
%type <stmt> top_statement statement function_declaration_statement
%type <expr> namespace_name name
%type <stmt> class_declaration_statement trait_declaration_statement interface_declaration_statement
%type <exprs> interface_extends_list /*
%type <ast> group_use_declaration inline_use_declarations inline_use_declaration
%type <ast> mixed_group_use_declaration use_declaration unprefixed_use_declaration*/
%type <stmt> /*unprefixed_use_declarations */ inner_statement declare_statement finally_statement 
%type <expr> const_decl expr optional_expr foreach_variable
%type <expr> unset_variable variable
%type <expr> extends_from parameter optional_type argument expr_without_variable global_var
%type <expr> static_var /*trait_adaptation trait_precedence trait_alias*/
%type <expr> /*absolute_trait_method_reference trait_method_reference */property echo_expr
%type <expr> new_expr /*anonymous_class*/ class_name class_name_reference simple_variable
%type <expr> internal_functions_in_yacc
%type <expr> exit_expr scalar lexical_var function_call member_name property_name
%type <expr> variable_class_name dereferencable_scalar constant dereferencable
%type <expr> callable_expr callable_variable static_member new_variable
%type <expr> encaps_var encaps_var_offset 
%type <stmts> top_statement_list /*use_declarations */ inner_statement_list case_list catch_list class_statement_list
%type <stmt> foreach_statement for_statement while_statement alt_if_stmt alt_if_stmt_without_else switch_case_list class_statement 
%type <exprs> isset_variables backticks_expr echo_expr_list unset_variables catch_name_list parameter_list implements_list property_list
%type <stmt> if_stmt if_stmt_without_else
%type <expr> argument_list class_const_decl 
%type <exprs> const_list static_var_list global_var_list for_exprs non_empty_parameter_list non_empty_argument_list non_empty_for_exprs name_list
%type <ast> /* trait_adaptations method_body */
%type <expr> ctor_arguments /*trait_adaptation_list*/ lexical_vars
%type <exprs> lexical_var_list encaps_list class_const_list
%type <exprs> array_pair non_empty_array_pair_list array_pair_list possible_array_pair variable_modifiers non_empty_member_modifiers class_modifiers  
%type <expr> isset_variable type return_type type_expr member_modifier class_modifier 

%type <expr> identifier 

%type <tok> function reserved_non_modifiers semi_reserved
%type <num> returns_ref is_reference is_variadic
%type <exprs> method_modifiers/*
%type <num> use_type 
*/

%type <expr> backup_doc_comment backup_fn_flags

%% /* Rules */

start:
	top_statement_list	{ 
        $$ = &ast.Program{Statements: $1}; 
        if l, ok := yylex.(*LexerWrapper); ok {
            l.program = $$
        }
    }
;

reserved_non_modifiers:
	  T_INCLUDE | T_INCLUDE_ONCE | T_EVAL | T_REQUIRE | T_REQUIRE_ONCE | T_LOGICAL_OR | T_LOGICAL_XOR | T_LOGICAL_AND
	| T_INSTANCEOF | T_NEW | T_CLONE | T_EXIT | T_IF | T_ELSEIF | T_ELSE | T_ENDIF | T_ECHO | T_DO | T_WHILE | T_ENDWHILE
	| T_FOR | T_ENDFOR | T_FOREACH | T_ENDFOREACH | T_DECLARE | T_ENDDECLARE | T_AS | T_TRY | T_CATCH | T_FINALLY
	| T_THROW | T_USE | T_INSTEADOF | T_GLOBAL | T_VAR | T_UNSET | T_ISSET | T_EMPTY | T_CONTINUE | T_GOTO
	| T_FUNCTION | T_CONST | T_RETURN | T_PRINT | T_YIELD | T_LIST | T_SWITCH | T_ENDSWITCH | T_CASE | T_DEFAULT | T_BREAK
	| T_ARRAY | T_CALLABLE | T_EXTENDS | T_IMPLEMENTS | T_NAMESPACE | T_TRAIT | T_INTERFACE | T_CLASS
	| T_CLASS_C | T_TRAIT_C | T_FUNC_C | T_METHOD_C | T_LINE | T_FILE | T_DIR | T_NS_C
;

semi_reserved:
	  reserved_non_modifiers
	| T_STATIC | T_ABSTRACT | T_FINAL | T_PRIVATE | T_PROTECTED | T_PUBLIC
;

identifier:
               T_STRING { $$ = ast.NewStringLiteral($1, $1.Literal); }
       |       semi_reserved  { $$ = ast.NewStringLiteral($1, $1.Literal); }
;

top_statement_list:
		top_statement_list top_statement { $$ = append($1, $2); }
	|	/* empty */ { $$ = []ast.Statement{} }
;


namespace_name:
		T_STRING								{ $$ = ast.NewStringLiteral($1, $1.Literal); }
	|	namespace_name T_NS_SEPARATOR T_STRING	{ $$ = ast.NewNamespaceExpression(nil, nil, $1, ast.NewStringLiteral($3, $3.Literal)) }
;

name:
		namespace_name								{ $$ = $1; }
	|	T_NAMESPACE T_NS_SEPARATOR namespace_name	{ $$ = ast.NewNamespaceExpression($1, $2, $3) }
	|	T_NS_SEPARATOR namespace_name				{ $$ = ast.NewNamespaceExpression(nil, $1, $2) }
;

top_statement:
		statement							{ $$ = $1; }
	|	function_declaration_statement		{ $$ = $1; }
	|	class_declaration_statement			{ $$ = $1; }
	|	trait_declaration_statement			{ $$ = $1; }
	|	interface_declaration_statement		{ $$ = $1; }/*
	|	T_HALT_COMPILER '(' ')' ';'
			{ $$ = zend_ast_create(ZEND_AST_HALT_COMPILER,
			      zend_ast_create_zval_from_long(zend_get_scanned_file_offset()));
			  zend_stop_lexing(); }
	|	T_NAMESPACE namespace_name ';'
			{ $$ = zend_ast_create(ZEND_AST_NAMESPACE, $2, NULL);
			  RESET_DOC_COMMENT(); }
	|	T_NAMESPACE namespace_name { RESET_DOC_COMMENT(); }
		'{' top_statement_list '}'
			{ $$ = zend_ast_create(ZEND_AST_NAMESPACE, $2, $5); }
	|	T_NAMESPACE { RESET_DOC_COMMENT(); }
		'{' top_statement_list '}'
			{ $$ = zend_ast_create(ZEND_AST_NAMESPACE, NULL, $4); }
	|	T_USE mixed_group_use_declaration ';'		{ $$ = $2; }
	|	T_USE use_type group_use_declaration ';'	{ $$ = $3; $$->attr = $2; }
	|	T_USE use_declarations ';'					{ $$ = $2; $$->attr = ZEND_SYMBOL_CLASS; }
	|	T_USE use_type use_declarations ';'			{ $$ = $3; $$->attr = $2; }
	|	T_CONST const_list ';'						{ $$ = $2; }*/
;

/*
use_type:
	 	T_FUNCTION 		{ $$ = ZEND_SYMBOL_FUNCTION; }
	| 	T_CONST 		{ $$ = ZEND_SYMBOL_CONST; }
;

group_use_declaration:
		namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
			{ $$ = zend_ast_create(ZEND_AST_GROUP_USE, $1, $4); }
	|	T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' unprefixed_use_declarations possible_comma '}'
			{ $$ = zend_ast_create(ZEND_AST_GROUP_USE, $2, $5); }
;

mixed_group_use_declaration:
		namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
			{ $$ = zend_ast_create(ZEND_AST_GROUP_USE, $1, $4);}
	|	T_NS_SEPARATOR namespace_name T_NS_SEPARATOR '{' inline_use_declarations possible_comma '}'
			{ $$ = zend_ast_create(ZEND_AST_GROUP_USE, $2, $5); }
;

possible_comma:
		/* empty *//*
	|	','
;

inline_use_declarations:
		inline_use_declarations ',' inline_use_declaration
			{ $$ = zend_ast_list_add($1, $3); }
	|	inline_use_declaration
			{ $$ = zend_ast_create_list(1, ZEND_AST_USE, $1); }
;

unprefixed_use_declarations:
		unprefixed_use_declarations ',' unprefixed_use_declaration
			{ $$ = zend_ast_list_add($1, $3); }
	|	unprefixed_use_declaration
			{ $$ = zend_ast_create_list(1, ZEND_AST_USE, $1); }
;

use_declarations:
		use_declarations ',' use_declaration
			{ $$ = zend_ast_list_add($1, $3); }
	|	use_declaration
			{ $$ = zend_ast_create_list(1, ZEND_AST_USE, $1); }
;

inline_use_declaration:
		unprefixed_use_declaration { $$ = $1; $$->attr = ZEND_SYMBOL_CLASS; }
	|	use_type unprefixed_use_declaration { $$ = $2; $$->attr = $1; }
;

unprefixed_use_declaration:
		namespace_name
			{ $$ = zend_ast_create(ZEND_AST_USE_ELEM, $1, NULL); }
	|	namespace_name T_AS T_STRING
			{ $$ = zend_ast_create(ZEND_AST_USE_ELEM, $1, $3); }
;

use_declaration:
		unprefixed_use_declaration                { $$ = $1; }
	|	T_NS_SEPARATOR unprefixed_use_declaration { $$ = $2; }
;
*/

const_list:
		const_list ',' const_decl { $$ = append($1, $3); }
	|	const_decl { $$ = append($$, $1); }
;

inner_statement_list:
		inner_statement_list inner_statement
			{ $$ = append($1, $2); }
	|	/* empty */
			{ $$ = []ast.Statement{}; }
;

inner_statement:
		statement { $$ = $1; }
	|	function_declaration_statement 		{ $$ = $1; }
	|	class_declaration_statement 		{ $$ = $1; }

	|	trait_declaration_statement			{ $$ = $1; }
	|	interface_declaration_statement		{ $$ = $1; }/*
	|	T_HALT_COMPILER '(' ')' ';'
			{ $$ = NULL; zend_error_noreturn(E_COMPILE_ERROR,
			      "__HALT_COMPILER() can only be used from the outermost scope"); }*/
;

statement:
	    '{' inner_statement_list '}' { $$ = ast.NewBlockStatement($2...); }
	|	if_stmt { $$ = $1; }
	|	alt_if_stmt { $$ = $1; }
	|	T_WHILE '(' expr ')' while_statement
			{ $$ = ast.NewWhileStatement($1, $3, $5); }
	|	T_DO statement T_WHILE '(' expr ')' ';'
			{ $$ = ast.NewDoWhileStatement($1, $5, $2); }
	|	T_FOR '(' for_exprs ';' for_exprs ';' for_exprs ')' for_statement
			{ $$ = ast.NewForStatement($1, $3, $5, $7, $9); }
	|	T_SWITCH '(' expr ')' switch_case_list
			{ $$ = ast.NewSwitchStatement($1, $3, $5); }
	|	T_BREAK optional_expr ';'		{ $$ = ast.NewBreakStatement($1, $2); }
	|	T_CONTINUE optional_expr ';'	{ $$ = ast.NewContinueStatement($1, $2); }
	|	T_RETURN optional_expr ';'		{ $$ = ast.NewReturnStatement($1, $2); }
	|	T_GLOBAL global_var_list ';'	{ $$ = ast.NewGlobalStatement($1, $2); }
	|	T_STATIC static_var_list ';'	{ $$ = ast.NewStaticStatement($1, $2); }
	|	T_ECHO echo_expr_list ';'		{ $$ = ast.NewEchoStatement($1, $2); }
	|	T_INLINE_HTML { $$ = ast.NewInlineHTMLStatement($1); }
	|	expr ';' { $$ = ast.NewExpressionStatement($1); }
	|	T_UNSET '(' unset_variables ')' ';' { $$ = ast.NewUnsetStatement($1, $3); }
	|	T_FOREACH '(' expr T_AS foreach_variable ')' foreach_statement
			{ $$ = ast.NewForeachStatement($1, $3, nil, $5, $7); }
	|	T_FOREACH '(' expr T_AS foreach_variable T_DOUBLE_ARROW foreach_variable ')' foreach_statement
			{ $$ = ast.NewForeachStatement($1, $3, $5, $7, $9); }
	|	T_DECLARE '(' const_list ')' declare_statement
			{ $$ = ast.NewDeclareStatement($1, $3, $5); }
	|	';'	/* empty statement */ { $$ = ast.NewEmptyStatement(); }
	|	T_TRY '{' inner_statement_list '}' catch_list finally_statement
			{ $$ = ast.NewTryStatement($1, $3, $5, $6); }
	|	T_THROW expr ';' { $$ = ast.NewThrowStatement($1, $2); }
	|	T_GOTO T_STRING ';' { $$ = ast.NewGotoStatement($1, ast.NewStringLiteral($2, $2.Literal)); }
	|	T_STRING ':' { $$ = ast.NewLabelStatement(ast.NewStringLiteral($1, $1.Literal)); }
;

catch_list:
		/* empty */
			{ $$ = []ast.Statement{}; }
	|	catch_list T_CATCH '(' catch_name_list T_VARIABLE ')' '{' inner_statement_list '}'
			{ $$ = append($$, ast.NewCatchListStatement($2, $4, ast.NewVariableLiteral($5, $5.Literal), $8)); }
;

catch_name_list:
		name { $$ = append($$, $1); }
	|	catch_name_list '|' name { $$ = append($1, $3); }
;

finally_statement:
		/* empty */ { $$ = nil; }
	|	T_FINALLY '{' inner_statement_list '}' { $$ = ast.NewFinallyStatement($1, $3); }
;

unset_variables:
		unset_variable { $$ = append($$, $1); }
	|	unset_variables ',' unset_variable { $$ = append($1, $3); }
;

unset_variable:
		variable { $$ = $1; }
;

function_declaration_statement:
	function returns_ref T_STRING backup_doc_comment '(' parameter_list ')' return_type
	backup_fn_flags '{' inner_statement_list '}' backup_fn_flags
		{ $$ = ast.NewFunctionStatement($1, $2, $3, $6, $8, $11); }
;

is_reference:
		/* empty */	{ $$ = 0; }
	|	'&'			{ $$ = 1; }
;

is_variadic:
		/* empty */ { $$ = 0; }
	|	T_ELLIPSIS  { $$ = 1; }
;

class_declaration_statement:
		class_modifiers T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
			{ $$ = ast.NewClassDeclarationStatement($1, $2, $3, $4, $5, $8); }
	|	T_CLASS T_STRING extends_from implements_list backup_doc_comment '{' class_statement_list '}'
			{ $$ = ast.NewClassDeclarationStatement([]ast.Expression{}, $1, $2, $3, $4, $7); }
;

class_modifiers:
		class_modifier 					{ $$ = append($$, $1); }
	|	class_modifiers class_modifier 	{ $$ = append($$, $2); }
;

class_modifier:
		T_ABSTRACT 		{ $$ = ast.NewAbstractLiteral($1, $1.Literal); }
	|	T_FINAL 		{ $$ = ast.NewFinalLiteral($1, $1.Literal); }
;


trait_declaration_statement:
		T_TRAIT T_STRING backup_doc_comment '{' class_statement_list '}'
			{ $$ = ast.NewTraitDeclarationStatement($1, $2, $5); }
;

interface_declaration_statement:
		T_INTERFACE T_STRING interface_extends_list backup_doc_comment '{' class_statement_list '}'
			{ $$ = ast.NewInterfaceDeclarationStatement($1, $2, $3, $6); }
;

extends_from:
		/* empty */		{ $$ = nil; }
	|	T_EXTENDS name	{ $$ = $2; }
;

interface_extends_list:
		/* empty */			{ $$ = []ast.Expression{}; }
	|	T_EXTENDS name_list	{ $$ = $2; }
;

implements_list:
		/* empty */				{ $$ = []ast.Expression{}; }
	|	T_IMPLEMENTS name_list	{ $$ = $2; }
;

foreach_variable:
		variable			{ $$ = $1; }
	|	'&' variable		{ $$ = ast.NewAmpersandLiteral($2); }
	|	T_LIST '(' array_pair_list ')' { $$ = ast.NewListExpression($1, $3...); }
	|	'[' array_pair_list ']' { $$ = ast.NewArrayExpression(ast.Short, $2...); }
;

for_statement:
		statement { $$ = $1; }
	|	':' inner_statement_list T_ENDFOR ';' { $$ = ast.NewAltForStatement($2); }
;

foreach_statement:
		statement { $$ = $1; }
	|	':' inner_statement_list T_ENDFOREACH ';' { $$ = ast.NewAltForeachStatement($2); }
;

declare_statement:
		statement { $$ = $1; }
	|	':' inner_statement_list T_ENDDECLARE ';' { $$ = ast.NewAltDeclareStatement($2); }
;

switch_case_list:
		'{' case_list '}'					{ $$ = ast.NewSwitchCaseListStatement($2, false); }
	|	'{' ';' case_list '}'				{ $$ = ast.NewSwitchCaseListStatement($3, true); }
	|	':' case_list T_ENDSWITCH ';'		{ $$ = ast.NewAltSwitchCaseListStatement($2, false); }
	|	':' ';' case_list T_ENDSWITCH ';'	{ $$ = ast.NewAltSwitchCaseListStatement($3, true); }
;

case_list:
		/* empty */ { $$ = []ast.Statement{}; }
	|	case_list T_CASE expr ':' inner_statement_list
			{ $$ = append($1, ast.NewCaseListStatement($2, $3, $5, false)); }
	|	case_list T_CASE expr ';' inner_statement_list
			{ $$ = append($1, ast.NewCaseListStatement($2, $3, $5, true)); }
	|	case_list T_DEFAULT ':' inner_statement_list
			{ $$ = append($1, ast.NewCaseListStatement($2, nil, $4, false)); }
	|	case_list T_DEFAULT ';' inner_statement_list
			{ $$ = append($1, ast.NewCaseListStatement($2, nil, $4, true)); }
;

while_statement:
		statement { $$ = $1; }
	|	':' inner_statement_list T_ENDWHILE ';' { $$ = ast.NewAltWhileStatement($3, $2...); }
;

if_stmt_without_else:
		T_IF '(' expr ')' statement
			{ $$ = ast.NewIfStatement($1, $3, $5, nil); }
	|	if_stmt_without_else T_ELSEIF '(' expr ')' statement
			{ $$ = ast.NewIfStatement($2, $4, $6, $$); }
;

if_stmt:
		if_stmt_without_else %prec T_NOELSE { $$ = $1; }
	|	if_stmt_without_else T_ELSE statement
			{ $$ = ast.NewIfStatement($2, nil, $3, $$); }
;


alt_if_stmt_without_else:
		T_IF '(' expr ')' ':' inner_statement_list
			{ $$ = ast.NewAltIfStatement($1, $3, $6, nil); }
	|	alt_if_stmt_without_else T_ELSEIF '(' expr ')' ':' inner_statement_list
			{ $$ = ast.NewAltIfStatement($2, $4, $7, $$); }
;

alt_if_stmt:
		alt_if_stmt_without_else T_ENDIF ';' 
            { 
                stmt := $1; 
                $$ = ast.NewAltIfStatement($2, nil, nil, stmt); 
            }
	|	alt_if_stmt_without_else T_ELSE ':' inner_statement_list T_ENDIF ';'
			{ 
                stmt := ast.NewAltIfStatement($2, nil, $4, $$);
                $$ = ast.NewAltIfStatement($5, nil, nil, stmt);
            }
;

parameter_list:
		non_empty_parameter_list { $$ = $1; }
	|	/* empty */	{ $$ = []ast.Expression{}; }
;

non_empty_parameter_list:
		parameter
			{ $$ = append($$, $1);}
	|	non_empty_parameter_list ',' parameter
			{ $$ = append($1, $3); }
;

parameter:
		optional_type is_reference is_variadic T_VARIABLE
			{ $$ = ast.NewParameterExpression($1, $2, $3, ast.NewVariableLiteral($4, $4.Literal), nil); }
	|	optional_type is_reference is_variadic T_VARIABLE '=' expr
			{ $$ = ast.NewParameterExpression($1, $2, $3, ast.NewVariableLiteral($4, $4.Literal), $6); }
;

optional_type:
		/* empty */	{ $$ = nil; }
	|	type_expr	{ $$ = $1; }
;

type_expr:
		type		{ $$ = $1; }
	|	'?' type	{ $$ = ast.NewOptionalTypeExpression($2); }
;

type:
		T_ARRAY		{ $$ = ast.NewTypeExpression($1); }
	|	T_CALLABLE	{ $$ = ast.NewTypeExpression($1); }
	|	name		{ $$ = $1; }
;

return_type:
		/* empty */	{ $$ = nil; }
	|	':' type_expr	{ $$ = ast.NewReturnTypeExpression($2); }
;

argument_list:
		'(' ')'	{ $$ = ast.NewArgumentListExpression(); }
	|	'(' non_empty_argument_list ')' { $$ = ast.NewArgumentListExpression($2...); }
;

non_empty_argument_list:
		argument
			{ $$ = append($$, $1); }
	|	non_empty_argument_list ',' argument
			{ $$ = append($1, $3); }
;

argument:
		expr			{ $$ = $1; }
	|	T_ELLIPSIS expr	{ $$ = ast.NewArgumentExpression($1, $2); }
;

global_var_list:
		global_var_list ',' global_var { $$ = append($1, $3); }
	|	global_var { $$ = append($$, $1); }
;

global_var:
	simple_variable
		{ $$ = $1; }
;

static_var_list:
		static_var_list ',' static_var { $$ = append($1, $3); }
	|	static_var { $$ = append($$, $1); }
;

static_var:
		T_VARIABLE			{ $$ = ast.NewVariableLiteral($1, $1.Literal); }
	|	T_VARIABLE '=' expr	{ $$ = ast.NewAssignExpression(ast.Equal, ast.NewVariableLiteral($1, $1.Literal), $3, false); }
;

class_statement_list:
		class_statement_list class_statement
			{ $$ = append($$, $2); }
	|	/* empty */
			{ $$ = []ast.Statement{} }
;

class_statement:
		variable_modifiers property_list ';'
			{ $$ = ast.NewClassStatement($1, $2); }
	|	method_modifiers T_CONST class_const_list ';'
			{ $$ = ast.NewClassStatement(append($1, ast.NewConstantEncapsedStringLiteral($2, $2.Literal)), $3); }/*
	|	T_USE name_list trait_adaptations
			{ $$ = zend_ast_create(ZEND_AST_USE_TRAIT, $2, $3); }
	|	method_modifiers function returns_ref identifier backup_doc_comment '(' parameter_list ')'
		return_type backup_fn_flags method_body backup_fn_flags
			{ $$ = zend_ast_create_decl(ZEND_AST_METHOD, $3 | $1 | $12, $2, $5,
				  zend_ast_get_str($4), $7, NULL, $11, $9); CG(extra_fn_flags) = $10; }*/
;

name_list:
		name { $$ = append($$, $1); }
	|	name_list ',' name { $$ = append($1, $3); }
;

/*
trait_adaptations:
		';'								{ $$ = NULL; }
	|	'{' '}'							{ $$ = NULL; }
	|	'{' trait_adaptation_list '}'	{ $$ = $2; }
;

trait_adaptation_list:
		trait_adaptation
			{ $$ = zend_ast_create_list(1, ZEND_AST_TRAIT_ADAPTATIONS, $1); }
	|	trait_adaptation_list trait_adaptation
			{ $$ = zend_ast_list_add($1, $2); }
;

trait_adaptation:
		trait_precedence ';'	{ $$ = $1; }
	|	trait_alias ';'			{ $$ = $1; }
;

trait_precedence:
	absolute_trait_method_reference T_INSTEADOF name_list
		{ $$ = zend_ast_create(ZEND_AST_TRAIT_PRECEDENCE, $1, $3); }
;

trait_alias:
		trait_method_reference T_AS T_STRING
			{ $$ = zend_ast_create_ex(ZEND_AST_TRAIT_ALIAS, 0, $1, $3); }
	|	trait_method_reference T_AS reserved_non_modifiers
			{ zval zv; zend_lex_tstring(&zv); $$ = zend_ast_create_ex(ZEND_AST_TRAIT_ALIAS, 0, $1, zend_ast_create_zval(&zv)); }
	|	trait_method_reference T_AS member_modifier identifier
			{ $$ = zend_ast_create_ex(ZEND_AST_TRAIT_ALIAS, $3, $1, $4); }
	|	trait_method_reference T_AS member_modifier
			{ $$ = zend_ast_create_ex(ZEND_AST_TRAIT_ALIAS, $3, $1, NULL); }
;

trait_method_reference:
		identifier
			{ $$ = zend_ast_create(ZEND_AST_METHOD_REFERENCE, NULL, $1); }
	|	absolute_trait_method_reference { $$ = $1; }
;

absolute_trait_method_reference:
	name T_PAAMAYIM_NEKUDOTAYIM identifier
		{ $$ = zend_ast_create(ZEND_AST_METHOD_REFERENCE, $1, $3); }
;

method_body:
		';' /* abstract method *//*		{ $$ = NULL; }
	|	'{' inner_statement_list '}'	{ $$ = $2; }
;
*/

variable_modifiers:
		non_empty_member_modifiers		{ $$ = $1; }
	|	T_VAR							{ $$ = []ast.Expression{ast.NewVarLiteral($1, $1.Literal)}; }
;

method_modifiers:
		/* empty */						{ $$ = []ast.Expression{}; }
	|	non_empty_member_modifiers
			{ $$ = $1 }
;

non_empty_member_modifiers:
		member_modifier			{ $$ = []ast.Expression{$1}; }
	|	non_empty_member_modifiers member_modifier
			{ $$ = append($1, $2); }
;

member_modifier:
		T_PUBLIC				{ $$ = ast.NewPublicLiteral($1, $1.Literal); }
	|	T_PROTECTED				{ $$ = ast.NewProtectedLiteral($1, $1.Literal);; }
	|	T_PRIVATE				{ $$ = ast.NewPrivateLiteral($1, $1.Literal);; }
	|	T_STATIC				{ $$ = ast.NewStaticLiteral($1, $1.Literal);; }
	|	T_ABSTRACT				{ $$ = ast.NewAbstractLiteral($1, $1.Literal); }
	|	T_FINAL					{ $$ = ast.NewFinalLiteral($1, $1.Literal); }
;

property_list:
		property_list ',' property { $$ = append($1, $3); }
	|	property { $$ = append($$, $1); }
;

property:
		T_VARIABLE backup_doc_comment
			{ $$ = ast.NewVariableLiteral($1, $1.Literal); }
	|	T_VARIABLE '=' expr backup_doc_comment
			{ $$ = ast.NewAssignExpression(ast.Equal, ast.NewVariableLiteral($1, $1.Literal), $3, false); }
;

class_const_list:
		class_const_list ',' class_const_decl { $$ = append($1, $3); }
	|	class_const_decl { $$ = append($$, $1); }
;

class_const_decl:
	identifier '=' expr backup_doc_comment { $$ = ast.NewAssignExpression(ast.Equal, $1, $3, false); }
;

const_decl:
	T_STRING '=' expr backup_doc_comment { $$ = ast.NewAssignExpression(ast.Equal, ast.NewStringLiteral($1, $1.Literal), $3, false); }
;

echo_expr_list:
		echo_expr_list ',' echo_expr { $$ = append($1, $3); }
	|	echo_expr { $$ = append($$, $1); }
;

echo_expr:
	expr { $$ = $1 }
;

for_exprs:
		/* empty */			{ $$ = []ast.Expression{}; }
	|	non_empty_for_exprs	{ $$ = $1; }
;

non_empty_for_exprs:
		non_empty_for_exprs ',' expr { $$ = append($1, $3); }
	|	expr { $$ = []ast.Expression{$1}; }
;

/*
anonymous_class:
        T_CLASS { $<num>$ = CG(zend_lineno); } ctor_arguments
		extends_from implements_list backup_doc_comment '{' class_statement_list '}' {
			zend_ast *decl = zend_ast_create_decl(
				ZEND_AST_CLASS, ZEND_ACC_ANON_CLASS, $<num>2, $6, NULL,
				$4, $5, $8, NULL);
			$$ = zend_ast_create(ZEND_AST_NEW, decl, $3);
		}
;
*/

new_expr:
		T_NEW class_name_reference ctor_arguments
			{ $$ = ast.NewNewExpression($1, $2, $3); }/*
	|	T_NEW anonymous_class
			{ $$ = $2; }*/
;

expr_without_variable:
		T_LIST '(' array_pair_list ')' '=' expr
			{ $$ = ast.NewAssignExpression(ast.Equal, ast.NewListExpression($1, $3...), $6, false); }
	|	'[' array_pair_list ']' '=' expr
			{ $$ = ast.NewAssignExpression(ast.Equal, ast.NewArrayExpression(ast.Short, $2...), $5, false); }
	|	variable '=' expr
			{ $$ = ast.NewAssignExpression(ast.Equal, $1, $3, false); }
	|	variable '=' '&' variable
			{ $$ = ast.NewAssignExpression(ast.Equal, $1, $4, true); }
	|	T_CLONE expr { $$ = ast.NewCloneExpression($1, $2); }
	|	variable T_PLUS_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.PlusEqual, $1, $3, false); }
	|	variable T_MINUS_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.MinusEqual, $1, $3, false); }
	|	variable T_MUL_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.MulEqual, $1, $3, false); }
	|	variable T_POW_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.PowEqual, $1, $3, false); }
	|	variable T_DIV_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.DivEqual, $1, $3, false); }
	|	variable T_CONCAT_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.ConcatEqual, $1, $3, false); }
	|	variable T_MOD_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.ModEqual, $1, $3, false); }
	|	variable T_AND_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.AndEqual, $1, $3, false); }
	|	variable T_OR_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.QrEqual, $1, $3, false); }
	|	variable T_XOR_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.XorEqual, $1, $3, false); }
	|	variable T_SL_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.SlEqual, $1, $3, false); }
	|	variable T_SR_EQUAL expr
			{ $$ = ast.NewAssignExpression(ast.SrEqual, $1, $3, false); }
	|	variable T_INC { $$ = ast.NewIncrementExpression(ast.PostInc, $1); }
	|	T_INC variable { $$ = ast.NewIncrementExpression(ast.PreInc, $2); }
	|	variable T_DEC { $$ = ast.NewDecrementExpression(ast.PostDec, $1); }
	|	T_DEC variable { $$ = ast.NewDecrementExpression(ast.PreDec, $2); }
	|	expr T_BOOLEAN_OR expr
			{ $$ = ast.NewInfixExpression(ast.BooleanOr, $1, $3); }
	|	expr T_BOOLEAN_AND expr
			{ $$ = ast.NewInfixExpression(ast.BooleanAnd, $1, $3); }
	|	expr T_LOGICAL_OR expr
			{ $$ = ast.NewInfixExpression(ast.LogicalOr, $1, $3); }
	|	expr T_LOGICAL_AND expr
			{ $$ = ast.NewInfixExpression(ast.LogicalAnd, $1, $3); }
	|	expr T_LOGICAL_XOR expr
			{ $$ = ast.NewInfixExpression(ast.LogicalXor, $1, $3); }
	|	expr '|' expr	{ $$ = ast.NewInfixExpression(ast.BwOr, $1, $3); }
	|	expr '&' expr	{ $$ = ast.NewInfixExpression(ast.BwAnd, $1, $3); }
	|	expr '^' expr	{ $$ = ast.NewInfixExpression(ast.BwXor, $1, $3); }
	|	expr '.' expr 	{ $$ = ast.NewInfixExpression(ast.Concat, $1, $3); }
	|	expr '+' expr 	{ $$ = ast.NewInfixExpression(ast.Add, $1, $3); }
	|	expr '-' expr 	{ $$ = ast.NewInfixExpression(ast.Sub, $1, $3); }
	|	expr '*' expr	{ $$ = ast.NewInfixExpression(ast.Mul, $1, $3); }
	|	expr T_POW expr	{ $$ = ast.NewInfixExpression(ast.Pow, $1, $3); }
	|	expr '/' expr	{ $$ = ast.NewInfixExpression(ast.Div, $1, $3); }
	|	expr '%' expr 	{ $$ = ast.NewInfixExpression(ast.Mod, $1, $3); }
	| 	expr T_SL expr	{ $$ = ast.NewInfixExpression(ast.Sl, $1, $3); }
	|	expr T_SR expr	{ $$ = ast.NewInfixExpression(ast.Sr, $1, $3); }
	|	'+' expr %prec T_INC { $$ = ast.NewPrefixExpression(ast.UnaryPlus, $2); }
	|	'-' expr %prec T_INC { $$ = ast.NewPrefixExpression(ast.UnaryMinus, $2); }
	|	'!' expr { $$ = ast.NewPrefixExpression(ast.BoolNot, $2); }
	|	'~' expr { $$ = ast.NewPrefixExpression(ast.BwNot, $2); }
	|	expr T_IS_IDENTICAL expr
			{ $$ = ast.NewInfixExpression(ast.IsIdentical, $1, $3); }
	|	expr T_IS_NOT_IDENTICAL expr
			{ $$ = ast.NewInfixExpression(ast.IsNotIdentical, $1, $3); }
	|	expr T_IS_EQUAL expr
			{ $$ = ast.NewInfixExpression(ast.IsEqual, $1, $3); }
	|	expr T_IS_NOT_EQUAL expr
			{ $$ = ast.NewInfixExpression(ast.IsNotEqual, $1, $3); }
	|	expr '<' expr
			{ $$ = ast.NewInfixExpression(ast.Smaller, $1, $3); }
	|	expr T_IS_SMALLER_OR_EQUAL expr
			{ $$ = ast.NewInfixExpression(ast.SmallerOrEqual, $1, $3); }
	|	expr '>' expr
			{ $$ = ast.NewInfixExpression(ast.Greater, $1, $3); }
	|	expr T_IS_GREATER_OR_EQUAL expr
			{ $$ = ast.NewInfixExpression(ast.GreaterOrEqual, $1, $3); }
	|	expr T_SPACESHIP expr
			{ $$ = ast.NewInfixExpression(ast.Spaceship, $1, $3); }
	|	expr T_INSTANCEOF class_name_reference
			{ $$ = ast.NewInfixExpression(ast.InstanceOf, $1, $3); }
	|	'(' expr ')' { $$ = ast.NewDereferencableExpression(ast.Wrapped, $2); }
	|	new_expr { $$ = $1; }
	|	expr '?' expr ':' expr
			{ $$ = ast.NewTernaryOperatorExpression($1, $3, $5); }
	|	expr '?' ':' expr
			{ $$ = ast.NewTernaryOperatorExpression($1, nil, $4); }
	|	expr T_COALESCE expr
			{ $$ = ast.NewInfixExpression(ast.Coalesce, $1, $3); }
	|	internal_functions_in_yacc { $$ = $1; }
	|	T_INT_CAST expr		{ $$ = ast.NewCastExpression($1, $2); }
	|	T_DOUBLE_CAST expr	{ $$ = ast.NewCastExpression($1, $2); }
	|	T_STRING_CAST expr	{ $$ = ast.NewCastExpression($1, $2); }
	|	T_ARRAY_CAST expr	{ $$ = ast.NewCastExpression($1, $2); }
	|	T_OBJECT_CAST expr	{ $$ = ast.NewCastExpression($1, $2); }
	|	T_BOOL_CAST expr	{ $$ = ast.NewCastExpression($1, $2); }
	|	T_UNSET_CAST expr	{ $$ = ast.NewCastExpression($1, $2); }
	|	T_EXIT exit_expr	{ $$ = ast.NewExitExpression($1, $2); }
	|	'@' expr			{ $$ = ast.NewPrefixExpression(ast.Silence, $2); }
	|	scalar { $$ = $1; }
	|	'`' backticks_expr '`' { $$ = ast.NewBackticksExpression($2...); }
	|	T_PRINT expr { $$ = ast.NewPrintExpression($1, $2); }
	|	T_YIELD { $$ = ast.NewYieldExpression($1, nil); }
	|	T_YIELD expr { $$ = ast.NewYieldExpression($1, $2); }
	|	T_YIELD expr T_DOUBLE_ARROW expr { $$ = ast.NewYieldExpression($1, ast.NewArrayPairExpression($2, $4, false)); }
	|	T_YIELD_FROM expr { $$ = ast.NewYieldExpression($1, $2); }
	|	function returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars return_type
		backup_fn_flags '{' inner_statement_list '}' backup_fn_flags
			{ $$ = ast.NewFunctionExpression($1, $2, $5, $7, $8, $11, false); }
	|	T_STATIC function returns_ref backup_doc_comment '(' parameter_list ')' lexical_vars
		return_type backup_fn_flags '{' inner_statement_list '}' backup_fn_flags
			{ $$ = ast.NewFunctionExpression($2, $3, $6, $8, $9, $12, true); }
;

function:
	T_FUNCTION { $$ = $1; }
;

backup_doc_comment:
	/* empty */ { $$ = nil; }
;

backup_fn_flags:
	/* empty */ { $$ = nil; }
;

returns_ref:
		/* empty */	{ $$ = 0; }
	|	'&'			{ $$ = 1; }
;

lexical_vars:
		/* empty */ { $$ = nil; }
	|	T_USE '(' lexical_var_list ')' { $$ = ast.NewLexicalVariableListExpression($1, $3...); }
;

lexical_var_list:
		lexical_var_list ',' lexical_var { $$ = append($1, $3); }
	|	lexical_var { $$ = append($$, $1); }
;

lexical_var:
		T_VARIABLE		{ $$ = ast.NewVariableLiteral($1, $1.Literal); }
	|	'&' T_VARIABLE	{ $$ = ast.NewAmpersandLiteral(ast.NewVariableLiteral($2, $2.Literal)) }
;

function_call:
		name argument_list
			{ $$ = ast.NewFunctionCallExpression(ast.Call, $1, nil, $2); }
	|	class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
			{ $$ = ast.NewFunctionCallExpression(ast.StaticCall, $1, $3, $4); }
	|	variable_class_name T_PAAMAYIM_NEKUDOTAYIM member_name argument_list
			{ $$ = ast.NewFunctionCallExpression(ast.StaticCall, $1, $3, $4); }
	|	callable_expr argument_list
			{ $$ = ast.NewFunctionCallExpression(ast.Call, $1, nil, $2); }
;

class_name:
		T_STATIC { $$ = ast.NewStaticLiteral($1, $1.Literal) }
	|	name { $$ = $1; }
;

class_name_reference:
		class_name		{ $$ = $1; }
	|	new_variable	{ $$ = $1; }
;

exit_expr:
		/* empty */				{ $$ = nil; }
	|	'(' optional_expr ')'	{ $$ = ast.NewDereferencableExpression(ast.Wrapped, $2); }
;

backticks_expr:
		/* empty */
			{ $$ = []ast.Expression{}; }
	|	T_ENCAPSED_AND_WHITESPACE { $$ = []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral($1, $1.Literal)}; }
	|	encaps_list { $$ = $1; }
;

ctor_arguments:
		/* empty */	{ $$ = nil; }
	|	argument_list { $$ = $1; }
;

dereferencable_scalar:
		T_ARRAY '(' array_pair_list ')'	{ $$ = ast.NewArrayExpression(ast.Long, $3...); }
	|	'[' array_pair_list ']'			{ $$ = ast.NewArrayExpression(ast.Short, $2...); }
	|	T_CONSTANT_ENCAPSED_STRING		{ $$ = ast.NewConstantEncapsedStringLiteral($1, $1.Literal); }
;

scalar:
		T_LNUMBER 	{ $$ = ast.NewIntegerLiteral($1, $1.Literal); }
	|	T_DNUMBER 	{ $$ = ast.NewDoubleLiteral($1, $1.Literal); }
	|	T_LINE 		{ $$ = ast.NewMagicConstLiteral($1, $1.Literal); }
	|	T_FILE 		{ $$ = ast.NewMagicConstLiteral($1, $1.Literal); }
	|	T_DIR   	{ $$ = ast.NewMagicConstLiteral($1, $1.Literal); }
	|	T_TRAIT_C	{ $$ = ast.NewMagicConstLiteral($1, $1.Literal); }
	|	T_METHOD_C	{ $$ = ast.NewMagicConstLiteral($1, $1.Literal); }
	|	T_FUNC_C	{ $$ = ast.NewMagicConstLiteral($1, $1.Literal); }
	|	T_NS_C		{ $$ = ast.NewMagicConstLiteral($1, $1.Literal); }
	|	T_CLASS_C	{ $$ = ast.NewMagicConstLiteral($1, $1.Literal); }
	|	T_START_HEREDOC T_ENCAPSED_AND_WHITESPACE T_END_HEREDOC { $$ = ast.NewHeredocExpression($1, $3, ast.NewEncapsedAndWhitespaceLiteral($2, $2.Literal)); }
	|	T_START_HEREDOC T_END_HEREDOC { $$ = ast.NewHeredocExpression($1, $2); }
	|	'"' encaps_list '"' 	{ $$ = ast.NewEncapsListExpression($2...); }
	|	T_START_HEREDOC encaps_list T_END_HEREDOC { $$ = ast.NewHeredocExpression($1, $3, $2...); }
	|	dereferencable_scalar	{ $$ = $1; }
	|	constant			{ $$ = $1; }
;

constant:
		name { $$ = $1; } 
	|	class_name T_PAAMAYIM_NEKUDOTAYIM identifier
			{ $$ = ast.NewConstantExpression($1, $3); }
	|	variable_class_name T_PAAMAYIM_NEKUDOTAYIM identifier
			{ $$ = ast.NewConstantExpression($1, $3); }
;

expr:
		variable					{ $$ = $1; }
	|	expr_without_variable		{ $$ = $1; }
;

optional_expr:
		/* empty */	{ $$ = nil; }
	|	expr		{ $$ = $1; }
;

variable_class_name:
	dereferencable { $$ = $1; }
;

dereferencable:
		variable				{ $$ = $1; }
	|	'(' expr ')'			{ $$ = ast.NewDereferencableExpression(ast.Wrapped, $2); }
	|	dereferencable_scalar	{ $$ = $1; }
;

callable_expr:
		callable_variable		{ $$ = $1; }
	|	'(' expr ')'			{ $$ = ast.NewDereferencableExpression(ast.Wrapped, $2);; }
	|	dereferencable_scalar	{ $$ = $1; }
;

callable_variable:
		simple_variable
			{ $$ = $1; }
	|	dereferencable '[' optional_expr ']'
			{ $$ = ast.NewCallableVariableExpression(ast.Dim, $1, $3); }
	|	constant '[' optional_expr ']'
			{ $$ = ast.NewCallableVariableExpression(ast.Dim, $1, $3); }
	|	dereferencable '{' expr '}'
			{ $$ = ast.NewCallableVariableExpression(ast.Curly, $1, $3); }
	|	dereferencable T_OBJECT_OPERATOR property_name argument_list
			{ $$ = ast.NewCallableVariableExpression(ast.Prop, $1, []ast.Expression{$3, $4}...); }
	|	function_call { $$ = $1; }
;

variable:
		callable_variable
			{ $$ = $1; }
	|	static_member
			{ $$ = $1; }
	|	dereferencable T_OBJECT_OPERATOR property_name
			{ $$ = ast.NewVariableExpression($1, $3); }
;

simple_variable:
		T_VARIABLE			{ $$ = ast.NewVariableLiteral($1, $1.Literal); }
	|	'$' '{' expr '}'	{ $$ = ast.NewSimpleVariableExpression(ast.CurlyOpen, $3); }
	|	'$' simple_variable	{ $$ = ast.NewSimpleVariableExpression(ast.Var, $2); }
;

static_member:
		class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
			{ $$ = ast.NewStaticMemberExpression($1, $3); }
	|	variable_class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
			{ $$ = ast.NewStaticMemberExpression($1, $3); }
;

new_variable:
		simple_variable
			{ $$ = $1; }
	|	new_variable '[' optional_expr ']'
			{ $$ = ast.NewNVariableExpression(ast.Dim, $1, $3); }
	|	new_variable '{' expr '}'
			{ $$ = ast.NewNVariableExpression(ast.Curly, $1, $3); }
	|	new_variable T_OBJECT_OPERATOR property_name
			{ $$ = ast.NewNVariableExpression(ast.Prop, $1, $3); }
	|	class_name T_PAAMAYIM_NEKUDOTAYIM simple_variable
			{ $$ = ast.NewStaticMemberExpression($1, $3); }
	|	new_variable T_PAAMAYIM_NEKUDOTAYIM simple_variable
			{ $$ = ast.NewStaticMemberExpression($1, $3); }
;

member_name:
		identifier { $$ = $1; }
	|	'{' expr '}'	{ $$ = ast.NewMemberNameExpression($2); }
	|	simple_variable	{ $$ = $1; }
;

property_name:
		T_STRING { $$ = ast.NewStringLiteral($1, $1.Literal); }
	|	'{' expr '}'	{ $$ = ast.NewPropertyNameExpression($2); }
	|	simple_variable	{ $$ = $1; }
;

array_pair_list:
		non_empty_array_pair_list
			{ /* allow single trailing comma */ $$ = $1; }
;

possible_array_pair:
		/* empty */ { $$ = []ast.Expression{}; }
	|	array_pair  { $$ = $1; }
;

non_empty_array_pair_list:
		non_empty_array_pair_list ',' possible_array_pair
			{ $$ = append($1, $3...); }
	|	possible_array_pair
			{ $$ = $1; }
;

array_pair:
		expr T_DOUBLE_ARROW expr
			{ $$ = []ast.Expression{ast.NewArrayPairExpression($1, $3, false)}; }
	|	expr
			{ $$ = []ast.Expression{$1}; }
	|	expr T_DOUBLE_ARROW '&' variable
			{ $$ = []ast.Expression{ast.NewArrayPairExpression($1, $4, true)}; }
	|	'&' variable
			{ $$ = []ast.Expression{ast.NewAmpersandLiteral($2)}; }
	|	expr T_DOUBLE_ARROW T_LIST '(' array_pair_list ')'
			{ $$ = []ast.Expression{ast.NewArrayPairExpression($1, ast.NewListExpression($3, $5...), false)}; }
	|	T_LIST '(' array_pair_list ')'
			{ $$ = []ast.Expression{ast.NewListExpression($1, $3...)} }
;

encaps_list:
		encaps_list encaps_var
			{ $$ = append($1, $2); }
	|	encaps_list T_ENCAPSED_AND_WHITESPACE
			{ $$ = append($1, ast.NewEncapsedAndWhitespaceLiteral($2, $2.Literal)); }
	|	encaps_var
			{ $$ = append($$, $1); }
	|	T_ENCAPSED_AND_WHITESPACE encaps_var
			{ $$ = append($$, []ast.Expression{ast.NewEncapsedAndWhitespaceLiteral($1, $1.Literal), $2}...); }
;

encaps_var:
		T_VARIABLE
			{ $$ = ast.NewVariableLiteral($1, $1.Literal); }
	|	T_VARIABLE '[' encaps_var_offset ']'
			{ $$ = ast.NewEncapsVarExpression($1, $1.Literal, ast.Dim, $3); }
	|	T_VARIABLE T_OBJECT_OPERATOR T_STRING
			{ $$ = ast.NewEncapsVarExpression($1, $1.Literal, ast.Prop, ast.NewStringLiteral($3, $3.Literal)); }
	|	T_DOLLAR_OPEN_CURLY_BRACES expr '}'
			{ $$ = ast.NewEncapsVarExpression($1, $1.Literal, ast.DollarOpenCurlyBraces, $2); }
	|	T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '}'
			{ $$ = ast.NewEncapsVarExpression($1, $1.Literal, ast.DollarOpenCurlyBraces, ast.NewStringLiteral($2, $2.Literal)); }
	|	T_DOLLAR_OPEN_CURLY_BRACES T_STRING_VARNAME '[' expr ']' '}'
			{ $$ = ast.NewEncapsVarExpression($1, $1.Literal, ast.DimInDollarOpenCurlyBraces, []ast.Expression{ast.NewStringLiteral($2, $2.Literal), $4}...); }
	|	T_CURLY_OPEN variable '}' { $$ = ast.NewEncapsVarExpression($1, $1.Literal, ast.CurlyOpen, $2); }
;

encaps_var_offset:
		T_STRING			{ $$ = ast.NewStringLiteral($1, $1.Literal); }
	|	T_NUM_STRING		{ $$ = ast.NewIntegerLiteral($1, $1.Literal); }
	|	'-' T_NUM_STRING 	{ $$ = ast.NewIntegerLiteral($2, "-" + $2.Literal); }
	|	T_VARIABLE			{ $$ = ast.NewVariableLiteral($1, $1.Literal); }
;


internal_functions_in_yacc:
		T_ISSET '(' isset_variables ')' { $$ = ast.NewIssetExpression($1, $3...); }
	|	T_EMPTY '(' expr ')' { $$ = ast.NewEmptyExpression($1, $3); }
	|	T_INCLUDE expr
			{ $$ = ast.NewIncludeExpression($1, $2); }
	|	T_INCLUDE_ONCE expr
			{ $$ = ast.NewIncludeExpression($1, $2); }
	|	T_EVAL '(' expr ')'
			{ $$ = ast.NewEvalExpression($1, $3); }
	|	T_REQUIRE expr
			{ $$ = ast.NewRequireExpression($1, $2); }
	|	T_REQUIRE_ONCE expr
			{ $$ = ast.NewRequireExpression($1, $2); }
;

isset_variables:
		isset_variable { $$ = append($$, $1); }
	|	isset_variables ',' isset_variable
			{ $$ = append($1, $3); }
;

isset_variable:
		expr { $$ = $1; }
;

%%

type LexerWrapper struct {
	l          *lexer.Lexer
	recentLit  string
	recentPos  token.Position
	program    *ast.Program
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
	log.Fatalf("Line %d, Column %d: %q %s\n%s", l.recentPos.Line, l.recentPos.Column, l.recentLit, e, l.l.Src())
}

func Parse(l *lexer.Lexer) *ast.Program {
	w := LexerWrapper{l: l}
	if yyParse(&w) != 0 {
		panic("Parse error")
	}
	return w.program
}
