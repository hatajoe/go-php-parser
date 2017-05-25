package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hatajoe/go-php-parser/lexer"
	"github.com/hatajoe/go-php-parser/parser"
)

func main() {
	for _, arg := range os.Args[1:] {
		l := lexer.Lexer{}
		body, err := ioutil.ReadFile(arg)
		if err != nil {
			panic(err)
		}
		l.Init(string(body))
		program := parser.Parse(&l)
		fmt.Printf("%#v\n", program)
		fmt.Println(program.String())
	}
}
