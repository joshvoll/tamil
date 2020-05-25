package token

// TokenType for different type of token
type TokenType string

// Token structures
type Token struct {
	Type    TokenType
	Literal string
}

// tokens definitions
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Ident and literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","
	SEMICOLON = ";"

	//  keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	FALSE    = "FALSE"
	TRUE     = "TRUE"
	RETURN   = "RETURN"
	LT       = "<"
	GT       = ">"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"false":  FALSE,
	"true":   TRUE,
	"return": RETURN,
}

// LookupIdent return the TokenType
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
