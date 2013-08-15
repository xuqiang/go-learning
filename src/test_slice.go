package main

func main() {
	// 创建slice
	sl := make([]int, 10)	
	println(sl)

	// 另外slice可以从array或者其他的slice中创建
	a := [...]int {1,2 ,3, 4, 5}
	// 2:4表示从a下标为2的位置开始，也就是3，到下标是4-1位置结束，也就是4结束
	s2 := a[2:4]
	for i := 0; i < len(s2); i++ {
		println(s2[i])
	}
	// slice拓展
	// 此时s2是[3, 4]
	s3 := [...]int {6, 7}	
	s4 := append(s2, 10)

	// copy函数可以用来拷贝slice
	var a = [...]int { 0, 1, 2, 3, 4, 5, 6, 7 }
	var s = make([]int, 6)
	n1 := copy(s, a[0:])
	n1 := copy(s, s[0:])
}

