package ast

import "github.com/joshvoll/tamil/token"

// Node hold the logic of the ast
type Node interface {
	TokenLiteral() string
}

// Expression define the expression method
type Expression interface {
	Node
	expressionNode()
}

// Statement define the statements method
type Statement interface {
	Node
	statementNode()
}

// Program definition
type Program struct {
	Statements []Statement
}

// TokenLiteral implements Node interface
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement definition
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// statementNode() implementation fro letStatement
func (ls *LetStatement) statementNode() {}

// TokenLiteral implementation for letStatement
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier definition
type Identifier struct {
	Token token.Token
	Value string
}

// statementNode() implements for Identifier
func (i *Identifier) expressionNode() {}

// TokenLiteral implements for Identifier
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
