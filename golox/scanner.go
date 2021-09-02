package main

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

func (s *Scanner) addToken(type_ TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	token := Token{type_, text, literal, s.line}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) advance() byte {
	s.current = s.current + 1
	return s.source[s.current-1]
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
