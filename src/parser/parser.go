package main

import (
	"fmt"
	"strconv"
	"strings"
	"tokenizer"
)

type GrammarProduction string

const (
	NumericLiteral    GrammarProduction = "NUMERIC_LITERAL"
	BinaryExpression  GrammarProduction = "BINARY_EXPRESSION"
	ADDITIVE_OPERATOR GrammarProduction = "ADDITIVE_OPERATOR"
)

type ASTree struct {
	Type string
	Body ASTNode
}
type ASTNode struct {
	Type     GrammarProduction
	Value    interface{}
	Operator string
	left     *ASTNode
	right    *ASTNode
}

type Parser interface {
	Parse(p string) *ASTNode
	BinaryExpression() *ASTNode
	NumericLiteral() *ASTNode
}

type StringParser struct {
	state     string
	tokenizer *tokenizer.Tokenizer
	lookAhead *tokenizer.Token
}

func NewStringParser() *StringParser {
	sp := &StringParser{
		state: "",
	}
	return sp
}

func (sp *StringParser) Parse(p string) *ASTree {
	sp.state = p
	sp.tokenizer = &tokenizer.Tokenizer{Stack: strings.Fields(p), Cursor: 0}
	sp.lookAhead = sp.tokenizer.GetNextToken()
	return &ASTree{Type: "AbstractSyntaxTree", Body: *sp.BinaryExpression()}
}

func (sp *StringParser) BinaryExpression() *ASTNode {
	left := sp.NumericLiteral()
	for sp.lookAhead != nil && sp.lookAhead.Type == "ADDITIVE_OPERATOR" {
		operator := sp.eat("ADDITIVE_OPERATOR")
		right := sp.NumericLiteral()

		left = &ASTNode{
			Type:     BinaryExpression,
			Operator: operator.Value.(string),
			left:     left,
			right:    right,
		}
	}
	return left
}

func (sp *StringParser) NumericLiteral() *ASTNode {
	token := sp.eat("NUMBER")
	if value, err := strconv.Atoi(token.Value.(string)); err == nil {
		return &ASTNode{
			Type:  NumericLiteral,
			Value: value,
		}
	}
	return nil

}

func (sp *StringParser) eat(tokenType string) *tokenizer.Token {
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

func (n *ASTNode) Evaluate() int {
	if n.Type == NumericLiteral {
		return n.Value.(int)
	} else if n.Operator == "+" {
		return n.left.Evaluate() + n.right.Evaluate()
	} else {
		return n.left.Evaluate() - n.right.Evaluate()
	}
}

func main() {
	fmt.Println("IN MAIN")
	p := NewStringParser()
	/* firstTree := p.Parse("42 + 13")
	fmt.Println(*firstTree.Body.left)
	fmt.Println(firstTree.Body.Operator)
	fmt.Println(*firstTree.Body.right) */
	secondTree := p.Parse("42 + 13 - 1")
	fmt.Println(*secondTree.Body.left)
	fmt.Println(secondTree.Body.Operator)
	fmt.Println(*secondTree.Body.right)
	fmt.Println(secondTree.Body.Evaluate())

}
