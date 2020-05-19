package token

// TokenType definition
type TokenType string

// Token struct definition
type Token struct {
	Type    TokenType
	Literal string
}

// Token list definition
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Operatos + Literals
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
