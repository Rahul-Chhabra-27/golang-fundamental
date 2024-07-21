package calculatorpackage

func Add(a int, b int) int {
    return a + b;
}

func Subtract(a int, b int) int {
    return a - b;
}

func Multiply(a int, b int) int {
    return a * b;
}

func divide(a int, b int) int {
    if b == 0 {
        return -1;
    } else {
        return a / b;
    }
}