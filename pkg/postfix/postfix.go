package postfix

import (
	"errors"
	"strconv"
	"strings"

	"github.com/mstreet3/rdp/pkg/entities"
	"github.com/mstreet3/rdp/pkg/tokenizer"
)

var (
	ErrInvalidPostfixExpression = errors.New("invalid postfix expression")
	ErrUnexpectedInput          = errors.New("unexpected input")
	ErrMismatchedParentheses    = errors.New("mismatched parentheses")
)

type expression struct {
	tokenizer tokenizer.Interface
}

var _ entities.Evaluator = (*expression)(nil)

type BasicParser struct{}

// Parse expects an infix expression with space separated tokens and returns a
// postfix expression or an error.
func (p *BasicParser) Parse(expr string) (entities.Evaluator, error) {
	postfix, err := parse(expr)
	if err != nil {
		return nil, err
	}
	return &expression{
		tokenizer: tokenizer.NewBasicTokenizer(postfix),
	}, nil
}

func (e *expression) Evaluate() (int, error) {
	var stack []int

	for e.tokenizer.HasMoreTokens() {
		next := e.tokenizer.GetNextToken()
		if next == nil {
			return 0, ErrUnexpectedInput
		}

		if next.Type == tokenizer.NUM {
			val, err := strconv.Atoi(next.Value)
			if err != nil {
				return 0, err
			}
			stack = append([]int{val}, stack...)
			continue
		}

		if next.Type == tokenizer.ADD_OP {
			if len(stack) < 2 {
				return 0, ErrInvalidPostfixExpression
			}

			var (
				right = stack[0]
				left  = stack[1]
				sum   int
			)

			if next.Value == "+" {
				sum = left + right
			}

			if next.Value == "-" {
				sum = left - right
			}

			stack = append([]int{sum}, stack[2:]...)
		}
	}

	if len(stack) == 1 {
		return stack[0], nil
	}

	return 0, nil
}

// parse implements a simplified version of Djikstra's "Shunting Yard" alg
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm#The_algorithm_in_detail
func parse(expr string) (string, error) {
	var (
		postfix []string
		opStack []string
	)
	tokens := tokenizer.NewBasicTokenizer(expr)

	for tokens.HasMoreTokens() {
		next := tokens.GetNextToken()
		if next == nil {
			return "", ErrUnexpectedInput
		}

		if next.Type == tokenizer.NUM {
			postfix = append(postfix, next.Value)
			continue
		}

		if next.Type == tokenizer.ADD_OP {
			// TODO: check operator precedence.  Okay for now
			// because only additive operators are allowed.  Any
			// additive operator in the stack should be pushed to
			// the postfix queue as they are left associative.
			for len(opStack) > 0 && opStack[0] != "(" {
				postfix = append(postfix, opStack[0])
				opStack = opStack[1:]
			}
			opStack = append([]string{next.Value}, opStack...)
			continue
		}

		if next.Type == tokenizer.O_PAREN {
			opStack = append([]string{next.Value}, opStack...)
			continue
		}

		if next.Type == tokenizer.C_PAREN {
			for len(opStack) > 0 && opStack[0] != "(" {
				if len(opStack) == 0 {
					return "", ErrMismatchedParentheses
				}
				postfix = append(postfix, opStack[0])
				opStack = opStack[1:]
			}
			if len(opStack) > 0 && opStack[0] != "(" {
				return "", ErrMismatchedParentheses
			}
			if len(opStack) > 0 {
				opStack = opStack[1:]
			}
		}
	}

	for _, token := range opStack {
		if token == string(tokenizer.O_PAREN) {
			return "", ErrMismatchedParentheses
		}
		postfix = append(postfix, token)
	}

	return strings.Join(postfix, " "), nil
}
