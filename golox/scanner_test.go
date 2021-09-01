package main

import "testing"

func TestNewScanner(t *testing.T) {
	source := "print \"Hello World!\""
	s := NewScanner(source)

	if s.source != source {
		t.Errorf("Source doesn't match")
	}

	if s.start != 0 {
		t.Errorf("start should equal 0")
	}

	if s.current != 0 {
		t.Errorf("current should equal 0")
	}

	if s.line != 1 {
		t.Errorf("s.line should equal 1")
	}
}

func TestAddToken(t *testing.T) {
	source := "print \"Hello World!\""
	s := NewScanner(source)
	s.addToken(SEMICOLON, nil)

	token := s.tokens[0]
	if token.Type != SEMICOLON {
		t.Errorf("Unexpected token type")
	}

	if token.Line != 1 {
		t.Errorf("Unexpected token line")
	}

	s.addToken(PRINT, nil)
	if len(s.tokens) != 2 {
		t.Errorf("Expected 2 tokens, got %v", len(s.tokens))
	}
}
