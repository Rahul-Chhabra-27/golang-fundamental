package main

import "fmt"
// global declaration of variables
var globalVar1 int = 10
var globalVar2 int = 20
var globalVar3 int = globalVar1 + globalVar2

// global declaration of variables with short hand
var (
	t1         int = 35
	t2         int = 35
	t3, t4, t6     = 23, 33, 43
)
// we can't use this way(num := 10) of declaraing variables outside the function body in golang 

func Datatypes() {
	
	// Ways to declare variables
	var a int = 10
	var b int = 20
	var c int = a + b

	fmt.Println("Sum of a and b is: ", c)
	// Short hand declaration
	d := 10
	e := 20
	f := d + e
	fmt.Println("Sum of d and e is: ", f)

	// Multiple variable declaration
	var g, h, i int = 10, 20, 30
	fmt.Println("Sum of g, h and i is: ", g+h+i)

	// Multiple variable declaration with short hand
	j, k, l := 10, 20, 30
	fmt.Println("Sum of j, k and l is: ", j+k+l)

	// printing global variables
	fmt.Println("Sum of globalVar1 and globalVar2 is: ", globalVar3);


	// string datatype
	var str string = "Hello, World!"
	fmt.Println(str)

	// short hand declaration of string
	str1 := "Hello, World!"
	fmt.Println(str1)


	// boolean datatype
	var bool1 bool = true
	fmt.Println(bool1)

	// short hand declaration of boolean
	bool2 := true
	fmt.Println(bool2)

	// float datatype
	var n float32 = 10.5
	fmt.Println(n)

	// short hand declaration of float
	o := 20.5
	fmt.Println(o)

	fmt.Println(t1 + t2 + t3 + t4 + t6)
}
