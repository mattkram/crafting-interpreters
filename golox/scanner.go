package main

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	source   string
	tokens   []Token
	start    int
	current  int
	line     int
	keywords map[string]TokenType
}

func NewScanner(source string) Scanner {
	keywords := map[string]TokenType{
		"and":    AND,
		"class":  CLASS,
		"else":   ELSE,
		"false":  FALSE,
		"fun":    FUN,
		"if":     IF,
		"nil":    NIL,
		"or":     OR,
		"print":  PRINT,
		"return": RETURN,
		"super":  SUPER,
		"this":   THIS,
		"true":   TRUE,
		"var":    VAR,
		"while":  WHILE,
	}
	s := Scanner{source: source, line: 1, keywords: keywords}
	return s
}

func (s *Scanner) scanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.start = s.current
	s.addToken(EOF, nil)
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN, nil)
	case ')':
		s.addToken(RIGHT_PAREN, nil)
	case '{':
		s.addToken(LEFT_BRACE, nil)
	case '}':
		s.addToken(RIGHT_BRACE, nil)
	case ',':
		s.addToken(COMMA, nil)
	case '.':
		s.addToken(DOT, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '+':
		s.addToken(PLUS, nil)
	case ';':
		s.addToken(SEMICOLON, nil)
	case '*':
		s.addToken(STAR, nil)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}
	case ' ', '\t', '\r':
		break
	case '"':
		s.string()
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			Error(s.line, fmt.Sprintf("Unexpected character: '%c'", c))
		}
	}
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line = s.line + 1
		}
		s.advance()
	}

	if s.isAtEnd() {
		Error(s.line, "Unterminated string")
		return
	}

	// The closing ".
	s.advance()

	value := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, value)
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	// Look for a fractional part.
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}

	value, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		Error(s.line, fmt.Sprintf("Error converting float"))
	}
	s.addToken(NUMBER, value)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	type_, ok := s.keywords[text]
	if !ok {
		type_ = IDENTIFIER
	}
	s.addToken(type_, nil)
}

func (s *Scanner) addToken(type_ TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	token := Token{type_, text, literal, s.line}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}

	s.current = s.current + 1
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\000'
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return '\000'
	}
	return s.source[s.current+1]
}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current = s.current + 1
	return c
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c == '_')
}

func isDigit(c byte) bool {
	return (c >= '0') && (c <= '9')
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
