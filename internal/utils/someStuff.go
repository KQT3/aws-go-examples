package utils

import "fmt"

func Daw() int {
	x := 10
	p := &x

	fmt.Println("Value of x:", x)
	fmt.Println("Value of p:", p)

	*p = 20
	fmt.Println("Value of x:", x)

	return x
}
