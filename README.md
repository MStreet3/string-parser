### expression parsing

implement a parser to read an expression as a string and return
the value of the expression as an integer.

develops a manual `recursive descent parser`.

to execute the tests

```bash
$ cd recursiveDescentParser
$ cd ./src/parser
$ go run parser.go
```

this `tokenizer` is rather basic and expects that all tokenizable `bytes` are
separated by white space. the specification is also limited.

```go
/* tokenizer specification */
var specification = []TokenSpecification{
	{regex: `^\d+`, name: "NUMBER"},
	{regex: `^[+\-]`, name: "ADDITIVE_OPERATOR"},
	{regex: `^\(`, name: "OPEN_PAREN"},
	{regex: `^\)`, name: "CLOSE_PAREN"},
}
```
