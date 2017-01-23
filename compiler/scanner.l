%{
package compiler

import (
	"math/big"
	"strings"
	"strconv"
)

func lex(y *yylexer, lval *yySymType) int {
	c := y.current
	if y.empty {
		c, y.empty = y.getc(), false
	}
%}

%yyc c
%yyn c = y.getc()

BASE ([2-9]|[1-2][0-9]|3[0-6])
BASEDIGIT [0-9a-zA-Z]

DIGIT [0-9_]

SIGN [+-]

%%
	y.buf = y.buf[:0]

{SIGN}?{BASEDIGIT}+\|{BASE}
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

{SIGN}?{DIGIT}+
	lval.string = string(y.buf)
	if val, ok := big.NewInt(0).SetString(strings.Replace(lval.string, "_", "", -1), 10); ok {
		lval.int = val
	} else {
		return 0
	}
	return tokInt

[ ]+

%%
	return int(c)
}
