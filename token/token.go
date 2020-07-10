package token

// TokenType definition
type TokenType string

// Token definition
type Token struct {
	Type    TokenType
	Literal string
}

// list of tokens
const (
	// Illegal and End of file
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Ident and Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"

	// Delimeters
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	EQ        = "=="
	NOT_EQ    = "!="
	COMMA     = ","
	SEMICOLON = ";"
	GT        = ">"
	LT        = "<"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

// keywors maps
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// LookupIdent return the token type of toke pass
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
