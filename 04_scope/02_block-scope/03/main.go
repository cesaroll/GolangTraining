package main

import "fmt"

func main() {

	x := 0

	//Anonimous method
	//Is a function without a name
	increment := func() int {

		x++
		return x

	}

	fmt.Println(increment())
	fmt.Println(increment())

}
