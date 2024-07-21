package main

import "fmt"

// function with return type
func greetMessage(name string) string {
    return "Hello, " + name + "!"
}
// function without return type
func printMessage(message string) {
    fmt.Println(message)
}
// funtions with defer keyword
func deferFunction1() {
    fmt.Println("Defer Function 1")
}
func deferFunction2() {
    fmt.Println("Defer Function 2")
}

func Caller() {
    result := greetMessage("World")
    fmt.Println("Result: ", result)
    // calling function without return type

    printMessage("Hello, World!")
    // Immediately Invoked Function Expression (IIFE)
    // This function will be called immediately after its declaration and definition 
    // and it will not be called again in the program execution 
    
func() {
        fmt.Println("IIFE")
    }()


    // defer keyword
    // defer keyword is used to delay the execution of a function until 
    // the surrounding function returns
    // defer functions are executed in Last In First Out (LIFO) order
    defer deferFunction1()
    defer deferFunction2()
}


