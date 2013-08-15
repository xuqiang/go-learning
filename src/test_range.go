package main

import "fmt"

func main() {
	list := []string { "a", "b", "c", "d" }
	for k, v := range list {
		// 其中k是下标，1,2，
		// v是遍历的具体指 a b 
		println(k)
		println(v)
	}
	
	s := "中国人"
	for pos, char := range s {
		fmt.Printf("char '%c' start at byte position %s\n", char, pos)
	}
}
