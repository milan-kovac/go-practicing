package main

import (
	"complexConcepts/myMath"
	"fmt"
)

func main() {
	// Pointer
	var x int = 55
	fmt.Println(&x); // this character & is a special way to specify the memory location of this variable
	var ptr *int = &x;
	fmt.Println(ptr); // this character * is a special character that indicates the type that the value will be a memory location 
	fmt.Println(*ptr) // if we add * in front of it, we will dereference and get the actual value and not the memory location

	var total = myMath.Sum(50, 50);
	fmt.Println(total)
}