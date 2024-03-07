package main

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
