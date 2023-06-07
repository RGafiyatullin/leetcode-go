package main

import "fmt"

func main() {
	s0 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s0)

	// Slice the slice to give it zero length.
	s1 := s0[:0]
	printSlice(s1)

	// Extend its length.
	s2 := s1[:4]
	printSlice(s2)

	// Drop its first two values.
	s3 := s2[2:]
	printSlice(s3)

	s0[0] = -2
	s0[5] = -13
	s4 := s3[0:4]
	fmt.Println(s4)
	fmt.Println(s3)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
