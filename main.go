package main

import (
	"fmt"
	"strings"
)



func vinay(input string) string {
	return strings.ToUpper(input)
}

func factorial(n int) int {
    if n <= 0 {
        return 1
    }
    return n * factorial(n-1)
}

func isEven(number int) bool {
    return number%2 == 0
}


func main() {
	lowercaseString := "vinay kapoor"
	uppercaseString := vinay(lowercaseString)
	fmt.Println(uppercaseString)

	result := factorial(5)
   	fmt.Printf("Factorial of 5 is %d\n", result)

	  num := 6
    if isEven(num) {
        fmt.Printf("%d is even.\n", num)
    } else {
        fmt.Printf("%d is odd.\n", num)
    }

}
