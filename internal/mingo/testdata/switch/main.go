package main

import "fmt"

func main() {
	i := 0

	switch i {
	case 0:
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	case 1, 2:
		fmt.Println("4")
		fmt.Println("5")
		fmt.Println("6")
	default:
		fmt.Println("7")
		fmt.Println("8")
		fmt.Println("9")
	}
}
