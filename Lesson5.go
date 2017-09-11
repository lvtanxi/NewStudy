// 数组
package main

import "fmt"

func main() {
	//申明
	var a [2]int
	b := [2]int{1, 1}
	fmt.Println(a, b)
	//指定某个元素的值
	c := [20]int{19: 12}
	fmt.Println(c)
	//不固定长度
	d := [...]int{1, 2, 3, 4}
	fmt.Println(d)

	//数组的指针
	var p *[20] int = &c
	fmt.Println(p)

	//指针数据
	x, y := 1, 2
	f := [...]*int{&x, &y}
	fmt.Println(f)

	fmt.Println(a == b)

	//数组的指针
	g := new([10]int)
	g[1] = 2
	fmt.Println(g)

	//多维数组
	gg := [2][3]int{{1, 1, 1}, {2, 2, 2}}
	fmt.Println(gg)

	ww := [...]int{5, 3, 2, 9}

	count := len(ww)
	for i := 0; i < count; i++ {
		for j := 0; j < i; j++ {
			if ww[j] > ww[i] {
				temp := ww[i]
				ww[i] = ww[j]
				ww[j] = temp
			}
		}
	}
	fmt.Println(ww)

}
