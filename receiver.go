package main
import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

//Area method with a value receiver of type Rectangle struct
// The receiver is a value receiver because the method does not modify the struct
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
// Scale method with a pointer receiver
// The receiver is a pointer receiver because the method modifies the struct
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func Receiver() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.Area()) 
	rect.Scale(2)
	fmt.Println("Area:", rect.Area()) 
	fmt.Println("Width:", rect.Width) 
	fmt.Println("Height:", rect.Height) 
}