package main

import "fmt"

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t Token) ToString() string {
	return fmt.Sprintf("%-15v %-10v %v", TokenTypeString(t.Type), t.Lexeme, t.Literal)
}
