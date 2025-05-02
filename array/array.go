package main

import "fmt"

/* Array in Go is a fixed-size collection of elements of the same type.
   Where as a slice is a dynamically-sized, flexible view into the elements of an array.
   So we will go through slices in depth later on.
*/

func main() {
	// array of integers
	var arr [3]int // in go by default arrays are initialized with 0 if int, 0.0 if float, "" if string, false if bool, nil if pointer and so on.
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// array of strings
	var arr2 [3]string
	arr2[0] = "Hello"
	arr2[1] = "World"
	arr2[2] = "!"
	fmt.Println(arr2)

	// array of bools
	var arr3 [3]bool
	arr3[0] = true
	arr3[1] = false
	arr3[2] = true
	fmt.Println(arr3)

	// array of floats
	var arr4 [3]float64
	arr4[0] = 1.1
	arr4[1] = 2.2
	arr4[2] = 3.3
	fmt.Println(arr4)

	// array of pointers
	var arr5 [3]*string // pointers to strings
	arr5[0] = new(string)
	*arr5[0] = "Hello"
	arr5[1] = new(string)
	*arr5[1] = "World"
	arr5[2] = new(string)
	*arr5[2] = "!"
	fmt.Println(arr5)

	// array of interfaces
	var arr6 [3]interface{}
	arr6[0] = 1
	arr6[1] = "Hello"
	arr6[2] = true
	fmt.Println(arr6)

	// array of structs
	type Person struct {
		name string
		age  int
	}
	var arr7 [3]Person
	arr7[0] = Person{"John", 30}
	arr7[1] = Person{"Jane", 25}
	arr7[2] = Person{"Bob", 40}
	fmt.Println(arr7)

	// 2d array
	var arr8 [2][3]int
	arr8[0][0] = 1
	arr8[0][1] = 2
	arr8[0][2] = 3
	arr8[1][0] = 4
	arr8[1][1] = 5
	arr8[1][2] = 6
	fmt.Println(arr8)

	// 3d array
	var arr9 [2][3][4]int
	arr9[0][0][0] = 1
	arr9[0][0][1] = 2
	arr9[0][0][2] = 3
	arr9[0][0][3] = 4
	arr9[0][1][0] = 5
	arr9[0][1][1] = 6
	arr9[0][1][2] = 7
	arr9[0][1][3] = 8
	arr9[1][0][0] = 9
	arr9[1][0][1] = 10
	arr9[1][0][2] = 11
	arr9[1][0][3] = 12
	arr9[1][1][0] = 13
	arr9[1][1][1] = 14
	arr9[1][1][2] = 15
	arr9[1][1][3] = 16
	fmt.Println(arr9)

	// array of arrays
	var arr10 [3][4]int
	arr10[0][0] = 1
	arr10[0][1] = 2
	arr10[0][2] = 3
	arr10[0][3] = 4
	arr10[1][0] = 5
	arr10[1][1] = 6
	arr10[1][2] = 7
	arr10[1][3] = 8
	arr10[2][0] = 9
	arr10[2][1] = 10
	arr10[2][2] = 11
	arr10[2][3] = 12
	fmt.Println(arr10)

}
