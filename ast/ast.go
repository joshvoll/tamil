package ast

import "github.com/joshvoll/tamil/token"

// Node representation
type Node interface {
	TokenLiteral() string
}

// Statement representation
type Statement interface {
	Node
	statementNode()
}

// Expression representation
type Expression interface {
	Node
	expressionNode()
}

// Program definition
type Program struct {
	Statements []Statement
}

// TokenLiteral implements Node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement represent the LET statment variable binding
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// statementNode helper method
func (ls *LetStatement) statementNode() {}

// TokenLiteral helper method
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier definition
type Identifier struct {
	Token token.Token
	Value string
}

// expresionNode implements Expression
func (i *Identifier) expressionNode() {}

// TokenLiteral implements Node
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
