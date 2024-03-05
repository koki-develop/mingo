package mingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Minify(t *testing.T) {
	testcases := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "if statements",
			src: `package main

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
`,
			want: `package main;import "fmt";func main(){if true{fmt.Println("TRUE");if true{fmt.Println("TRUE")}};if false{fmt.Println("TRUE")}else{fmt.Println("FALSE")};if false{fmt.Println("TRUE")}else if true{fmt.Println("TRUE")};if err:=run1();err!=nil{fmt.Println(err)};if _,err:=run2();err!=nil{fmt.Println(err)}};func run1()error{return nil};func run2()(string,error){return "",nil};`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Minify("main.go", []byte(tc.src))

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
