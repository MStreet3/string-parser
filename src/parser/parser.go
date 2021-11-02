package main

import (
	"fmt"
	"strconv"
	"strings"
	"tokenizer"
)

type Node interface {
	Evaluate() int
}

type GrammarProduction string

const (
	NumericLiteral   GrammarProduction = "NUMERIC_LITERAL"
	BinaryExpression GrammarProduction = "BINARY_EXPRESSION"
)

type ASTree struct {
	Root ASTNode
}
type ASTNode struct {
	Type     GrammarProduction
	Value    int
	Operator string
	left     *ASTNode
	right    *ASTNode
}

type Parser interface {
	Parse(expr string) *ASTree
	BinaryExpression() *ASTNode
	PrimaryExpression() *ASTNode
	ParenthesizedExpression() *ASTNode
	NumericLiteral() *ASTNode
	eat(tokenType tokenizer.TokenName) *tokenizer.Token
}
type BasicParser struct {
	state     string
	tokenizer *tokenizer.Tokenizer
	lookAhead *tokenizer.Token
}

func NewBasicParser() *BasicParser {
	sp := &BasicParser{
		state: "",
	}
	return sp
}

/* BasicParser accepts a string and returns the abstract syntax tree
representation of the string while assuming that all tokenizable bytes
are separated by a space. */
func (sp *BasicParser) Parse(p string) *ASTree {
	sp.state = p
	sp.tokenizer = &tokenizer.Tokenizer{Stack: strings.Fields(p), Cursor: 0}
	sp.lookAhead = sp.tokenizer.GetNextToken()
	return &ASTree{Root: *sp.BinaryExpression()}
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
func Calculate(expr string) int {
	sp := NewBasicParser()
	ast := sp.Parse(expr)
	return ast.Root.Evaluate()
}

func main() {
	cases := []string{
		"1 - 2 + 3",
		"1 - ( 2 + 3 )",
		"1 + 2",
		"( 1 )",
		"( 1 - 2 ) + ( 3 + 3 )",
		"0",
		"( ( 1 - 5 ) + 4 ) + ( 4 - 1 )",
		"( ( 1 - 5 ) + ( 4 + ( 3 ) ) ) + ( 4 - ( ( 1 ) ) )",
	}
	expected := []int{
		2,
		-4,
		3,
		1,
		5,
		0,
		3,
		6,
	}

	passed := true
	for i, test := range cases {
		result := Calculate(test)
		if result != expected[i] {
			fmt.Printf("failed case %d: %s\n", i, test)
			passed = false
		}

	}
	if passed {
		fmt.Println("all test cases passed!")
	}

}
