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
	F7 func(f1 func() string)
}

func main() {
	type S3 struct {
		S1

		A string
		B int
		C bool
	}
}
`,
			want: `package main;type S1 struct{S struct{A string;B int;C bool};F1 func()string;F2 func()(string,error);F3 func()(s string,err error)};type S2 struct{S1;A string;B int;C bool;a string;b int;c bool;F4 func(s string);F5 func(string);F6 func(s1,s2 string);F7 func(f1 func()string)};func main(){type S3 struct{S1;A string;B int;C bool};};`,
		},
		{
			name: "interface",
			src: `package main

type I1 interface {
	F1() string
	F2() (string, error)
	F3() (s string, err error)
}

type I2 interface {
	I1

	F4(s string)
	F5(string)
	F6(s1, s2 string)
	F7(s1 string, s2 string)
	F8(s1, s2 string, i int)
	F9(f func() string)
}
`,
			want: `package main;type I1 interface{F1()string;F2()(string,error);F3()(s string,err error)};type I2 interface{I1;F4(s string);F5(string);F6(s1,s2 string);F7(s1 string,s2 string);F8(s1,s2 string,i int);F9(f func()string)};`,
		},
		{
			name: "go",
			src: `package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello, 世界")
	time.Sleep(1 * time.Second)
}

func main() {
	go func() {
		fmt.Println("Hello, 世界")
		time.Sleep(1 * time.Second)
	}()

	go hello()
}
`,
			want: `package main;import("fmt";"time");func hello(){fmt.Println("Hello, 世界");time.Sleep(1*time.Second)};func main(){go func(){fmt.Println("Hello, 世界");time.Sleep(1*time.Second)}();go hello()};`,
		},
		{
			name: "func",
			src: `package main

import "fmt"

func hello() {
	func() {
		fmt.Println("Hello")
	}()

	func(s string) {
		fmt.Printf("Hello, %s\n", s)
	}("world")

	if err := func() error {
		return nil
	}(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	go hello()
}
`,
			want: `package main;import "fmt";func hello(){func(){fmt.Println("Hello")}();func(s string){fmt.Printf("Hello, %s\n",s)}("world");if err:=func()error{return nil}();err!=nil{fmt.Println(err)}};func main(){go hello()};`,
		},
		{
			name: "goto",
			src: `package main

import "fmt"

func main() {
	fmt.Println("A")
	goto label
	fmt.Println("B")
label:
	fmt.Println("C")
}
`,
			want: `package main;import "fmt";func main(){fmt.Println("A");goto label;fmt.Println("B");label:fmt.Println("C")};`,
		},
		{
			name: "generic",
			src: `package main

import "fmt"

func Ptr[T any](p T) *T {
	return &p
}

func Equals[T, U comparable](t T, u U) bool {
	return t == u
}

func Hoge[T comparable, U fmt.Stringer](s []T, e U) bool {
	return true
}

func main() {
	fmt.Println(Ptr(1))
}
`,
			want: `package main;import "fmt";func Ptr[T any](p T)*T{return &p};func Equals[T,U comparable](t T,u U)bool{return t==u};func Hoge[T comparable,U fmt.Stringer](s []T,e U)bool{return true};func main(){fmt.Println(Ptr(1))};`,
		},
		{
			name: "const",
			src: `package main

const (
	A           = 1
	B       int = 2
	C, D    int = 3, 4
	E, F, G     = 3, 4, 5
)

const (
	_ = iota
	a
	b
	c
)

const w = 1
const x int = 2
const y, z = 3, 4

func main() {
	const a = 1
	const b, c = 2, 3
	const (
		d = 4
		e = 5
	)
}
`,
			want: `package main;const(A=1;B int=2;C,D int=3,4;E,F,G=3,4,5);const(_=iota;a;b;c);const w=1;const x int=2;const y,z=3,4;func main(){const a=1;const b,c=2,3;const(d=4;e=5);};`,
		},
		{
			name: "var",
			src: `package main

var (
	A           = 1
	B       int = 2
	C, D    int = 3, 4
	E, F, G     = 3, 4, 5
)

var w = 1
var x int = 2
var y, z = 3, 4

func main() {
	var a int = 1
	var b, c int = 2, 3
	var (
		d = 4
		e = 5
	)
}
`,
			want: `package main;var(A=1;B int=2;C,D int=3,4;E,F,G=3,4,5);var w=1;var x int=2;var y,z=3,4;func main(){var a int=1;var b,c int=2,3;var(d=4;e=5);};`,
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
