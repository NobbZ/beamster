package compiler

import (
	"bufio"
	"fmt"
)

type yylexer struct{
	src     *bufio.Reader
	buf     []byte
	empty   bool
	current byte
}

func (y *yylexer) Lex(lval *yySymType) int {
	return lex(y, lval)
}

func (y *yylexer) getc() byte {
	fmt.Println("Start", y.current)
	if y.current != 0 {
		y.buf = append(y.buf, y.current)
	}
	y.current = 0
	if b, err := y.src.ReadByte(); err == nil {
		y.current = b
	}
	fmt.Println("Stop", y.current)
	return y.current
}
