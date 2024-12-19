Backus Naur Form (BNF) for 'mul(X,Y)' expression where X, Y are ints in
range [-999, 999]:

---

```
<program> ::= <instruction> | <instruction> <program>
<instruction> ::= "mul" "(" <number> "," <number> ")"
<number> ::= <digit> | <digit> <digit> | <digit> <digit> <digit>
<digit> ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```

---
