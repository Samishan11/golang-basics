package main

import "fmt"

func main() {

	/* for loop traditional way of writing a loop in Go */
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/*In Go there is no while loop but you can use a for loop to achieve the same thing as a while loop */
	fmt.Println("While loop--")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	/*In Go there is no do while loop but you can use a for loop to achieve the same thing as a do while loop */
	fmt.Println("Do while loop--")
	i = 0
	for {
		fmt.Println(i)
		i++
		if i > 5 {
			break
		}
	}

	/* Range loop in Go is used to iterate over a slice, map, or channel and perform an action on each element. Note: We will go through map, slice and channel in depth later on. */
	fmt.Println("Range loop--")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // slice of integers
	for _, value := range slice {
		fmt.Println(value)
	}
}
