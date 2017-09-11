//运算符
package main

import "fmt"

/**
	四个比较特殊的运算符号
		6: 0110
	  11: 1011
		& 0010 2
		| 1111 15
		^ 1101 13
	  &^ 0100  4
 */

func main() {
	fmt.Println(^2)
	fmt.Println(1 ^ 2)
	fmt.Println(2 << 3)
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)
}
