package parser

import (
	"github.com/joshvoll/tamil/token"

	"github.com/joshvoll/tamil/lexer"
)

// Parser struct definition
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
