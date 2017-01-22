// CAUTION: Generated file - DO NOT EDIT.

package compiler

import "fmt"

func lex(y *yylexer, lval *yySymType) int {
	fmt.Println(y.src)
	c := y.current
	if y.empty {
		c, y.empty = y.getc(), false
	}

yystate0:

	y.buf = y.buf[:0]

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = y.getc()
yystart1:
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate2
	case c >= '0' && c <= '9':
		goto yystate3
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule2
	case c == ' ':
		goto yystate2
	}

yystate3:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c >= '0' && c <= '9':
		goto yystate3
	}

yyrule1: // {DIGIT}+
	{

		fmt.Println("DIGIT", y.buf)
		return tokInt
	}
yyrule2: // [ ]+

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	return int(c)
}
