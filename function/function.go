package main

import "fmt"

func add(a int, b int) int { // a and b are parameters of the function and int is the return type
	return a + b
}

func addStrings(a string, b string) string {
	return a + b
}

func addAndSubtract(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func combineStrings(a string, b string) string {
	if a < b {
		return a + " " + b
	}
	return "Strings not combined"
}

// function with conditions
func function(val1, val2 interface{}) interface{} {
	switch v1 := val1.(type) {
	case int:
		if v2, ok := val2.(int); ok {
			if v1 > v2 {
				return v1
			}
		}
	case string:
		if v2, ok := val2.(string); ok {
			if v1 < v2 {
				return v1 + " " + v2
			}
		}
	}
	return "Something else"
}

func main() {
	result := add(1, 2)
	fmt.Println(result)

	result2 := addStrings("Hello ", "World!")
	fmt.Println(result2)

	result3, _ := addAndSubtract(5, 2)
	fmt.Println(result3)

	result4 := function(10, 5)
	fmt.Println(result4)

	combined := combineStrings("apple", "banana")
	fmt.Println("Combined strings:", combined)

	result5 := function("Hello", "World!")
	fmt.Println(result5)

}
