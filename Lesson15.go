//反射
package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Id   string
	Name string
	Age  int
}

func (u user) Hello() {
	fmt.Println("Hello world")
}

func main() {

}

func info(o interface{}) {
	//获取传入类型
	t :=reflect.TypeOf(o)
	//获取类型名称
	fmt.Println("Type:",t.Name())
	//获取属性集合
	v :=reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		//获取字段
		f:=t.Field(i)
		//获取字段值
		val :=v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v \n",f.Name,f.Type,val)
	}
}
