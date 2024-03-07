package main

import "fmt"

func main() {
	s := []string{"A", "B", "C"}

	fmt.Println(s[0])
	fmt.Println(s[0:1])
	fmt.Println(s[0:])
	fmt.Println(s[:1])
	fmt.Println(s[:])
}
