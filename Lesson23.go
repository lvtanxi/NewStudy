/**  
* Date: 2017-10-09 
* Time: 17:46 
* Description: string操作
*/
package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	testStrings()
	testStrconv()
}

//strconv 包中的基本运用
func testStrconv() {
	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中
	str := make([]byte, 0, 100)
	str = strconv.AppendBool(str, false)
	fmt.Println(string(str))
	str = strconv.AppendInt(str, 4567, 10)
	fmt.Println(string(str))
	str = strconv.AppendQuote(str, "append test")
	fmt.Println(string(str))
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	//Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1024)
	fmt.Println(a, b, c, d, e)

	//Parse 系列函数把字符串转换为其他
	aa, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println(err)
	}
	bb, err := strconv.ParseFloat("123.23", 64)
	if err != nil {
		fmt.Println(err)
	}
	cc, err := strconv.ParseInt("7890", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	dd, err := strconv.ParseUint("12345", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ee, err := strconv.Atoi("1024")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aa, bb, cc, dd, ee)

}

//strings 包中的基本运用
func testStrings() {
	//字符串s中是否包含substr，返回bool值
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	//字符串链接，把slice a通过sep链接起来
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ","))
	//在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	//重复s字符串count次，最后返回重复的字符串
	fmt.Println("ba" + strings.Repeat("na", 2))
	//在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moon", -1))
	//把s字符串按照sep分割，返回slice
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	//在s字符串中去除cutset指定的字符串
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "!  "))
	//去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Printf("Fields are : %q\n", strings.Fields("   foo  bar  baz  "))
}
