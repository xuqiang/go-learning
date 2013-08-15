package main

func main() {
	var arr[10] int;
	arr[0] = 0
	arr[1] = 1
	
	// 或者是
	arr3 := [...]int{1, 2,3}
	// 这是go将自动推导类型，并指定长度
	println(len(arr3))

	// 二维数组
	a := [2][2] int{ [...]int{1, 2}, [...]int{3, 4} }
	println(a)
	b := [2][2] int{ {1, 2}, {3, 4} }
	println (b)
}
