package main

import (
	"testing"
)

func TestTokenTypeIsInt(t *testing.T) {
	// Check existance of all tokens, which are ints >= 0.
	tokenTypes := []TokenType{
		LEFT_PAREN, RIGHT_PAREN, LEFT_BRACE, RIGHT_BRACE,
		COMMA, DOT, MINUS, PLUS, SEMICOLON, SLASH, STAR,

		BANG, BANG_EQUAL, EQUAL, EQUAL_EQUAL,
		GREATER, GREATER_EQUAL, LESS, LESS_EQUAL,

		IDENTIFIER, STRING, NUMBER,

		AND, CLASS, ELSE, FALSE, FUN, FOR, IF, NIL,
		OR, PRINT, RETURN, SUPER, THIS, TRUE, VAR, WHILE,

		EOF,
	}

	for _, tokenType := range tokenTypes {
		if tokenType < 0 {
			t.Errorf("tokenType must be an integer >= 0")
		}
	}
}
