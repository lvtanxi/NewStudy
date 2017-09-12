//函数
package main

import "fmt"

func main() {
	gg(1, 2, 3)
	a, b := 1, 2
	changeValue(a, b)
	fmt.Println(a, b)

	c := "京哈"
	changeStringValue(c)
	fmt.Println(c)

	s1 := []int{1, 2}
	changeArrayValue(s1)
	fmt.Println(s1)

	//函数当成变量
	aa := aa
	aa()

	//匿名海曙
	bbq := func() { fmt.Println("我是匿名海曙") }
	bbq()

	fmt.Println(closure(1)(2))

}

//闭包函数
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

//无参函数,无返回值
func aa() {
	fmt.Println("我这里是函数变量")
}

//有参数，无返回值
func bb(a int, b string) {

}

//有参数，单返回值
func cc(a int, b string) int {
	return 0
}

//有参数，多返回值
func dd(a int, b string) (int, string) {
	return 0, "1"
}

//有参数，多命名返回值
func ee(a, b, c int) (d, e, f int) {
	return a, b, c
}

//有参数，多命名返回值(一般不这样)
func ff(a, b, c int) (d, e, f int) {
	fmt.Println(a, b, c)
	d, e, f = 1, 2, 3
	return
}

//不定长参数
func gg(param ... int) {
	fmt.Println(param)
}

//修改基本数据，函数调用处的值不受影响
func changeValue(param ... int) {
	param[0] = 3
	param[1] = 4
	fmt.Println(param)
}

func changeStringValue(param string) {
	param = "修改了数据"
	fmt.Println(param)
}

//修改引用类型，函数调用处的值受影响
func changeArrayValue(param []int) {
	param[0] = 3
	param[1] = 4
	fmt.Println(param)
}
