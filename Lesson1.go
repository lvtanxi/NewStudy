//当前程序的包名
package main
//导入其他的包名
import "fmt"

//常量的定义
const PI  = 2.14

//全局变量
var name  = "gopher"

//一般类型申明
type newType int

//结构的申明
type goher struct {

}

//接口的什么
type golang interface {

}

//常量组的定义
const (
	lv = 123
	tan ="123"
	xi = 456
)
//全局变量组的定义
var (
	name0 = "123"
	name1 = "123"
	name2 = "123"
)

//一般类型组的申明

type (
	type0 int
	type1 float32
	type2 string
)


//主程序入口与package main 联合使用
func main() {
	fmt.Println("Hello world ! 我是你爸爸~！")
}
