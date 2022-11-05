package ast

import (
	"strconv"

	"github.com/mstreet3/rdp/pkg/entities"
	"github.com/mstreet3/rdp/pkg/tokenizer"
)

type GrammarProduction string

const (
	NumericLiteral   GrammarProduction = "NUMERIC_LITERAL"
	BinaryExpression GrammarProduction = "BINARY_EXPRESSION"
)

type BasicParser struct {
	tokenizer tokenizer.Interface
	lookAhead *tokenizer.Token
}

// BasicParser accepts a string and returns the abstract syntax tree
// representation of the string while assuming that all tokenizable bytes
// are separated by a space.
func (bp *BasicParser) Parse(p string) (entities.Evaluator, error) {
	bp.tokenizer = tokenizer.NewBasicTokenizer(p)
	bp.lookAhead = bp.tokenizer.GetNextToken()
	return NewASTree(bp.BinaryExpression()), nil
}

// BinaryExpression
//
//	: PrimaryExpression
//	| PrimaryExpression ADDITIVE_OPERATOR PrimaryExpression
//	;
func (bp *BasicParser) BinaryExpression() *ASTNode {
	left := bp.PrimaryExpression()
	for bp.lookAhead != nil && bp.lookAhead.Type == tokenizer.ADD_OP {
		operator := bp.eat(tokenizer.ADD_OP)
		right := bp.PrimaryExpression()
		left = &ASTNode{
			Type:     BinaryExpression,
			Operator: operator.Value,
			left:     left,
			right:    right,
		}
	}
	return left
}

// PrimaryExpression
//
//	: NumericLiteral
//	| ParenthesizedExpression
//	;
func (bp *BasicParser) PrimaryExpression() *ASTNode {
	switch bp.lookAhead.Type {
	case tokenizer.O_PAREN:
		return bp.ParenthesizedExpression()
	default:
		return bp.NumericLiteral()
	}
}

// ParenthesizedExpression
//
//	: "(" BinaryExpression ")"
//	;
func (bp *BasicParser) ParenthesizedExpression() *ASTNode {
	// eat and discard the parentheses, returning the expression
	bp.eat(tokenizer.O_PAREN)
	expression := bp.BinaryExpression()
	bp.eat(tokenizer.C_PAREN)
	return expression
}

// NumericLiteral
//
//	: NUMBER
//	;
func (bp *BasicParser) NumericLiteral() *ASTNode {
	token := bp.eat(tokenizer.NUM)
	if value, err := strconv.Atoi(token.Value); err == nil {
		return &ASTNode{
			Type:  NumericLiteral,
			Value: value,
		}
	}
	return nil

}

// eat is a helper function that validates the token from the tokenizer
// and steps the parser forward
func (bp *BasicParser) eat(tokenType tokenizer.TokenName) *tokenizer.Token {
	token := bp.lookAhead
	if token == nil {
		panic("unexpected input")
	}
	if token.Type != tokenType {
		panic("unexpected token type")
	}
	bp.lookAhead = bp.tokenizer.GetNextToken()
	return token
}
