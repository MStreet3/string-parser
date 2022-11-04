package tokenizer

import (
	"regexp"
	"strings"
)

type Interface interface {
	HasMoreTokens() bool
	GetNextToken() *Token
}

type TokenName string

const (
	NUM     TokenName = "NUMBER"
	ADD_OP  TokenName = "ADDITIVE_OPERATOR"
	O_PAREN TokenName = "OPEN_PAREN"
	C_PAREN TokenName = "CLOSE_PAREN"
)

type TokenSpecification struct {
	regex string
	name  TokenName
}

type basicTokenizer struct {
	Stack  []string
	Cursor int
	spec   []TokenSpecification
}

type Token struct {
	Type  TokenName
	Value string
}

func NewBasicTokenizer(expr string) *basicTokenizer {
	// tokenizer specification
	var spec = []TokenSpecification{
		{regex: `^(\-)?\d+`, name: NUM},
		{regex: `^[+\-]`, name: ADD_OP},
		{regex: `^\(`, name: O_PAREN},
		{regex: `^\)`, name: C_PAREN},
	}

	return &basicTokenizer{
		Stack:  strings.Fields(expr),
		Cursor: 0,
		spec:   spec,
	}
}

func (t *basicTokenizer) HasMoreTokens() bool {
	return t.Cursor < len(t.Stack)
}

// GetNextToken returns the token for the current raw value cursor or nil advancing cursor
// if token is found.
func (t *basicTokenizer) GetNextToken() *Token {
	if !t.HasMoreTokens() {
		return nil
	}
	raw := t.Stack[t.Cursor]
	for _, s := range t.spec {
		if matched, _ := regexp.MatchString(s.regex, raw); matched {
			t.Cursor++
			res := &Token{
				Type:  s.name,
				Value: raw,
			}
			return res
		}
	}
	return nil
}
