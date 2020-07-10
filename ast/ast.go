package ast

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
