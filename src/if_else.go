package main

import "fmt"

func main() {
	var x int = 10
	if x >= 10 { /* {必须和if同行 */
		fmt.Printf("x >= 10")
	} else { /* 这里需要在同一行 */
		fmt.Printf("x < 10")
	}
}
