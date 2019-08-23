package lexer

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {

	lex := new(Lexer)
	srtipts := []string{
		"int age = 45;",
		"inta age = 45;",
		"in age = 45;",

		"age >= 45",
		"age > 45;",
	}

	for _, script := range srtipts {
		fmt.Printf("[ %s ]\n", script)
		tokenReader := lex.tokenize(script)
		fmt.Println(tokenReader.Dump())
	}
}
