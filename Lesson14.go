//接口interface
package main

import "fmt"

//定义接口

type Usb interface {
	//接口中方法的申明
	getName() string
	//嵌入接口
	Connecter
}

type Connecter interface {
	connect()
}

//定义结构
type PhoneConnecter struct {
	name string
}

type TvConnecter struct {
	name string
}

//实现接口
func (pc PhoneConnecter) getName() string {
	return pc.name
}

func (pc PhoneConnecter) connect() {
	fmt.Println("connect", pc.name)
}

func (tv TvConnecter) connect() {
	fmt.Println("connect", tv.name)
}

func main() {
	pc := PhoneConnecter{"PhoneConnecter"}
	a := Connecter(pc)

	a.connect()
	disConnect2(a)

	tv := TvConnecter{"TvConnecter"}
	tv.name = "测阿斯顿"
	b := Connecter(tv)
	b.connect()
	disConnect2(b)

	var aa interface{}
	fmt.Println(aa == nil)

	var p *int = nil
	aa = p
	fmt.Println(aa == nil)

}

//简单的类型检查
func disConnect(usb Usb) {
	//利用ok pattern进行类型检查
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("disConnect", pc.name)
		return
	}

	fmt.Println("UnKnown decive")
}

//利用空接口来进行类型检查
func disConnect2(usb interface{}) {
	//利用type switch 类型检查
	switch target := usb.(type) {
	case PhoneConnecter:
		fmt.Println("disConnect", target.name)
	case TvConnecter:
		fmt.Println("disConnect", target.name)
	default:
		fmt.Println("UnKnown decive")
	}

}
