package token

// TokenType for multiple types of the tokens
type TokenType string

// Token definitions
type Token struct {
	Type    TokenType
	Literal string
}

// tokens defintitions
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASSIGN   = "="
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// Delimeters
	SEMICOLON = ";"
	COMMA     = ","
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	GT        = ">"
	LT        = "<"
	EQ        = "=="
	NOT_EQ    = "!="

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

// keywords properties
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// LookupIdent return the TokenType of a token
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
