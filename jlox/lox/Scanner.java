package lox;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static lox.TokenType.*;

class Scanner {
    // The main scanner/tokenizer, responsible for converting the source
    // code string into a series of Tokens.

    private final String source;
    private final List<Token> tokens = new ArrayList<>();
    private int start = 0;   // the first char of the lexeme being scanned
    private int current = 0; // the current character being considered
    private int line = 1;    // the current line being scanned

    Scanner(String source) {
        this.source = source;
    }

    List<Token> scanTokens() {
        // Scan the source code one token at a time, add a special EOF token
        // at the end of file to denote completion.
        while (!isAtEnd()) {
            // We are at the beginning of the next lexeme.
            start = current;
            scanToken();
        }

        tokens.add(new Token(EOF, "", null, line));
        return tokens;
    }

    private void scanToken() {
        char c = advance();
        switch (c) {
            case '(': addToken(LEFT_PAREN); break;
            case ')': addToken(RIGHT_PAREN); break;
            case '{': addToken(LEFT_BRACE); break;
            case '}': addToken(RIGHT_BRACE); break;
            case ',': addToken(COMMA); break;
            case '.': addToken(DOT); break;
            case '-': addToken(MINUS); break;
            case '+': addToken(PLUS); break;
            case ';': addToken(SEMICOLON); break;
            case '*': addToken(STAR); break;
            case '!':
                addToken(match('=') ? BANG_EQUAL : BANG);
                break;
            case '=':
                addToken(match('=') ? EQUAL_EQUAL : EQUAL);
                break;
            case '<':
                addToken(match('=') ? LESS_EQUAL : LESS);
                break;
            case '>':
                addToken(match('=') ? GREATER_EQUAL : GREATER);
                break;
            default:
                Lox.error(line, "Unexpected character: '" + c + "'.");
                break;
        }
    }

    private boolean match(char expected) {
        // Check whether current char matches expected, and if so advance.
        if (isAtEnd()) return false;
        if (source.charAt(current) != expected) return false;

        current++;
        return true;
    }

    private boolean isAtEnd() {
        // True if we have reached the end of the source.
        return current >= source.length();
    }

    private char advance() {
        // Return character at current location, and advance one character.
        return source.charAt(current++);
    }

    private void addToken(TokenType type) {
        // Add a token with no literal.
        addToken(type, null);
    }

    private void addToken(TokenType type, Object literal) {
        // Add a new token.
        String text = source.substring(start, current);
        tokens.add(new Token(type, text, literal, line));
    }
}
