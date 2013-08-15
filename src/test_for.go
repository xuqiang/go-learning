package main

func main() {
	s := 0
	for i := 1; i <= 10; i++ {
		s += i
	}

	for i := 1; i< 10; i++ {
		if i > 5  {
			continue
		}
		println(i)
	}
}
