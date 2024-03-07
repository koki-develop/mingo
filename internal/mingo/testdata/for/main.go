package main

func main() {
	for {
		break
	}

	for i := 0; i < 10; i++ {
		println(i)
	}

	for i := 0; i < 10; {
		println(i)
		i++
	}

	for i := 0; ; i++ {
		if i >= 10 {
			break
		}
		println(i)
	}

	i := 0
	for ; i < 10; i++ {
		println(i)
	}

	for i < 20 {
		println(i)
		i++
	}

	for true {
		break
	}
}
