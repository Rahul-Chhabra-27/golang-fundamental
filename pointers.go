package main

import "fmt"
// syntax for declaring pointers
// var pointer *datatype
// var ptr *int
// var ptr1 *string
// In functions we can pass the address of a variable as an argument to a function
// and the function can change the value of the variable at that address
// syntax for passing pointers to a function
// func functionName(pointer *datatype)
func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func double(x *int) *int {
	result := *x * 2
	return &result
}
// pointer with struct
type person struct {
	name string
	age  int
}

// function that takes a pointer to a struct as an argument
func changeName(p *person) {
	p.name = "Alice"
}
func Pointers() {
	x := 1
	y := 2
	fmt.Println(x, y) // output: 1 2
	swap(&x, &y)
	fmt.Println(x, y) // output: 2 1
	ptr := double(&x)
	fmt.Println(*ptr) // output: 4

	// pointer with struct
	p := person{name: "Bob", age: 30}
	fmt.Println(p) // output: {Bob 30}
	changeName(&p)
	fmt.Println(p) // output: {Alice 30}
}
