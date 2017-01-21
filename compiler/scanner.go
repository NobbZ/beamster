// CAUTION: Generated file - DO NOT EDIT.

package compiler

func tokenize() {

yystate0:

	buf = buf[:0]

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	yyn
yystart1:
	switch {
	default:
		goto yyabort
	case yyc == ' ':
		goto yystate2
	}

yystate2:
	yyn
	switch {
	default:
		goto yyrule1
	case yyc == ' ':
		goto yystate2
	}

yyrule1: // [ ]+

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
}
