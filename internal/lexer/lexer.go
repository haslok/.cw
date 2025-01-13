package lexer

import (
	"unicode"
)

// TokenType represents the type of token.
type TokenType int

const (
	// Define token types
	EOF TokenType = iota
	IDENTIFIER
	NUMBER
	STRING
	PLUS
	MINUS
	ASTERISK
	SLASH
	LPAREN
	RPAREN
)

// Token represents a single token.
type Token struct {
	Type    TokenType
	Literal string
}

// Lexer is responsible for tokenizing the input source code.
type Lexer struct {
	input        string
	position     int  // current position in the input
	readPosition int  // current reading position in the input
	ch          byte // current character
}

// NewLexer initializes a new Lexer.
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads the next character from the input.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = Token{Type: PLUS, Literal: string(l.ch)}
	case '-':
		tok = Token{Type: MINUS, Literal: string(l.ch)}
	case '*':
		tok = Token{Type: ASTERISK, Literal: string(l.ch)}
	case '/':
		tok = Token{Type: SLASH, Literal: string(l.ch)}
	case '(':
		tok = Token{Type: LPAREN, Literal: string(l.ch)}
	case ')':
		tok = Token{Type: RPAREN, Literal: string(l.ch)}
	case 0:
		tok.Type = EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = IDENTIFIER
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = NUMBER
			return tok
		} else {
			tok = Token{Type: EOF, Literal: ""}
		}
	}

	l.readChar()
	return tok
}

// skipWhitespace skips over whitespace characters.
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}

// readIdentifier reads an identifier from the input.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a number from the input.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter checks if a character is a letter.
func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

// isDigit checks if a character is a digit.
func isDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}