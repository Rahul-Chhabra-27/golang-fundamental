package main

import "fmt"

func ArraySlices() {
    // Declaring an array of integers with a fixed length of 5
    var a [5]int
    // Initializing an array with values
    b := [5]int{1, 2, 3, 4, 5}
    // Partial initialization (remaining elements are zero-valued)
    c := [5]int{1, 2, 3}
    // Accessing and modifying elements
    a[0] = 10
    fmt.Println(a[0]) 
    fmt.Println(b) 
    fmt.Println(c) 

	// slices 
	// Declaring a slice
	var d []int
	// Initializing a slice using make
	e := make([]int, 5)
	// Initializing a slice using a slice literal
	f := []int{1, 2, 3, 4, 5}
	// Accessing and modifying elements
	d = append(d, 10)

	fmt.Println(d[0])
	fmt.Println(e)
	fmt.Println(f)
}