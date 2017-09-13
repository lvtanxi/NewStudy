//反射
package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Id   int
	Name string
	Age  int
}

func (u user) Say(name string) {
	fmt.Println("Hello", name, "my name is", u.Name)
}

type manger struct {
	user
	title string
}

func (u user) Hello() {
	fmt.Println("Hello world")
}

func main() {
	u := user{1, "lv", 28}
	info(u)
	changeElemValue(&u)
	fmt.Println(u)
	u.Say("test")
	callMethod(&u)
}

func callMethod(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr {
		fmt.Println("不满足方法调用的条件")
		return
	}
	mv := v.MethodByName("Say")
	if !mv.IsValid() {
		fmt.Println("bad method")
		return
	}
	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}

func changeElemValue(o interface{}) {
	//获取结构的值
	v := reflect.ValueOf(o)
	//判断是否是指针接口 和属性能否被修改
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("不满足修改属性值的条件")
		return
	}
	//获取实际对象
	v = v.Elem()
	//根据名称获取属性
	f := v.FieldByName("Name")
	//判断属性是否攒竹
	if !f.IsValid() {
		fmt.Println("bad field")
		return
	}
	//判断
	if f.Kind() == reflect.String {
		f.SetString("修改了-byebye")
	}
}

//通过反射获取结构的信息
func info(o interface{}) {
	//获取传入类型
	t := reflect.TypeOf(o)
	//获取类型名称
	fmt.Println("Type:", t.Name())
	//判断是否是结构体，反射只能针对值，不能针对指针
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("o is not interface")
		return
	}

	//获取属性集合
	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		//获取字段
		f := t.Field(i)
		//获取字段值
		val := v.Field(i).Interface()
		fmt.Printf("%s %v = %v \n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		//获取方法
		m := t.Method(i)
		//打印方法的名称和类型
		fmt.Printf("%s %v \n", m.Name, m.Type)
	}
}
