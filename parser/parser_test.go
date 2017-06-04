package parser

import (
	"testing"

	"github.com/hatajoe/go-php-parser/lexer"
)

func TestSimpleEcho(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`<?php echo 1;`,
			`echo 1;`,
		},
		{
			`<?php 
			
echo 1;`,
			`echo 1;`,
		},
		{
			`<?php echo 1,2,3,4;`,
			`echo 1, 2, 3, 4;`,
		},
		{
			`<?php echo 1.0;`,
			`echo 1.0;`,
		},
		{
			`<?php echo 1.2;`,
			`echo 1.2;`,
		},
		{
			`<?php echo 1.2,3;`,
			`echo 1.2, 3;`,
		},
		{
			`<?php echo __LINE__;`,
			`echo __LINE__;`,
		},
		{
			`<?php echo __FILE__;`,
			`echo __FILE__;`,
		},
		{
			`<?php echo __DIR__;`,
			`echo __DIR__;`,
		},
		{
			`<?php echo __TRAIT__;`,
			`echo __TRAIT__;`,
		},
		{
			`<?php echo __METHOD__;`,
			`echo __METHOD__;`,
		},
		{
			`<?php echo __FUNCTION__;`,
			`echo __FUNCTION__;`,
		},
		{
			`<?php echo __NAMESPACE__;`,
			`echo __NAMESPACE__;`,
		},
		{
			`<?php echo __CLASS__;`,
			`echo __CLASS__;`,
		},
		{
			`<?php echo <<<EOL
foo
bar
bars
	hoge
EOL;`,
			`echo <<<EOL
foo
bar
bars
	hoge
EOL;`,
		},
		{
			`<?php echo <<<EOL
EOL;`,
			`echo <<<EOL
EOL;`,
		},
		{
			`<?php echo "hello $foo";`,
			`echo "hello $foo";`,
		},
		{
			`<?php echo "hello $foo world";`,
			`echo "hello $foo world";`,
		},
		{
			`<?php echo "hello $foo world $bar";`,
			`echo "hello $foo world $bar";`,
		},
		{
			`<?php echo "hello
$foo
world

	$bar";`,
			`echo "hello
$foo
world

	$bar";`,
		},
		{
			`<?php echo <<<EOL
hello $foo
EOL;`,
			`echo <<<EOL
hello $foo
EOL;`,
		},
		{
			`<?php echo <<<EOL
hello $foo world
EOL;`,
			`echo <<<EOL
hello $foo world
EOL;`,
		},
		{
			`<?php echo <<<EOL
hello $foo world $bar
EOL;`,
			`echo <<<EOL
hello $foo world $bar
EOL;`,
		},
		{
			`<?php echo <<<EOL
hello
$foo
world

	$bar
EOL;`,
			`echo <<<EOL
hello
$foo
world

	$bar
EOL;`,
		},
		{
			`<?php echo "hello $foo[true]";`,
			`echo "hello $foo[true]";`,
		},
		{
			`<?php echo "hello $foo[1]";`,
			`echo "hello $foo[1]";`,
		},
		{
			`<?php echo "hello $foo[-1]";`,
			`echo "hello $foo[-1]";`,
		},
		{
			`<?php echo "hello $foo[$bar]";`,
			`echo "hello $foo[$bar]";`,
		},
		{
			`<?php echo "hello $foo->bar";`,
			`echo "hello $foo->bar";`,
		},
		{
			`<?php echo "hello ${1}";`,
			`echo "hello ${1}";`,
		},
		{
			`<?php echo "hello ${a}";`,
			`echo "hello ${a}";`,
		},
		{
			`<?php echo "hello ${$foo}";`,
			`echo "hello ${$foo}";`,
		},
		{
			`<?php echo "hello ${a[1]}";`,
			`echo "hello ${a[1]}";`,
		},
		{
			`<?php echo "hello {$foo}";`,
			`echo "hello {$foo}";`,
		},
		{
			`<?php echo "hello {${$foo}}";`,
			`echo "hello {${$foo}}";`,
		},
		{
			`<?php echo "hello {${$$$foo}}";`,
			`echo "hello {${$$$foo}}";`,
		},
		{
			`<?php echo "hello {${$$$foo}}";`,
			`echo "hello {${$$$foo}}";`,
		},
		{
			`<?php echo "hello {$foo[]}";`,
			`echo "hello {$foo[]}";`,
		},
		{
			`<?php echo "hello {$foo[$bar]}";`,
			`echo "hello {$foo[$bar]}";`,
		},
		{
			`<?php echo "hello {()[$bar]}";`,
			`echo "hello {()[$bar]}";`,
		},
	}

	for idx, test := range tests {
		l := &lexer.Lexer{}
		l.Init(test.input)
		program := Parse(l)
		if program.String() != test.expected {
			t.Errorf("test %d failed. expected=`%s`, got=`%s`", idx, test.expected, program.String())
		}
	}
}
