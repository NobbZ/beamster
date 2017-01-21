%{
package compiler

func lex(y *yylexer, lval *yySymType) int {
	c := y.current
%}

%yyc c
%yyn c = y.getc()

%%
	y.buf = y.buf[:0]

[ ]+

%%
	return int(c)
}
