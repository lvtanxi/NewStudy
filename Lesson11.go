//结构struct
package main

import "fmt"

//结构的申明
type person struct {
	name string
	age  int
	//匿名结构嵌套
	contact struct {
		phone, city string
	}
	//匿名字段
	string
}

func main() {
	//实例化一个结构体
	per := person{}
	fmt.Println(per)

	//属性操作
	per.name = "joe"
	per.age = 19
	fmt.Println(per)

	//实例化一个结构体和赋值
	a := person{name: "lvt", age: 123}
	fmt.Println(a)

	//一般情况，结构也是值类型
	changeStruceValue(a)
	fmt.Println(a)

	//通过指针的方式可以修改结构的某些属性
	changeStruceAddressValue(&a)
	fmt.Println(a)

	//一般在结构的初始化的时候添加&符号
	b := &person{name: "哈哈"}
	b.name = "ok"
	changeStruceAddressValue(b)
	fmt.Println(b)

	//匿名结构
	c := struct {
		name string
		age  int
	}{name: "匿名结构", age: 33}
	fmt.Println(c)

	//结构体内部匿名结构
	d :=person{name:"结构体内部匿名结构",age:12}
	//初始化结构体内部匿名结构
	d.contact.phone="13212"
	d.contact.city="成都"
	//匿名字段
	d.string="匿名字段"
	fmt.Println(d)
}

func changeStruceValue(person person) {
	person.age = 25
	fmt.Println("changeStruceValue", person)
}

func changeStruceAddressValue(person *person) {
	person.age = 25
	fmt.Println("changeStruceValue", person)
}
