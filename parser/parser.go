package parser

import (
	"github.com/joshvoll/tamil/ast"
	"github.com/joshvoll/tamil/token"

	"github.com/joshvoll/tamil/lexer"
)

// Parser struct definition
// The Parser has three fields: l, curToken and peekToken. l is a pointer to an instance of the lexer,
// on which we repeatedly call NextToken() to get the next token in the input.
// curToken and peekToken act exactly like the two “pointers” our lexer has: position and readPosition.
// But instead of pointing to a character in the input, they point to the current and the next token.
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New constructor function fo the Parser
// Read two tokens, so curToken and peekToken are both set
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}
	p.nextToken()
	p.nextToken()
	return p
}

// nextToken take peekToken and NextToken
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram define the pont to the ast.Program
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
