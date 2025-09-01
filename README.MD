# Golang

## [Golang Cohort Link ](https://petal-estimate-4e9.notion.site/Golang-cohort-1257dfd1073580238258fe25973c319b)

## Go Semicolon Rules

### 1. Automatic Semicolons

Go **adds semicolons automatically** at the end of a line if it ends with:

- **Identifiers** → `fmt.Println(name)`
- **Literals** → `x := 10`
- **Keywords** → `return`, `break`, `continue`, `fallthrough`
- **Symbols** → `)`, `]`, `}`, `++`, `--`

---

### 2. No Semicolon When Statement Continues

If the next token is on the **same line**, Go **doesn't insert** one:

```go
if x < 5 {
    fmt.Println("Less") // semicolon added
} else {               // no semicolon here
    fmt.Println("More") // semicolon added
}                       // semicolon added
```

### 3. Multiple Statements on One Line

You **must add semicolons manually**:

```go
x := 5; y := 10; fmt.Println(x + y)
```
