//Golang 数据类型与初始值
package main

import (
	"fmt"
	"math"
	"strconv"
)

const la = 1

const lb = "A"

const (
	lc     = 2
	ld     = la + 1
	lw
	lz
	le, lf = 1, 2
)
const (
	//第一个常量不可省略表达式
	Monday    = iota
	Tuesady
	Webnesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	xa = "A"
	xb
	xc = iota
	xd
)

const (
	b  = 1 << (iota * 10)
	kb
	mb
	gb
)

func main() {
	fmt.Println(b,kb, mb, gb)
	fmt.Println(xa, xb, xc, xd)

	var a int
	fmt.Println(a)

	var b string
	fmt.Println(b)

	var c []int
	fmt.Println(c)

	var d bool
	fmt.Println(d)

	var e byte
	fmt.Println(e)

	fmt.Println(math.MinInt8, math.MaxInt8)

	var f = "jaychou"
	fmt.Println(f)

	g := "测试赋值"
	fmt.Println(g)

	aa, bb := 65, 2
	fmt.Println(aa, bb)
	var cc, dd, ee int
	fmt.Println(cc, dd, ee)

	ff := 1.1
	fmt.Println(int(ff))

	fmt.Println(strconv.Itoa(aa))
	fmt.Println(strconv.Atoi("65"))
	fmt.Println(lz, lw, Sunday)

}
