package token

// TokenType definition
type TokenType string

// Token struct definition
type Token struct {
	Type    TokenType
	Literal string
}

// all the tokens lists
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Ident and Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// Delimeters
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","
	SEMICOLON = ";"
	GT        = ">"
	LT        = "<"
	EQ        = "=="
	NOT_EQ    = "!="

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

// keywords array
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// LookupIdent retur the token type of the given token
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
