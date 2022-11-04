### basic parsing

implements a manual `recursive descent parser` to read an expression as a string and return
the value of the expression as an integer.

```go
calculator := NewBasicCalculator()
calculator.Calculate("1 - 2 + 3")      //  2
calculator.Calculate("1 - ( 2 + 3 )")  // -4
calculator.Calculate("1 - ( ( 2 + 3 ) + 1 )")  // -5
```

The `tokenizer` is basic and expects that all tokenizable elements of the expression are
separated by white space. any tokenizer, however, that fulfills the `Tokenizer` interface can
be used.

```go
type Tokenizer interface {
	HasMoreTokens() bool
	GetNextToken() *Token
}
```

The token specification used by the parser is limited to four distinct `runes`:

```go
var specification = []TokenSpecification{
	{regex: `^(\-)?\d+`, name: NUM},
	{regex: `^[+\-]`, name: ADD_OP},
	{regex: `^\(`, name: O_PAREN},
	{regex: `^\)`, name: C_PAREN},
}
```

### tests

To execute the tests

```bash
$ cd ./string-parser
$ go test ./... -v
```

### References

Content in this repo is inspired and guided by the [Parser from Scratch](http://dmitrysoshnikov.com/courses/parser-from-scratch/)
course in JavaScript.
