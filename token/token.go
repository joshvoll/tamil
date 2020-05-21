package token

// TokenType is for have multiple type of values
type TokenType string

// Token model definition
type Token struct {
	Type    TokenType
	Literal string
}

// tokens definitions
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT" // add, foo,
	INT   = "INT"   // 1234,4343,454

	// Operators
	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// keywords properties just got all tne keyword in one slice
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent is going to return the right tokentype of the keywords
// checks the keywords table to see whether the given identifier is in fact a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
