package main;
import "fmt"

// Structs are user-defined data types that are used to store a collection of data fields.
// Person is a struct that has four fields: FirstName, LastName, Email, and Age.
// The type of each field is specified after the field name.
// The fields of a struct can be accessed using the dot operator.
type Person struct {
	FirstName string
	LastName  string
	Email     string
	Age       uint
}

// NewPerson is a function that creates a new Person struct and returns it.
// The function takes four parameters: firstName, lastName, email, and age.
// The function returns a Person struct and an error.
func NewPerson(firstName, lastName, email string, age uint) (Person , error) {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Age:       age,
	}, nil;
}
func Structs() {
	person1 := Person{FirstName: "John", LastName: "Doe", Email: "johndoe@mail.com", Age: 43}
	person2 := Person{FirstName: "Alex", LastName: "Carey", Email: "carrybusinessmail.com", Age: 56}

	fmt.Printf("FirstName of person1 is %s \n", person1.FirstName)
	fmt.Printf("LastName of the person2 is %s\n", person2.LastName)

	// Structs can also be initialized without specifying the field names.
	person3 := Person{"Jane", "Doe", "jane@gmail.com", 30}
	fmt.Printf("Email of person3 is %s\n", person3.Email)

	// Can Change the value of a field
	person3.Email = "janedoe@gmail.com";
	fmt.Printf("Updated Email of person3 is %s\n", person3.Email)

	// Create a new Person struct using the NewPerson function.
	person4, _ := NewPerson("Alice", "Smith", "alice@mial.com", 25)
	fmt.Printf("FirstName of person4 is %s\n", person4.FirstName)
}