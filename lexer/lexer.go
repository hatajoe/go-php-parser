package lexer

import (
	"bytes"
	"encoding/json"
	"io"
	"os/exec"
	"path"
	"runtime"

	"github.com/hatajoe/go-php-parser/token"
)

type Lexer struct {
	src    string
	tokens token.Tokens
	offset int
}

func (s *Lexer) Init(src string) {

	s.src = src

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	echo := exec.Command("echo", src)
	php := exec.Command("php", path.Dir(filename)+"/lexer.php")

	r, w := io.Pipe()
	echo.Stdout = w
	php.Stdin = r

	var out bytes.Buffer
	php.Stdout = &out

	if err := echo.Start(); err != nil {
		panic(err)
	}
	if err := php.Start(); err != nil {
		panic(err)
	}
	if err := echo.Wait(); err != nil {
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}
	if err := php.Wait(); err != nil {
		panic(err)
	}

	tokens := token.Tokens{}
	if err := json.Unmarshal(out.Bytes(), &tokens); err != nil {
		panic(err)
	}
	s.tokens = tokens
}

func (l *Lexer) Scan() *token.Token {
	if l.offset >= len(l.tokens) {
		return &token.Token{Type: token.EOF}
	}
	token := l.tokens[l.offset]
	l.offset++
	return token
}

func (l *Lexer) Src() string {
	return l.src
}
