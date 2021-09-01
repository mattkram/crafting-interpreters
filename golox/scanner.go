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

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c == '_')
}

func isDigit(c rune) bool {
	return (c >= '0') && (c <= '9')
}

func isAlphaNumeric(c rune) bool {
	return isAlpha(c) || isDigit(c)
}
