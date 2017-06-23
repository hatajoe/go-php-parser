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
		{
			`<?php echo array();`,
			`echo array();`,
		},
		{
			`<?php echo [];`,
			`echo [];`,
		},
		{
			`<?php echo array(1=>2);`,
			`echo array(
    1 => 2,
);`,
		},
		{
			`<?php echo [1=>2];`,
			`echo [
    1 => 2,
];`,
		},
		{
			`<?php echo array(1=>2,3=>4);`,
			`echo array(
    1 => 2,
    3 => 4,
);`,
		},
		{
			`<?php echo [1=>2,3=>4];`,
			`echo [
    1 => 2,
    3 => 4,
];`,
		},
		{
			`<?php echo array(1,2,3,4,5,6,7,8,9,10);`,
			`echo array(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);`,
		},
		{
			`<?php echo [1,2,3,4,5,6,7,8,9,10];`,
			`echo [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];`,
		},
		{
			`<?php echo array(1,2,3,4,5,6,7,8,9,10,11);`,
			`echo array(
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
    10,
    11,
);`,
		},
		{
			`<?php echo [1,2,3,4,5,6,7,8,9,10,11];`,
			`echo [
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
    10,
    11,
];`,
		},
		{
			`<?php echo array(1=>&$foo,3=>&$bar);`,
			`echo array(
    1 => &$foo,
    3 => &$bar,
);`,
		},
		{
			`<?php echo [1=>&$foo,3=>&$bar];`,
			`echo [
    1 => &$foo,
    3 => &$bar,
];`,
		},
		{
			`<?php echo array(&$foo,&$bar);`,
			`echo array(&$foo, &$bar);`,
		},
		{
			`<?php echo [&$foo,&$bar];`,
			`echo [&$foo, &$bar];`,
		},
		{
			`<?php echo array(1=>list($foo,$bar));`,
			`echo array(
    1 => list($foo, $bar),
);`,
		},
		{
			`<?php echo [1=>list($foo,$bar)];`,
			`echo [
    1 => list($foo, $bar),
];`,
		},
		{
			`<?php echo array(list($foo,$bar),list(1,2));`,
			`echo array(list($foo, $bar), list(1, 2));`,
		},
		{
			`<?php echo [list($foo,$bar),list(1,2)];`,
			`echo [list($foo, $bar), list(1, 2)];`,
		},
		{
			`<?php echo "hello world";`,
			`echo "hello world";`,
		},
		{
			`<?php echo $foo[0];`,
			`echo $foo[0];`,
		},
		{
			`<?php echo $foo[0];`,
			`echo $foo[0];`,
		},
		{
			`<?php echo foo[0];`,
			`echo foo[0];`,
		},
		{
			`<?php echo $foo{0};`,
			`echo $foo{0};`,
		},
		{
			`<?php echo $foo->bar();`,
			`echo $foo->bar();`,
		},
		{
			`<?php echo $foo->{"bar"}();`,
			`echo $foo->{"bar"}();`,
		},
		{
			`<?php echo $foo->$bar();`,
			`echo $foo->$bar();`,
		},
		{
			`<?php echo $foo->bar(1);`,
			`echo $foo->bar(1);`,
		},
		{
			`<?php echo $foo->bar(1,2);`,
			`echo $foo->bar(1, 2);`,
		},
		{
			`<?php echo $foo->bar(1,2);`,
			`echo $foo->bar(1, 2);`,
		},
		{
			`<?php echo $foo->bar(... $baz);`,
			`echo $foo->bar(...$baz);`,
		},
		{
			`<?php echo Foo\Bar\Baz;`,
			`echo Foo\Bar\Baz;`,
		},
		{
			`<?php echo namespace \Foo\Bar\Baz;`,
			`echo namespace \Foo\Bar\Baz;`,
		},
		{
			`<?php echo \Foo\Bar\Baz;`,
			`echo \Foo\Bar\Baz;`,
		},
		{
			`<?php echo foo();`,
			`echo foo();`,
		},
		{
			`<?php echo foo(1,2);`,
			`echo foo(1, 2);`,
		},
		{
			`<?php echo static::bar(1,2);`,
			`echo static::bar(1, 2);`,
		},
		{
			`<?php echo Foo::bar(1,2);`,
			`echo Foo::bar(1, 2);`,
		},
		{
			`<?php echo Foo::{"bar"}(1,2);`,
			`echo Foo::{"bar"}(1, 2);`,
		},
		{
			`<?php echo Foo::$bar(1,2);`,
			`echo Foo::$bar(1, 2);`,
		},
		{
			`<?php echo $foo::$bar(1,2);`,
			`echo $foo::$bar(1, 2);`,
		},
		{
			`<?php echo ($foo)->$bar(1,2);`,
			`echo ($foo)->$bar(1, 2);`,
		},
		{
			`<?php echo [$foo](1,2);`,
			`echo [$foo](1, 2);`,
		},
		{
			`<?php echo Foo::$bar;`,
			`echo Foo::$bar;`,
		},
		{
			`<?php echo $foo::$bar;`,
			`echo $foo::$bar;`,
		},
		{
			`<?php echo $foo->bar;`,
			`echo $foo->bar;`,
		},
		{
			`<?php echo Foo::Bar;`,
			`echo Foo::Bar;`,
		},
		{
			`<?php echo $foo::Bar;`,
			`echo $foo::Bar;`,
		},
		{
			`<?php echo list($a, $b) = $foo->bar();`,
			`echo list($a, $b) = $foo->bar();`,
		},
		{
			`<?php echo [$a, $b] = $foo->bar();`,
			`echo [$a, $b] = $foo->bar();`,
		},
		{
			`<?php echo $a = 1;`,
			`echo $a = 1;`,
		},
		{
			`<?php echo $a = &$b;`,
			`echo $a = &$b;`,
		},
		{
			`<?php echo $a += 1;`,
			`echo $a += 1;`,
		},
		{
			`<?php echo $a -= 1;`,
			`echo $a -= 1;`,
		},
		{
			`<?php echo $a *= 1;`,
			`echo $a *= 1;`,
		},
		{
			`<?php echo $a **= 1;`,
			`echo $a **= 1;`,
		},
		{
			`<?php echo $a /= 1;`,
			`echo $a /= 1;`,
		},
		{
			`<?php echo $a .= "1";`,
			`echo $a .= "1";`,
		},
		{
			`<?php echo $a %= 1;`,
			`echo $a %= 1;`,
		},
		{
			`<?php echo $a &= 1;`,
			`echo $a &= 1;`,
		},
		{
			`<?php echo $a |= 1;`,
			`echo $a |= 1;`,
		},
		{
			`<?php echo $a ^= 1;`,
			`echo $a ^= 1;`,
		},
		{
			`<?php echo $a <<= 1;`,
			`echo $a <<= 1;`,
		},
		{
			`<?php echo $a >>= 1;`,
			`echo $a >>= 1;`,
		},
		{
			`<?php echo $a || $b;`,
			`echo $a || $b;`,
		},
		{
			`<?php echo $a && $b;`,
			`echo $a && $b;`,
		},
		{
			`<?php echo $a or $b;`,
			`echo $a or $b;`,
		},
		{
			`<?php echo $a and $b;`,
			`echo $a and $b;`,
		},
		{
			`<?php echo $a xor $b;`,
			`echo $a xor $b;`,
		},
		{
			`<?php echo $a | $b;`,
			`echo $a | $b;`,
		},
		{
			`<?php echo $a & $b;`,
			`echo $a & $b;`,
		},
		{
			`<?php echo $a ^ $b;`,
			`echo $a ^ $b;`,
		},
		{
			`<?php echo $a . $b;`,
			`echo $a . $b;`,
		},
		{
			`<?php echo $a + $b;`,
			`echo $a + $b;`,
		},
		{
			`<?php echo $a - $b;`,
			`echo $a - $b;`,
		},
		{
			`<?php echo $a * $b;`,
			`echo $a * $b;`,
		},
		{
			`<?php echo $a ** $b;`,
			`echo $a ** $b;`,
		},
		{
			`<?php echo $a / $b;`,
			`echo $a / $b;`,
		},
		{
			`<?php echo $a % $b;`,
			`echo $a % $b;`,
		},
		{
			`<?php echo $a << $b;`,
			`echo $a << $b;`,
		},
		{
			`<?php echo $a >> $b;`,
			`echo $a >> $b;`,
		},
		{
			`<?php echo $a === $b;`,
			`echo $a === $b;`,
		},
		{
			`<?php echo $a !== $b;`,
			`echo $a !== $b;`,
		},
		{
			`<?php echo $a == $b;`,
			`echo $a == $b;`,
		},
		{
			`<?php echo $a != $b;`,
			`echo $a != $b;`,
		},
		{
			`<?php echo $a < $b;`,
			`echo $a < $b;`,
		},
		{
			`<?php echo $a <= $b;`,
			`echo $a <= $b;`,
		},
		{
			`<?php echo $a > $b;`,
			`echo $a > $b;`,
		},
		{
			`<?php echo $a >= $b;`,
			`echo $a >= $b;`,
		},
		{
			`<?php echo $a <=> $b;`,
			`echo $a <=> $b;`,
		},
		{
			`<?php echo $a instanceof Foo;`,
			`echo $a instanceof Foo;`,
		},
		{
			`<?php echo $a instanceof $foo;`,
			`echo $a instanceof $foo;`,
		},
		{
			`<?php echo $a instanceof $foo[$bar];`,
			`echo $a instanceof $foo[$bar];`,
		},
		{
			`<?php echo $a instanceof $foo{"foo"+$bar};`,
			`echo $a instanceof $foo{"foo" + $bar};`,
		},
		{
			`<?php echo $a instanceof $foo->bar;`,
			`echo $a instanceof $foo->bar;`,
		},
		{
			`<?php echo $a instanceof Foo::$bar;`,
			`echo $a instanceof Foo::$bar;`,
		},
		{
			`<?php echo $a instanceof $foo::$bar;`,
			`echo $a instanceof $foo::$bar;`,
		},
		{
			`<?php echo (1 + 2);`,
			`echo (1 + 2);`,
		},
		{
			`<?php echo clone $a;`,
			`echo clone $a;`,
		},
		{
			`<?php echo (int)$a;`,
			`echo (int)$a;`,
		},
		{
			`<?php echo (integer)$a;`,
			`echo (integer)$a;`,
		},
		{
			`<?php echo (real)$a;`,
			`echo (real)$a;`,
		},
		{
			`<?php echo (double)$a;`,
			`echo (double)$a;`,
		},
		{
			`<?php echo (float)$a;`,
			`echo (float)$a;`,
		},
		{
			`<?php echo (string)$a;`,
			`echo (string)$a;`,
		},
		{
			`<?php echo (array)$a;`,
			`echo (array)$a;`,
		},
		{
			`<?php echo (object)$a;`,
			`echo (object)$a;`,
		},
		{
			`<?php echo (bool)$a;`,
			`echo (bool)$a;`,
		},
		{
			`<?php echo (boolean)$a;`,
			`echo (boolean)$a;`,
		},
		{
			`<?php echo (unset)$a;`,
			`echo (unset)$a;`,
		},
		{
			`<?php echo exit;`,
			`echo exit;`,
		},
		{
			`<?php echo exit(1);`,
			`echo exit(1);`,
		},
		{
			`<?php echo +$a;`,
			`echo +$a;`,
		},
		{
			`<?php echo -$a;`,
			`echo -$a;`,
		},
		{
			`<?php echo !$a;`,
			`echo !$a;`,
		},
		{
			`<?php echo ~$a;`,
			`echo ~$a;`,
		},
		{
			`<?php echo @$a;`,
			`echo @$a;`,
		},
		{
			`<?php echo $a++;`,
			`echo $a++;`,
		},
		{
			`<?php echo ++$a;`,
			`echo ++$a;`,
		},
		{
			`<?php echo $a--;`,
			`echo $a--;`,
		},
		{
			`<?php echo --$a;`,
			`echo --$a;`,
		},
		{
			`<?php echo $a > 1 ? $b : $c;`,
			`echo $a > 1 ? $b : $c;`,
		},
		{
			`<?php echo $a > 1 ?: $b;`,
			`echo $a > 1 ?: $b;`,
		},
		{
			"<?php echo ``;",
			"echo ``;",
		},
		{
			"<?php echo `$a`;",
			"echo `$a`;",
		},
		{
			"<?php echo `$a -la`;",
			"echo `$a -la`;",
		},
		{
			`<?php echo print "foo";`,
			`echo print "foo";`,
		},
		{
			`<?php echo new Foo;`,
			`echo new Foo;`,
		},
		{
			`<?php echo new Foo();`,
			`echo new Foo();`,
		},
		{
			`<?php echo new Foo(1, 2);`,
			`echo new Foo(1, 2);`,
		},
		{
			`<?php echo isset($a);`,
			`echo isset($a);`,
		},
		{
			`<?php echo isset($a,$b);`,
			`echo isset($a, $b);`,
		},
		{
			`<?php echo empty($a);`,
			`echo empty($a);`,
		},
		{
			`<?php echo include "foo/bar.php";`,
			`echo include "foo/bar.php";`,
		},
		{
			`<?php echo include_once "foo/bar.php";`,
			`echo include_once "foo/bar.php";`,
		},
		{
			`<?php echo eval($a);`,
			`echo eval($a);`,
		},
		{
			`<?php echo require "foo/bar.php";`,
			`echo require "foo/bar.php";`,
		},
		{
			`<?php echo require_once "foo/bar.php";`,
			`echo require_once "foo/bar.php";`,
		},
		{
			`<?php echo yield;`,
			`echo yield;`,
		},
		{
			`<?php echo yield foo();`,
			`echo yield foo();`,
		},
		{
			`<?php echo yield $foo => $bar;`,
			`echo yield $foo => $bar;`,
		},
		{
			`<?php echo yield from $foo;`,
			`echo yield from $foo;`,
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

func TestBlockStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`<?php {}`,
			`{}`,
		},
		{
			`<?php {
    $a = 1;
    $b = 2;
}`,
			`{
    $a = 1;
    $b = 2;
}`,
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

func TestIfStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`<?php if (true) $a = 1;`,
			`if (true) $a = 1;`,
		},
		{
			`<?php if (true) {
    $a = 1;
    $b = 2;
}`,
			`if (true) {
    $a = 1;
    $b = 2;
}`,
		},
		{
			`<?php if (true) {
    $a = 1;
    $b = 2;
} elseif ($a > 1) {
    $c = 1;
    $d = 2;
} elseif ($b < 1) {
    $e = 1;
    $f = 2;
}`,
			`if (true) {
    $a = 1;
    $b = 2;
} elseif ($a > 1) {
    $c = 1;
    $d = 2;
} elseif ($b < 1) {
    $e = 1;
    $f = 2;
}`,
		},
		{
			`<?php if (true) {
    $a = 1;
    $b = 2;
} elseif ($a > 1) {
    $c = 1;
    $d = 2;
} elseif ($b < 1) {
    $e = 1;
    $f = 2;
} else {
    $g = 3;
}`,
			`if (true) {
    $a = 1;
    $b = 2;
} elseif ($a > 1) {
    $c = 1;
    $d = 2;
} elseif ($b < 1) {
    $e = 1;
    $f = 2;
} else {
    $g = 3;
}`,
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

func TestAltIfStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`<?php if (true): $a = 1; endif;`,
			`if (true):
    $a = 1;
endif;`,
		},
		{
			`<?php if (true):
    $a = 1;
    $b = 2;
endif;`,
			`if (true):
    $a = 1;
    $b = 2;
endif;`,
		},
		{
			`<?php if (true):
    $a = 1;
    $b = 2;
elseif ($a > 1):
    $c = 1;
    $d = 2;
elseif ($b < 1):
    $e = 1;
    $f = 2;
endif;`,
			`if (true):
    $a = 1;
    $b = 2;
elseif ($a > 1):
    $c = 1;
    $d = 2;
elseif ($b < 1):
    $e = 1;
    $f = 2;
endif;`,
		},
		{
			`<?php if (true):
    $a = 1;
    $b = 2;
elseif ($a > 1):
    $c = 1;
    $d = 2;
elseif ($b < 1):
    $e = 1;
    $f = 2;
else:
    $g = 3;
endif;`,
			`if (true):
    $a = 1;
    $b = 2;
elseif ($a > 1):
    $c = 1;
    $d = 2;
elseif ($b < 1):
    $e = 1;
    $f = 2;
else:
    $g = 3;
endif;`,
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

func TestWhileStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`<?php while (true) {
    $a = 1;
    $b = 2;
}`,
			`while (true) {
    $a = 1;
    $b = 2;
}`,
		},
		{
			`<?php while (true) :
    $a = 1;
    $b = 2;
endwhile;`,
			`while (true) :
    $a = 1;
    $b = 2;
endwhile;`,
		},
		{
			`<?php do {
    $a = 1;
    $b = 2;
} while (true);`,
			`do {
    $a = 1;
    $b = 2;
} while (true);`,
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

func TestForStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`<?php for (;;) {
    $a = 1;
    $b = 2;
}`,
			`for (;;) {
    $a = 1;
    $b = 2;
}`,
		},
		{
			`<?php for ($i=0,$j=0; $i<0,$j<0; $i++,$j++) {
    $a = 1;
    $b = 2;
}`,
			`for ($i = 0, $j = 0; $i < 0, $j < 0; $i++, $j++) {
    $a = 1;
    $b = 2;
}`,
		},
		{
			`<?php for ($i=0; $i<0; $i++) {
    $a = 1;
    $b = 2;
}`,
			`for ($i = 0; $i < 0; $i++) {
    $a = 1;
    $b = 2;
}`,
		},
		{
			`<?php for (; $i<0; $i++) {
    $a = 1;
    $b = 2;
}`,
			`for (; $i < 0; $i++) {
    $a = 1;
    $b = 2;
}`,
		},
		{
			`<?php for ($i=0;; $i++) {
    $a = 1;
    $b = 2;
}`,
			`for ($i = 0;; $i++) {
    $a = 1;
    $b = 2;
}`,
		},
		{
			`<?php for ($i=0; $i<0;) {
    $a = 1;
    $b = 2;
}`,
			`for ($i = 0; $i < 0;) {
    $a = 1;
    $b = 2;
}`,
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
