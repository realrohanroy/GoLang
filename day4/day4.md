
# Day 4 Summary - Go Loops & Strings

## 🚀 What I Learned Today

### 1. Loops in Go
- **Classic for loop**
  ```go
  for i := 0; i < 5; i++ {
      fmt.Println(i)
  }
  ```

- **While-like loop**
  ```go
  i := 0
  for i < 5 {
      fmt.Println(i)
      i++
  }
  ```

- **Infinite loop**
  ```go
  for {
      fmt.Println("Runs forever!")
      break // use break to stop
  }
  ```

- **Range loop (for collections)**
  ```go
  nums := []int{10, 20, 30}
  for index, value := range nums {
      fmt.Println(index, value)
  }
  ```

### 2. Strings in Go
- Strings are **immutable** (cannot be changed directly).
- Stored as bytes, but for Unicode (like emojis/accents), we use **runes**.

- **Iterating over a string**
  ```go
  s := "Go💖"
  for i, r := range s {
      fmt.Printf("%d: %c
", i, r)
  }
  ```

- **String functions**
  - `len(s)` → number of **bytes**.
  - `utf8.RuneCountInString(s)` → number of **characters/runes**.
  - `strings.Contains(s, "Go")`
  - `strings.HasPrefix(s, "Go")`
  - `strings.Split(s, ",")`
  - `strings.Join([]string{"Go", "Lang"}, " ")`

### 3. Best Practices
- Always use `range` when working with strings/maps.
- Use `rune` for Unicode-safe iteration.
- Keep loops clean, avoid complex conditions inside.
- Prefer readability over clever tricks.

---

## ✅ Confidence Boost
- I can now **write any type of loop in Go**.
- I can safely **work with strings** including Unicode.
- I know **helper functions** in `strings` package.

---

## 🔥 Next Steps
Tomorrow’s mini-project will:
- Use loops + strings in real code.
- Combine slices, maps, and functions.
- Reinforce everything learned so far.

---

