package main

import "fmt"

func main() {
	fmt.Println("stacking")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
