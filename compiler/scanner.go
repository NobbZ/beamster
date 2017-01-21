// CAUTION: Generated file - DO NOT EDIT.

package compiler

func lex(y *yylexer, lval *yySymType) int {
	c := y.current

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
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c == ' ':
		goto yystate2
	}

yyrule1: // [ ]+

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	return int(c)
}
