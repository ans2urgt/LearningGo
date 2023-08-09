package main

import "fmt"

func main() {

	// regular for loop
	for i := 0; i < 9; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 3 {
		fmt.Println(j)
		j += 1
	}

	// for each in array
	var a = []int{1,2,3,4,5}
	fmt.Println(a)
	for i,s := range a {
		fmt.Println(i,s)
	}

	// for each in map
	fmt.Println("Map")
	m := map[string]int{
		"1": 1,
		"2": 2,
	}

	for x,s := range m {
		fmt.Println(x,s)
	}

}