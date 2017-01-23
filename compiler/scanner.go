// CAUTION: Generated file - DO NOT EDIT.

package compiler

import (
	"math/big"
	"strconv"
	"strings"
)

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
	case c == '+' || c == '-':
		goto yystate3
	case c == '_':
		goto yystate11
	case c >= '0' && c <= '9':
		goto yystate4
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate5
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
		goto yyabort
	case c == '_':
		goto yystate11
	case c >= '0' && c <= '9':
		goto yystate4
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate5
	}

yystate4:
	c = y.getc()
	switch {
	default:
		goto yyrule2
	case c == '_':
		goto yystate11
	case c == '|':
		goto yystate6
	case c >= '0' && c <= '9':
		goto yystate4
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate5
	}

yystate5:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '|':
		goto yystate6
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate5
	}

yystate6:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '1':
		goto yystate7
	case c == '2':
		goto yystate9
	case c == '3':
		goto yystate10
	case c >= '4' && c <= '9':
		goto yystate8
	}

yystate7:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate8
	}

yystate8:
	c = y.getc()
	goto yyrule1

yystate9:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c >= '0' && c <= '9':
		goto yystate8
	}

yystate10:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c >= '0' && c <= '6':
		goto yystate8
	}

yystate11:
	c = y.getc()
	switch {
	default:
		goto yyrule2
	case c >= '0' && c <= '9' || c == '_':
		goto yystate11
	}

yyrule1: // {SIGN}?{BASEDIGIT}+\|{BASE}
	{

		lval.string = string(y.buf)
		parts := strings.Split(lval.string, "|")
		number := parts[0]
		base, _ := strconv.ParseInt(parts[1], 10, 32)
		if val, ok := big.NewInt(0).SetString(number, int(base)); ok {
			lval.int = val
		} else {
			return 0
		}
		return tokInt
	}
yyrule2: // {SIGN}?{DIGIT}+
	{

		lval.string = string(y.buf)
		if val, ok := big.NewInt(0).SetString(strings.Replace(lval.string, "_", "", -1), 10); ok {
			lval.int = val
		} else {
			return 0
		}
		return tokInt
	}
yyrule3: // [ ]+

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	return int(c)
}
