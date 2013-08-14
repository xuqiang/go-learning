package main

func main() {
	var i int = 0
LABEL:
	println(i)	
	i++
	if i <= 10 {
		goto LABEL
	}
}
