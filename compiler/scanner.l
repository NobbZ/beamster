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

DIGIT [0-9]

%%
	y.buf = y.buf[:0]

{DIGIT}+
	lval.string = string(y.buf)
	return tokInt

[ ]+

%%
	return int(c)
}
