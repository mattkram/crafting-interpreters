package main

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	tokens := []Token{
		{STRING, "\"Hello\"", "Hello", 1},
		{NUMBER, "1.25", 1.25, 1},
		{LEFT_PAREN, "(", nil, 1},
	}

	// TODO: Add actual test
	for _, token := range tokens {
		fmt.Println(token.ToString())
	}
}
