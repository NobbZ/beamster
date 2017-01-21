package compiler

import "bufio"

type yylexer struct{
	src     *bufio.Reader
	buf     []byte
	empty   bool
	current byte
}
