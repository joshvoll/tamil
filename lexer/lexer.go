package lexer

import (
	"fmt"

	"github.com/joshvoll/tamil/token"
)

// Lexer struct defintiion
type Lexer struct {
	input        string
	position     int  // the current position of the character
	readPosition int  // current reading position in input (after current character)
	ch           byte // current character under examination
}

// New Constructor of the lexer
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

// readChar is going to give us the next character and advance our position in the input string
// The first thing it does is to check whether we have reached the end of input
// If that’s the case it sets l.ch to 0, which is the ASCII code for the "NUL"
// character and signifies either “we haven’t read anything yet” or “end of file” for us
// But if we haven’t reached the end of input yet it sets l.ch to the next character
// by accessing l.input[l.readPosition]
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	fmt.Println("Position ", l.position, " readPosition ", l.readPosition, " character ", l.ch)
	l.position = l.readPosition
	l.readPosition++
}

// NextToken going to assign the nexto token to the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

// newToken get the token type and character byte and return the type
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
