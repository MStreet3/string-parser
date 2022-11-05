package ast

import (
	"errors"

	"github.com/mstreet3/rdp/pkg/entities"
)

var (
	ErrInvalidNodeType = errors.New("invalid node type")
	ErrInvalidOperator = errors.New("invalid operator")
)

type ASTree struct {
	root entities.Evaluator
}

func NewASTree(n entities.Evaluator) *ASTree {
	return &ASTree{
		root: n,
	}
}

func (ast *ASTree) Evaluate() (int, error) {
	return ast.root.Evaluate()
}

type ASTNode struct {
	Type     GrammarProduction
	Value    int
	Operator string
	left     *ASTNode
	right    *ASTNode
}

// Evaluate returns the evaluated expression of the AST
func (n *ASTNode) Evaluate() (int, error) {
	if n.Type == NumericLiteral {
		return n.Value, nil
	}

	if n.Type == BinaryExpression {
		left, err := n.left.Evaluate()
		if err != nil {
			return 0, err
		}

		right, err := n.right.Evaluate()
		if err != nil {
			return 0, err
		}

		switch n.Operator {
		case "+":
			return left + right, nil
		case "-":
			return left - right, nil
		default:
			return 0, ErrInvalidOperator
		}

	}

	return 0, ErrInvalidNodeType
}
