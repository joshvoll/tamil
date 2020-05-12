package token

// TokenType  is defined as string that allows us to create mandy different values
type TokenType string

// Token definition
type Token struct {
	Type    TokenType
	Literal string
}

// WE need to define the token defintion list, this is key for every language
// Here is examples how the golang team is doing that definition
// https://github.com/golang/go/blob/master/src/go/token/token.go
// For now we're going to scape some language feature like comments and numbers of line
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and Literals
	IDENT = "IDENT" // add(), server(), add , etc...
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
