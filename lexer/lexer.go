package lexer

// Lexer definition
type Lexer struct {
	input        string
	position     int  // current position in input (points to the curren chararacter)
	readPosition int  // current reading position in input (after current character)
	ch           byte // current character under examination
}

// New constructor function for lexer
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

// readChar going to use position and readPosition to undestand what is next to read
// position is the current character and readPosition read to the next character
// if the readPosition is greater than the input length the character is back to zero
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
