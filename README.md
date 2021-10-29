### basic parsing

expose a `Calculate` method that accepts a string of addition / subtraction operations
and also parentheses to indicate order of operations.

- assumes all string characters are separated by white space

develops a manual `recursive descent parser` to read an expression as a string and return
the value of the expression as an integer.

the `tokenizer` is rather basic and expects that all tokenizable `bytes` are
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

### tests

to execute the tests

```bash
$ cd recursiveDescentParser
$ cd ./src/parser
$ go run parser.go
```

content in this repo is inspired and guided by the [Parser from Scratch](http://dmitrysoshnikov.com/courses/parser-from-scratch/)
course in JavaScript.
