package compiler

import (
	"bufio"
	"fmt"
	"math/big"
	"strings"
	"testing"
)

func TestScanInteger(t *testing.T) {
	testData := []struct {
		input  string
		expect *big.Int
	}{
		{"1", toBigInt("1", 10)},
		{"5|6", toBigInt("5", 6)},
		{"1_000", toBigInt("1000", 10)},
		{"-1", toBigInt("-1", 10)},
		{"-5|6", toBigInt("-5", 6)},
		{"-1_000", toBigInt("-1000", 10)},
		{"+1", toBigInt("+1", 10)},
		{"+5|6", toBigInt("+5", 6)},
		{"+1_000", toBigInt("+1000", 10)},
	}

	for _, data := range testData {
		lexer := yylexer{
			src:   bufio.NewReader(strings.NewReader(data.input)),
			empty: true,
		}

		lval := yySymType{}

		(&lexer).Lex(&lval)

		if lval.int.Cmp(data.expect) != 0 {
			t.Errorf("Got %#v, expected %#v.\n", lval.int, data.expect)
		}
	}
}

func toBigInt(s string, base int) *big.Int {
	result, ok := big.NewInt(0).SetString(s, base)
	if ok {
		return result
	}
	error := fmt.Sprint("Can't parse ", s, " with base ", base, " into BigInt")
	panic(error)
}
