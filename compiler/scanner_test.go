package compiler

import (
	"bufio"
	"strings"
	"testing"
)

func TestScannerTokenRecognitionWithoutValue(t *testing.T) {
	testData := []struct {
		input  string
		expect int
	}{
		{"1", tokInt},
		{"5|6", tokInt},
		{"1_000", tokInt},
		{"-1", tokInt},
		{"-5|6", tokInt},
		{"-1_000", tokInt},
		{"+1", tokInt},
		{"+5|6", tokInt},
		{"+1_000", tokInt},
	}

	for _, data := range testData {
		lexer := yylexer{
			src:   bufio.NewReader(strings.NewReader(data.input)),
			empty: true,
		}

		lval := yySymType{}

		result := (&lexer).Lex(&lval)

		if result != data.expect {
			t.Error("Got", result,
				"expected", data.expect,
				"input was", data.input,
			)
		}
	}
}

func TestScannerTokenRecognitionWithValue(t *testing.T) {
	testData := []struct{ input, expect string }{
		{"1", "1"},
		{"5|3", "5|3"},
		{"1_000", "1_000"},
		{"-1", "-1"},
		{"-5|3", "-5|3"},
		{"-1_000", "-1_000"},
		{"+1", "+1"},
		{"+5|3", "+5|3"},
		{"+1_000", "+1_000"},
	}

	for _, data := range testData {
		lexer := yylexer{
			src:   bufio.NewReader(strings.NewReader(data.input)),
			empty: true,
		}

		lval := yySymType{}

		(&lexer).Lex(&lval)

		if lval.string != data.expect {
			t.Errorf("Got %#v, expected %#v.\n", lval.string, data.expect)
		}
	}
}

func BenchmarkIntegerWithoutBase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lval := yySymType{}
		lexer := &yylexer{
			src:   bufio.NewReader(strings.NewReader("500")),
			empty: true,
		}
		lexer.Lex(&lval)
	}
}

func BenchmarkIntegerWithBase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lval := yySymType{}
		lexer := &yylexer{
			src:   bufio.NewReader(strings.NewReader("500|32")),
			empty: true,
		}
		lexer.Lex(&lval)
	}
}
