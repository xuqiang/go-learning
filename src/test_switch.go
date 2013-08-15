package main


func main() {
	i := 10
	switch i {
		case 0:
			println("i==0")
		case 1:
			println("i==1")
			// 没有i匹配的值, 不执行任何操作
	}

	switch i {
	case 1 : fallthrough
	case 10 :
		// 使用fallthrough关键字，i==1和i==10共享处理block
		println("i==1 or i==10")
	}

	switch i {
	case 1, 2 : 
		// case中可以并列条件
		println("i==1 or i==2")
	}

	switch i {
	case 1:
		println("i==1")
	default:
		println("default action")
	}
}
