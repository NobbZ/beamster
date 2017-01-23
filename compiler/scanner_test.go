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
	}

	for _, data := range testData {
		lexer := yylexer{
			src:     bufio.NewReader(strings.NewReader(data.input)),
			buf:     make([]byte, 100),
			empty:   true,
			current: 0,
		}

		lval := yySymType{}

		result := (&lexer).Lex(&lval)

		if result != data.expect {
			t.Error("Got", result,
				"expected", data.expect)
		}
	}
}

func TestScannerTokenRecognitionWithValue(t *testing.T) {
	testData := []struct{ input, expect string }{
		{"1", "1"},
	}

	for _, data := range testData {
		lexer := yylexer{
			src:     bufio.NewReader(strings.NewReader(data.input)),
			buf:     make([]byte, 100),
			empty:   true,
			current: 0,
		}

		lval := yySymType{}

		(&lexer).Lex(&lval)

		if lval.string != data.expect {
			t.Errorf("Got %#v, expected %#v.\n", lval.string, data.expect)
		}
	}
}
