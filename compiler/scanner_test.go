package compiler

import (
	"bufio"
	"strings"
	"testing"
)

func TestScannerTokenRecognitionWithoutValue(t *testing.T) {
	testData := []struct{input string; expect int}{
		{"1", tokInt},
	}

	for _, data := range testData {
		lexer := yylexer{
			src: bufio.NewReader(strings.NewReader(data.input)),
			buf: make([]byte, 1024),
			empty: true,
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
