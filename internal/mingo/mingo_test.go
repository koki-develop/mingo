package mingo

import (
	"go/format"
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
			name: "if",
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
		{
			name: "inc dec",
			src: `package main

import "fmt"

func main() {
	i := 1
	i++
	i--
	i++
	fmt.Println(i)
}
`,
			want: `package main;import "fmt";func main(){i:=1;i++;i--;i++;fmt.Println(i)};`,
		},
		{
			name: "switch",
			src: `package main

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
`,
			want: `package main;import "fmt";func main(){i:=0;switch i{case 0:fmt.Println("1");fmt.Println("2");fmt.Println("3");case 1,2:fmt.Println("4");fmt.Println("5");fmt.Println("6");default:fmt.Println("7");fmt.Println("8");fmt.Println("9")}};`,
		},
		{
			name: "struct",
			src: `package main

type S1 struct {
	S struct {
		A string
		B int
		C bool
	}

	F1 func() string
	F2 func() (string, error)
	F3 func() (s string, err error)
}

type S2 struct {
	S1

	A string
	B int
	C bool

	a string
	b int
	c bool

	F4 func(s string)
	F5 func(string)
	F6 func(s1, s2 string)
}`,
			want: `package main;type S1 struct{S struct{A string;B int;C bool};F1 func()string;F2 func()(string,error);F3 func()(s string,err error)};type S2 struct{S1;A string;B int;C bool;a string;b int;c bool;F4 func(s string);F5 func(string);F6 func(s1,s2 string)};`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Minify("main.go", []byte(tc.src))

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)

			// check syntax
			_, err = format.Source([]byte(got))
			assert.NoError(t, err)
		})
	}
}
