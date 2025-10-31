# Golang

## [Golang Cohort Link ](https://petal-estimate-4e9.notion.site/Golang-cohort-1257dfd1073580238258fe25973c319b)

## Go Semicolon Rules --- Quick Notes 📝

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

## Go Variables --- Quick Notes 📝

Variables in Go can be declared in **three main ways**.

---

## 1. Using `var` with Type (Explicit Declaration)

```go
var name string = "Manikanta Korimilli"
fmt.Println(name)
fmt.Printf("The type: %T\n", name)
```

- Declares a variable **with a fixed type**.
- You **must** assign the value at declaration.

---

## 2. Using `var` without Type (Type Inference)

```go
var totalAmount = 40
fmt.Println(totalAmount)
fmt.Printf("The type: %T\n", totalAmount)
```

- Go **infers the type** automatically based on the value.

---

## 3. Using `:=` Operator (Short Variable Declaration)

```go
isLoggin := true
fmt.Println(isLoggin)
fmt.Printf("The type: %T\n", isLoggin)
```

- Works **only inside functions**.

---

## 4. Constants in Go

```go
const apiKey = "69:69"
fmt.Println(apiKey)
fmt.Printf("The type: %T\n", apiKey)
```

- Use `const` for values that **never change**.
- Constants **must** be assigned when declared.

---

## Key Takeaways ✅

- **`var name string = "Go"`** → Declare variable **with type**.
- **`var name = "Go"`** → Go **infers type** automatically.
- **`name := "Go"`** → Short variable declaration (**inside functions
  only**).
- **`const apiKey = "12345"`** → Declare a constant.

## [userInput](/03userinput/main.go)

## [conversion](/04conversion/main.go)

## [Time](/05learn/main.go)

## [Pointers](/06mypointers/main.go)

## [Arrays](/07myarrays/main.go)

## [Slices](/08myslices/main.go)

## [Maps](/09mymaps/main.go)

## [Structs](/10mystructs/main.go)

## [IfElse](/11IfElse/main.go)

## [Loops In Go](/12Loops/main.go)

## [Functions](/13Functions/main.go)

## [Methods](/14Methods/main.go)

## [Defter](/15Defer/main.go)

## [File Operations](/16Files/main.go)

## [Web Requests](/17WebRequests/main.go)

## [URL Handling](/18HandlingURL/main.go)

## [Get Request](/20GetRequest/main.go)

## [Post Request](/21PostRequest/main.go)

## [Post Form Request](/22PostFormRequest/main.go)

# 🔍 Key Differences: Gorilla Mux vs Chi Router

This document highlights the key differences between **Gorilla Mux** and **Chi** routers in Go.

---

## ⚙️ Comparison Table

| Feature / Concept         | Gorilla Mux 🦍                                     | Chi 🐹                                        | Notes                                             |
| ------------------------- | -------------------------------------------------- | --------------------------------------------- | ------------------------------------------------- |
| **Create a router**       | `r := mux.NewRouter()`                             | `r := chi.NewRouter()`                        | Both create a new router instance                 |
| **Get URL params**        | `mux.Vars(r)["id"]`                                | `chi.URLParam(r, "id")`                       | Chi uses a simpler function call                  |
| **Register GET route**    | `r.HandleFunc("/path", handler).Methods("GET")`    | `r.Get("/path", handler)`                     | Chi has built-in HTTP method helpers              |
| **Register POST route**   | `r.HandleFunc("/path", handler).Methods("POST")`   | `r.Post("/path", handler)`                    | Cleaner syntax                                    |
| **Register PUT route**    | `r.HandleFunc("/path", handler).Methods("PUT")`    | `r.Put("/path", handler)`                     | Consistent style                                  |
| **Register DELETE route** | `r.HandleFunc("/path", handler).Methods("DELETE")` | `r.Delete("/path", handler)`                  | Consistent style                                  |
| **Subrouting / grouping** | `r.PathPrefix("/api").Subrouter()`                 | `r.Route("/api", func(r chi.Router) { ... })` | Chi uses function-based grouping                  |
| **Middleware usage**      | `r.Use(middleware)`                                | `r.Use(middleware)`                           | Same concept, but Chi provides many built-in ones |
| **Project status**        | ❌ Archived                                        | ✅ Actively maintained                        | Gorilla Mux is no longer maintained               |
| **Performance**           | Slower (reflection heavy)                          | Faster, lightweight                           | Chi is optimized for speed and memory             |

---

## 🧩 Example Comparison

### Gorilla Mux

```go
r := mux.NewRouter()
r.HandleFunc("/course/{id}", getCourse).Methods("GET")

vars := mux.Vars(r)
id := vars["id"]
```

# Let's Go..!
