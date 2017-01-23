// CAUTION: Generated file - DO NOT EDIT.

package compiler

func lex(y *yylexer, lval *yySymType) int {
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
	case c == '_':
		goto yystate10
	case c >= '0' && c <= '9':
		goto yystate3
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate4
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule3
	case c == ' ':
		goto yystate2
	}

yystate3:
	c = y.getc()
	switch {
	default:
		goto yyrule2
	case c == '_':
		goto yystate10
	case c == '|':
		goto yystate5
	case c >= '0' && c <= '9':
		goto yystate3
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate4
	}

yystate4:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '|':
		goto yystate5
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate4
	}

yystate5:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '1':
		goto yystate6
	case c == '2':
		goto yystate8
	case c == '3':
		goto yystate9
	case c >= '4' && c <= '9':
		goto yystate7
	}

yystate6:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate7:
	c = y.getc()
	goto yyrule1

yystate8:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate9:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c >= '0' && c <= '6':
		goto yystate7
	}

yystate10:
	c = y.getc()
	switch {
	default:
		goto yyrule2
	case c >= '0' && c <= '9' || c == '_':
		goto yystate10
	}

yyrule1: // {BASEDIGIT}+\|{BASE}
	{

		lval.string = string(y.buf)
		return tokInt
	}
yyrule2: // {DIGIT}+
	{

		lval.string = string(y.buf)
		return tokInt
	}
yyrule3: // [ ]+

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	return int(c)
}
