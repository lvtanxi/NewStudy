// 控制语句
package main

import (
	"fmt"
)

func main() {
	a := true
	if a, b, c := 1, 2, 3; a+b+c > 6 {
		fmt.Println("大于6")
	} else {
		fmt.Println("小于等于6")
		fmt.Println(a)
	}
	fmt.Println(a)

	// 无限循环
	aa := 1
	for {
		aa ++
		if aa > 3 {
			break
		}
		fmt.Println(aa)
	}
	fmt.Println("无限循环over")

	//简化版
	aa = 1
	for aa <= 3 {
		aa ++
		fmt.Println(aa)
	}
	fmt.Println("简化版over")

	//普通版
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("普通版over", aa)

	//普通版本
	switch aa {
	case 0:
		fmt.Println("aa=0")
	case 1:
		fmt.Println("aa=1")
	default:
		fmt.Println("None")
	}

	// case + 表达式 +fallthrough
	switch {
	case aa >= 0:
		fmt.Println("a>=0")
		fallthrough
	case aa >= 3:
		fmt.Println("a>=3")
	default:
		fmt.Println("None")
	}

	// 赋值语句

	switch a := 3; {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 3:
		fmt.Println("a>=3")
	default:
		fmt.Println("None")
	}
	fmt.Println(a)

LABLE1:
	for {
		for i := 0; i < 10; i++ {
			fmt.Println("before break ", i)
			if i > 3 {
				break LABLE1
			}
		}
	}
	fmt.Println("break 跳出了无限多层循环")

	for {
		for i := 0; i < 10; i++ {
			fmt.Println("before goto ", i)
			if i > 3 {
				goto LABLE2
			}
		}
	}
LABLE2:
	fmt.Println("goto 跳出了无限多层循环")

LABLE3:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println("before continue", i)
			continue LABLE3
		}
	}
	fmt.Println("continue 跳出了无限多层循环")
}
