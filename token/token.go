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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent checks the keywords table to see whether the given identifier is in fact a keyword.
// If it is, it returns the keyword’s TokenType constant
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
