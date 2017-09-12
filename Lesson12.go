//嵌入结构struct,模拟继承
package main

import "fmt"

//公共部分
type human struct {
	sex int
}

type teacher struct {
	human
	name string
	age  int
}

type student struct {
	human //嵌入struct
	name string
	age  int
}

type bbb struct {
	human
	sex int
}

func main() {
	//第一种初始化嵌入struct，匿名的默认就是结构名名称作为字段名称
	a := teacher{name: "teacher", age: 33, human: human{sex: 0}}
	a.human.sex = 0

	//第二种初始化嵌入struct，可以直接获取嵌入嵌入struct(前提是没有相同的属性名)
	b := student{name: "teacher", age: 33}
	b.sex=1

	fmt.Println(a,b)

	//不同级别中相同属性字段
	bbq :=bbb{sex:12,human:human{sex:15}}
	fmt.Println(bbq.sex,bbq.human.sex)


}
