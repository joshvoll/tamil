package lexer

// Lexer struct defintiion
type Lexer struct {
	input        string
	position     int  // the current position of the character
	readPosition int  // the current position of reading the character
	ch           byte // position for byte
}

// New Constructor of the lexer

// readChar is going to give us the next character and advance our position in the input string
// The first thing it does is to check whether we have reached the end of input
// If that’s the case it sets l.ch to 0, which is the ASCII code for the "NUL"
// character and signifies either “we haven’t read anything yet” or “end of file” for us
// But if we haven’t reached the end of input yet it sets l.ch to the next character
// by accessing l.input[l.readPosition]

// NextToken is going to get the next token from each sentence
// We look at the current character under examination (l.ch)
// and return a token depending on which character it is
// What our lexer needs to do is recognize
// whether the current character is a letter and if so

// newToken is a helper function that return the right token

// readIdentifier helper
// it needs to read the rest of the identifier/keyword until it encounters a non-letter-character
// Having read that identifier/keyword, we then need to find out if it is a identifier or a keyword
