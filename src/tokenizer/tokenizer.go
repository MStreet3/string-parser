package tokenizer

import (
	"regexp"
)

type TokenSpecification struct {
	regex string
	name  string
}

/* tokenizer specification */
var specification = []TokenSpecification{
	{regex: `^\d+`, name: "NUMBER"},
	{regex: `^[+\-]`, name: "ADDITIVE_OPERATOR"},
	{regex: `^\(`, name: "OPEN_PAREN"},
	{regex: `^\)`, name: "CLOSE_PAREN"},
}

type Tokenizer struct {
	Stack  []string
	Cursor int
}
type Token struct {
	Type  string
	Value interface{}
}

func (t *Tokenizer) HasMoreTokens() bool {
	return t.Cursor < len(t.Stack)
}

func (t *Tokenizer) GetNextToken() *Token {
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
