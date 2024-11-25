package main

import "fmt"

func main() {
    // Arrays have a fixed length while slices have a dynamic length

	// defining array
	fixedLength := [5]int{0, 1, 2, 3, 4}
	fmt.Println(fixedLength)

	// defining slices
	dynamicLength := []int{0, 1, 2, 3, 4, 8}
	fmt.Println(dynamicLength)
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// a[low : high]
	var sliceA []int = dynamicLength[1:4] // [1,2,3] FROM 1 TO 4 but without 4
	fmt.Println(sliceA)
	var sliceB []int = dynamicLength[:4] // [0,1,2,3] FROM FIRST TO 4 but without 4
	fmt.Println(sliceB)
	var sliceC []int = dynamicLength[3:] // [3,4,8] FROM 3 TO LAST
	fmt.Println(sliceC)

}