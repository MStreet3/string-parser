package main

import (
	"strconv"
	"tokenizer"
)

type Node interface {
	Evaluate() int
}
type Parser interface {
	Parse(expr string) *ASTree
}
type Calculate func(expr string) int
type Calculator interface {
	Calculate(expr string) int
}

type GrammarProduction string

const (
	NumericLiteral   GrammarProduction = "NUMERIC_LITERAL"
	BinaryExpression GrammarProduction = "BINARY_EXPRESSION"
)

type ASTree struct {
	Root *ASTNode
}
type ASTNode struct {
	Type     GrammarProduction
	Value    int
	Operator string
	left     *ASTNode
	right    *ASTNode
}

type BasicParser struct {
	tokenizer tokenizer.Tokenizer
	lookAhead *tokenizer.Token
}

/* BasicParser accepts a string and returns the abstract syntax tree
representation of the string while assuming that all tokenizable bytes
are separated by a space. */
func (sp *BasicParser) Parse(p string) *ASTree {
	sp.tokenizer = tokenizer.NewBasicTokenizer(p)
	sp.lookAhead = sp.tokenizer.GetNextToken()
	return &ASTree{
		Root: sp.BinaryExpression(),
	}
}

/*
BinaryExpression
	: PrimaryExpression
	| PrimaryExpression ADDITIVE_OPERATOR PrimaryExpression
	;
*/
func (sp *BasicParser) BinaryExpression() *ASTNode {
	left := sp.PrimaryExpression()
	for sp.lookAhead != nil && sp.lookAhead.Type == tokenizer.ADD_OP {
		operator := sp.eat(tokenizer.ADD_OP)
		right := sp.PrimaryExpression()
		left = &ASTNode{
			Type:     BinaryExpression,
			Operator: operator.Value,
			left:     left,
			right:    right,
		}
	}
	return left
}

/*
PrimaryExpression
	: NumericLiteral
	| ParenthesizedExpression
	;
*/
func (sp *BasicParser) PrimaryExpression() *ASTNode {
	switch sp.lookAhead.Type {
	case tokenizer.O_PAREN:
		return sp.ParenthesizedExpression()
	default:
		return sp.NumericLiteral()
	}
}

/*
ParenthesizedExpression
	: "(" BinaryExpression ")"
	;
*/
func (sp *BasicParser) ParenthesizedExpression() *ASTNode {
	// eat and discard the parentheses, returning the expression
	sp.eat(tokenizer.O_PAREN)
	expression := sp.BinaryExpression()
	sp.eat(tokenizer.C_PAREN)
	return expression
}

/*
NumericLiteral
	: NUMBER
	;
*/
func (sp *BasicParser) NumericLiteral() *ASTNode {
	token := sp.eat(tokenizer.NUM)
	if value, err := strconv.Atoi(token.Value); err == nil {
		return &ASTNode{
			Type:  NumericLiteral,
			Value: value,
		}
	}
	return nil

}

/*
eat is a helper function that validates the token from the tokenizer
and steps the parser forward
*/
func (sp *BasicParser) eat(tokenType tokenizer.TokenName) *tokenizer.Token {
	token := sp.lookAhead
	if token == nil {
		panic("unexpected input")
	}
	if token.Type != tokenType {
		panic("unexpected token type")
	}
	sp.lookAhead = sp.tokenizer.GetNextToken()
	return token
}

/* Evaluate returns the evaluated expression of the AST */
func (n *ASTNode) Evaluate() int {
	if n.Type == NumericLiteral {
		return n.Value
	} else if n.Operator == "+" {
		return n.left.Evaluate() + n.right.Evaluate()
	} else {
		return n.left.Evaluate() - n.right.Evaluate()
	}
}

/*
Calculate accepts a string of addition / subtraction operations
and also parentheses to indicate order of operations.
- assumes all string characters are separated by white space
*/

func NewCalculate(p Parser) Calculate {
	calculate := func(expr string) int {
		ast := p.Parse(expr)
		return ast.Root.Evaluate()
	}
	return calculate
}

type ParsingCalculator struct {
	parser Parser
}

func (pc ParsingCalculator) Calculate(expr string) int {
	ast := pc.parser.Parse(expr)
	return ast.Root.Evaluate()
}

func NewCalculator(p Parser) Calculator {
	return ParsingCalculator{parser: p}
}

func NewBasicParsingCalculator() Calculator {
	return ParsingCalculator{
		parser: &BasicParser{},
	}
}
