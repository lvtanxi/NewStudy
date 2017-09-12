//方法method
package main

import "fmt"

type ma struct {
	name string
}

type mbb struct {
	name string
}

type tz int




func main() {
	a := ma{}
	a.print()
	fmt.Println(a)

	b := mbb{}
	b.print()
	fmt.Println(b)

	var t  tz
	t.print()
	t.add(100)
	fmt.Println(t)

	(*tz).print(&t)

}

func (tz *tz) print()  {
	fmt.Println("method tz")
}

func (t *tz) add(num int)  {
	*t += tz(num)
}


//为某个结构添加方法
func (ma *ma) print() {
	ma.name="AA"
	fmt.Println("this is my first method",ma)
}


func (mbb mbb) print() {
	mbb.name="BB"
	fmt.Println("不同的结构体可以有相同的方法名")
}
