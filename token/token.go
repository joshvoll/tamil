package token

// TokenType defenition
type TokenType string

// Token definition
type Token struct {
	Type    TokenType
	Literal string
}

// Tokens Lists
const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Ident and Literals
	INT   = "INT"
	IDENT = "IDENT"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASSGIN   = "="
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimeters
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","
	SEMICOLON = ";"
	LT        = "<"
	GT        = ">"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

// map of tokens
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// LookupIdent method to return the token types
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
