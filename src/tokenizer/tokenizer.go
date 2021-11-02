package tokenizer

import (
	"regexp"
	"strings"
)

type Tokenizer interface {
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

/* tokenizer specification */
var specification = []TokenSpecification{
	{regex: `^\d+`, name: NUM},
	{regex: `^[+\-]`, name: ADD_OP},
	{regex: `^\(`, name: O_PAREN},
	{regex: `^\)`, name: C_PAREN},
}

type BasicTokenizer struct {
	Stack  []string
	Cursor int
}
type Token struct {
	Type  TokenName
	Value string
}

func NewBasicTokenizer(expr string) *BasicTokenizer {
	return &BasicTokenizer{
		Stack:  strings.Fields(expr),
		Cursor: 0,
	}
}

func (t *BasicTokenizer) HasMoreTokens() bool {
	return t.Cursor < len(t.Stack)
}

/* returns token at current cursor or nil advancing cursor
if token is found. */
func (t *BasicTokenizer) GetNextToken() *Token {
	if !t.HasMoreTokens() {
		return nil
	}
	raw := t.Stack[t.Cursor]
	for _, s := range specification {
		matched, _ := regexp.MatchString(s.regex, raw)
		if matched {
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
