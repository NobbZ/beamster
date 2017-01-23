%{
package compiler

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

%%
	y.buf = y.buf[:0]

{BASEDIGIT}+\|{BASE}
	lval.string = string(y.buf)
	return tokInt

{DIGIT}+
	lval.string = string(y.buf)
	return tokInt

[ ]+

%%
	return int(c)
}
