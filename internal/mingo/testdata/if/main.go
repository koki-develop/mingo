package main

import "fmt"

func main() {
	if true {
		fmt.Println("TRUE")
		if true {
			fmt.Println("TRUE")
		}
	}

	if false {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}

	if false {
		fmt.Println("TRUE")
	} else if true {
		fmt.Println("TRUE")
	}

	if err := run1(); err != nil {
		fmt.Println(err)
	}

	if _, err := run2(); err != nil {
		fmt.Println(err)
	}
}

func run1() error {
	return nil
}

func run2() (string, error) {
	return "", nil
}
