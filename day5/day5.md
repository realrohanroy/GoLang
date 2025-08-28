# Day 5 - Go Learning Journey

## Topics Covered
- **Functions in Go**
  - Function declaration & parameters
  - Returning multiple values
  - Error handling with `error` type
- **Practice**
  - Implemented `avg()` function to calculate sum & average of a slice
  - Handled edge case of empty slice with `fmt.Errorf`
  - Returned `sum`, `average`, and `error`

## Code Example
```go
package main

import "fmt"

func avg(num []int) (int, int, error) {
	sum := 0
	n := len(num)
	if n == 0 {
		return 0, 0, fmt.Errorf("empty slice")
	}
	for _, i := range num {
		sum = sum + i
	}
	average := sum / n
	return sum, average, nil
}

func main() {
	num := []int{}
	sum, avg, err := avg(num)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Sum:", sum, "Average:", avg)
}
