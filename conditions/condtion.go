package main

import "fmt"

func main() {

	// if else statement
	age := 20
	if age == 18 {
		fmt.Println("Congratulations! You can vote")
	} else if age > 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}

	// switch statement
	switch age {
	case 18:
		fmt.Println("You can vote")
	case 21:
		fmt.Println("You can vote")
	default:
		fmt.Println("You can't vote")
	}

	var typeSwitch interface{} = "Golang" // interface{} is a type that can hold any value(something like type any in TypeScript)

	// type switch statement : check the type of the variable and execute the corresponding case
	switch typeSwitch.(type) {
	case int:
		fmt.Println("This is an int")
	case float64:
		fmt.Println("This is a float")
	case string:
		fmt.Println("This is a string")
	case bool:
		fmt.Println("This is a bool")
	default:
		fmt.Println("This is something else")
	}

	// fallthrough statement
	switch age {
	case 18:
		fmt.Println("You can vote")
		fallthrough
	case 21:
		fmt.Println("You can vote")
	default:
		fmt.Println("You can't vote")
	}

}
