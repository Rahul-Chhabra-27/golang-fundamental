package main


import (
    "fmt"
    "time"
)


func cook(name string,message chan string) {
    fmt.Println("Cooking Routine started:", name)
    // Sleep pauses the current goroutine for at least the duration d.
    // A negative or zero duration causes Sleep to return immediately.
    time.Sleep(2 * time.Second /** duration d**/)
    fmt.Println("Cooking Routine finished:", name)
	message <- "Pizza is ready"
}


func GoRoutines() {
    fmt.Printf("Starting of Main Routine\n")
	fmt.Println("Cooking started: Pasta")
	message := make(chan string)
    go cook("Pizza",message) // Start cooking pizza in a new go routine
	fmt.Println("Message from the cook:", <-message)
    fmt.Printf("Completing the Main Routine\n")
}
