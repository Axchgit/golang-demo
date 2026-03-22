package main

import "fmt"

func main() {

	fmt.Println(123)

	mp := map[int]string{
		1: "a",
		2: "b",
	}

	for index, value := range mp {
		fmt.Println(index, value)
	}

	fmt.Println(mp)

}
