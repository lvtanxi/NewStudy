//defer
package main

import "fmt"

func main() {
	/*deferOrder()*/
	/*	deferParam()*/

	/*	deferParam2()//*/

	deferAll()

}
func deferAll() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}
	for _, f := range fs {
		f()
	}
}

func deferParam2() {
	for i := 0; i < 3; i++ {
		fmt.Printf("函数中的i = %d \n", i)
		defer func(y int) { fmt.Println(y) }(i) //由于i是参数，所以获取的是i当时的值
	}
}

// defer 调用函数
func deferParam() {
	for i := 0; i < 3; i++ {
		fmt.Printf("函数中的i = %d \n", i)
		defer func() { fmt.Println(i) }() //由于i不是参数，所以获取的是一个地址，所以i为最后的3
	}
}

//查看defer执行顺序
func deferOrder() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
