package token

// TokenType is for define lot of token types
type TokenType string

// Token definition
type Token struct {
	Type    TokenType
	Literla string
}

// Tokens definitions
const (
	// Genral operation
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Ident and Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "ASSIGN"
	PLUS   = "+"

	// Delimeters
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","
	SEMICOLON = ";"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
