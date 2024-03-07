package main

func main() {
	for range 10 {
		break
	}

	for i := range 10 {
		println(i)
	}

	for _, s := range []string{"a", "b", "c"} {
		println(s)
	}

	for i, s := range []string{"a", "b", "c"} {
		println(i, s)
	}

	for i := range []string{"a", "b", "c"} {
		println(i)
	}
}
