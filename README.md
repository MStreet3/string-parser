### expression parsing

implement a parser to read an expression as a string and return
the value of the expression as an integer.

develops a manual `recursive descent parser`.

### abstract syntax trees

a json object with a `type`, `operator`, `left` and `right`.

```json
{
  "type": "BinaryExpression",
  "operator": "+",
  "left": {
    "type": "NumericalLiteral",
    "value": 3
  },
  "right": {
    "type": "NumericalLiteral",
    "value": 4
  }
}
```
