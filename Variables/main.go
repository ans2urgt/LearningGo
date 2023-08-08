package main

import "fmt"

func main() {
	// create two int variables and add
	var a int = 20
	var b int = 30
	fmt.Println(a + b)
	// create two float variables and add
	const x float64 = 10.01
	const y float64 = 20.2
	fmt.Println(x + y)
	// create two string variables and conctate
	var c string = "Hello "
	var d string = "World"
	fmt.Println(c + d)
	// create an array and print
	var j [10]string
	j[0] = "Hi"
	j[1] = "Ho"
	fmt.Println(j[0])
	fmt.Println(j[1])
}