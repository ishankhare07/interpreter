package ast

import "gitlab.com/ishankhare07/monkey-lang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// define type LetStatement

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// implement Node interface

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// implement Statement interface

func (ls *LetStatement) statementNode() {}

// define Identifier type

type Identifier struct {
	Token token.Token
	Value string
}

// implement Node interface

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// implement Expression interface

func (i *Identifier) expressionNode() {}

// define type ReturnStatement

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// implement Node interface
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// implement Statement interface

func (rs *ReturnStatement) statementNode() {}
