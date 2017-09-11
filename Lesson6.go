//切片 Slice
package main

import "fmt"

func main() {
	// slice的普通申明
	var s1 [] int
	fmt.Println(s1)

	a :=[10]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(a)

	//获取一个slice
	s2 := a[2]
	fmt.Println(s2)
	//从数组中获取切片 A:B,是一个前闭后开的一个区间 ,如果B不写表示从A到数组的最后长度，如果A不写就表示从0到B
	s3 :=a[5:]
	fmt.Println(s3)
	s4 :=a[:5]
	fmt.Println(s4)

	//利用make创建 slice
	s5 :=make([]int,3,10)
	fmt.Println(len(s5),cap(s5))

	s6 :=a[2:5]
	fmt.Println(s6)

	s7 :=a[3:8]
	fmt.Println(s7)

	//改变原数组的值，切片中的值随着改变,相反也成立
	a[6]=100
	fmt.Println(s7,a)

	s7[1]=200
	fmt.Println(a)

	//添加元素
	s7 =append(s7,1,2,3)
	fmt.Println(s7)


	//拷贝
	s8 :=[]int{1,2,3,4,5,6}
	s9 :=[]int{7,8,9}
	copy(s8,s9)
	fmt.Println(s8)

	s10 :=[]int{1,2,3,4,5,6}
	s11 :=[]int{7,8,9}
	copy(s11[2:4],s10[1:3])
	fmt.Println(s11)


}
