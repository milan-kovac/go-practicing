package main

import "fmt"

/*
In the Go programming language, at least one package must have a main function,
and it is the function that represents the entry point of the program's execution.
*/
func main(){
    // Variable types
    const price = 500; // the value does not change during program execution
    var quantity = 5; // the value changes during program execution

    var wallet int = 5000;
    
    total := price * quantity; // := is a shorthand way to declare a new variable and assign it a value at the same time.
    
    // IF - ELSE IF / ELSE
    if total > wallet {
        fmt.Println("You don't have enough money.")
    } else {
        wallet -= total;
        fmt.Println("The remaining amount in the wallet is :" + fmt.Sprint(wallet))
    }

    // Basic types

    // int can be negative and uint cannot.
    const a int = 10; // int8 - int16 - int32 - int64
    const b uint = 10; // uint8 - uint16 - uint32 - uint64
    const c float32 = 10.5; // float64
     
    /*
    In Go programming, complex numbers consist of two parts: a real part and an imaginary part, 
    typically written as a + bi, where a is the real part, b is the imaginary part, and i is the imaginary unit satisfying 𝑖 2 = − 1 i 2 =−1.
    */
    const d complex64 = 1 + 2i; // complex128 
    const e string = "string";
    const f bool = true;
    
    /*
    Byte is an 8-bit unsigned type that represents a single character in ASCII encoding, 
    while rune is a 32-bit signed type that represents a single Unicode character, allowing support for a wider range of characters.
    */
    const g byte = 'A';
    const h rune = '我'

    var ab string; // This will be nil, a non-existent value
    print(ab)
    
    // Calling a function
    print("Text")
}

func print(text string){
    fmt.Println(text);
}

// In order for a method to have a return value, a type must be defined
func sum(a int, b int) int {
    return a + b
}