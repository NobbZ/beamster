%{
package compiler

import "math/big"
%}

%union {
	string string
	int *big.Int
}

%token tokInt

%%

integer: tokInt {
	$$ = $1
}


%%
